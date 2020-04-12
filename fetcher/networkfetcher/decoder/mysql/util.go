package mysql

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

func getPktCyclically(r io.Reader, isC2S bool, pktChan chan<- *packet, closeChan chan<- struct{}) {
	for {
		pkt := resolveReader2Pkt(r, isC2S)
		if pkt == nil {
			if closeChan != nil {
				closeChan <- struct{}{}
			}
			break
		}
		pktChan <- pkt
	}
	logrus.WithField("isC2S", isC2S).Warn("end reader")
}

func resolveReader2Pkt(r io.Reader, isC2S bool) *packet {
	var payload bytes.Buffer
	var seq uint8
	var err error
	if seq, err = resolvePacketTo(r, &payload); err != nil {
		fmt.Println(err)
		return nil
	}

	//close stream
	if err == io.EOF {
		return nil
	} else if err != nil {
		logrus.WithError(err).Warn("mysql decoder err")
	}

	//generate new packet
	var pk = packet{
		seq:     int(seq),
		length:  payload.Len(),
		payload: payload.Bytes(),
		ts:      time.Now().UnixNano() / (1000 * 1000),
	}
	if !isC2S {
		pk.isClient2ServerFlow = false
	} else {
		pk.isClient2ServerFlow = true
	}

	return &pk
}

func resolvePacketTo(r io.Reader, w io.Writer) (uint8, error) {

	header := make([]byte, 4)
	if n, err := io.ReadFull(r, header); err != nil {
		if n == 0 && err == io.EOF {
			return 0, io.EOF
		}
		return 0, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)

	var seq uint8
	seq = header[3]

	if n, err := io.CopyN(w, r, int64(length)); err != nil {
		return 0, err
	} else if n != int64(length) {
		return 0, err
	} else {
		return seq, nil
	}

	return seq, nil
}

func LengthEncodedString(b []byte) ([]byte, bool, int, error) {

	num, isNull, n := LengthEncodedInt(b)
	if num < 1 {
		return nil, isNull, n, nil
	}

	n += int(num)

	if len(b) >= n {
		return b[n-int(num) : n], false, n, nil
	}
	return nil, false, n, io.EOF
}

func LengthEncodedInt(input []byte) (num uint64, isNull bool, n int) {

	switch input[0] {

	case 0xfb:
		n = 1
		isNull = true
		return
	case 0xfc:
		num = uint64(input[1]) | uint64(input[2])<<8
		n = 3
		return
	case 0xfd:
		num = uint64(input[1]) | uint64(input[2])<<8 | uint64(input[3])<<16
		n = 4
		return
	case 0xfe:
		num = uint64(input[1]) | uint64(input[2])<<8 | uint64(input[3])<<16 |
			uint64(input[4])<<24 | uint64(input[5])<<32 | uint64(input[6])<<40 |
			uint64(input[7])<<48 | uint64(input[8])<<56
		n = 9
		return
	}

	num = uint64(input[0])
	n = 1
	return
}
