package redis

import (
	"ant/core"
	base2 "ant/fetcher/base"
	"ant/fetcher/networkfetcher/base"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"strings"
	"time"
)

var redisTable core.Table

func init() {
	base.RegisterTextBasedDecoder("\r\n", New)
	fs := make([]*core.Field, 0)
	fs = append(fs, &core.Field{Name: "sourceIpAndPort", DataType: core.String})
	fs = append(fs, &core.Field{Name: "dstIpAndPort", DataType: core.String})
	fs = append(fs, &core.Field{Name: "cmd", DataType: core.String})
	fs = append(fs, &core.Field{Name: "params", DataType: core.String})
	fs = append(fs, &core.Field{Name: "response", DataType: core.String})
	fs = append(fs, &core.Field{Name: "rt", DataType: core.Number})
	fs = append(fs, &core.Field{Name: "requestSize", DataType: core.Number})
	fs = append(fs, &core.Field{Name: "responseSize", DataType: core.Number})
	redisTable = core.Table{
		Name:   "redis",
		Desc:   "redis",
		Fields: fs,
	}
}

func New(backend *base2.FetcherBackend) base.Interface {
	decoder := &redis{backend: backend}
	decoder.requestChan = make(chan *redisWithTime)
	decoder.responseChan = make(chan *redisWithTime)

	return decoder
}

type redis struct {
	backend      *base2.FetcherBackend
	requestChan  chan *redisWithTime
	responseChan chan *redisWithTime
}

type redisWithTime struct {
	time     time.Time
	request  *redisRequest
	response *redisEntry
}

func (r *redis) startCombineRequestAndResponse() {
	for {
		request := <-r.requestChan
		response := <-r.responseChan
		if response.time.After(request.time) {
			response = <-r.responseChan
		}
		redisRequest := request.request
		httpResponse := response.response
		rt := response.time.Sub(request.time).Milliseconds()
		row := r.buildBaseRow()
		row.Content = append(row.Content, "")
		row.Content = append(row.Content, "")
		row.Content = append(row.Content, redisRequest.cmd)
		row.Content = append(row.Content, redisRequest.params)
		//todo
		row.Content = append(row.Content, nil)
		row.Content = append(row.Content, rt)
		row.Content = append(row.Content, len(redisRequest.cmd)+len(redisRequest.params))
		row.Content = append(row.Content, httpResponse.size())
		r.backend.ReportRow(row)
	}
}

func (r *redis) Protocol() string {
	return "REDIS"
}

func (r *redis) IsValidBeginningBytes(bytes []byte, streamType base.StreamType) ([]byte, bool) {
	if streamType == base.Client2ServerStream {
		if bytes[0] != '*' {
			return nil, false
		}
		l := string(bytes[1:])
		_, err := strconv.Atoi(l)
		if err != nil {
			return nil, false
		}
		return bytes, true
	} else if streamType == base.Server2ClientStream {
		switch bytes[0] {
		case byte(SIMPLE_STRING):
			return bytes, true
		case byte(ERR_MSG):
			return bytes, true
		case byte(INTEGER):
			_, err := strconv.Atoi(string(bytes[1:]))
			if err != nil {
				return nil, false
			}
			return bytes, true
		case byte(BUIK_STRING):
			_, err := strconv.Atoi(string(bytes[1:]))
			if err != nil {
				return nil, false
			}
			return bytes, true
		case byte(ARRAY):
			_, err := strconv.Atoi(string(bytes[1:]))
			if err != nil {
				return nil, false
			}
			return bytes, true
		default:
			return nil, false
		}
	}
	return bytes, true
}

func (r *redis) ParseClient2ServerStream(decider *base.Decider) error {
	reader := bufio.NewReader(decider)
	for {
		cmd, para, err := parseRequest(reader)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "EOF") {
				break
			} else {
				return base.NewMismatchedDecoderError(r.Protocol(), err, nil)
			}
		}

		if cmd == "" {
			logrus.Warnf("redis request cmd is empty try parse again")
			continue
		}
		rr := &redisRequest{cmd, para}
		r.requestChan <- &redisWithTime{decider.TimeOfLastByte(), rr, nil}
	}
	return nil
}

