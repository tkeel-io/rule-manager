package utils

import (
	"fmt"
	"strings"

	"git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
)

// topicTemp :="/sys/{ThingID}/{DeviceID}/thing/event/property/post"
const TopicFormat = "/sys/%s/thing/%s/+/post"
const TopicFormatMsg = "/sys/%s/thing/%s/base/post"

func GenerateRuleql(rule *dao.Rule) string {

	if !CheckRule(rule) {
		return ""
	}
	//rule.ShortTopic = thingId/deviceId

	sql := fmt.Sprintf("select %s from %s", rule.SelectText, GenerateTopic(rule))
	//conditions.
	if rule.WhereText != "" {
		sql = fmt.Sprintf("%s where %s", sql, rule.WhereText)
	}
	return sql
}

func getEventIdentifier(rule *dao.Rule) string {
	var eventCount int
	ret := "+"
	for _, field := range rule.SelectFields {
		if !strings.HasSuffix(field.Expr, "()") {
			if eventCount > 0 {
				ret = "+"
			} else {
				ret = field.Expr
			}
			eventCount += 1
		}
	}
	if eventCount > 1 {
		log.ErrorWithFields("too much event, must be one.", log.Fields{
			"user_id":       rule.UserID,
			"rule_id":       rule.ID,
			"select_fields": rule.SelectFields,
		})
	}
	return ret
}

// func GenerateRuleqlById(ctx context.Context, id, userId string) string {
// 	rule, err := daoutils.QueryRule(ctx, id, userId)
// 	if nil == err {
// 		return GenerateRuleql(rule)
// 	}
// 	return ""
// }

func GenerateTopic(rule *dao.Rule) string {
	if !CheckRule(rule) {
		return ""
	}
	var topic string
	switch rule.TopicType {
	case constant.TopicTypeRaw:
		return rule.ShortTopic
	case constant.TopicTypeProperty:
		topic = fmt.Sprintf(TopicFormat, rule.ShortTopic, rule.TopicType)
	case constant.TopicTypeEvent:
		topic = fmt.Sprintf("/sys/%s/thing/%s/%s/post", rule.ShortTopic, rule.TopicType, getEventIdentifier(rule))
	default:
	}
	//rule.ShortTopic = thingId/deviceId
	return topic
}

func GenerateTopicMsg(thingId, deviceId, topicType string) string {
	return fmt.Sprintf(TopicFormatMsg, fmt.Sprintf("%s/%s", thingId, deviceId), topicType)
}

func CheckRule(rule *dao.Rule) bool {
	flag1 := CheckSelectText(rule.SelectText)
	flag2 := true // CheckShortTopic(rule.ShortTopic)
	flag3 := CheckWhereText(rule.WhereText)
	flag4 := CheckTopicType(rule.TopicType)
	return flag1 && flag2 && flag3 && flag4
}

func CheckShortTopic(t string) bool {
	ss := strings.Split(t, "/")
	if len(ss) == 2 {
		if "+" != ss[0] {
			return true
		}
	}
	return false
}

func CheckTopicType(t string) bool {
	switch t {
	case constant.TopicTypeProperty:
		return true
	case constant.TopicTypeEvent:
		return true
	case constant.TopicTypeAll:
		return false
	case constant.TopicTypeRaw:
		return true
	default:
		//rule.TopicType = "+"
	}
	return false
}

func CheckSelectText(s string) bool {
	return "" != s
}

func CheckWhereText(where string) bool {
	return true
}

func isFunc(str string) bool {
	return strings.HasSuffix(str, "()")
}
