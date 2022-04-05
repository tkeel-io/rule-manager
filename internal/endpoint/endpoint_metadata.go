package endpoint

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/tkeel-io/kit/log"
	"go.uber.org/zap"

	xutils "github.com/tkeel-io/rule-manager/internal/endpoint/utils"
	xgrpc "github.com/tkeel-io/rule-manager/internal/transport/grpc"
	metapb "github.com/tkeel-io/rule-util/metadata/v1"
)

//1. 创建Rule，在metadata中创建rule
//	1. 直接创建Rule
//	2. 为Rule追加Action
//2. 删除Rule，在metadata中删除Rule
//	1. 直接删除Rule
//	2. 删除Action的时候重新生成Rule

var metaEnd *MetadataEndpoint

const metadataLogTitle = "[MetadataClient]"

type MetadataEndpoint struct {
	Client   *xgrpc.Client
	packedId int64
}

func NewMetadataEndpoint() *MetadataEndpoint {
	cli := xgrpc.NewClient()
	if cli == nil {
		log.Error("new client error")
		return nil
	}
	return &MetadataEndpoint{
		Client: cli,
	}
}

func GetMetadataEndpointIns() *MetadataEndpoint {
	return metaEnd
}

func initMetaEnd() {
	metaEnd = NewMetadataEndpoint()
}

// type EventRuleData struct {
// 	Type        int32
// 	UserId      string
// 	RuleId      string
// 	Identifier  string
// 	NoticeList  string
// 	NoticeCount int32
// 	Name        string
// 	Level       string
// 	TargetType  int32
// 	Description string
// }

// type IndexRuleData struct {
// 	Type        int32
// 	UserId      string
// 	RuleId      string
// 	Items       []IndexRuleDataItem
// 	Opterator   string
// 	NoticeList  string
// 	NoticeCount int32
// 	Name        string
// 	Level       string
// 	TargetType  int32
// 	Period      int32
// 	Description string
// }

// type IndexRuleDataItem struct {
// 	Meter         string
// 	ConditionType string
// 	Thresholds    float32
// }

func logRule(rule *metapb.RuleQL) {
	log.Debug(zap.Any("[RuleAddDebug]", map[string]interface{}{
		"Id":             rule.Id,
		"UserId":         rule.UserId,
		"Body":           string(rule.Body),
		"TopicFilter":    rule.TopicFilter,
		"Count(Actions)": len(rule.Actions),
	}))
	for _, ac := range rule.Actions {
		log.Debug(zap.Any("[RuleAddDebug]", map[string]interface{}{
			"Id":        ac.Id,
			"Type":      ac.Type,
			"Sink":      ac.Sink,
			"Metadata":  ac.Metadata,
			"Body":      string(ac.Body),
			"ErrorFlag": ac.ErrorFlag,
		}))
	}
}

func (c *MetadataEndpoint) AddRule(ctx context.Context, rule *xutils.Rule) error {

	if nil == rule {
		return nil
	}

	var flagStart bool
	rule.UserId = "admin"
	packetId := atomic.AddInt64(&c.packedId, 1)
	ruleql := &metapb.RuleQL{
		Id:          rule.Id,
		UserId:      rule.UserId,
		Body:        rule.Body,
		TopicFilter: rule.Topic,
		Actions:     make([]*metapb.Action, 0),
		RefreshTime: time.Now().Unix(),
	}
	for _, ac := range rule.Actions {
		if !ac.ErrorFlag {
			flagStart = true
		}
		ruleql.Actions = append(ruleql.Actions, &metapb.Action{
			Id:        ac.Id,
			Type:      ac.ActionType,
			Sink:      ac.Sink,
			Metadata:  ac.Metadata,
			Body:      ac.Body,
			ErrorFlag: ac.ErrorFlag,
		})
	}

	//如果没有可用的数据转发action就禁止启动
	if !flagStart {
		err := errors.New("no available action, please check actions.")
		log.Error(zap.Any(metadataLogTitle, map[string]interface{}{
			"code":  err.Error(),
			"error": err.Error(),
		}))
		return err
	}

	logRule(ruleql)

	_, err := c.Client.Rule().AddRule(ctx, &metapb.RuleRequest{
		Header: &metapb.RequestHeader{
			UserId: rule.UserId,
		},
		PacketId: packetId,
		Rule:     ruleql,
	})
	if err != nil {
		log.Error(zap.Any(metadataLogTitle, map[string]interface{}{
			"call":    "AddRule",
			"rule_id": rule.Id,
			"user_id": rule.UserId,
			"error":   err,
		}))
		return err
	}
	// log.DebugWithFields(metadataLogTitle, log.Fields{
	// 	"call":     "AddRule",
	// 	"rule_id":  rule.Id,
	// 	"user_id":  rule.UserId,
	// 	"Header":   resp.Header,
	// 	"PacketId": resp.PacketId,
	// 	//"RuleQl":   ruleql,
	// })
	return nil
}

//require Ruleql{user_id&id}
func (c *MetadataEndpoint) DelRule(ctx context.Context, rule *xutils.Rule) error {

	packetId := atomic.AddInt64(&c.packedId, 1)
	rule.UserId = "admin"
	resp, err := c.Client.Rule().DelRule(ctx, &metapb.RuleRequest{
		Header: &metapb.RequestHeader{
			UserId: rule.UserId,
		},
		PacketId: packetId,
		Rule: &metapb.RuleQL{
			Id:     rule.Id,
			UserId: rule.UserId,
		},
	})
	if err != nil {
		log.Error(zap.Any(metadataLogTitle, map[string]interface{}{
			"call":    "DelRule",
			"rule_id": rule.Id,
			"user_id": rule.UserId,
			"error":   err,
		}))
		return err
	}
	log.Info(zap.Any(metadataLogTitle, map[string]interface{}{
		"call":     "DelRule",
		"Header":   resp.Header,
		"PacketId": resp.PacketId,
		"RuleQl":   resp.RuleQl,
	}))
	return nil
}