func (r *redis) ParseServer2ClientStream(decider *base.Decider) error {
	reader := bufio.NewReader(decider)
	for {
		entry, err := parseRedis(reader)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "EOF") {
				break
			} else {
				return base.NewMismatchedDecoderError(r.Protocol(), err, nil)
			}
		}
		r.requestChan <- &redisWithTime{decider.TimeOfLastByte(), nil, entry}
	}
	return nil
}

type redisRequest struct {
	cmd    string
	params string
}

type redisEntry struct {
	cmdType redisCmdType
	length  int
	content interface{}
}

func parseRequest(reader *bufio.Reader) (string, string, error) {
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return "", "", err
		}

		//Filtering useless data
		if !strings.HasPrefix(string(line), "*") {
			continue
		}

		//run
		l := string(line[1:])
		cmdCount, err := strconv.Atoi(l)
		if err != nil {
			return "", "", err
		}
		if cmdCount == 0 {
			return "", "", nil
		}

		cmd, err := readBuilkString(reader)
		if err != nil {
			return "", "", err
		}
		params := ""
		for j := 1; j < cmdCount; j++ {
			content, err := readBuilkString(reader)
			if err != nil {

			}
			params += content + " "
		}
		return cmd, params, nil
	}
}

func parseRedis(reader *bufio.Reader) (*redisEntry, error) {
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	switch line[0] {
	case byte(SIMPLE_STRING):
		str := string(line[1:])
		return newRedisEntry(SIMPLE_STRING, len(line), str), nil
	case byte(ERR_MSG):
		err := string(line[1:])
		return newRedisEntry(ERR_MSG, len(line), err), nil
	case byte(INTEGER):
		number, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return nil, err
		}
		return newRedisEntry(ERR_MSG, len(line), number), nil
	case byte(BUIK_STRING):
		strLen, _ := strconv.Atoi(string(line[1:]))
		if strLen < 0 {
			return newRedisEntry(BUIK_STRING, len(line), "nil"), nil
		}
		if strLen == 0 {
			return newRedisEntry(BUIK_STRING, len(line), ""), nil
		}
		content := readStringWithSpeLength(reader, strLen)
		return newRedisEntry(BUIK_STRING, strLen+len(line), content), nil
	case byte(ARRAY):
		arrLen, _ := strconv.Atoi(string(line[1:]))
		arr := make([]*redisEntry, 0)
		for arrLen > 0 {
			entry, err := parseRedis(reader)
			if err != nil {
				return nil, err
			}
			arr = append(arr, entry)
			arrLen--
		}
		return newRedisEntry(BUIK_STRING, len(line), arr), nil
	default:
		fmt.Printf("unknown line format %s", string(line))
		return nil, nil
	}
}

func readBuilkString(reader *bufio.Reader) (string, error) {
	li, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	if li[0] != byte(BUIK_STRING) {
		return "", fmt.Errorf("attent to read a build string from {%s}", string(li[0]))
	}
	strLen, _ := strconv.Atoi(string(li[1:]))
	if strLen < 0 {
		return "nil", nil
	}
	if strLen == 0 {
		return "", nil
	}
	return readStringWithSpeLength(reader, strLen), nil
}

func readStringWithSpeLength(reader *bufio.Reader, strLen int) string {
	readed := 0
	content := ""
	for {
		l, _, _ := reader.ReadLine()
		content += string(l)
		readed += len(l)
		if readed >= strLen {
			break
		}
		//BUIK_STRING contains crlf
		readed += 2
		content += "\r\n"
	}
	return content
}

func newRedisEntry(tp redisCmdType, len int, content interface{}) *redisEntry {
	return &redisEntry{
		cmdType: tp,
		length:  len,
		content: content,
	}
}

func (r *redisEntry) size() int {
	if r.cmdType != ARRAY {
		return r.length
	} else {
		arr, ok := r.content.([]*redisEntry)
		if !ok {
			return -100
		}
		size := 0
		for _, entry := range arr {
			size += entry.size()
		}
		return size
	}
}

func (r *redis) buildBaseRow() *core.Row {
	res := new(core.Row)
	res.Meta = &redisTable
	res.Content = make([]interface{}, 0)
	return res
}
