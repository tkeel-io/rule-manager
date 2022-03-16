package event

import (
	"context"
)

/*

event用来同步，rule在metadata的执行状态。


1. event解析mq的消息
2. event只是做数据转发解耦，不做消息的缓存。

*/

func Collect(msg Message) {
	Dispatch(context.TODO(), msg)
	//manager.Send(event)
}

type Subscriber interface {
	Name() string
	Id() string
	OnEvent(IEvent)
}

type IEvent interface {
	Id() string
	Name() string
	Type() string
	Marshal() []byte
}

//---------------------------------------------- base event.

//定义消息格式
type EventBase struct {
	MessageType string `json:"type"`
}

func (this EventBase) Type() string {
	return this.MessageType
}

//Message.
type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

const (
	MessageTypeRuleStatus = "rule-status"
	MessageTypeRuleActive = "rule-active"
)

type ErrorDesc struct {
	Error     string `json:"error"`
	Status    string `json:"status"`
	TimeStamp int64  `json:"time_stamp"`
}
