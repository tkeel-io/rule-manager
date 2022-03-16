package qmq

import (
	"context"

	"git.internal.yunify.com/manage/common/log"

	"github.com/tkeel-io/rule-manager/internal/source"
)

const protocol = "mq"
const sourceDriverQMQLogTitle = "[SourceQMQ]"

type QmQDriver struct {
}

func (m *QmQDriver) NewSourceTransport(ctx context.Context, endpoint string, name string) (source.SourceTransport, error) {

	conf := configT{
		name:  name,
		sink:  endpoint,
		topic: "",
	}
	s := NewQmqSourceTransport(&conf)
	err := s.Open(ctx)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func Init() {

	log.InfoWithFields(sourceDriverQMQLogTitle, log.Fields{
		"desc":     "source register driver.",
		"protocol": protocol,
	})
	drivers := source.GetSourceManager().SourceDrivers
	if _, ok := drivers[protocol]; !ok {
		drivers[protocol] = &QmQDriver{}
	}
	source.GetSourceManager().SourceDrivers = drivers
}
