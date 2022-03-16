package qmq

import (
	"context"

	"git.internal.yunify.com/MDMP2/iot/transport"
	"github.com/tkeel-io/rule-manager/internal/event"
)

type listener struct {
	ID string
}

func (l *listener) Invoke(ctx context.Context, message transport.Message) error {

	event.Collect(event.Message{
		Type: event.MessageTypeRuleStatus,
		Data: event.NewRuleStatusEvent(message.TargetEntity(), message.Data()),
	})
	// log.InfoWithFields(sourceDriverQMQLogTitle, log.Fields{
	// 	"desc":         "recv message.",
	// 	"topic":        message.Topic(),
	// 	"method":       message.Method(),
	// 	"targetEntity": message.TargetEntity(),
	// })
	return nil
}

func (*listener) Flush(ctx context.Context) error {
	// TODO Flush Data
	return nil
}
