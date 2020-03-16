package decoder

import (
	"ant/core"
	"ant/fetcher/networkfetcher/base"
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"regexp"
	"strconv"
	"time"
)

var httpRequestFirstLineReg *regexp.Regexp
var httpResponseFirstLineReg *regexp.Regexp
var httpTable core.Table

func init() {
	httpRequestFirstLineReg, _ = regexp.Compile(`((GET)|(HEAD)|(POST)|(PUT)|(DELETE)|(PATCH))\s.*\sHTTP\/\d\.\d`)
	httpResponseFirstLineReg, _ = regexp.Compile(`HTTP\/\d\.\d\s(\d{0,3})\s\w+`)
	RegisterDecoder("http", NewHttpDecoder)
	fs := make([]*core.Field, 0)
	fs = append(fs, &core.Field{Name: "domain", DataType: core.String})
	fs = append(fs, &core.Field{Name: "version", DataType: core.Number})
	fs = append(fs, &core.Field{Name: "method", DataType: core.String})
	fs = append(fs, &core.Field{Name: "path", DataType: core.String})
	fs = append(fs, &core.Field{Name: "code", DataType: core.String})
	fs = append(fs, &core.Field{Name: "queryParameters", DataType: core.Map})
	fs = append(fs, &core.Field{Name: "requestHeaders", DataType: core.Map})
	fs = append(fs, &core.Field{Name: "responseHeaders", DataType: core.Map})
	fs = append(fs, &core.Field{Name: "requestBody", DataType: core.String})
	fs = append(fs, &core.Field{Name: "responseBody", DataType: core.String})
	fs = append(fs, &core.Field{Name: "responseTime", DataType: core.Number})
	httpTable = core.Table{
		Name:   "http",
		Desc:   "http",
		Fields: fs,
	}
}
func parseHttpReq2Row(rep *http.Response, rt int) *core.Row {
	//todo how to deal with media type http request ？
	res := new(core.Row)
	res.Meta = &httpTable
	content := make([]interface{}, 0)
	req := rep.Request
	url := req.URL
	content = append(content, req.Host)
	content = append(content, req.Proto)
	content = append(content, req.Method)
	content = append(content, url.Path)
	content = append(content, strconv.Itoa(rep.StatusCode))
	content = append(content, parseMap2Map(url.Query()))
	content = append(content, parseMap2Map(req.Header))
	content = append(content, parseMap2Map(rep.Header))
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		content = append(content, "")
		log.WithError(err).Error("fail to read request body")
	} else {
		content = append(content, string(reqBody))
	}
	repBody, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		content = append(content, "")
		log.WithError(err).Error("fail to read responce body")
	} else {
		content = append(content, string(repBody))
	}
	content = append(content, rt)
	res.Content = content
	return res
}
func parseMap2Map(source map[string][]string) map[string]string {
	query := make(map[string]string)
	for k, v := range source {
		if len(v) >= 1 {
			query[k] = v[0]
		} else {
			query[k] = ""
		}
	}
	return query
}

type httpDecoder struct {
}

func (h *httpDecoder) Parse(conn *base.TCPConn, sendBackChan chan<- *core.Row) {
	type resonseWithRt struct {
		rep *http.Response
		rt  int64
	}

	//err := h.LocateToValidStartInS2CStream(conn.S2CStream())
	//if err != nil {
	//	log.WithError(err).Warn("fail to locate start in s2c Stream")
	//	return
	//}
	//log.Info("success locate start in s2c Stream")
	//
	//err = h.LocateToValidStartInC2SStream(conn.C2SStream())
	//if err != nil {
	//	log.WithError(err).Warn("fail to locate start in c2s Stream")
	//	return
	//}
	//log.Info("success locate start in c2s Stream")

	r5 := bufio.NewReader(newNamedReader("http", conn.C2SStream()))
	r6 := bufio.NewReader(newNamedReader("http", conn.S2CStream()))
	for {
		//log.WithField("clientAddr", conn.GetClientAddr().String()).Info("start parse http request")
		req, err := http.ReadRequest(r5)
		if err != nil {
			//log.WithFields(log.Fields{"errMsg": err.Error(), "clientAddr": conn.GetClientAddr().String()}).Warn("parse http request err")
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return
			}
			continue
		}
		if req.Body != nil {
			ioutil.ReadAll(req.Body)
		}
		//log.WithField("clientAddr", conn.GetClientAddr().String()).Info("parse http request success and send 2 channel")
		reqSendTs := time.Now().UnixNano()
		//log.WithField("clientAddr", conn.GetClientAddr().String()).Info("start parse http response")
		rep, err := http.ReadResponse(r6, req)
		if err != nil {
			//log.WithFields(log.Fields{"errMsg": err.Error(), "clientAddr": conn.GetClientAddr().String(),}).Warn("parse http response err")
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return
			}
			continue
		}
		//log.WithField("clientAddr", conn.GetClientAddr().String()).Info("parse http response success and send 2 channel")
		rt := (time.Now().UnixNano() - reqSendTs) / 1000000 // unit ms
		log.Info("耗時", rt)
		row := parseHttpReq2Row(rep, int(rt))
		sendBackChan <- row
	}
}
func (h *httpDecoder) Name() string {
	return "http"
}

//TODO 定位到一次请求的结束位置
func (h *httpDecoder) LocateToValidStartInC2SStream(reader *base.ConcurrentReader) error {
	var totalCount = 0
	br := bufio.NewReader(nil)
	treader := textproto.NewReader(br)
	for {
		line, err := treader.ReadLine()
		if err != nil {
			if err == base.InternalReaderNil {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		if !checkIsHttpRequestStart(line) {
			totalCount += len(line) + 2
			continue
		}
		break
	}
	//reader.SeekTo(totalCount)
	return nil
}
func (h *httpDecoder) LocateToValidStartInS2CStream(reader *base.ConcurrentReader) error {
	var totalCount = 0
	br := bufio.NewReader(nil)
	treader := textproto.NewReader(br)
	for {
		line, err := treader.ReadLine()
		if err != nil {
			if err == base.InternalReaderNil {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		if !checkIsHttpResponceStart(line) {
			totalCount += len(line) + 2
			continue
		}
		break
	}
	//reader.SeekTo(totalCount)
	return nil
}
func (h *httpDecoder) ValidateType() byte {
	return 1
}
func checkIsHttpRequestStart(line string) bool {
	return httpRequestFirstLineReg.MatchString(line)
}
func checkIsHttpResponceStart(line string) bool {
	return httpResponseFirstLineReg.MatchString(line)
}
func NewHttpDecoder() Interface {
	return new(httpDecoder)
}
