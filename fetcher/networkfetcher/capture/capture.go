package capture

import (
	"ant/fetcher/base"
	base2 "ant/fetcher/networkfetcher/base"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
	"net"
)

type Type int8
type captureInternalStatus int8
type Status int8

const (
	_ Type = iota
	ListenStreamCapture
	ConnectStreamCapture

	_ captureInternalStatus = iota
	Ready
	Working
	Paused

	_ Status = iota
	NewCreated
	//Ready
	Staring
	Running
	Stoping
	Stoppted
)

func (c Status) String() string {
	if c == 1 {
		return "NewCreated"
	} else if c == 2 {
		return "Ready"
	} else if c == 3 {
		return "Staring"
	} else if c == 4 {
		return "Running"
	} else if c == 5 {
		return "Stoping"
	} else if c == 6 {
		return "Stoppted"
	} else {
		return fmt.Sprintf("unknown status with id = %d", int8(c))
	}
}

func (c Type) String() string {
	if c == ListenStreamCapture {
		return "ListenStreamCapture"
	} else if c == ConnectStreamCapture {
		return "ConnectStreamCapture"
	}
	return "unexpected capture type"
}

type Capture struct {
	Name          string
	DeviceName    string
	FetcherCtx    *base.FetcherCtx
	streamFactory base2.StreamFactory
	captureType   Type
	addr          []net.Addr
	bpfStr        string
	status        Status
	interStatus   captureInternalStatus
	oldStatus     Status
	handle        *pcap.Handle
}

func (c *Capture) SetAddrs(addrs []net.Addr) {
	//todo 根据是否变化来设置状态
	// 如果正在运行中，改变bpf的值
	if c.handle == nil {
		return
	}
	c.addr = addrs
	c.addr = addrs
	newBpf := base2.GenerateBPFStr(addrs, c.captureType == ListenStreamCapture)
	//todo to be verified later
	c.handle.SetBPFFilter(newBpf)
}

func (c *Capture) ensureRunning() error {
	if len(c.bpfStr) == 0 {
		return nil
	}
	if c.handle == nil {
		handle, err := pcap.OpenLive(c.DeviceName, 65535, false, pcap.BlockForever)
		if err != nil {
			return err
		}
		err = handle.SetBPFFilter(c.bpfStr)
		if err != nil {
			return err
		}
		c.handle = handle
		src := gopacket.NewPacketSource(handle, handle.LinkType())
		packets := src.Packets()

		go func() {
			assembler := base2.NewAssembler(base2.NewStreamPool(c.streamFactory))
			for {
				select {
				case packet := <-packets:
					//todo 这一块是不是后面要通过锁的机制来保证可见性？应该是需要的，后面测试一下看下情况
					if c.interStatus == Paused {
						// 通过查看状态的方式来启动或者停止抓包。这里的停止抓包
						// 依然会轮训 packets ，只是在有包的时候忽略这个包。这样设计的主要考虑是
						// 避免底层往packets塞数据被阻塞，packets的默认大小是1000。如果后面有需求不能这么搞
						// 那么在这里直接跳出这个循环就好了。跳出这个循环不能break，这样只能跳出select，需要给
						// 外层for循环一个标记
						// 另外这里对status这个变量的读取应该会有可见性问题，严格来说需要加锁，
						// 但是太损耗性能了，而且暂停抓包这种场景应该用的也不多。而且即便用，出现了
						// 可见行问题，也不会对系统造成不可预估的影响，故而这里就不加锁了
						continue
					}
					if packet.NetworkLayer() == nil ||
						packet.TransportLayer() == nil ||
						packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
						continue
					}
					tcp := packet.TransportLayer().(*layers.TCP)
					if len(tcp.Payload) > 10 {
						log.Trace("success capture a package")
					}
					//fmt.Println(len(tcp.LayerPayload()))
					//fmt.Println(len(tcp.LayerContents()))
					assembler.AssembleWithTimestamp(
						packet.NetworkLayer().NetworkFlow(),
						tcp, packet.Metadata().Timestamp,
					)
					//case <-ticker:
					//	assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
				}
			}
		}()
		c.interStatus = Working
	}
	if c.interStatus != Working {
		c.interStatus = Working
	}
	return nil
}

func (c *Capture) ensureStopped() {
	//todo
}

func (c *Capture) SetStatus(status Status) error {
	if status == Running {
		return c.ensureRunning()
	}
	//todo 其他状态设置
	return nil
}

func (c *Capture) updateCaptureStatus(msg *StatusSetMsg) {
	c.oldStatus, c.status = c.status, msg.NewStatus
	switch c.status {
	//case common.NewCreated:
	//	//why? do nothing
	//case common.Ready:
	//	//why? do nothing
	//case common.Staring:
	//	//why? do nothing
	case Running:
		//
		err := c.ensureRunning()
		if err != nil {
			log.WithField("errMsg", err.Error()).Warn("start capture failed")
		}
	//case common.Stoping:
	//	//why? do nothing
	case Stoppted:
		c.ensureStopped()
	default:
		log.WithField("status", c.status.String()).Warn("unSupported status set action")
	}
}

func NewCapture(name, deviceName string, addr []net.Addr, captureType Type, ctx *base.FetcherCtx) *Capture {
	//todo
	isListen := true
	if captureType == ConnectStreamCapture {
		isListen = false
	}
	c := &Capture{
		Name:        name,
		DeviceName:  deviceName,
		FetcherCtx:  ctx,
		addr:        addr,
		captureType: captureType,
		bpfStr:      base2.GenerateBPFStr(addr, isListen),
		status:      NewCreated,
		oldStatus:   0,
		interStatus: Ready,
	}
	sf := newDSF(c)
	c.streamFactory = sf
	return c
}

func GenerateCaptureName(deviceName string, catogery Type) string {
	return fmt.Sprintf("{%s}-{%s}", deviceName, catogery.String())
}
