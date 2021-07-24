package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpReqReg(t *testing.T) {
	str := `GET /MFEwTzBNMEswSTAJBgUrDgMCGgUABBTrJdiQ%2Ficg9B19asFe73bPYs%2BreAQUdXGnGUgZvJ2d6kFH35TESHeZ03kCEFslzmkHxCZVZtM5DJmpVK0%3D HTTP/1.1`
	res := checkIsHttpRequestStart(str)
	if res != true {
		t.Errorf("unexpected")
	}
}

func TestHttpRepReg(t *testing.T) {
	//str := `12412412414HTTP/1.1 200 OK`
	//res := httpResponseFirstLineReg.Find([]byte(str))
	//fmt.Println(string(res))
	//res := checkIsHttpResponceStart(str)
	//if res != true {
	//	t.Errorf("unexpected")
	//}
	//h := make(map[int][]string)
	//h[1] = append(h[1], "123")
	//fmt.Println(h)

	rep, err := http.Get("http://p.3.cn/prices/mgets?skuIds")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bs, err := ioutil.ReadAll(rep.Body)
	fmt.Println(string(bs))
}
