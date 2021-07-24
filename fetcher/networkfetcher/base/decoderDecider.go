package base

import (
	"ant/core"
	"ant/fetcher/base"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

// key is line splitor
var textBasedDecoders = make(map[string][]DecoderFactoryFunc)

// key is valid length
var binaryBasedDecoders = make(map[int][]DecoderFactoryFunc)

func RegisterTextBasedDecoder(splitStr string, decoder DecoderFactoryFunc) {
	textBasedDecoders[splitStr] = append(textBasedDecoders[splitStr], decoder)
}

func RegisterBinaryBasedDecoder(length int, decoder DecoderFactoryFunc) {
	binaryBasedDecoders[length] = append(binaryBasedDecoders[length], decoder)
}

type DecoderGroup interface {
	findAdaptiveDecoder(bytes []byte) (decoder Interface, find bool)
	getValidRow() []byte
	reset()
}

type decoderGroup struct {
	currentRow []byte
	decoders   []Interface
	decider    *Decider
}

type TextBasedDecoderGroup struct {
	decoderGroup
	splitBytes []byte
}

func NewTextBasedDecoderGroup(splitStr string, decider *Decider, decoders []Interface) DecoderGroup {
	r := new(TextBasedDecoderGroup)
	r.currentRow = make([]byte, 0, 20)
	r.decoders = decoders
	r.decider = decider
	r.splitBytes = []byte(splitStr)
	return r
}

func (c *decoderGroup) getValidRow() []byte {
	return c.currentRow
}

func (c *decoderGroup) reset() {
	c.currentRow = make([]byte, 0, 20)
}

func (t *TextBasedDecoderGroup) findAdaptiveDecoder(bytes []byte) (decoder Interface, find bool) {
	sl := len(t.splitBytes)
	var i = 0
	for i < len(bytes) {
		if i+sl > len(bytes) {
			//剩余字符不足以形成一个有效的切割字符, 直接添加到当前行
			t.currentRow = append(t.currentRow, bytes[i])
			i++
		} else if BytesEquals(bytes[i:i+sl], t.splitBytes) {
			//当前行有效，校验当前行是否能匹配上某个decoder
			for _, decoder := range t.decoders {
				if processed, success := decoder.IsValidBeginningBytes(t.currentRow, t.decider.streamType); success {
					//某个decoder校验成功，返回该decoder
					t.currentRow = append(processed, bytes[i:]...)
					return decoder, true
				}
			}
			//如果都无法成功匹配，则清空并重新开始。
			i += sl
			t.currentRow = t.currentRow[:0]
		} else {
			t.currentRow = append(t.currentRow, bytes[i])
			i++
		}
	}
	return nil, false
}

type BinaryBasedDecoderGroup struct {
	decoderGroup
	length int
}

func NewBinaryBasedDecoderGroup(length int, decider *Decider, decoders []Interface) DecoderGroup {
	r := new(BinaryBasedDecoderGroup)
	r.currentRow = make([]byte, 0, 20)
	r.decoders = decoders
	r.decider = decider
	r.length = length
	return r
}

func (t *BinaryBasedDecoderGroup) findAdaptiveDecoder(bytes []byte) (decoder Interface, find bool) {
	for i, b := range bytes {
		l := len(t.currentRow)
		if l == t.length {
			t.currentRow = t.currentRow[0 : l-1]
		}
		t.currentRow = append(t.currentRow, b)
		if l < t.length-1 {
			continue
		}
		for _, decoder := range t.decoders {
			if processed, success := decoder.IsValidBeginningBytes(t.currentRow, t.decider.streamType); success {
				//某个decoder校验成功，返回该decoder
				t.currentRow = append(processed, bytes[i+1:]...)
				return decoder, true
			}
		}
	}
	return nil, false
}

type Decider struct {
	clientAddr *net.TCPAddr
	serverAddr *net.TCPAddr
	streamType StreamType
	reader     *StreamReader
	backend    *base.FetcherBackend
	groups     []DecoderGroup
	buffer     []byte
}

func NewDecider(clientAddr, serverAddr *net.TCPAddr, streamType StreamType, reader *StreamReader, backend *base.FetcherBackend) *Decider {
	dc := new(Decider)
	dc.clientAddr = clientAddr
	dc.serverAddr = serverAddr
	dc.streamType = streamType
	dc.reader = reader
	dc.backend = backend
	dc.groups = make([]DecoderGroup, 0, 10)
	for k, fus := range textBasedDecoders {
		dcs := make([]Interface, 0, len(fus))
		for _, fu := range fus {
			dcs = append(dcs, fu(backend))
		}
		dc.groups = append(dc.groups, NewTextBasedDecoderGroup(k, dc, dcs))
	}
	for k, fus := range binaryBasedDecoders {
		dcs := make([]Interface, 0, len(fus))
		for _, fu := range fus {
			dcs = append(dcs, fu(backend))
		}
		dc.groups = append(dc.groups, NewBinaryBasedDecoderGroup(k, dc, dcs))
	}
	return dc
}

func (d *Decider) StartDecode() {
	var adaptiveDecoder Interface
	logrus.WithFields(logrus.Fields{"ca": d.clientAddr.String(), "sa": d.serverAddr.String(),
		"st": d.streamType.String()}).Debug("start decide real decoder")
	for _, dg := range d.groups {
		dg.reset()
	}
outer:
	for {
		buffer := make([]byte, 1024)
		count, err := d.reader.Read(buffer)
		if err != nil {
			logrus.WithField("err", err.Error()).Warn("read data failed")
			break
		}
		if count == 0 {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		buffer = buffer[0:count]
		for _, dg := range d.groups {
			if de, success := dg.findAdaptiveDecoder(buffer); success {
				adaptiveDecoder = de
				d.buffer = dg.getValidRow()
				break outer
			}
		}
	}
	if adaptiveDecoder == nil {
		logrus.WithFields(logrus.Fields{"ca": d.clientAddr.String(), "sa": d.serverAddr.String(),
			"st": d.streamType.String()}).Warn("fail to find adaptiveDecoder for stream")
		return
	}
	logrus.WithFields(logrus.Fields{"protocol": adaptiveDecoder.Protocol(), "ca": d.clientAddr.String(),
		"sa": d.serverAddr.String(), "st": d.streamType.String()}).Debug("find adaptiveDecoder for stream success ")
	var err error
	if d.streamType == Client2ServerStream {
		err = adaptiveDecoder.ParseClient2ServerStream(d)
	} else {
		err = adaptiveDecoder.ParseServer2ClientStream(d)
	}

	if err != nil {
		logrus.WithField("error", err).Warn("decode error")
		if _, success := err.(*MismatchedDecoderError); success {
			d.StartDecode()
		}
	}
}

func (d *Decider) Read(bytes []byte) (int, error) {
	if len(d.buffer) > 0 {
		copied := copy(bytes, d.buffer)
		d.buffer = d.buffer[copied:]
		return copied, nil
	}
	return d.reader.Read(bytes)
}

func (d *Decider) TimeOfLastByte() time.Time {
	return d.reader.LastPackageTime()
}

func (d *Decider) ReportRow(row *core.Row) {
	d.backend.ReportRow(row)
}

func BytesEquals(bs1, bs2 []byte) bool {
	if len(bs1) != len(bs2) {
		return false
	}
	for i, b := range bs1 {
		if bs2[i] != b {
			return false
		}
	}
	return true
}
