package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	"gorm.io/gorm"
)

const convertLogTitle = "[ActionConvert]"

type Action struct {
	Id         string
	UserId     string
	ActionType string
	Sink       string
	ErrorFlag  bool
	Metadata   map[string]string
	Body       []byte
}

type Field struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ActionOption struct {
	Version string                 `json:"version"`
	Option  map[string]interface{} `json:"option"`
}

//---------------------------------convert action
func ConvertAction(ctx context.Context, actionId, ruleId uint) *Action {
	targetConnd := &dao.Target{RuleID: ruleId}

	targets := make([]*dao.Target, 0)
	tx := dao.DB().Model(targetConnd).Where(targetConnd)
	result := tx.Find(&targets)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("query error", result.Error)
		return nil
	}

	for _, ac := range targets {
		return ConvertActionZ(ac)
	}
	return nil
}

func errorAction(ac *dao.Target, err error) {

}
func ConvertActionZ(ac *dao.Target) *Action {

	const version = "v0.0.1"
	var options = make(map[string]interface{})
	switch ac.Type {
	case constant.ActionType_Kafka:
		sink := ac.Host
		topic := ac.Value
		options["sink"] = xutils.GenerateUrlKafka(sink, "user", "password", topic)
		options["topic"] = topic
	default:
	}

	act := &Action{
		Id:         "id",
		UserId:     ac.Rule.UserID,
		ActionType: "kafka",
		Sink:       options["sink"].(string),
		Metadata: map[string]string{
			"version": version,
			"option":  xutils.Encode2String(options),
		},
		ErrorFlag: false,
	}
	return act
}

func getString(m map[string]interface{}, key string) string {
	if value, ok := m[key]; ok {
		if v, ok := value.(string); ok {
			return v
		}
	}
	return ""
}

func ruleId2RulexID(id uint) string {
	return fmt.Sprintf("rule-%d", id)
}
