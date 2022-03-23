package event

import (
	"encoding/json"

	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/kit/log"
)

type RuleActiveEvent struct {
	EventBase
	eventType constant.EventType
	ruleId    string
	userId    string
	actionIds []string
}

func NewRuleActiveEvent(t constant.EventType, rid, uid string, aids []string) *RuleActiveEvent {
	e := &RuleActiveEvent{
		EventBase: EventBase{
			MessageType: MessageTypeRuleActive,
		},
		eventType: t,
		ruleId:    rid,
		userId:    uid,
		actionIds: aids,
	}
	return e
}

func (this *RuleActiveEvent) Id() string {
	return this.ruleId
}

func (this *RuleActiveEvent) RuleId() string {
	return this.ruleId
}
func (this *RuleActiveEvent) UserId() string {
	return this.userId
}
func (this *RuleActiveEvent) ActionIds() []string {
	return this.actionIds
}
func (this *RuleActiveEvent) EventType() constant.EventType {
	return this.eventType
}

func (this *RuleActiveEvent) Name() string {
	return constant.EventNameRuleActive
}

func (this *RuleActiveEvent) Marshal() []byte {
	buf, err := json.Marshal(this)
	if nil != err {
		log.Error(err)
		return nil
	}
	return buf
}
