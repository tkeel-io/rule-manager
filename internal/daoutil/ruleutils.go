package daoutil

import (
	"context"
	"fmt"
	"strings"

	"git.internal.yunify.com/manage/common/db"
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/rule-manager/internal/dao"
)

func QueryRule(ctx context.Context, ruleId, userId string) (*dao.Rule, error) {

	rule := &dao.Rule{
		ID:     ruleId,
		UserID: userId,
	}
	return rule, rule.Select(ctx)
}

func UpdateStatusPG(ctx context.Context, id, userId, status string) (err error) {
	log.InfoWithFields("UpdateStatusPG", log.Fields{
		"rule_id": id,
		"user_id": userId,
		"status":  status,
	})
	var tx *pg.Tx
	rule := dao.Rule{}
	if tx, err = db.GetTransaction(); nil == err {
		_, err = rule.Update(ctx, tx, dao.RuleUpdateCondition{
			ID:     id,
			UserID: userId,
			Status: status,
		})

		//commit transaction.
		err = CommitTransaction(tx, err, "[UpdateRuleStatus]", log.Fields{
			"desc":    "update rule status sucessaful.",
			"rule_id": id,
			"status":  status,
			"error":   err,
		})
	}
	return err
}

func GenerateSelectText(topicType string, fields []*dao.SelectField) string {
	switch topicType {
	case constant.TopicTypeProperty:
		return GenerateSelectTextP(fields)
	case constant.TopicTypeEvent:
		return GenerateSelectTextE(fields)
	case constant.TopicTypeRaw:
		return "*"
	default:

	}
	return ""
}

func GenerateSelectTextP(fields []*dao.SelectField) string {

	//generate alias.
	elems := make([]string, 0)
	for _, field := range fields {
		if "" == field.Alias {
			if !strings.HasSuffix(field.Expr, "()") {
				if "*" != field.Expr {
					elems = append(elems, fmt.Sprintf("params.%s", field.Expr))
				} else {
					elems = append(elems, field.Expr)
				}
			} else {
				alias := strings.TrimSuffix(field.Expr, "()")
				elems = append(elems, fmt.Sprintf("%s as %s", field.Expr, alias))
			}
		} else {
			var prefix string
			if !strings.HasSuffix(field.Expr, "()") {
				prefix = "params."
			}
			elems = append(elems, fmt.Sprintf("%s%s as %s", prefix, field.Expr, field.Alias))
		}
	}
	return strings.Join(elems, ", ")
}

func GenerateSelectTextE(fields []*dao.SelectField) string {

	elems := []string{"params"}
	for _, field := range fields {
		if strings.HasSuffix(field.Expr, "()") {
			alias := field.Alias
			if "" == alias {
				alias = strings.TrimSuffix(field.Expr, "()")
			}
			elems = append(elems, fmt.Sprintf("%s as %s", field.Expr, alias))
		}
	}
	return strings.Join(elems, ", ")
}
