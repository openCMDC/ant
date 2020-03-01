package networkfetcher

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"testing"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func TestNetworkDataFetcher_all(t *testing.T) {
	n := NewInstance()
	addr, _ := ParseIpAndPort2TCPAddr("180.101.212.32", "80")
	c := CaptureConf{DeviceName: `\Device\NPF_{7D0089A9-CD14-4C31-926D-B5D61B002A60}`, ListenAddrs: []net.Addr{}, ConnAddrs: []net.Addr{addr}}
	n.UpdateConf(&c)
	ch, _ := n.GetDataChan()
	for r := range ch {
		fmt.Println(r)
	}
}
