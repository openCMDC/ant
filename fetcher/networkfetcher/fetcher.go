package networkfetcher

import (
	"ant/core/err"
	"ant/core/types"
	"ant/core/v1/resource"
	"ant/fetcher/base"
	base2 "ant/fetcher/networkfetcher/base"
	"ant/fetcher/networkfetcher/capture"
	"fmt"
	"github.com/goinggo/mapstructure"
	log "github.com/sirupsen/logrus"
	"time"
)

type NetworkFetcherRuntime struct {
	fetcher    types.NetworkFetcher
	captures   map[string]*capture.Capture
	fetcherCtx *base.FetcherBackend
	TotalBytes string
	ValidBytes string
}

func (*NetworkFetcherRuntime) Name() string {
	return "NetworkFetcherRuntime"
}

func (fetcher *NetworkFetcherRuntime) Start() error {
	for _, c := range fetcher.captures {
		err := c.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (fetcher *NetworkFetcherRuntime) updateConfPeriodically() {
	go func() {
		//todo
		pid := "123"
		for {
			time.Sleep(time.Second * 5)
			log.WithField("targetPid", pid).Trace("start get new net work conf info")
			if pid == "" {
				log.Warn("stop update networkfetcher conf because targetPid is null")
				continue
			}
			lisAddr, connAddr := base2.FetchAppNetworkInfo(pid)
			log.WithFields(map[string]interface{}{
				"lisAddr":  lisAddr,
				"connAddr": connAddr,
			}).Debug("success get net work info")
			dn, _, err := base2.GetNetDeviceName()
			if err != nil {
				log.WithError(err).Warn("stop update networkfetcher")
				continue
			}
			log.WithField("deviceName", dn).Debug("success get device name")
			conf := &capture.Conf{
				DeviceName: dn,
			}
			fetcher.UpdateConf(conf)
			time.Sleep(time.Hour * 5)
		}
	}()
}

func (fetcher *NetworkFetcherRuntime) Pause() error {
	panic("implement me")
}

func (fetcher *NetworkFetcherRuntime) Stop() error {
	panic("implement me")
}

func (fetcher *NetworkFetcherRuntime) UpdateConf(conf *capture.Conf) {
	log.WithFields(log.Fields{"device name": conf.DeviceName, "bpfStr": conf.BpfStr}).
		Trace("start updateConf")
	device := conf.DeviceName
	bpf := conf.DeviceName
	if len(device) == 0 {
		log.Warn("empty device name")
		return
	}
	//c.captureConfs[device] = &conf.Conf

	c, exist := fetcher.captures[device]
	if exist {
		log.WithField("deviceName", device).Trace("start update existed capture")
		c.UpdateConf(conf)
		log.WithField("deviceName", device).Trace("update existed capture success")
		//ctx.Send(listenCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
		//ctx.Send(connCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
	} else {
		log.WithField("deviceName", device).Trace("start instantiate capture")
		c = capture.NewCapture(fetcher.fetcherCtx, device, bpf, nil)
		fetcher.captures[device] = c
		log.WithField("deviceName", device).Trace("instantiate capture success")
	}
}

func (fetcher *NetworkFetcherRuntime) InitCapture() {
	fe := fetcher.fetcher
	log.WithFields(log.Fields{"device name": fe.DeviceName, "bpfStr": fe.BpfFilter}).
		Trace("start updateConf")
	device := fe.DeviceName
	bpf := fe.BpfFilter
	la := fe.LocalAddr
	if len(device) == 0 {
		log.Warn("empty device name")
		return
	}
	//c.captureConfs[device] = &conf.Conf

	c, exist := fetcher.captures[device]
	if exist {
		log.WithField("deviceName", device).Trace("start update existed capture")
		//c.UpdateConf(conf)
		log.WithField("deviceName", device).Trace("update existed capture success")
		//ctx.Send(listenCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
		//ctx.Send(connCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
	} else {
		log.WithField("deviceName", device).Trace("start instantiate capture")
		c = capture.NewCapture(fetcher.fetcherCtx, device, bpf, la)
		fetcher.captures[device] = c
		log.WithField("deviceName", device).Trace("instantiate capture success")
	}
}

//deprecated
func (fetcher *NetworkFetcherRuntime) ConsumeData(resourceName, action string, payload interface{}) {
	if action == resource.NetworkDataFetcher {
		nfc := new(types.NetworkFetcher)
		err := mapstructure.Decode(payload, nfc)
		if err != nil {
			log.Warn("decode payload to task failed")
			return
		}
		capConf := &capture.Conf{nfc.DeviceName, nfc.BpfFilter}
		fetcher.UpdateConf(capConf)
	} else {
		log.WithField("action", action).Warn("unsupported action")
	}
}

func (fetcher *NetworkFetcherRuntime) ConfChange(newConf interface{}) error {
	if conf, ok := newConf.(*capture.Conf); !ok {
		return err.NewTypeNotExpectedError(newConf)
	} else {
		fetcher.UpdateConf(conf)
		return nil
	}
}

func (fetcher *NetworkFetcherRuntime) StatusChange(newStatus interface{}) error {
	if status, ok := newStatus.(*types.StatusSetCommand); !ok {
		return err.NewTypeNotExpectedError(newStatus)
	} else {
		if status.Status == types.Running {
			return fetcher.Start()
		} else if status.Status == types.Stopped {
			return fetcher.Stop()
		} else {
			return fmt.Errorf("unsppored status %s", status.Status)
		}
	}
}

func NewInstance(ctx *base.FetcherBackend) (*NetworkFetcherRuntime, error) {
	fetcher := new(NetworkFetcherRuntime)
	fetcher.fetcherCtx = ctx
	fetcher.captures = make(map[string]*capture.Capture)
	dn, addrs, err := base2.GetNetDeviceName()
	if err != nil {
		return nil, err
	}
	fetcher.fetcher = types.NetworkFetcher{DeviceName: dn, LocalAddr: addrs, BpfFilter: "tcp"}
	//fetcher.updateConfPeriodically()
	//tcp, _ := net.ResolveTCPAddr("tcp", "39.103.21.100:6379")
	//conf := &capture.Conf{
	//	DeviceName:  `\Device\NPF_{7D0089A9-CD14-4C31-926D-B5D61B002A60}`,
	//}
	//fetcher.UpdateConf(conf)
	return fetcher, nil
}
