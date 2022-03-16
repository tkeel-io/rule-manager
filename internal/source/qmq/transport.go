package qmq

import (
	"context"

	"git.internal.yunify.com/MDMP2/api/pkg/log"
	"git.internal.yunify.com/MDMP2/iot/transport"
	"github.com/tkeel-io/rule-manager/internal/source"
)

type QmqSourceTransport struct {
	source transport.StreamClient
	config source.Configuration
}

func NewQmqSourceTransport(config source.Configuration) *QmqSourceTransport {
	return &QmqSourceTransport{
		config: config,
	}
}

func (this *QmqSourceTransport) Open(ctx context.Context) error {
	var err error
	log.InitLogger("A", "B")
	//new source.
	this.source, err = transport.NewSource(context.Background(), this.config.GetString("sink"))
	if err != nil {
		panic(err)
	}
	return err
}
func (this *QmqSourceTransport) Run(ctx context.Context) error {
	this.source.StartReceiver(ctx, &listener{"A"})
	return nil
}

func (this *QmqSourceTransport) Close(ctx context.Context) error {
	return nil
}
