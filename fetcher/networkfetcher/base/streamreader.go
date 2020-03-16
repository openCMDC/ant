package base

import (
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	log "github.com/sirupsen/logrus"
	"io"
)

type StreamReader struct {
	LossErrors   bool
	reassembled  chan []tcpassembly.Reassembly
	done         chan bool
	current      []tcpassembly.Reassembly
	closed       bool
	lossReported bool
	first        bool
	initiated    bool

	//added
	netLayer   gopacket.Flow
	transLayer gopacket.Flow
}

func NewReaderStream(netLayer, transLayer gopacket.Flow) StreamReader {
	r := StreamReader{
		reassembled: make(chan []tcpassembly.Reassembly),
		done:        make(chan bool),
		first:       true,
		initiated:   true,
		netLayer:    netLayer,
		transLayer:  transLayer,
	}
	return r
}

// Reassembled implements capturer.Stream's Reassembled function.
func (r *StreamReader) Reassembled(reassembly []tcpassembly.Reassembly) {
	log.WithFields(log.Fields{"netLayer": r.netLayer, "transportLayer": r.transLayer}).Trace("Reassembled")
	if !r.initiated {
		panic("ReaderStream not created via NewReaderStream")
	}
	r.reassembled <- reassembly
	<-r.done
}

// ReassemblyComplete implements capturer.Stream's ReassemblyComplete function.
func (r *StreamReader) ReassemblyComplete() {
	log.WithFields(log.Fields{"netLayer": r.netLayer, "transportLayer": r.transLayer}).Trace("read completely")
	close(r.reassembled)
	close(r.done)
}

// stripEmpty strips empty reassembly slices off the front of its current set of
// slices.
func (r *StreamReader) stripEmpty() {
	for len(r.current) > 0 && len(r.current[0].Bytes) == 0 {
		r.current = r.current[1:]
		r.lossReported = false
	}
}

// DataLost is returned by the ReaderStream's Read function when it encounters
// a Reassembly with Skip != 0.
var DataLost = errors.New("lost data")

// Read implements io.reader's Read function.
// Given a byte slice, it will either copy a non-zero number of bytes into
// that slice and return the number of bytes and a nil error, or it will
// leave slice p as is and return 0, io.EOF.

func (r *StreamReader) Read(p []byte) (int, error) {
	if !r.initiated {
		panic("ReaderStream not created via NewReaderStream")
	}
	var ok bool
	r.stripEmpty()
	for !r.closed && len(r.current) == 0 {
		if r.first {
			r.first = false
		} else {
			r.done <- true
		}
		if r.current, ok = <-r.reassembled; ok {
			r.stripEmpty()
		} else {
			r.closed = true
		}
	}
	if len(r.current) > 0 {
		current := &r.current[0]
		if r.LossErrors && !r.lossReported && current.Skip != 0 {
			r.lossReported = true
			return 0, DataLost
		}
		length := copy(p, current.Bytes)
		current.Bytes = current.Bytes[length:]
		return length, nil
	}
	return 0, io.EOF
}

// Close implements io.Closer's Close function, making ReaderStream a
// io.ReadCloser.  It discards all remaining bytes in the reassembly in a
// manner that's safe for the assembler (IE: it doesn't block).
func (r *StreamReader) Close() error {
	log.WithFields(log.Fields{"netLayer": r.netLayer, "transportLayer": r.transLayer}).Trace("Close streamreader")
	r.current = nil
	r.closed = true
	for {
		if _, ok := <-r.reassembled; !ok {
			return nil
		}
		r.done <- true
	}
}
