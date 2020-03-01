package networkfetcher

import (
	"fmt"
	"net"
	"testing"
)

func TestNetInterfacesfunc(t *testing.T) {
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("start parse type %T", i)
		//net, success := i.(*net.IPNet)
		//net.
		//addrs, err := i()
		//if err != nil {
		//	panic(err)
		//}
		//
		//for _, addr := range addrs {
		//	if a, ok := addr.(*net.IPNet); ok {
		//		if ip == a.IP.String() {
		//			return i.HardwareAddr.String(), nil
		//		}
		//	}
		//}
		fmt.Println(i)
		fmt.Println("end")
	}
}

func TestGetDeviceName(t *testing.T) {
	deviceName, err := GetNetDeviceName()
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(deviceName)
}
