package remote

import (
	"ant/core"
	topic2 "ant/core/topic"
	"ant/core/types"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type BaseRemote struct {
	msgConsumers map[string]*msgConsumers
	rpcExecutors map[string]*rpcExecutors
}

type MqttRemote struct {
	BaseRemote
	client mqtt.Client
}

type FileRemote struct {
	BaseRemote
	filePath string
}

type msgConsumers struct {
	version   string
	consumers map[string][]core.MsgConsumer
}

type rpcExecutors struct {
	version   string
	executors map[string]core.RpcExecutor
}

func (m *msgConsumers) registerConsumer(topic string, consumer core.MsgConsumer) {
	ca := m.consumers[topic]
	ca = append(ca, consumer)
	m.consumers[topic] = ca
}

func (m *msgConsumers) onMsg(topic string, msg interface{}) {
	if ca, ok := m.consumers[topic]; ok {
		log.WithField("topic", topic).Warn("same rpc topic has multiple executor")
		for _, c := range ca {
			err := c(topic, msg)
			if err != nil {
				log.Warn("consume msg failed", err)
			}
		}
	}
}

func (m *rpcExecutors) registerExecutor(topic string, executor core.RpcExecutor) {
	if _, ok := m.executors[topic]; ok {
		log.WithField("topic", topic).Warn("same rpc topic has multiple executor")
		m.executors[topic] = executor
	} else {
		m.executors[topic] = executor
	}
}

func (m *rpcExecutors) onRpc(topic string, rpc interface{}) (interface{}, error) {
	if ex, ok := m.executors[topic]; ok {
		log.WithField("topic", topic).Warn("same rpc topic has multiple executor")
		return ex(topic, rpc)
	} else {
		return nil, fmt.Errorf("no excutor found for topic %s", topic)
	}
}

func (s *BaseRemote) RegisterConsumer(topic string, version string, consumer core.MsgConsumer) {
	if mc, ok := s.msgConsumers[version]; ok {
		mc.registerConsumer(topic, consumer)
	} else {
		mc := new(msgConsumers)
		mc.version = version
		mc.consumers = make(map[string][]core.MsgConsumer)
		mc.registerConsumer(topic, consumer)
		s.msgConsumers[version] = mc
	}
}

func (s *BaseRemote) RegisterExecutor(topic string, version string, executor core.RpcExecutor) {
	if re, ok := s.rpcExecutors[version]; ok {
		re.registerExecutor(topic, executor)
	} else {
		re := new(rpcExecutors)
		re.version = version
		re.executors = make(map[string]core.RpcExecutor)
		re.registerExecutor(topic, executor)
		s.rpcExecutors[version] = re
	}
}

func (s *BaseRemote) Dispatch(topic string, payload []byte) {
	fields := strings.Split(topic, "/")
	if len(fields) < 3 {
		log.WithField("topic", topic).Warn("invalid topic pattern")
		return
	}
	if fields[1] != "topic" {
		log.WithField("topic", topic).Warn("invalid topic prefix")
		return
	}
	version := fields[2]
	msgType := fields[3]
	if msgType == "msg" {
		s.dispatchMsg(version, topic, payload)
	} else if msgType == "rpc" {
		s.dispatchRpc(version, topic, payload)
	} else {
		log.WithField("msgType", msgType).Warn("unsupported msg type")
	}
}

func (s *MqttRemote) SendSync(topic string, msg interface{}) {
	bs, err := json.Marshal(msg)
	if err != nil {
		log.WithField("error", err).Warn("marshal msg to bytes failed")
		return
	}
	token := s.client.Publish(topic, 0, false, bs)
	if token.WaitTimeout(time.Second * 5) {
		return
	} else {
		if err := token.Error(); err != nil {
			log.WithField("err", err).Warn("send data failed")
		}
	}
}

func (s *FileRemote) SendAsync(topic string, msg interface{}) {
	//todo
	fmt.Println(topic, msg)
}

func (s *FileRemote) SendSync(topic string, msg interface{}) {
	s.SendAsync(topic, msg)
}

func (s *MqttRemote) SendAsync(topic string, msg interface{}) {
	//todo
}

func (s *BaseRemote) dispatchMsg(version, topic string, payload []byte) {
	if mc, ok := s.msgConsumers[version]; !ok {
		log.WithField("version", version).Warn("unsupported version")
	} else {
		msg := new(types.Msg)
		err := json.Unmarshal(payload, msg)
		if err != nil {
			log.WithField("payload", string(payload)).Warn("unmarshal data to types.Msg failed")
			return
		}
		mc.onMsg(topic, msg.Payload)
	}
}

func (s *FileRemote) ReadAndDispatch() {
	file, err := os.Open(s.filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		fs := strings.Split(line, "#")
		s.Dispatch(fs[0], []byte(fs[1]))
	}
}

func (s *MqttRemote) dispatchRpc(version, topic string, payload []byte) {
	if re, ok := s.rpcExecutors[version]; !ok {
		log.WithField("version", version).Warn("unsupported version")
	} else {
		rpc := new(types.Rpc)
		err := json.Unmarshal(payload, rpc)
		if err != nil {
			log.WithField("payload", string(payload)).Warn("unmarshal data to types.Msg failed")
			return
		}
		meta := rpc.Meta
		start := time.Now()
		meta.ReceiveTime = start.Unix()
		result, err := re.onRpc(topic, rpc.Payload)
		end := time.Now()
		if err != nil {
			meta.Success = false
			meta.ErrMsg = err.Error()
			meta.ExecutionTime = int64(end.Sub(start))
		} else {
			meta.Success = true
			meta.ExecutionTime = int64(end.Sub(start))
		}
		rpcResult := new(types.Rpc)
		rpcResult.Meta = meta
		rpcResult.Payload = result
		s.SendAsync(topic2.RpcResponseTopic, rpcResult)
	}
}

func (s *BaseRemote) dispatchRpc(version, topic string, payload []byte) {
	if re, ok := s.rpcExecutors[version]; !ok {
		log.WithField("version", version).Warn("unsupported version")
	} else {
		rpc := new(types.Rpc)
		err := json.Unmarshal(payload, rpc)
		if err != nil {
			log.WithField("payload", string(payload)).Warn("unmarshal data to types.Msg failed")
			return
		}
		meta := rpc.Meta
		start := time.Now()
		meta.ReceiveTime = start.Unix()
		result, err := re.onRpc(topic, rpc.Payload)
		end := time.Now()
		if err != nil {
			meta.Success = false
			meta.ErrMsg = err.Error()
			meta.ExecutionTime = int64(end.Sub(start))
		} else {
			meta.Success = true
			meta.ExecutionTime = int64(end.Sub(start))
		}
		rpcResult := new(types.Rpc)
		rpcResult.Meta = meta
		rpcResult.Payload = result
		fmt.Println(rpcResult)
	}
}

func Conn(endpoint, userName, password string) (*MqttRemote, error) {
	//todo
	remote := new(MqttRemote)
	remote.msgConsumers = make(map[string]*msgConsumers)
	remote.rpcExecutors = make(map[string]*rpcExecutors)
	opts := mqtt.NewClientOptions().AddBroker(endpoint).SetClientID("sample")
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		remote.Dispatch(msg.Topic(), msg.Payload())
	})
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.WaitTimeout(10*time.Second) || token.Error() != nil {
		if token.Error() != nil {
			log.Error("Connect failed", token.Error())
		}
		return nil, fmt.Errorf("Connect to %s failed", endpoint)
	}
	remote.client = c
	return remote, nil
}

func NewFileRemote(path string, delay time.Duration) (*FileRemote, error) {
	//todo
	remote := new(FileRemote)
	remote.msgConsumers = make(map[string]*msgConsumers)
	remote.rpcExecutors = make(map[string]*rpcExecutors)
	remote.filePath = path
	go func() {
		<-time.After(delay)
		remote.ReadAndDispatch()
	}()
	return remote, nil
}
