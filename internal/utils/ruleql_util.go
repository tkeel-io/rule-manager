package utils

import (
	"fmt"
	"strings"

	dao "github.com/tkeel-io/rule-manager/internal/dao"
)

func GenerateSqlExpr(rule *dao.Rule) (string, string, string) {
	selectExpr := rule.SelectExpr
	if selectExpr == "" {
		selectExpr = "SELECT *"
	}
	fromExpr := "FROM " + GenerateTopic(rule)
	return selectExpr, rule.WhereExpr, fromExpr
}

func GenerateRuleql(rule *dao.Rule) string {
	var expr []string
	//rule.ShortTopic = thingId/deviceId
	selectExpr, whereExpr, fromExpr := GenerateSqlExpr(rule)
	if whereExpr != "" {
		expr = make([]string, 3)
		expr[0] = selectExpr
		expr[1] = fromExpr
		expr[2] = whereExpr
	} else {
		expr = make([]string, 2)
		expr[0] = selectExpr
		expr[1] = fromExpr
	}
	//return fmt.Sprintf("select %s from %s", "*", GenerateTopic(rule))
	return strings.Join(expr, " ")
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
	topic := fmt.Sprintf("rulex/rule-%d", rule.ID)
	return topic
}
