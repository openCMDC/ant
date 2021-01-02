package types

import "github.com/google/gopacket/pcap"

type Ant struct {
	Uid   string
	Label map[string]string
}

type Task struct {
	Id       string
	Name     string
	SqlStr   string
	Interval int //unit ms
	Table    string
	Status   Status //created running stopped
}

type NetworkFetcher struct {
	DeviceName string
	LocalAddr  []pcap.InterfaceAddress
	BpfFilter  string
}
