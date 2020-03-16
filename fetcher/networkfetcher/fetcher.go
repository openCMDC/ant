package networkfetcher

import (
	"ant/fetcher/base"
	base2 "ant/fetcher/networkfetcher/base"
	"ant/fetcher/networkfetcher/capture"
	"ant/storage"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

type NetworkDataFetcher struct {
	captures   map[string][2]*capture.Capture
	fetcherCtx *base.FetcherCtx
	storage    storage.Interface
}

func (*NetworkDataFetcher) Name() string {
	return "NetworkDataFetcher"
}

func (fetcher *NetworkDataFetcher) Start() error {
	fetcher.updateStatus(&capture.StatusSetMsg{
		CaptureIds: nil,
		NewStatus:  capture.Running,
		SetAll:     true,
	})
	return nil
}

func (fetcher *NetworkDataFetcher) updateConfPeriodically() {
	go func() {
		pid := fetcher.fetcherCtx.AntCtx.TargetProcessID()
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
			dn, err := base2.GetNetDeviceName()
			if err != nil {
				log.WithError(err).Warn("stop update networkfetcher")
				continue
			}
			log.WithField("deviceName", dn).Debug("success get device name")
			conf := &capture.Conf{
				DeviceName:  dn,
				ListenAddrs: lisAddr,
				ConnAddrs:   connAddr,
			}
			fetcher.UpdateConf(conf)
			time.Sleep(time.Hour * 5)
		}
	}()
}

func (fetcher *NetworkDataFetcher) Pause() error {
	panic("implement me")
}

func (fetcher *NetworkDataFetcher) Stop() error {
	panic("implement me")
}

func (fetcher *NetworkDataFetcher) UpdateConf(conf *capture.Conf) {
	log.WithFields(log.Fields{"device name": conf.DeviceName,
		"listen_ip_and_port":  conf.ListenAddrs,
		"connect_ip_and_port": conf.ConnAddrs}).Trace("start updateConf")
	device := conf.DeviceName
	if len(device) == 0 {
		log.Warn("empty device name")
		return
	}
	//c.captureConfs[device] = &conf.Conf

	captures, exist := fetcher.captures[device]
	if exist {
		log.WithField("deviceName", device).Trace("start update existed capture")
		listenCapture, connCapture := captures[0], captures[1]
		if listenCapture == nil || connCapture == nil {
			log.WithFields(log.Fields{"deviceName": device}).Warn("why?")
			return
		}
		listenCapture.SetAddrs(conf.ListenAddrs)
		listenCapture.SetAddrs(conf.ConnAddrs)
		log.WithField("deviceName", device).Trace("update existed capture success")
		//ctx.Send(listenCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
		//ctx.Send(connCapture.Pid, AddrUpdateMsg(conf.ListenAddrs))
	} else {
		log.WithField("deviceName", device).Trace("start instantiate capture")
		lcn := capture.GenerateCaptureName(device, capture.ListenStreamCapture)
		lc := capture.NewCapture(lcn, device, conf.ListenAddrs, capture.ListenStreamCapture, fetcher.fetcherCtx)
		ccn := capture.GenerateCaptureName(device, capture.ConnectStreamCapture)
		cc := capture.NewCapture(ccn, device, conf.ConnAddrs, capture.ConnectStreamCapture, fetcher.fetcherCtx)
		fetcher.captures[device] = [...]*capture.Capture{lc, cc}
		log.WithField("deviceName", device).Trace("instantiate capture success")
	}
}

func (fetcher *NetworkDataFetcher) updateStatus(msg *capture.StatusSetMsg) {
	log.WithFields(log.Fields{"setAll": msg.SetAll, "newStatus": msg.NewStatus, "names": msg.CaptureIds}).Trace("start updateStatus")
	if msg.SetAll {
		for _, cs := range fetcher.captures {
			if cs[0] != nil {
				cs[0].SetStatus(msg.NewStatus)
			}
			if cs[1] != nil {
				cs[1].SetStatus(msg.NewStatus)
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
		a.SetStatus(msg.NewStatus)
	}
}
func (fetcher *NetworkDataFetcher) findCaptureByName(name string) *capture.Capture {
	for _, cs := range fetcher.captures {
		if cs[0].Name == name {
			return cs[0]
		}
		if cs[1].Name == name {
			return cs[1]
		}
	}
	return nil
}

func NewInstance(ctx *base.FetcherCtx) *NetworkDataFetcher {
	fetcher := new(NetworkDataFetcher)
	fetcher.fetcherCtx = ctx
	fetcher.captures = make(map[string][2]*capture.Capture)
	//fetcher.updateConfPeriodically()
	tcp, _ := net.ResolveTCPAddr("tcp", "120.92.182.58:80")
	conf := &capture.Conf{
		DeviceName:  `\Device\NPF_{7D0089A9-CD14-4C31-926D-B5D61B002A60}`,
		ListenAddrs: nil,
		ConnAddrs:   []net.Addr{tcp},
	}
	fetcher.UpdateConf(conf)
	return fetcher
}
