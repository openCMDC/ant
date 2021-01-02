package remote

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Session interface {
	/**
	 */
	RemoteLabels() map[string]string
	RemoteId() string
	SendMqttMsg(topic string, payload []byte)
}

type ClientSideSession struct {
	status int8 //-1: invalid, 0:connected, 1:BaseRemote id and labels awared
	id     string
	labels map[string]string
	client mqtt.Client
}

type ServerSideSession struct {
}

func (c *ClientSideSession) RemoteLabels() map[string]string {
	return c.labels
}

func (c *ClientSideSession) RemoteId() string {
	return c.id
}

//this call won't be blocked
//but it should be aware of which msg is succeed sent and which is not
func (c *ClientSideSession) SendMqttMsg(topic string, payload []byte) {
	//todo store token to check whether this msg is success sent later
	c.client.Publish(topic, 0, false, payload)
}

func NewClientSideSession() *ClientSideSession {
	return &ClientSideSession{
		status: -1,
		labels: make(map[string]string),
	}
}
