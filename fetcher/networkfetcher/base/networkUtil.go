package base

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func FetchAppNetworkInfo(pid string) ([]net.Addr, []net.Addr) {
	if runtime.GOOS == "linux" {
		return fetchInfo4Linux(pid)
	} else if runtime.GOOS == "darwin" {
		return fetchInfo4Darwin(pid)
	} else if runtime.GOOS == "windows" {
		return fetchInfo4Windows(pid)
	} else {
		log.WithField("os", runtime.GOOS).Errorf("unSupported OS type")
		return []net.Addr{}, []net.Addr{}
	}
}

func fetchInfo4Linux(pid string) ([]net.Addr, []net.Addr) {
	path := fmt.Sprintf("/proc/%s/fd/", pid)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Warnf("FetchAppNetworkInfo failed of {%s}", err.Error())
		return []net.Addr{}, []net.Addr{}
	}
	socketInodes := make(map[string]bool)
	for _, f := range files {
		l, err := os.Readlink(path + f.Name())
		if err != nil {
			log.Warnf("read link for file {%s} failed of {%s}", path+f.Name(), err.Error())
			continue
		}
		if strings.HasPrefix(l, "socket:") {
			//if there exist chinese character, this code doesn't work
			sInode := l[8 : len(l)-1]
			socketInodes[sInode] = true
		}
	}

	listenArr, connArr, err := parseTcpFile(socketInodes, "/proc/net/tcp")
	if err != nil {
		log.Warnf("read tcp4 connections failed of {%s}", err.Error())
		return []net.Addr{}, []net.Addr{}
	}
	//listenArr, connArr = temp1, temp2 ignore ipv6
	//temp1, temp2, err = parseTcpFile(listenArr, connArr, socketInodes, "/proc/net/tcp6")
	//if err != nil {
	//	log.Warnf("read tcp6 connections failed of {%s}", err.Error())
	//	return []net.Addr{}, []net.Addr{}
	//}
	//listenArr, connArr = temp1, temp2
	listenports := make(map[int]bool)
	for _, listen := range listenArr {
		tcp, success := listen.(*net.TCPAddr)
		if !success {
			continue
		}
		listenports[tcp.Port] = true
	}
	temp := map[string]bool{}
	var connectIpAndPort []net.Addr
	for _, conn := range connArr {
		tcp, success := conn[0].(*net.TCPAddr)
		if !success {
			continue
		}
		localPort := tcp.Port
		if !listenports[localPort] {
			//说明此时这个socket连接是本机作为client连接的远端server
			//将对端的Ip地址和端口号存储起来
			//connectIpAndPort = append(connectIpAndPort, conn[1])
			temp[conn[1].String()] = true
		}
	}
	for k, _ := range temp {
		id, err := net.ResolveTCPAddr("", k)
		if err != nil {
			log.Warnf("wrong formate of tcp addr {%s}", k)
			continue
		}
		connectIpAndPort = append(connectIpAndPort, id)
	}
	return listenArr, connectIpAndPort
}

//临时方案
func fetchInfo4Darwin(pid string) ([]net.Addr, []net.Addr) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -n -P -p %s | grep TCP", pid))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	res := string(out.Bytes())
	lines := strings.Split(res, "\n")
	listenAddr := make(map[string]bool)
	var listenArr []net.Addr
	var connPair [][2]net.Addr
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 10 {
			continue
		}
		if fields[9] == "(LISTEN)" {
			i := strings.LastIndex(fields[8], ":")
			var addrStr string
			if fields[8][0:i] == `*` {
				addrStr = "0.0.0.0:" + fields[8][i+1:]
			} else {
				addrStr = fields[8][0:i] + ":" + fields[8][i+1:]
			}
			addr, err := net.ResolveTCPAddr("tcp", addrStr)
			if err != nil {
				log.WithField("srcStr", addrStr).
					Warn("fail to parse string to tcp address")
				continue
			}
			listenArr = append(listenArr, addr)
			listenAddr[addr.String()] = true
		} else {
			ipPorts := strings.Split(fields[8], "->")
			lisAddrStr := ipPorts[0]
			connAddrStr := ipPorts[1]
			lisAddr, err := net.ResolveTCPAddr("tcp", lisAddrStr)
			if err != nil {
				log.WithField("srcStr", lisAddrStr).
					Warn("fail to parse string to tcp address")
				continue
			}
			connAddr, err := net.ResolveTCPAddr("tcp", connAddrStr)
			if err != nil {
				log.WithField("srcStr", connAddrStr).
					Warn("fail to parse string to tcp address")
				continue
			}
			connPair = append(connPair, [...]net.Addr{lisAddr, connAddr})
		}
	}
	var connectArr []net.Addr
	for _, conn := range connPair {
		tcp, success := conn[0].(*net.TCPAddr)
		if !success {
			continue
		}

		if listenAddr[tcp.String()] != true {
			connectArr = append(connectArr, conn[1])
		}
	}
	return listenArr, connectArr
}

