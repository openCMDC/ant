package redis

import (
	"ant/core"
	"ant/fetcher/networkfetcher/base"
	"ant/fetcher/networkfetcher/decoder"
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var redisTable core.Table

func init() {
	decoder.RegisterDecoder(redisName, NewRedisDecoder)
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

type redis struct {
}

type redisEntry struct {
	cmdType redisCmdType
	length  int
	content interface{}
}

func (r *redis) Name() string {
	return redisName
}

func (r *redis) Parse(conn *base.TCPConn, senBackChan chan<- *core.Row) {
	c2sReader := bufio.NewReader(decoder.NewNamedReader(redisName, conn.C2SStream()))
	s2cReader := bufio.NewReader(decoder.NewNamedReader(redisName, conn.S2CStream()))

	type request struct {
		cmd    string
		params string
		ts     int64
	}

	//size set to 1024 in case of pipe mode
	rqCh := make(chan *request, 1024)
	go func() {
		for {
			cmd, para, err := parseRequest(c2sReader)
			if err != nil {
				logrus.Warnf("parse redis request failed {%s}", err.Error())
				break
			}

			if cmd == "" {
				logrus.Warnf("redis request cmd is empty try parse again")
				continue
			}
			ts := time.Now().UnixNano() / 1000 / 1000
			rqCh <- &request{cmd, para, ts}
		}
		close(rqCh)
	}()

	for rq := range rqCh {
		entry, err := parseRedis(s2cReader)
		if err != nil {
			logrus.Warnf("parse redis response failed {%s}", err.Error())
		}
		row := r.buildBaseRow()
		row.Content = append(row.Content, conn.GetClientAddr().String())
		row.Content = append(row.Content, conn.GetServerAddr().String())
		row.Content = append(row.Content, rq.cmd)
		row.Content = append(row.Content, rq.params)
		//todo
		row.Content = append(row.Content, nil)
		row.Content = append(row.Content, time.Now().UnixNano()/1000/1000-rq.ts)
		row.Content = append(row.Content, len(rq.cmd)+len(rq.params))
		row.Content = append(row.Content, entry.size())
		senBackChan <- row
	}

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

func NewRedisDecoder() decoder.Interface {
	return new(redis)
}
