package utils

import (
	"context"
	"errors"

	"github.com/tkeel-io/kit/log"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const convertRuleLogTitle = "[RuleConvert]"

type Rule struct {
	UserId  string
	Id      string
	Body    []byte
	Topic   string
	Actions []*Action
}

func ConvertRule(ctx context.Context, id uint, userId string) (*Rule, error) {
	return genRule(ctx, id, userId)
}

func genRule(ctx context.Context, id uint, userId string) (*Rule, error) {
	// query rule.
	var (
		err error
	)

	rule := &dao.Rule{
		ID:     uint(id),
		UserID: userId,
	}
	if result := rule.SelectFirst(); result.Error != nil {
		return nil, result.Error
	}

	res := &Rule{
		Id:      ruleId2RulexID(rule.ID),
		UserId:  rule.TenantID,
		Topic:   xutils.GenerateTopic(rule),
		Body:    []byte(xutils.GenerateRuleql(rule)),
		Actions: make([]*Action, 0),
	}
	// query actions, update rule from metadata
	targetConnd := &dao.Target{RuleID: rule.ID}
	targets := make([]*dao.Target, 0)
	tx := dao.DB().Model(targetConnd).Where(targetConnd)
	result := tx.Find(&targets)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("query error", result.Error)
		return nil, result.Error
	}
	for _, ac := range targets {
		// check action status.
		if act := ConvertActionZ(ac); nil != act {
			res.Actions = append(res.Actions, act)
		} else {
			log.Error(zap.Any(convertRuleLogTitle, map[string]interface{}{
				"call":           "ConvertRule",
				"rule_id":        res.Id,
				"user_id":        res.UserId,
				"topic":          res.Topic,
				"count(actions)": len(targets),
				"count(actives)": len(res.Actions),
			}))
		}
	}

	if rule.SubID > 0 {
		// TODO
		errorSinkHost := "tkeel-middleware-kafka-headless:9092"
		errorTopic := rule.SubEndpoint
		const version = "v0.0.1"
		options := make(map[string]interface{})
		errorSink := xutils.GenerateUrlKafka(errorSinkHost, "user", "password", errorTopic)
		options["sink"] = errorSink
		options["topic"] = errorTopic
		errorAction := &Action{
			Id:         ruleId2RulexID(rule.ID),
			UserId:     res.UserId,
			ActionType: "kafka",
			Sink:       errorSink,
			ErrorFlag:  true,
			Metadata: map[string]string{
				"version": version,
				"option":  xutils.Encode2String(options),
			},
			Body: []byte{},
		}
		res.Actions = append(res.Actions, errorAction)
	}

	log.Info(zap.Any(convertRuleLogTitle, map[string]interface{}{
		"call":           "ConvertRule",
		"rule_id":        res.Id,
		"user_id":        res.UserId,
		"topic":          res.Topic,
		"count(actions)": len(targets),
		"count(actives)": len(res.Actions),
	}))

	return res, err
}