func fetchInfo4Windows(pid string) ([]net.Addr, []net.Addr) {
	cmd := exec.Command("netstat", "-n", "-o", "-a", "-p", "tcp")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	lines := strings.Split(out.String(), "\r\n")
	var listenArr []net.Addr
	listenAddr := make(map[string]bool)
	var connPair [][2]net.Addr
	for _, line := range lines {
		line := strings.Trim(line, " ")
		if !strings.HasPrefix(line, "TCP") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 5 {
			continue
		}
		if fields[4] != pid {
			continue
		}
		if fields[3] == "LISTENING" {
			addr, err := net.ResolveTCPAddr("tcp", fields[1])
			if err != nil {
				log.WithField("srcStr", addr).
					Warn("fail to parse string to tcp address")
				continue
			}
			listenArr = append(listenArr, addr)
			listenAddr[addr.String()] = true
		} else {
			la, err := net.ResolveTCPAddr("tcp", fields[1])
			ra, err := net.ResolveTCPAddr("tcp", fields[2])
			if err != nil {
				log.WithFields(map[string]interface{}{
					"localAddr":  fields[1],
					"remoteAddr": fields[2],
				}).Warn("fail to parse string to tcp address")
				continue
			}
			connPair = append(connPair, [...]net.Addr{la, ra})
		}
	}
	var connectArr []net.Addr
	for _, conn := range connPair {
		tcp, success := conn[0].(*net.TCPAddr)
		if !success {
			continue
		}
		if listenAddr[tcp.String()] != true {
			connectArr = append(connectArr, conn[1])
		}
	}
	return listenArr, connectArr
}

func parseTcpFile(socketInodes map[string]bool, path string) ([]net.Addr, [][2]net.Addr, error) {
	listenArr := []net.Addr{}
	connArr := [][2]net.Addr{}
	tcp, err := os.Open(path)
	if err != nil {
		log.Warnf("read file content failed of {%s}", err.Error())
	} else {
		tcp4Reader := bufio.NewReader(tcp)
		firstLine := true
		for {
			line, _, err := tcp4Reader.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				return listenArr, connArr, fmt.Errorf("reading   content failed of {%s}", err.Error())
			}
			if firstLine {
				firstLine = false
				continue
			}
			ss := strings.Fields(string(line))
			if _, e := socketInodes[ss[9]]; !e {
				continue
			} else {
				local, err := parseIpAddr(ss[1])
				if err != nil {
					log.Warnf(err.Error())
					continue
				}
				remote, err := parseIpAddr(ss[2])
				if err != nil {
					log.Warnf(err.Error())
					continue
				}
				if strings.EqualFold(ss[3], "0A") {
					listenArr = append(listenArr, local)
				} else {
					temp := [...]net.Addr{local, remote}
					connArr = append(connArr, temp)
				}
			}
		}
	}
	return listenArr, connArr, nil
}

//解析一个ip:port字符串 形如969310AC:A180 或者 0000000000000000FFFF0000969310AC:A180
func parseIpAddr(content string) (*net.TCPAddr, error) {
	ipPort := strings.Split(content, ":")
	if len(ipPort) != 2 {
		return nil, fmt.Errorf("parseIpAddr failed, {%s} is not ip:port format", content)
	}
	//port, err := strconv.ParseInt(ipPort[1], 16, 32)
	//if err != nil {
	//	return nil, fmt.Errorf("parse port str {%s} failed of {%s} ", ipPort[1], err.Error())
	//}
	var ip net.IP
	var err error
	if len(ipPort[0]) == 32 {
		//ipv6场景，这种场景下，即便监听的是ipv6地址，
		//大多数情况下，客户端连接过来依然是用的ipv4地址
		//操作系统做了一定的转换，所以这里暂时只解析出ipv4
		//的地址
		ip, err = parseIpstr(ipPort[0][24:32])
	} else if len(ipPort[0]) == 8 {
		//ipv4场景
		ip, err = parseIpstr(ipPort[0])
	} else {
		return nil, fmt.Errorf("ipPort str {%s} is unexpcted fomate", ipPort[0])
	}
	if err != nil {
		return nil, fmt.Errorf("parse ip str {%s} failed of {%s} ", ipPort[0][24:32], err.Error())
	}
	return net.ResolveTCPAddr("tcp", ip.String()+":"+ipPort[1])
}

//将一个IP地址字符串转换承点分十进制表示的数组，
//输入形如 969310AC 对应的实际地址为
func parseIpstr(ipStr string) (net.IP, error) {
	ip := [4]byte{}
	if len(ipStr) != 8 {
		return nil, fmt.Errorf("ipstr {%s} is not in right formate ", ipStr)
	}
	i0, err := strconv.ParseInt(ipStr[6:8], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("parse ip str {%s} failed of {%s} ", ipStr[6:8], err.Error())
	}
	ip[0] = byte(i0)

	i1, err := strconv.ParseInt(ipStr[4:6], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("parse ip str {%s} failed of {%s} ", ipStr[4:6], err.Error())
	}
	ip[1] = byte(i1)

	i2, err := strconv.ParseInt(ipStr[2:4], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("parse ip str {%s} failed of {%s} ", ipStr[2:4], err.Error())
	}
	ip[2] = byte(i2)

	i3, err := strconv.ParseInt(ipStr[0:2], 16, 16)
	if err != nil {
		return nil, fmt.Errorf("parse ip str {%s} failed of {%s} ", ipStr[0:2], err.Error())
	}
	ip[3] = byte(i3)
	return net.IPv4(ip[0], ip[1], ip[2], ip[3]), nil
}
