package base

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var InternalReaderNil = errors.New("InternalReaderNil")

func ParseIpAndPort2TCPAddr(ip string, port string) (*net.TCPAddr, error) {
	ip1 := net.ParseIP(ip)
	port1, err := strconv.Atoi(port)
	if ip1 == nil || err != nil {
		return nil, fmt.Errorf("parser %s to ip failed or parser %s to port failed", ip, port)
	}
	return &net.TCPAddr{IP: ip1, Port: port1}, nil
}

func GenerateBPFStr(addrs []net.Addr, isListenAddr bool) string {
	//todo
	return generateBPFV2(addrs)
}

func generateBPFV2(addrs []net.Addr) string {
	inBpf := make([]string, 0)
	outBpf := make([]string, 0)
	inbpf := "(dst host %s and dst port %s)"
	outbpf := "(src host %s and src port %s)"
	for _, a := range addrs {
		p := a.String()
		i := strings.LastIndex(p, ":")
		inBpf = append(inBpf, fmt.Sprintf(inbpf, p[0:i], p[i+1:]))
		outBpf = append(outBpf, fmt.Sprintf(outbpf, p[0:i], p[i+1:]))
	}
	inRes := strings.Join(inBpf, " or ")
	outRes := strings.Join(outBpf, " or ")
	if len(inRes) == 0 {
		return outRes
	}
	if len(outRes) == 0 {
		return inRes
	}
	return inRes + " or " + outRes
}

func GoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

var deviceName string
var addrs []pcap.InterfaceAddress
var deviceNameMutext sync.Mutex

//get all devices traffic rate and chose the biggest one
// if there is a agent , this method may return a wrong DeviceName (all traffic will go into a local agent first which use a loop(127.0.0.1) device)
func GetNetDeviceName() (string, []pcap.InterfaceAddress, error) {
	deviceNameMutext.Lock()
	deviceNameMutext.Unlock()
	if deviceName != "" {
		return deviceName, addrs, nil
	}
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return "", nil, err
	}
	type result struct {
		bytes int64
		name  string
		addrs []pcap.InterfaceAddress
	}
	closeChan := make(chan struct{}, 1)
	resultChan := make(chan result, len(devices))
	var counter int32 = 0
	for _, device := range devices {
		temp := device
		go func() {
			handle, err := pcap.OpenLive(temp.Name, 65535, false, pcap.BlockForever)
			if err != nil {
				return
			}
			atomic.AddInt32(&counter, 1)
			defer handle.Close()
			//fmt.Println("-->" + temp.Name)
			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
			var total int64
			for {
				select {
				case pkg := <-packetSource.Packets():
					total += int64(len(pkg.Data()))
				case <-closeChan:
					resultChan <- result{total, temp.Name, temp.Addresses}
					return
				}
			}
		}()
	}
	time.Sleep(time.Second * 5)
	close(closeChan)
	chosen := result{-1, "", nil}
	for ; counter > 0; counter-- {
		re := <-resultChan
		if re.bytes > chosen.bytes {
			chosen = re
		}
	}
	deviceNameMutext.Lock()
	deviceName = chosen.name
	deviceNameMutext.Unlock()
	return chosen.name, chosen.addrs, nil
}
