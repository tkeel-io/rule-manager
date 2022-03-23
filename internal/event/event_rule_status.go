package event

import (
	"encoding/json"
	"time"

	"git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/rule-manager/constant"
)

/*

	metadata.Message:
		消息类型存放在msg.Metadata: "type": value

		status的其他字段存储在RawData中：

				type RulexStatus struct {
					UserID    string `json:"user_id"`
					RulexId   string `json:"rulex_id"`
					Status    Status `json:"status"`
					RuleID    string `json:"rule_id"`
					ActionId  string `json:"action_id"`
					Desc      string `json:"desc"`
					Error     string `json:"error"`
					Timestamp int64  `json:"time_stamp"`
				}


*/

// msg := map[string]interface{
// 	"rule_id": "12345678",
// 	"status": "runing",
// 	"desc": "rule is running."
// 	"err":"",
// 	"time_stamp":177758945,
// }

// metadata.ProtoMessage {
//     ID:	65589,
//     RawData: msg.Marshal(),
//     Metadata: map[string]string {
// 		"type":"rule-status",
// 	},
// }

//-----------------------------------------------rule status event---------------------------------------------------
type RuleStatusEvent struct {
	EventBase
	targetEntity string
	Data         map[string]interface{} `json:"data"`
	rawData      []byte
}

func NewRuleStatusEvent(targetEntity string, data []byte) *RuleStatusEvent {
	return &RuleStatusEvent{
		EventBase: EventBase{
			MessageType: MessageTypeRuleStatus,
		},
		targetEntity: targetEntity,
		rawData:      data,
	}
}

func (this *RuleStatusEvent) Id() string {
	return getString(this.Data, "rule_id")
}

func (this *RuleStatusEvent) RuleId() string {
	return getString(this.Data, "rule_id")
}

func (this *RuleStatusEvent) ActionId() string {
	return getString(this.Data, "action_id")
}

func (this *RuleStatusEvent) Name() string {
	return constant.EventNameRuleStatus
}

func (this *RuleStatusEvent) Status() string {
	return getString(this.Data, "status")
}

func (this *RuleStatusEvent) SetStatus(status string) {
	this.Data["status"] = status
}

func (this *RuleStatusEvent) SetRawData(data []byte) {
	this.rawData = data
}

func (this *RuleStatusEvent) GetRawData() []byte {
	return this.rawData
}

func (this *RuleStatusEvent) SetData(data map[string]interface{}) {
	this.Data = data
}

func (this *RuleStatusEvent) UserId() string {
	return getString(this.Data, "user_id")
}

func (this *RuleStatusEvent) TargetEntity() string {
	return this.targetEntity
}

func (this *RuleStatusEvent) TimeStamp() int64 {
	if tt, ok := this.Data["time_stamp"]; ok {
		if t, ok := tt.(int64); ok {
			return t
		}
	}
	return time.Now().Unix()
}

func (this *RuleStatusEvent) Error() string {
	return getString(this.Data, "error")
}

func (this *RuleStatusEvent) Marshal() []byte {
	buf, err := json.Marshal(this)
	if nil != err {
		log.Error(err)
		return nil
	}
	return buf
}

func getString(m map[string]interface{}, key string) string {
	if value, ok := m[key]; ok {
		if v, ok := value.(string); ok {
			return v
		}
	}
	return ""
}
