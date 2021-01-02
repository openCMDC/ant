package types

import "ant/core"

type MsgMeta struct {
	SendTime int64 `json:"sendTime"` //when queen send out this command
}

type RpcMeta struct {
	MsgId         string
	Success       bool
	ErrMsg        string
	SendTime      int64 //when queen send out this command
	ReceiveTime   int64 //when queen receive this command from ant
	ExecutionTime int64 //command execute cost, (unit ms)
}

type Msg struct {
	Meta    MsgMeta     `json:"meta"`
	Payload interface{} `json:"payload"`
}

type Rpc struct {
	Meta    RpcMeta
	Payload interface{}
}

type Status string

const Unknown Status = "unknown"
const Running Status = "running"
const Initialized Status = "initialized"
const Stopped Status = "stopped"

type StatusSetCommand struct {
	Status Status //started, running,stopped
}

type TaskStatusSetCommand struct {
	Status Status //started, running,stopped
	Ids    []string
}

type TaskResult struct {
	Id   string
	Rows []*core.Row
}
