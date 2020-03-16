package base

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

//copied from tcpreader.ReaderStream

type wrapper struct {
	slice []byte
	left  int
	err   error
}

type ConcurrentReader struct {
	reader          *StreamReader
	readerBroadcast chan struct{}
	wrapperChan     chan *wrapper
	notifyChan      chan struct{}
	lastWrapper     *wrapper
	//todo this splice may increase to very large, how to release it ?
	cache       []byte
	cacheMutex  sync.Mutex
	ownerId     string
	readIndices map[string]int

	//as for the first bit, 0 means reader is nil and 1 means reader is not nil
	//as for the second bit, 0 means Concurrent Read is allowed, 1 means Concurrent Read is not allowed
	//as for the third bit, 0 means this stream is not validated, 1 means this stream has been validated (validate means whether user has discard invalid bytes by some way)
	//the following bits is un used for now
	status      uint32
	statusMutex sync.Mutex
}

var ReadBlockNotified = errors.New("this error means your read is notified, you should try read again")

type ReadNotAllowedErr struct {
	callerId string
	ownerId  string
}

func (r *ReadNotAllowedErr) Error() string {
	return fmt.Sprintf("{%s} is not allowed to read because oowner is {%s}", r.callerId, r.ownerId)
}

func newReadBlockNotified(callerId string, ownerId string) *ReadNotAllowedErr {
	return &ReadNotAllowedErr{
		callerId: callerId,
		ownerId:  ownerId,
	}
}

func (m *ConcurrentReader) Read(p []byte, callerId string) (n int, err error) {
	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()
	if m.allowConcurrentRead() {
		return m.concurrentRead(p, callerId)
	} else {
		return m.singleRead(p, callerId)
	}
}

func (m *ConcurrentReader) concurrentRead(p []byte, callerId string) (n int, err error) {
	i := m.readIndices[callerId]
	if i < len(m.cache) {
		c := copy(p, m.cache[i:])
		m.readIndices[callerId] += c
		return c, nil
	}
	c, err := m.doRead(p, callerId)
	if err != nil {
		if c > 0 {
			m.cache = append(m.cache, p[0:c]...)
		}
		return c, err
	}
	m.cache = append(m.cache, p[0:c]...)
	m.readIndices[callerId] += c
	return c, nil
}

func (m *ConcurrentReader) singleRead(p []byte, callerId string) (n int, err error) {
	checkSuccess := m.checkReadPermission(callerId)
	if checkSuccess == false {
		return 0, newReadBlockNotified(callerId, m.ownerId)
	}

	i := m.readIndices[callerId]
	if i < len(m.cache) {
		n = copy(p, m.cache[i:])
		m.readIndices[callerId] += n
	} else {
		// cache is useless
		m.cache = make([]byte, 0)
		n, err = m.doRead(p, callerId)
	}
	return
}

func (m *ConcurrentReader) doRead(p []byte, callerId string) (n int, err error) {
	if m.lastWrapper == nil {
		return m.readNewWrapper(p, callerId)
	}
	w := m.lastWrapper
	if w.left <= 0 {
		return m.readNewWrapper(p, callerId)
	}
	copied := copy(p, w.slice[len(w.slice)-w.left:])
	w.left -= copied
	return copied, nil
}

func (m *ConcurrentReader) readNewWrapper(p []byte, callerId string) (n int, err error) {
	var w *wrapper
	select {
	case w = <-m.wrapperChan:
	case _ = <-m.notifyChan:
		if !m.checkReadPermissionStrictly(callerId) {
			return 0, ReadBlockNotified
		}
	}
	if w == nil {
		// if notified bug this callerId has permit to read, try read again
		w = <-m.wrapperChan
	}

	if w == nil {
		return 0, io.EOF
	}
	if w.err != nil {
		return 0, err
	}
	copied := copy(p, w.slice)
	w.left -= copied
	m.lastWrapper = w
	return copied, nil
}

