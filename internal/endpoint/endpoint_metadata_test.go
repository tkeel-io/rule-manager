package endpoint

import (
	"context"
	"testing"

	"github.com/tkeel-io/rule-manager/config"
	xutils "github.com/tkeel-io/rule-manager/internal/endpoint/utils"
	//	xgrpc "github.com/tkeel-io/rule-manager/internal/transport/grpc"
)

func TestMetadataEndpoint_AddRule(t *testing.T) {
	config.InitConfig("/Users/sc/liuzhen/rule-manager/config.yml")
	metaEnd := NewMetadataEndpoint()
	rule := &xutils.Rule{
		UserId:  "user-eventaaaaa",
		Id:      "abc7",
		Body:    []byte("select * from abcde"),
		Topic:   "abcde",
		Actions: []*xutils.Action{},
	}
	rule.Actions = append(rule.Actions, &xutils.Action{
		Id:         "action7",
		UserId:     "user-eventaaaaa",
		ActionType: "kafka",
		Sink:       "kafka://127.0.0.1:9092/topic2",
		ErrorFlag:  false,
		Metadata:   map[string]string{"option": `{"sink": "kafka://127.0.0.1:9092/topic2/group1"}`},
		Body:       []byte{},
	})
	metaEnd.AddRule(context.Background(), rule)
}

func TestMetadataEndpoint_DelRule(t *testing.T) {
	config.InitConfig("/Users/sc/liuzhen/rule-manager/config.yml")
	metaEnd := NewMetadataEndpoint()
	rule := &xutils.Rule{
		UserId:  "user-eventaaaaa",
		Id:      "abc7",
		Body:    []byte("select * from abcde"),
		Topic:   "abcd",
		Actions: []*xutils.Action{},
	}

	metaEnd.DelRule(context.Background(), rule)
}
