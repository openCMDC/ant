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
	se.conns = make(map[string]*base2.TCPConn)
	se.conns[conn.GetClientAddr().String()] = conn
	se.fistConn = true
	se.decoderName = decoder2.DefaultDecoderName
	se.fetcherCtx = ctx
	se.startParseTcpConn(conn)
	//the first conn is used to chose proper decoder
	//deprecated 2020.3.9
	//go func() {
	//	ds := decoders.GetDecoders()
	//	var wg sync.WaitGroup
	//	wg.Add(len(ds))
	//	for _, de := range ds {
	//		d := de
	//		go func() {
	//			defer wg.Done()
	//			vt := d.ValidateType()
	//			validateSuccess := false
	//			log.WithField("decoderName", d.Name()).Trace("start check decoder work?")
	//			if vt == 1 || vt == 2 {
	//				var count int
	//				var err error
	//				var stream *ConcurrentReader
	//				if vt == 1 {
	//					stream = conn.C2SStream()
	//					count, err = d.LocateToValidStartInC2SStream(newXReader(stream))
	//				} else {
	//					stream = conn.S2CStream()
	//					count, err = d.LocateToValidStartInS2CStream(newXReader(stream))
	//				}
	//				if err != nil {
	//					log.WithField("decoderName", d.Name()).Warn("validate decoder failed")
	//					return
	//				}
	//				log.WithField("decoderName", d.Name()).Trace("validate decoder success")
	//				stream.SetValidStatus()
	//				stream.seekTo(count)
	//				log.Trace("start forbidden later concurrent read and interrupt concurrent reading")
	//				stream.forbiddenConcurrentRead()
	//				stream.interruptReading()
	//				log.Trace("success")
	//				validateSuccess = true
	//			} else if vt == 3 {
	//				c1, err := d.LocateToValidStartInC2SStream(newXReader(conn.C2SStream()))
	//				if err != nil {
	//					log.WithField("decoderName", d.Name()).Warn("validate decoder failed")
	//					return
	//				}
	//				c2, err := d.LocateToValidStartInS2CStream(newXReader(conn.S2CStream()))
	//				if err != nil {
	//					log.WithField("decoderName", d.Name()).Warn("validate decoder failed")
	//					return
	//				}
	//				conn.C2SStream().SetValidStatus()
	//				conn.C2SStream().seekTo(c1)
	//				conn.C2SStream().forbiddenConcurrentRead()
	//				conn.C2SStream().interruptReading()
	//				conn.S2CStream().SetValidStatus()
	//				conn.S2CStream().seekTo(c2)
	//				conn.S2CStream().forbiddenConcurrentRead()
	//				conn.S2CStream().interruptReading()
	//				validateSuccess = true
	//			} else {
	//				log.Warn("Unsupported validate type")
	//				return
	//			}
	//			se.mutex.Lock()
	//			defer se.mutex.Unlock()
	//			if validateSuccess {
	//				se.decoderName = d.Name()
	//			}
	//		}()
	//	}
	//	//todo may wait forever
	//	wg.Wait()
	//	close(se.decoderReady)
	//}()
	return se
}
func (s *ServerEndPoint) AddStream(ca *net.TCPAddr, sa *net.TCPAddr, reader *base2.StreamReader, st StreamType) {
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
	if st == Client2ServerStream {
		c.SetC2SStream(reader)
	} else {
		c.SetS2CStream(reader)
	}
}

func (s *ServerEndPoint) startParseTcpConn(conn *base2.TCPConn) {
	log.Debug("start parse stream")
	storage := s.fetcherCtx.Storage
	name := GetDefaultDecoder(conn.GetServerAddr().String())
	d := decoder2.GetDecoder(name)
	if d != nil {
		log.WithFields(log.Fields{"decoder": d.Name(), "serverAddr": conn.GetServerAddr().String()}).Debug("find default decoder")
		queue := make(chan *core.Row, 10)
		go func() {
			conn.C2SStream().SetOwner(name)
			conn.S2CStream().SetOwner(name)
			d.Parse(conn, queue)
		}()
		go func() {
			for r := range queue {
				storage.StoreRow(r)
			}
		}()
		return
	}
	decoders := decoder2.GetDecoders()
	var chanMap sync.Map
	for _, de := range decoders {
		decoder := de
		go func() {
			name := decoder.Name()
			log.WithField("decoder", decoder.Name()).Trace("try to parse tcp connection")
			queue := make(chan *core.Row, 10)
			decodeGoroutingId := base2.GoID()
			chanMap.Store(name, queue)
			decoder.Parse(conn, queue)
			go func() {
				firstRow := true
				for r := range queue {
					AddDefaultDecoder(conn.GetServerAddr().String(), decoder.Name())
					if firstRow {
						log.WithField("decoder", decoder.Name()).Debug("parse success")
						conn.S2CStream().SetOwner(name)
						conn.C2SStream().SetOwner(name)
						//close other chan
						chanMap.Range(func(key, value interface{}) bool {
							if key == decodeGoroutingId {
								return true
							}
							ch, success := value.(chan *core.Row)
							if success {
								close(ch)
							}
							return true
						})
						firstRow = false
					}
					storage.StoreRow(r)
				}
				log.WithField("decoder", decoder).Debug("parse failed")
			}()
		}()
	}
}
