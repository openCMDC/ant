package http

import (
	"ant/core"
	base2 "ant/fetcher/base"
	"ant/fetcher/networkfetcher/base"
	"bufio"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Http struct {
	requestChan  chan *httpWithTime
	responseChan chan *httpWithTime
	backend      *base2.FetcherBackend
}

type httpWithTime struct {
	time     time.Time
	request  *http.Request
	response *http.Response
}

var httpRequestFirstLineReg, _ = regexp.Compile(`((GET)|(HEAD)|(POST)|(PUT)|(DELETE)|(PATCH))\s.*\sHTTP\/\d\.\d`)
var httpResponseFirstLineReg, _ = regexp.Compile(`HTTP\/\d\.\d\s(\d{0,3})\s\w+`)
var httpTable core.Table

func init() {
	base.RegisterTextBasedDecoder("\r\n", New)
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
		Name:   "HTTP",
		Desc:   "HTTP",
		Fields: fs,
	}
}

func (h *Http) Protocol() string {
	return "HTTP"
}

func (h *Http) IsValidBeginningBytes(bytes []byte, streamType base.StreamType) ([]byte, bool) {
	if streamType == base.Client2ServerStream {
		res := httpRequestFirstLineReg.Find(bytes)
		if res != nil {
			return res, true
		}
	} else if streamType == base.Server2ClientStream {
		res := httpResponseFirstLineReg.Find(bytes)
		if res != nil {
			return res, true
		}
	}
	return nil, false
}

func (h *Http) ParseClient2ServerStream(decider *base.Decider) error {
	reader := bufio.NewReader(decider)
	for {
		request, err := http.ReadRequest(reader)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "EOF") {
				break
			} else {
				return base.NewMismatchedDecoderError(h.Protocol(), err, nil)
			}
		}
		h.requestChan <- &httpWithTime{decider.TimeOfLastByte(), request, nil}
	}
	return nil
}

func (h *Http) ParseServer2ClientStream(decider *base.Decider) error {
	reader := bufio.NewReader(decider)
	for {
		response, err := http.ReadResponse(reader, &http.Request{})
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "EOF") {
				break
			} else {
				return base.NewMismatchedDecoderError(h.Protocol(), err, nil)
			}
		}
		_, err = ioutil.ReadAll(response.Body)
		if err != nil {
			logrus.Warn(err)
			continue
		}
		h.responseChan <- &httpWithTime{decider.TimeOfLastByte(), nil, response}
	}
	return nil
}

func (h *Http) startCombineRequestAndResponse() {
	for {
		request := <-h.requestChan
		response := <-h.responseChan
		if response.time.Before(request.time) {
			response = <-h.responseChan
		}
		httpRequest := request.request
		httpResponse := response.response
		rt := response.time.Sub(request.time).Milliseconds()
		httpResponse.Request = httpRequest
		row := parseHttpReq2Row(httpResponse, int(rt))
		h.backend.ReportRow(row)
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
		logrus.WithError(err).Error("fail to read request body")
	} else {
		content = append(content, string(reqBody))
	}
	repBody, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		content = append(content, "")
		logrus.WithError(err).Error("fail to read response body")
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

var requestChan = make(chan *httpWithTime)
var responseChan = make(chan *httpWithTime)

func New(backend *base2.FetcherBackend) base.Interface {
	decoder := new(Http)
	decoder.backend = backend
	decoder.requestChan = requestChan
	decoder.responseChan = responseChan
	go decoder.startCombineRequestAndResponse()
	return decoder
}
