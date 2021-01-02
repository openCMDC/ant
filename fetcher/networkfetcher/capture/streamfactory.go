package capture

import (
	"ant/fetcher/networkfetcher/base"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
	"sync"
)

type DefaultStreamFactory struct {
	capture         *Capture
	serverEndpoints sync.Map
}

func (sf *DefaultStreamFactory) New(netLayer, transLayer gopacket.Flow) tcpassembly.Stream {
	reader := base.NewReaderStream(netLayer, transLayer)
	//log.Debug("# Start new stream:", netLayer, transLayer)
	log.Debug("# Start new stream  ", transLayer)
	ca, sa, st, err := sf.GetInfo(netLayer, transLayer)
	if err != nil {
		log.WithField("errMsg", err.Error()).Warn("stream factory produce a invalid reader")
		return &reader
	}
	endPoint, ok := sf.serverEndpoints.Load(sa.String())
	if !ok {
		conn := base.NewTCPConn(ca, sa, nil, nil)
		endPoint = NewServerEndpoint(conn, sf.capture.FetcherCtx)
		sf.serverEndpoints.Store(sa.String(), endPoint)
	}
	ep := endPoint.(*ServerEndPoint)
	ep.AddStream(ca, sa, &reader, st)
	return &reader
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
	clientAddr, serverAddr := distinguishStreamDrc(srcAddr, dstAddr, sf.capture.ListeningAddr, sf.capture.LocalAddr)
	if clientAddr == nil || serverAddr == nil {
		log.WithFields(log.Fields{"srcAddr": srcAddr.String(), "dstAddr": dstAddr.String(), "addrs": sf.capture.ListeningAddr}).Warn("can't determine client server addr")
		return nil, nil, -1,
			fmt.Errorf("can't distinguish server and client between {%s} {%s} by clues {%s}", clientAddr, serverAddr, sf.capture.ListeningAddr)
	}
	if clientAddr == srcAddr {
		st = Client2ServerStream
	} else {
		st = Server2ClientStream
	}
	return clientAddr, serverAddr, st, nil
}
func distinguishStreamDrc(src, dst *net.TCPAddr, addrs []net.Addr, localAddr []pcap.InterfaceAddress) (clientAddr, serverAddr *net.TCPAddr) {
	//本机作为server端的情况
	for _, add := range addrs {
		if strings.Contains(add.String(), src.String()) {
			return dst, src
		}
		if strings.Contains(add.String(), dst.String()) {
			return src, dst
		}
	}

	//本机作为客户端的情况
	for _, add := range localAddr {
		local := add.IP.String()
		if strings.Contains(src.String(), local) {
			return src, dst
		}
		if strings.Contains(dst.String(), local) {
			return dst, src
		}
	}
	return nil, nil
}

func newDSF(c *Capture) *DefaultStreamFactory {
	f := &DefaultStreamFactory{
		capture: c,
	}
	return f
}
