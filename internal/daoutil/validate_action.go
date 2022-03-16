package daoutil

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
)

func ValidateAction(ctx context.Context, action *dao.Action) (bool, error) {

	// ID              string            `sql:"id,type:varchar(255), notnull, unique"`
	// UserID          string            `sql:"user_id,type:varchar(255)"`
	// RuleID          string            `sql:"rule_id,type:varchar(255), notnull"`
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
		if constant.ActionType_Republish != action.ActionType {
			err = newErr("error dispatch action must be republish.")
			log.ErrorWithFields("[ValidateAction]", log.Fields{
				"error": err,
			})
			return false, err
		}
	}

	switch action.ActionType {
	case constant.ActionType_Republish:
	case constant.ActionType_Kafka:
	case constant.ActionType_Bucket:
	case constant.ActionType_Chronus:
	//check actionSinkId
	// if "" == action.ActionSinkId {
	// 	err = errors.New("action not bind table.")
	// 	log.ErrorWithFields("[ValidateAction]", log.Fields{
	// 		"error": err,
	// 	})
	// 	return false, err
	// }
	case constant.ActionType_MYSQL:
	case constant.ActionType_POSTGRESQL:
	default:
		err = newErr(fmt.Sprintf("field ActionType must belong (%s)", strings.Join(constant.ActionTypes, ",")))
		log.ErrorWithFields("[ValidateAction]", log.Fields{
			"error": err,
		})
		return false, err
	}
	flag, err = checkConfiguration(action.ActionType, action.Configuration)
	if nil != err {
		log.ErrorWithFields("[ValidateAction]", log.Fields{
			"error": err,
		})
	}
	return flag, err
}

func checkConfiguration(actionType string, conf map[string]interface{}) (bool, error) {
	if nil == conf {
		return false, newErr("field Configuration is nil")
	}
	switch actionType {
	case constant.ActionType_Republish:
		//约束key: topic
		if _, ok := conf["topic"]; !ok {
			return false, newErr("Configuration should include field: topic")
		}
		return true, nil
	case constant.ActionType_Kafka:
		//约束key： topic, sink
		if _, ok := conf["topic"]; !ok {
			return false, newErr("Configuration should include field: topic")
		}
		if _, ok := conf["sink"]; !ok {
			return false, newErr("Configuration should include field: sink")
		}
	case constant.ActionType_Bucket:
	case constant.ActionType_Chronus:
	case constant.ActionType_MYSQL:
	case constant.ActionType_POSTGRESQL:
	default:
		return false, errors.New("invalid sink type")
	}
	return true, nil
}