//after this method is called, only the Gorouting with the id equals goroutingId can read data from this reader
func (m *ConcurrentReader) SetOwner(callerId string) {
	if !m.allowConcurrentRead() {
		// meas this method has been called, should't be called again
		log.WithFields(map[string]interface{}{
			"newId":     callerId,
			"existedId": m.ownerId,
		}).Error("SetOwner should not be called muti times")
		return
	}
	m.statusMutex.Lock()
	defer m.statusMutex.Unlock()
	//first set ownerId
	//atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.ownerId)), &goroutingId)
	m.ownerId = callerId
	//forbiden concurrent reader
	m.forbiddenConcurrentRead()
	//interrupt gorouting who are reading data now
	m.interruptReading()
}

func (m *ConcurrentReader) SetOwnerWithoutInterrupt(callerId string) {
	if !m.allowConcurrentRead() {
		// meas this method has been called, should't be called again
		log.WithFields(map[string]interface{}{
			"newId":     callerId,
			"existedId": m.ownerId,
		}).Error("SetOwner should not be called muti times")
		return
	}
	m.statusMutex.Lock()
	defer m.statusMutex.Unlock()
	//atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&m.ownerId)), &goroutingId)
	m.ownerId = callerId
	//forbiden concurrent reader
	m.forbiddenConcurrentRead()
}

func (m *ConcurrentReader) checkReadPermission(callerId string) bool {
	m.statusMutex.Lock()
	defer m.statusMutex.Unlock()

	//ownerId is not set , every caller has permission to read
	if m.ownerId == "" {
		return true
	}
	//ownerId is set, only caller with the same callerId has permission to read
	if m.ownerId == callerId {
		return true
	}
	return false
}

func (m *ConcurrentReader) checkReadPermissionStrictly(callerId string) bool {
	m.statusMutex.Lock()
	defer m.statusMutex.Unlock()
	//ownerId is set, only caller with the same callerId has permission to read
	if m.ownerId == callerId {
		return true
	}
	return false
}

func (m *ConcurrentReader) interruptReading() {
	m.notifyChan <- struct{}{}
	defer func() {
		time.Sleep(100 * time.Millisecond)
		if len(m.notifyChan) == 1 {
			<-m.notifyChan
		}
	}()
}

//seek start index to index,
func (m *ConcurrentReader) SeekTo(index int, callerId string) {
	m.readIndices[callerId] = index
	//m.cache = m.cache[index:]
}

const concurrentReadAllowedChecker = 0x20000000

func (m *ConcurrentReader) forbiddenConcurrentRead() {
	for {
		oldVal := atomic.LoadUint32(&m.status)
		i := oldVal & concurrentReadAllowedChecker
		if i != 0 {
			return
		}

		newVal := oldVal | concurrentReadAllowedChecker
		if success := atomic.CompareAndSwapUint32(&m.status, oldVal, newVal); success {
			return
		}
	}
}

func (m *ConcurrentReader) allowConcurrentRead() bool {
	val := atomic.LoadUint32(&m.status)
	i := val & concurrentReadAllowedChecker
	if i == 0 {
		return true
	} else {
		return false
	}
}

func (m *ConcurrentReader) waitReaderReady() {
	<-m.readerBroadcast
}

func (m *ConcurrentReader) SetReaderReady() {
	close(m.readerBroadcast)
}

func (m *ConcurrentReader) startRead() {
	go func() {
		container := make([]byte, 1024)
		for {
			m.waitReaderReady()
			count, err := m.reader.Read(container)
			if err != nil {
				m.wrapperChan <- &wrapper{nil, 0, err}
				break
			}
			if count <= 0 {
				log.Warn("does this situation exist?")
				continue
			}
			temp := make([]byte, count)
			copy(temp, container[0:count])
			m.wrapperChan <- &wrapper{temp, count, nil}
		}
		close(m.wrapperChan)
	}()
}

func NewConcurrentReader(reader *StreamReader) *ConcurrentReader {
	r := new(ConcurrentReader)
	r.reader = reader
	r.readerBroadcast = make(chan struct{})
	r.wrapperChan = make(chan *wrapper, 2048)
	r.notifyChan = make(chan struct{}, 1)
	r.cache = make([]byte, 0, 1024)
	r.readIndices = make(map[string]int, 8)
	r.status = 0x0000
	r.startRead()
	return r
}
