package daoutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
	commonlog "github.com/tkeel-io/rule-util/pkg/commonlog"
)

func ValidateAction(ctx context.Context, action *dao.Target) (bool, error) {

	// Id              string            `sql:"id,type:varchar(255), notnull, unique"`
	// UserId          string            `sql:"user_id,type:varchar(255)"`
	// RuleId          string            `sql:"rule_id,type:varchar(255), notnull"`
	// Name            string            `sql:"name,type:varchar(255), notnull"`
	// Status          bool              `sql:"status"`
	// Configuration   map[string]string `sql:"configuration, notnull"`
	// ActionType      string            `sql:"action_type, notnull"`
	// ActionSinkId    string            `sql:"sink_id"`
	// ErrorActionFlag bool              `sql:"error_action_flag"`
	// if "" == action.Name {
	// 	return false, newErr("action name is empty")
	// }

	var (
		err  error
		flag bool
	)
	if action.ErrorActionFlag {
		if constant.Action1Type_Republish != action.SinkType {
			err = errors.New("error dispatch action must be republish.")
			commonlog.ErrorWithFields("[ValidateAction]", commonlog.Fields{
				"error": err,
			})
			return false, err
		}
	}

	switch action.SinkType {
	case constant.Action1Type_Republish:
	case constant.Action1Type_Kafka:
	case constant.Action1Type_Bucket:
	case constant.Action1Type_Chronus:
	//check actionSinkId
	// if "" == action.ActionSinkId {
	// 	err = errors.New("action not bind table.")
	// 	log.ErrorWithFields("[ValidateAction]", log.Fields{
	// 		"error": err,
	// 	})
	// 	return false, err
	// }
	case constant.Action1Type_MYSQL:
	case constant.Action1Type_POSTGRESQL:
	default:
		err = errors.New(fmt.Sprintf("field ActionType must belong (%s)", strings.Join(constant.Action1Types, ",")))
		commonlog.ErrorWithFields("[ValidateAction]", commonlog.Fields{
			"error": err,
		})
		return false, err
	}

	actionConfig := make(map[string]interface{})
	if action.Ext == nil {
		err = errors.New("action config error")
		return false, err
	}
	err = json.Unmarshal([]byte(*action.Ext), &actionConfig)
	if err != nil {
		return false, err
	}

	flag, err = checkConfiguration(action.SinkType, actionConfig)
	if nil != err {
		commonlog.ErrorWithFields("[ValidateAction]", commonlog.Fields{
			"error": err,
		})
	}
	return flag, err
}

func checkConfiguration(actionType string, conf map[string]interface{}) (bool, error) {
	if nil == conf {
		return false, errors.New("field Configuration is nil")
	}
	switch actionType {
	case constant.Action1Type_Republish:
		//约束key: topic
		if _, ok := conf["topic"]; !ok {
			return false, errors.New("Configuration should include field: topic")
		}
		return true, nil
	case constant.Action1Type_Kafka:
		//约束key： topic, sink
		if _, ok := conf["topic"]; !ok {
			return false, errors.New("Configuration should include field: topic")
		}
		if _, ok := conf["sink"]; !ok {
			return false, errors.New("Configuration should include field: sink")
		}
	case constant.Action1Type_Bucket:
	case constant.Action1Type_Chronus:
	case constant.Action1Type_MYSQL:
	case constant.Action1Type_POSTGRESQL:
	default:
		return false, errors.New("invalid sink type")
	}
	return true, nil
}
