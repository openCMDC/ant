package decoder

import (
	"ant/core"
	"ant/fetcher/networkfetcher/base"
	log "github.com/sirupsen/logrus"
	"time"
)

var factoryFunctions = make(map[string]func() Interface, 8)

type Interface interface {
	Name() string
	Parse(conn *base.TCPConn, senBackChan chan<- *core.Row)
}

type defaultDecoder struct {
}

var DefaultDecoder = defaultDecoder{}
var DefaultDecoderName = "defaultDecoder"
var emptyRow = new(core.Row)

func init() {
	RegisterDecoder("dummy", func() Interface {
		return DefaultDecoder
	})
}

func (d defaultDecoder) Name() string {
	return "dummy"
}

// if 5 minutes has passed and 2MB bytes are read, this decoder will be the decoder associated with this conn
func (d defaultDecoder) Parse(conn *base.TCPConn, senBackChan chan<- *core.Row) {
	temp := make([]byte, 1024, 1024)
	start := time.Now()
	var totalRead1, totalRead2 int64
	var cansubmit = false
	for {
		if conn.CheckBothStreamClosed() {
			break
		}
		var c1, c2 = 0, 0
		var err error
		if conn.CheckC2SStreamExist() {
			c1, err = conn.C2SStream().Read(temp, "DefaultEmptyDecoder")
			if err != nil {
				log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "errMsg": err.Error()}).Debug("default decoder read client to server stream failed")
			}
			totalRead1 += int64(c1)
		}
		if conn.CheckS2CStreamExist() {
			c2, err = conn.C2SStream().Read(temp, "DefaultEmptyDecoder")
			if err != nil {
				log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "errMsg": err.Error()}).Debug("default decoder read server to client stream failed")
			}
			totalRead2 += int64(c2)
		}
		if c1 == 0 && c2 == 0 {
			time.Sleep(30 * time.Second)
		} else {
			time.Sleep(3 * time.Second)
		}

		if cansubmit {
			senBackChan <- emptyRow
		} else if start.Add(5*time.Minute).Before(time.Now()) && (totalRead1+totalRead2 > 1024*1024*2) {
			cansubmit = true
		}
	}
	log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String()}).Debug("default decoder decode end")
}

func RegisterDecoder(name string, factoryFunc func() Interface) {
	factoryFunctions[name] = factoryFunc
}
func instanceDefaultDecode() Interface {
	return DefaultDecoder
}
func GetDecoders() []Interface {
	ds := make([]Interface, 0, 8)
	for _, fu := range factoryFunctions {
		ds = append(ds, fu())
	}
	return ds
}
func GetDecoder(name string) Interface {
	fu, _ := factoryFunctions[name]
	if fu == nil {
		return nil
	}
	return fu()
}
