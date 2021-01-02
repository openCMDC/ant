package core

type MsgConsumer func(topic string, payload interface{}) error

type RpcExecutor func(topic string, payload interface{}) (interface{}, error)
