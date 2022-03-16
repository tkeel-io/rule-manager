package source

import (
	"context"
	"errors"
	"fmt"
)

var sourceManager *SourceManager = NewSourceManager()

func GetSourceManager() *SourceManager {
	return sourceManager
}

type SourceManager struct {
	SourceConfigs map[string]*SourceConfig
	SourceServer  map[string]SourceTransport
	SourceDrivers map[string]Driver //map[protocol]driver.
}

func NewSourceManager() *SourceManager {
	return &SourceManager{
		SourceConfigs: make(map[string]*SourceConfig),
		SourceServer:  make(map[string]SourceTransport),
		SourceDrivers: make(map[string]Driver),
	}
}
func (sm *SourceManager) Register(ctx context.Context, name, endpoint, protoType string) error {
	//exist?
	if _, has := sm.SourceConfigs[name]; has {
		return errors.New("source existed aready.")
	}
	sm.SourceConfigs[name] = &SourceConfig{
		Name:      name,
		Endpoint:  endpoint,
		ProtoType: protoType,
	}
	return nil
}

func (sm *SourceManager) Run(ctx context.Context) {
	for name, conf := range sm.SourceConfigs {
		if driver, ok := sm.SourceDrivers[conf.ProtoType]; ok {
			s, err := driver.NewSourceTransport(ctx, conf.Endpoint, name)
			if nil == err {
				err = s.Run(ctx)
				if nil == err {
					sm.SourceServer[name] = s
					return
				}
				panic(err)
			} else {
				panic(err)
			}
		} else {
			panic(fmt.Sprintf("driver<%s> not exists.", conf.ProtoType))
		}
	}
}

func (sm *SourceManager) Stop() {
	for _, ser := range sm.SourceServer {
		ser.Close(context.TODO())
	}
	sm.SourceConfigs = make(map[string]*SourceConfig)
	sm.SourceServer = make(map[string]SourceTransport)
}
