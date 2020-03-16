package decoder

import (
	"ant/core"
	"ant/fetcher/networkfetcher/base"
	log "github.com/sirupsen/logrus"
)

type DecodeCtx struct {
	Conn         *base.TCPConn
	SendBackChan chan core.Row
	CtlChan      chan interface{}
}

var factoryFunctions = make(map[string]func() Interface, 8)

type Interface interface {
	Name() string
	Parse(conn *base.TCPConn, senBackChan chan<- *core.Row)
}
type defaultDecoder struct {
}

var DefaultDecoder = defaultDecoder{}
var DefaultDecoderName = "defaultDecoder"

func (d defaultDecoder) Name() string {
	return "DefaultEmptyDecoder"
}
func (d defaultDecoder) ValidateType() byte {
	return 1
}
func (d defaultDecoder) Parse(conn *base.TCPConn, senBackChan chan<- *core.Row) {
	temp := make([]byte, 1024, 1024)
	c2sStreamBroken, s2cStreamBroken := false, false
	if c2sStreamBroken && s2cStreamBroken {
		return
	}
	if !c2sStreamBroken {
		c, err := conn.C2SStream().Read(temp, "yummy")
		if err != nil {
			log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "errMsg": err.Error()}).Debug("default decoder read client to server stream failed")
			c2sStreamBroken = true
		}
		content := temp[0:c]
		log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "count": c, "content": content}).Debug("default decoder read  client to server stream success")
	}
	if !s2cStreamBroken {
		c, err := conn.S2CStream().Read(temp, "yummy")
		if err != nil {
			log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "errMsg": err.Error()}).Debug("default decoder read server to client stream failed")
			c2sStreamBroken = true
		}
		content := temp[0:c]
		log.WithFields(log.Fields{"clientAddr": conn.GetClientAddr().String(), "serverAddr": conn.GetServerAddr().String(), "count": c, "content": content}).Debug("default decoder read  server to clientr stream success")
	}
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
