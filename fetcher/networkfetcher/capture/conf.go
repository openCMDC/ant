package capture

import "net"

type Conf struct {
	DeviceName  string
	ListenAddrs []net.Addr
	ConnAddrs   []net.Addr
}

type StatusSetMsg struct {
	CaptureIds []string
	NewStatus  Status
	SetAll     bool
}
