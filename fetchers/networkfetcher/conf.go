package networkfetcher

import "net"

type CaptureConf struct {
	DeviceName  string
	ListenAddrs []net.Addr
	ConnAddrs   []net.Addr
}

type CaptureStatusSetMsg struct {
	CaptureIds []string
	NewStatus  CaptureStatus
	SetAll     bool
}
