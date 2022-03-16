package daoutils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	"git.internal.yunify.com/manage/common/log"
)

// /*
// //验证的设计应该是分为两步操作的：（不算上gateway的binding）
// 	1. 创建时验证：验证字段值的合法性
// 	2. 启动时验证，验证运行的可行性

// 	创建规则和action是是否执行运行时检查：尚未确定

// */
// */

func ValidateRule(ctx context.Context, rule *dao.Rule) (bool, error) {
	// Id           string         `sql:"id,type:varchar(255), notnull, unique"`
	// UserId       string         `sql:"user_id,type:varchar(255)"`
	// Name         string         `sql:"name,type:varchar(255), unique"`
	// Status       string         `sql:"statustype:varchar(31)"`
	// RuleDesc     string         `sql:"rule_desc"`
	// DataType     uint8          `sql:"data_type, notnull"`
	// SelectText   string         `sql:"select_text,type:varchar(1024), notnull"`
	// SelectFields []*SelectField `sql:"select_fields"`
	// TopicType    string         `sql:"topic_type,type:varchar(255), notnull"`
	// ShortTopic   string         `sql:"short_topic,type:varchar(255), notnull"`
	// WhereText    string         `sql:"where_text,type:varchar(1024)"`
	// Ruleql       string         `sql:"ruleql, type:varchar(2048)"`
	// RawRequest   string         `sql:"raw"`

	//check Ruleql
	var err error
	if "" == xutils.GenerateRuleql(rule) {
		err = newErr("please check rule, Field<select_text, short_topic, topic_type, where_text>, then update this rule.")
		log.ErrorWithFields("[ValidateAction]", log.Fields{
			"error": err,
		})
		return false, err
	}
	if "" == rule.SelectText {
		err = newErr("field SelectText is empty")
		log.ErrorWithFields("[ValidateRule]", log.Fields{
			"error": err,
		})
		return false, err
	}
	if "" == rule.Name {
		err = newErr("field Name is empty")
		log.ErrorWithFields("[ValidateRule]", log.Fields{
			"error": err,
		})
		return false, err
	}
	if !inSlice(constant.TopicTypes, rule.TopicType) {
		err = newErr(fmt.Sprintf("field TopicType must belong (%s)", strings.Join(constant.TopicTypes, ",")))
		log.ErrorWithFields("[ValidateRule]", log.Fields{
			"error": err,
		})
		return false, err
	}
	// if len(strings.Split(rule.ShortTopic, "/")) != 2 {
	// 	//   short_topic = thingid/deviceid
	// 	err = newErr("field ShortTopic format: thingid/deviceid")
	// 	log.ErrorWithFields("[ValidateRule]", log.Fields{
	// 		"error": err,
	// 	})
	// 	return false, err
	// }
	return true, nil
}

func newErr(s string) error {
	return errors.New(s)
}

func CheckRule(ctx context.Context, ruleId, userId string) (bool, error) {

	if rule, err := QueryRule(ctx, ruleId, userId); nil == err {
		if _, err = ValidateRule(ctx, rule); nil != err {
			log.ErrorWithFields("[ValidateRule]", log.Fields{
				"error": err,
			})
			return false, err
		}
		var actions []*dao.Action
		if actions, err = QueryActions(ctx, &ruleId, nil, userId); nil != err {
			log.ErrorWithFields("[ValidateRule]", log.Fields{
				"error": err,
			})
			return false, err
		}
		for _, ac := range actions {
			if flag, err := ValidateAction(ctx, ac); nil != err || !flag {
				log.ErrorWithFields("[ValidateRule]", log.Fields{
					"error": err,
				})
				return false, err
			}
		}
		return true, nil
	} else {
		log.ErrorWithFields("[ValidateRule]", log.Fields{
			"error": err,
		})
		return false, err
	}
}
