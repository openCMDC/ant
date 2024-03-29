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
	fetcherCtx   *base.FetcherBackend
}

func NewServerEndpoint(conn *base2.TCPConn, ctx *base.FetcherBackend) *ServerEndPoint {
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
func (s *ServerEndPoint) AddStream(ca *net.TCPAddr, sa *net.TCPAddr, reader *base2.StreamReader, st base2.StreamType) {
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
		if st == base2.Client2ServerStream {
			c.SetC2SStream(reader)
		} else {
			c.SetS2CStream(reader)
		}
		return
	}
	if st == base2.Client2ServerStream {
		c.SetC2SStream(reader)
	} else {
		c.SetS2CStream(reader)
	}
}

func (s *ServerEndPoint) startParseTcpConn(conn *base2.TCPConn) {
	log.Trace("start parse stream")
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
				s.fetcherCtx.ReportRow(r)
			}
		}()
		return
	}
	decoders := decoder2.GetDecoders()
	//todo 这里要不要加个超时机制，如果同时满足超时一定时间并且流过的流量也足够多两个条件，但是仍然
	// 没有decoder解析成功，就认为这个流无法被解析。设置默认decoder或者设置bpf。
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
				s.fetcherCtx.ReportRow(r)
			}
			//log.WithField("decoder", decoder).Debug("parse end")
		}()
	}
}

func (s *ServerEndPoint) startParseTcpConnV2(conn *base2.TCPConn) {
	log.Trace("start parse stream")
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
				s.fetcherCtx.ReportRow(r)
			}
		}()
		return
	}
	decoders := decoder2.GetDecoders()
	//todo 这里要不要加个超时机制，如果同时满足超时一定时间并且流过的流量也足够多两个条件，但是仍然
	// 没有decoder解析成功，就认为这个流无法被解析。设置默认decoder或者设置bpf。
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
				s.fetcherCtx.ReportRow(r)
			}
			//log.WithField("decoder", decoder).Debug("parse end")
		}()
	}
}
