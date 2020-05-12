package capture

import (
	"ant/core"
	"ant/fetcher/base"
	base2 "ant/fetcher/networkfetcher/base"
	decoder2 "ant/fetcher/networkfetcher/decoder"
	log "github.com/sirupsen/logrus"
	"net"
	"sync"
)

type StreamType int8

const (
	_ StreamType = iota
	Client2ServerStream
	Server2ClientStream
)

func (t StreamType) String() string {
	if t == Client2ServerStream {
		return "Client2ServerStream"
	} else if t == Server2ClientStream {
		return "Server2ClientStream"
	} else {
		return "unKnown stream type"
	}
}

var connDecoderMapping sync.Map

func AddDefaultDecoder(serverAddr, decoderName string) {
	connDecoderMapping.Store(serverAddr, decoderName)
}

func GetDefaultDecoder(serverAddr string) string {
	res, ok := connDecoderMapping.Load(serverAddr)
	if !ok {
		return ""
	}
	return res.(string)
}

type ServerEndPoint struct {
	serverAddr   *net.TCPAddr
	decoderReady chan struct{}
	fistConn     bool
	conns        map[string]*base2.TCPConn
	decoderName  string
	mutex        sync.RWMutex
	fetcherCtx   *base.FetcherCtx
}

func NewServerEndpoint(conn *base2.TCPConn, ctx *base.FetcherCtx) *ServerEndPoint {
	se := new(ServerEndPoint)
	se.serverAddr = conn.GetServerAddr()
	se.decoderReady = make(chan struct{})
	se.conns = make(map[string]*base2.TCPConn, 10)
	se.conns[conn.GetClientAddr().String()] = conn
	se.fistConn = true
	se.decoderName = decoder2.DefaultDecoderName
	se.fetcherCtx = ctx
	se.startParseTcpConn(conn)
	return se
}
func (s *ServerEndPoint) AddStream(ca *net.TCPAddr, sa *net.TCPAddr, reader *base2.StreamReader, st StreamType) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if sa.String() != s.serverAddr.String() {
		log.WithFields(log.Fields{"addr1": s.serverAddr, "addr2": sa}).Warn("can not add a tcp connection to server endpoint with different server address")
		return
	}
	c, ok := s.conns[ca.String()]
	if !ok {
		s.conns[ca.String()] = base2.NewTCPConn(ca, sa, nil, nil)
		c = s.conns[ca.String()]
		s.startParseTcpConn(c)
	}
	//if st == Client2ServerStream {
	//	c.SetC2SStream(reader)
	//} else {
	//	c.SetS2CStream(reader)
	//}

	if c.HasClosedStream() {
		//parse gorouting for this conn should end, recreate this conn again
		delete(s.conns, ca.String())
		s.conns[ca.String()] = base2.NewTCPConn(ca, sa, nil, nil)
		c = s.conns[ca.String()]
		s.startParseTcpConn(c)
		if st == Client2ServerStream {
			c.SetC2SStream(reader)
		} else {
			c.SetS2CStream(reader)
		}
		return
	}
	if st == Client2ServerStream {
		c.SetC2SStream(reader)
	} else {
		c.SetS2CStream(reader)
	}
}

func (s *ServerEndPoint) startParseTcpConn(conn *base2.TCPConn) {
	log.Trace("start parse stream")
	storage := s.fetcherCtx.Storage
	name := GetDefaultDecoder(conn.GetServerAddr().String())
	d := decoder2.GetDecoder(name)
	if d != nil {
		//log.WithFields(log.Fields{"decoder": d.Name(), "serverAddr": conn.GetServerAddr().String()}).Debug("find default decoder")
		queue := make(chan *core.Row)
		go func() {
			conn.C2SStream().SetOwner(name)
			conn.S2CStream().SetOwner(name)
			d.Parse(conn, queue)
			close(queue)
		}()
		go func() {
			for r := range queue {
				storage.InsertRow(r)
			}
		}()
		return
	}
	decoders := decoder2.GetDecoders()
	for _, de := range decoders {
		decoder := de
		queue := make(chan *core.Row, 10)
		go func() {
			//log.WithField("decoder", decoder.Name()).Debug("try to parse tcp connection")
			decoder.Parse(conn, queue)
			close(queue)
		}()
		go func() {
			fst := true
			for r := range queue {
				AddDefaultDecoder(conn.GetServerAddr().String(), decoder.Name())
				if fst {
					conn.C2SStream().SetOwner(decoder.Name())
					conn.S2CStream().SetOwner(decoder.Name())
					fst = false
				}

				storage.InsertRow(r)
			}
			//log.WithField("decoder", decoder).Debug("parse end")
		}()
	}
}
