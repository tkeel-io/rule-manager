package utils

import (
	"fmt"

	dao "github.com/tkeel-io/rule-manager/internal/dao"
)

func GenerateRuleql(rule *dao.Rule) string {

	//rule.ShortTopic = thingId/deviceId

	sql := fmt.Sprintf("select %s from %s", "*", GenerateTopic(rule))
	return sql
}

// func GenerateRuleqlById(ctx context.Context, id, userId string) string {
// 	rule, err := daoutils.QueryRule(ctx, id, userId)
// 	if nil == err {
// 		return GenerateRuleql(rule)
// 	}
// 	return ""
// }

func GenerateTopic(rule *dao.Rule) string {
	//rule.ShortTopic = thingId/deviceId
	topic := "rulex"
	return topic
}
