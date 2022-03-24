package event

import (
	"context"
	"encoding/json"

	"github.com/tkeel-io/rule-manager/constant"
)

const eventDispatchLogTitle = "[EventDispatch]"

func Dispatch(ctx context.Context, msg Message) {
	//metadata.Message -> RuleStatus
	switch msg.Type {
	case MessageTypeRuleStatus:

		if ev := HandleRuleStatus(msg); nil != ev {
			//dispatch target entity.
			switch ev.TargetEntity() {
			case constant.ActionType_Republish:
				//对于不同action type的事情分别做不同的回调
			case constant.ActionType_Kafka:
			case constant.ActionType_Bucket:
			case constant.ActionType_Chronus:
			}

			GetEventManager().Send(ev)
			return
		}
	case MessageTypeRuleActive:
		//用户对rule的更新停止或启动....
		if ev := HandleRuleActive(msg); nil != ev {
			log.DebugWithFields(eventDispatchLogTitle, log.Fields{
				"type":       msg.Type,
				"event_type": ev.EventType(),
				"rule_id":    ev.RuleId(),
			})
			GetEventManager().Send(ev)
			return
		}
	default:
		log.ErrorWithFields(eventDispatchLogTitle, log.Fields{
			"type": msg.Type,
			"desc": "not spport message type.",
		})
		return
	}

	//error ----
	log.ErrorWithFields(eventDispatchLogTitle, log.Fields{
		"type":  msg.Type,
		"error": "handle message error.",
	})
}

func HandleRuleStatus(msg Message) *RuleStatusEvent {

	if e, ok := msg.Data.(*RuleStatusEvent); ok {

		data := e.rawData
		ev := make(map[string]interface{})
		if err := json.Unmarshal(data, &ev); nil != err {
			log.ErrorWithFields(eventDispatchLogTitle, log.Fields{
				"type":   msg.Type,
				"entity": e.TargetEntity(),
				"error":  err,
			})
			return nil
		}
		e.SetData(ev)
		e.SetRawData(nil)
		return e
	}
	return nil
}

func HandleRuleActive(msg Message) *RuleActiveEvent {
	if ev, ok := msg.Data.(*RuleActiveEvent); ok {
		return ev
	}
	return nil
}
