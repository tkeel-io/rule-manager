package daoutil

import (
	"context"
	"reflect"
	"time"

	"git.internal.yunify.com/manage/common/db"
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
)

func QueryActions(ctx context.Context, rid, aid string, userId string) ([]*dao.Action, error) {
	action := dao.Action{}
	return action.Query(ctx, dao.ActionQueryCondition{
		RuleID: rid,
		ID:     aid,
		UserID: userId,
	})
}

func UpdateActionSink(ctx context.Context, actionId, sinkId, userId string) error {
	var err, warn error
	var tx *pg.Tx
	var action = &dao.Action{
		ID:     actionId,
		UserID: userId,
	}
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	tx, err = db.GetTransaction()
	if nil != err {
		log.Error(err)
	}

	defer func() {
		//commit transacation.
		err = CommitTransaction(tx, err, "[UpdateRuleStatus]", log.Fields{
			"desc":      "update Action-sink successful.",
			"call":      "UpdateActionSink",
			"action_id": actionId,
			"sink_id":   sinkId,
			"user_id":   userId,
		})
	}()

	//select action.
	if err = action.Select(ctx); nil != err {
		return err
	}
	//validate.
	action.ActionSinkId = sinkId
	action.ConfigStatus, warn = ValidateAction(ctx, action)
	if nil != warn {
		log.WarnWithFields("[UpdateRuleStaus]", log.Fields{
			"desc":           "action update sink.",
			"call":           "UpdateActionSink",
			"action_id":      actionId,
			"user_id":        userId,
			"change_status":  action.ConfigStatus,
			"validate_error": warn,
		})
	}

	actionUpdateReq := dao.ActionUpdateCondition{
		ID:     actionId,
		UserID: userId,
		SinkID: sinkId,
		//update ConfigStatus.
		ConfigStatus: &action.ConfigStatus,
	}

	//update table.
	_, err = action.Update(ctx, tx, actionUpdateReq)
	if nil != err {
		log.Error(err)
	}
	return err
}

func UpdateStatusPGA(ctx context.Context, rid, aid, userId, status string) (err error) {
	log.InfoWithFields("UpdateStatusPGA", log.Fields{
		"rule_id":   rid,
		"action_id": aid,
		"user_id":   userId,
		"status":    status,
	})
	var tx *pg.Tx
	action := dao.Action{}
	if tx, err = db.GetTransaction(); nil == err {
		_, err = action.UpdateByUser(ctx, tx, dao.ActionUpdateCondition{
			ID:     aid,
			RuleID: rid,
			UserID: userId,
			Status: status,
		})

		//commit transacation.
		err = CommitTransaction(tx, err, "[UpdateActionStatus]", log.Fields{
			"desc":      "update action status successful.",
			"action_id": aid,
			"rule_id":   rid,
			"user_id":   userId,
			"status":    status,
			"error":     err,
		})
	}
	return err
}

func MapCatRecursive(m1, m2 map[string]interface{}) map[string]interface{} {

	if nil == m1 {
		m1 = make(map[string]interface{})
	}
	for key, value := range m2 {
		//判断value是否为map
		v1, has := m1[key]
		if !has {
			m1[key] = value
		} else {
			//检查类型...强类型一致
			if reflect.TypeOf(v1).Name() == reflect.TypeOf(value).Name() {
				switch v1.(type) {
				case map[string]interface{}:
					vv1, ok1 := v1.(map[string]interface{})
					vv2, ok2 := value.(map[string]interface{})
					if ok1 && ok2 {
						m1[key] = MapCatRecursive(vv1, vv2)
					} else {
						log.ErrorWithFields("[MapCatRecursive]", log.Fields{
							"desc":   "unkown error.",
							"call":   "MapCatRecursive",
							"source": m1,
							"dest":   m2,
							"key":    key,
						})
					}
				default:
					m1[key] = value
				}
			} else {
				log.ErrorWithFields("[MapCatRecursive]", log.Fields{
					"desc":         "source key asset  failed.",
					"call":         "MapCatRecursive",
					"source_value": v1,
					"dest_value":   value,
					"source":       m1,
					"dest":         m2,
					"key":          key,
				})
			}
		}
	}
	return m1
}

func UpdateSinkPGA(ctx context.Context, aid, userId string, sinkId string, configuration map[string]interface{}, configStatus bool) (err error) {

	var tx *pg.Tx
	//var warn error
	action := &dao.Action{
		// ID:     aid,
		// UserID: userId,
	}
	if tx, err = db.GetTransaction(); nil == err {

		// //select action.
		// if err = action.Select(ctx); nil != err {
		// 	log.ErrorWithFields("[UpdateRuleStaus]", log.Fields{
		// 		"desc":           "action update sink.",
		// 		"call":           "UpdateActionSink",
		// 		"action_id":      aid,
		// 		"user_id":        userId,
		// 		"change_status":  action.ConfigStatus,
		// 		"validate_error": err,
		// 	})
		// 	return err
		// }
		// //validate.
		// configuration = MapCatRecursive(action.Configuration, configuration)
		// action.Configuration = configuration
		// action.ConfigStatus, warn = ValidateAction(ctx, action)
		// if nil != warn {
		// 	log.ErrorWithFields("[UpdateRuleStaus]", log.Fields{
		// 		"desc":           "action update sink.",
		// 		"call":           "UpdateActionSink",
		// 		"action_id":      aid,
		// 		"user_id":        userId,
		// 		"change_status":  action.ConfigStatus,
		// 		"validate_error": warn,
		// 	})
		// }
		_, err = action.Update(ctx, tx, dao.ActionUpdateCondition{
			ID:            aid,
			UserID:        userId,
			SinkID:        sinkId,
			Configuration: configuration,
			ConfigStatus:  &configStatus,
		})

		err = CommitTransaction(tx, err, "[UpdateSinkPGA]", log.Fields{
			"desc":      "update action SinkConfig(sinkId,configuration) successful.",
			"action_id": aid,
			"user_id":   userId,
			"status":    configuration,
			"error":     err,
		})
	}
	return err
}
