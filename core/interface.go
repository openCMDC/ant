package core

type Remote interface {
	RegisterConsumer(topic, version string, consumer MsgConsumer)

	RegisterExecutor(topic, version string, executor RpcExecutor)

	SendSync(topic string, msg interface{})

	SendAsync(topic string, msg interface{})
}
