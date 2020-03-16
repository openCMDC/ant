package capture

import (
	"ant/fetcher/networkfetcher/base"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

type DefaultStreamFactory struct {
	capture         *Capture
	serverEndpoints map[string]*ServerEndPoint
}

func (sf *DefaultStreamFactory) New(netLayer, transLayer gopacket.Flow) tcpassembly.Stream {
	reader := base.NewReaderStream(netLayer, transLayer)
	log.Debug("# Start new stream:", netLayer, transLayer)
	//reader := tcpreader.NewReaderStream()
	ca, sa, st, err := sf.GetInfo(netLayer, transLayer)
	if err != nil {
		log.WithField("errMsg", err.Error()).Warn("stream factory produce a invalid reader")
		return &reader
	}
	endPoint, ok := sf.serverEndpoints[sa.String()]
	if !ok {
		conn := base.NewTCPConn(ca, sa, nil, nil)
		endPoint = NewServerEndpoint(conn, sf.capture.FetcherCtx)
		sf.serverEndpoints[sa.String()] = endPoint
	}
	endPoint.AddStream(ca, sa, &reader, st)
	//srcAddr := netLayer.Src().String()
	//bio := bufio.NewReader(&reader)
	//if srcAddr == "192.168.31.102" {
	//	go func() {
	//		for {
	//			_, err := http.ReadRequest(bio)
	//			if err != nil {
	//				if err != io.EOF{
	//					continue
	//				}
	//				fmt.Println("req err =>" + err.Error())
	//				return
	//			}
	//			fmt.Println("req success")
	//		}
	//	}()
	//} else {
	//	go func() {
	//		for {
	//			_, err := http.ReadResponse(bufio.NewReader(bio), nil)
	//			if err != nil {
	//				if err != io.EOF {
	//					continue
	//				}
	//				fmt.Println("rep err =>" + err.Error())
	//				return
	//			}
	//			fmt.Println("rep success")
	//		}
	//	}()
	//}
	return &reader
	//actorCtx := sf.actorTreeContext.GetActorContext()
	//decoderManager := sf.actorTreeContext.GetDecoderManager()
	//
	//srcAddr, err := common.ParseIpAndPort2TCPAddr(netLayer.Src().String(), transLayer.Src().String())
	//if err != nil {
	//	log.WithField("reason", err.Error()).Warn("parser to tcp addr failed")
	//	return &reader
	//}
	//
	//dstAddr, err := common.ParseIpAndPort2TCPAddr(netLayer.Dst().String(), transLayer.Dst().String())
	//if err != nil {
	//	log.WithField("reason", err.Error()).Warn("parser to tcp addr failed")
	//	return &reader
	//}
	//clientAddr, serverAddr := distinguishStreamDrc(srcAddr, dstAddr, sf.actor.addr)
	//if clientAddr == nil || serverAddr == nil {
	//	log.WithFields(log.Fields{"srcAddr": srcAddr.String(), "dstAddr": dstAddr.String()}).Warn("fail to recognise client addr and server addr")
	//	return &reader
	//}
	////
	////   src          dst    src         dst
	////client --------> server1 --------> server2
	////client <-------- server1 <-------- server2
	////   dst          src    dst         src
	////   if src == clientAddr
	//var st msg2.StreamType
	//if clientAddr == srcAddr {
	//	st = msg2.Client2ServerStream
	//} else {
	//	st = msg2.Server2ClientStream
	//}
	//msg := &msg2.StreamCreatedMsg{
	//	Client: clientAddr,
	//	Server: serverAddr,
	//	Stream: &reader,
	//	Type:   st,
	//	Device: sf.actor.DeviceName,
	//}
	//actorCtx.Send(decoderManager, msg)
	//return &reader
}

func (sf *DefaultStreamFactory) GetInfo(netLayer, transLayer gopacket.Flow) (cAddr, sAddr *net.TCPAddr, st StreamType, err error) {
	tcpStr := fmt.Sprintf("%s:%s", netLayer.Src().String(), transLayer.Src().String())
	srcAddr, err := net.ResolveTCPAddr("tcp", tcpStr)
	if err != nil {
		return nil, nil, -1,
			fmt.Errorf("parser {%s} to tcp addr failed of {%s}", tcpStr, err.Error())
	}
	tcpStr = fmt.Sprintf("%s:%s", netLayer.Dst().String(), transLayer.Dst().String())
	dstAddr, err := net.ResolveTCPAddr("tcp", tcpStr)
	if err != nil {
		return nil, nil, -1,
			fmt.Errorf("parser {%s} to tcp addr failed of {%s}", tcpStr, err.Error())
	}
	clientAddr, serverAddr := distinguishStreamDrc(srcAddr, dstAddr, sf.capture.addr)
	if clientAddr == nil || serverAddr == nil {
		log.WithFields(log.Fields{"srcAddr": srcAddr.String(), "dstAddr": dstAddr.String(), "addrs": sf.capture.addr}).Warn("can't determine client server addr")
		return nil, nil, -1,
			fmt.Errorf("can't distinguish server and client between {%s} {%s} by clues {%s}", clientAddr, serverAddr, sf.capture.addr)
	}
	if clientAddr == srcAddr {
		st = Client2ServerStream
	} else {
		st = Server2ClientStream
	}
	return clientAddr, serverAddr, st, nil
}
func distinguishStreamDrc(src, dst *net.TCPAddr, addrs []net.Addr) (clientAddr, serverAddr *net.TCPAddr) {
	for _, add := range addrs {
		if strings.Contains(add.String(), src.String()) {
			return dst, src
		}
		if strings.Contains(add.String(), dst.String()) {
			return src, dst
		}
	}
	return nil, nil
}

func newDSF(c *Capture) *DefaultStreamFactory {
	f := &DefaultStreamFactory{
		capture:         c,
		serverEndpoints: make(map[string]*ServerEndPoint, 100),
	}
	return f
}
