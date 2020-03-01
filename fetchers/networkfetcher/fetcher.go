package networkfetcher

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
	"mantis/core"
	"net"
	"time"
)

type captureType int8
type captureInternalStatus int8
type CaptureStatus int8

const (
	_ captureType = iota
	listenStreamCapture
	connectStreamCapture

	_ captureInternalStatus = iota
	Ready
	Working
	Paused

	_ CaptureStatus = iota
	NewCreated
	//Ready
	Staring
	Running
	Stoping
	Stoppted
)

func (c CaptureStatus) String() string {
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

func (c captureType) String() string {
	if c == listenStreamCapture {
		return "listenStreamCapture"
	} else if c == connectStreamCapture {
		return "connectStreamCapture"
	}
	return "unexpected capture type"
}

type capture struct {
	name          string
	deviceName    string
	streamFactory StreamFactory
	captureType   captureType
	addr          []net.Addr
	bpfStr        string
	status        CaptureStatus
	interStatus   captureInternalStatus
	oldStatus     CaptureStatus
	handle        *pcap.Handle
}

type NetworkDataFetcher struct {
	captures     map[string][2]*capture
	sendBackChan chan core.Row
	targetPid    string
}

func (*NetworkDataFetcher) Name() string {
	return "NetworkDataFetcher"
}

func (fetcher *NetworkDataFetcher) Start(ant core.AntContext) error {
	fetcher.targetPid = ant.TargetProcessID
	fetcher.updateStatus(&CaptureStatusSetMsg{
		CaptureIds: nil,
		NewStatus:  Running,
		SetAll:     true,
	})
	return nil
}

func (fetcher *NetworkDataFetcher) updateConfPeriodically() {
	ch := time.Tick(time.Second * 5)
	go func() {
		for range ch {
			if fetcher.targetPid == "" {
				log.Warn("stop update networkfetcher conf because targetPid is null")
				continue
			}
			log.WithField("targetPid", fetcher.targetPid).Debug("start get new net work conf info")
			lisAddr, connAddr := FetchAppNetworkInfo(fetcher.targetPid)
			log.WithFields(map[string]interface{}{
				"lisAddr":  lisAddr,
				"connAddr": connAddr,
			}).Debug("success get new net work")
			dn, err := GetNetDeviceName()
			if err != nil {
				log.WithError(err).Warn("stop update networkfetcher")
				continue
			}
			log.WithField("deviceName", dn).Debug("success get device name")
			conf := &CaptureConf{
				DeviceName:  dn,
				ListenAddrs: lisAddr,
				ConnAddrs:   connAddr,
			}
			fetcher.UpdateConf(conf)
		}
	}()
}

func (fetcher *NetworkDataFetcher) Pause() error {
	panic("implement me")
}

func (fetcher *NetworkDataFetcher) Stop() error {
	panic("implement me")
}

func (fetcher *NetworkDataFetcher) GetDataChan() (chan core.Row, error) {
	return fetcher.sendBackChan, nil
}

func (fetcher *NetworkDataFetcher) UpdateConf(conf *CaptureConf) {
	log.WithFields(log.Fields{"device name": conf.DeviceName,
		"listen_ip_and_port":  conf.ListenAddrs,
		"connect_ip_and_port": conf.ConnAddrs}).Trace("start updateConf")
	device := conf.DeviceName
	if len(device) == 0 {
		log.Warn("empty device name")
		return
	}
	//c.captureConfs[device] = &conf.CaptureConf

	captures, exist := fetcher.captures[device]
	if exist {
		listenCapture, connCapture := captures[0], captures[1]
		if listenCapture == nil || connCapture == nil {
			log.WithFields(log.Fields{"deviceName": device}).Warn("why?")
			return
		}
		listenCapture.setAddrs(conf.ListenAddrs)
		listenCapture.setAddrs(conf.ConnAddrs)
		//ctx.Send(listenCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
		//ctx.Send(connCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
	} else {
		lcn := generateCaptureName(device, listenStreamCapture)
		lc := newCapture(lcn, device, conf.ListenAddrs, listenStreamCapture)
		ccn := generateCaptureName(device, connectStreamCapture)
		cc := newCapture(ccn, device, conf.ConnAddrs, connectStreamCapture)

		fetcher.captures[device] = [...]*capture{lc, cc}
	}
}

func (fetcher *NetworkDataFetcher) updateStatus(msg *CaptureStatusSetMsg) {
	log.WithFields(log.Fields{"setAll": msg.SetAll, "newStatus": msg.NewStatus, "names": msg.CaptureIds}).Trace("start updateStatus")
	if msg.SetAll {
		for _, cs := range fetcher.captures {
			if cs[0] != nil {
				cs[0].setStatus(msg.NewStatus)
			}
			if cs[1] != nil {
				cs[1].setStatus(msg.NewStatus)
			}
		}
		return
	}
	ids := msg.CaptureIds
	for _, id := range ids {
		a := fetcher.findCaptureByName(id)
		if a == nil {
			log.WithField("name", id).Warn("can't find specify capture")
			continue
		}
		a.setStatus(msg.NewStatus)
	}
}
func (fetcher *NetworkDataFetcher) findCaptureByName(name string) *capture {
	for _, cs := range fetcher.captures {
		if cs[0].name == name {
			return cs[0]
		}
		if cs[1].name == name {
			return cs[1]
		}
	}
	return nil
}

func (c *capture) updateCaptureStatus(msg *CaptureStatusSetMsg) {
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

func (c *capture) setStatus(status CaptureStatus) error {
	if status == Running {
		return c.ensureRunning()
	}
	//todo 其他状态设置
	return nil
}

func (c *capture) setAddrs(addrs []net.Addr) {
	//todo 根据是否变化来设置状态
	// 如果正在运行中，改变bpf的值
	c.addr = addrs
	c.addr = addrs
	newBpf := GenerateBPFStr(addrs, c.captureType == listenStreamCapture)
	//todo to be verified later
	c.handle.SetBPFFilter(newBpf)
}

func (c *capture) ensureRunning() error {
	if len(c.bpfStr) == 0 {
		return nil
	}
	if c.handle == nil {
		handle, err := pcap.OpenLive(c.deviceName, 65535, false, pcap.BlockForever)
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
			assembler := NewAssembler(NewStreamPool(c.streamFactory))
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
						fmt.Printf("capture a package %s, payload=%d, content=%d\n", packet.TransportLayer().TransportFlow().String(), len(tcp.LayerPayload()), len(tcp.LayerContents()))
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

func (c *capture) ensureStopped() {
	//todo
}

func newCapture(name, deviceName string, addr []net.Addr, catogry captureType) *capture {
	//todo
	isListen := true
	if catogry == connectStreamCapture {
		isListen = false
	}
	c := &capture{
		name:        name,
		deviceName:  deviceName,
		addr:        addr,
		captureType: catogry,
		bpfStr:      GenerateBPFStr(addr, isListen),
		status:      NewCreated,
		oldStatus:   0,
		interStatus: Ready,
	}
	sf := newDSF(c)
	c.streamFactory = sf
	return c
}

func NewInstance() *NetworkDataFetcher {
	fetcher := new(NetworkDataFetcher)
	fetcher.updateConfPeriodically()
	fetcher.captures = make(map[string][2]*capture)
	fetcher.sendBackChan = make(chan core.Row, 4016)
	return fetcher
}
