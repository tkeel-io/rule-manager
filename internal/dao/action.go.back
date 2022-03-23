package dao

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"git.internal.yunify.com/manage/common/db"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/tkeel-io/rule-manager/constant"
)

type Action struct {
	ID              string                 `sql:"id,type:varchar(255), notnull, unique"`
	UserID          string                 `sql:"user_id,type:varchar(255)"`
	RuleID          string                 `sql:"rule_id,type:varchar(255), notnull"`
	Name            string                 `sql:"name,type:varchar(255), notnull"`
	Status          string                 `sql:"status, type:varchar(32)"`
	ConfigStatus    bool                   `sql:"config_status"`
	Configuration   map[string]interface{} `sql:"configuration"`
	ActionType      string                 `sql:"action_type, type:varchar(31), notnull"`
	ActionSinkId    string                 `sql:"sink_id, type:varchar(2048)"`
	ErrorActionFlag bool                   `sql:"error_action_flag"`
	CreateTime      int64                  `sql:"create_time"`
	UpdateTime      int64                  `sql:"update_time"`
	TableName       struct{}               `sql:"mdmp_actions"`
}

func checkAction(actions []*Action) error {
	for _, action := range actions {
		// if nil == action.Configuration {
		// 	return errors.New("action error: configuration is nil.")
		// }
		if "" == action.Name {
			action.Name = "action-"
		}

		if func(actionType string, ats []string) bool {
			at := strings.ToLower(actionType)
			for _, t := range ats {
				if at == t {
					return true
				}
			}
			return false
		}(action.ActionType, constant.ActionTypes) == false {
			return errors.New(fmt.Sprintf("action error: ActionType must %v.", constant.ActionTypes))
		}
	}
	return nil
}

func (action *Action) Insert(ctx context.Context, tx *pg.Tx, actions []*Action) (affected int, err error) {
	if nil == actions || len(actions) < 1 {
		return 0, nil
	}
	if err = checkAction(actions); nil == err {
		res, err := tx.ModelContext(ctx, &actions).Insert(&actions)
		if nil != err {
			return 0, err
		} else {
			return res.RowsAffected(), nil
		}
	}
	return 0, err
}

func (action *Action) Update(ctx context.Context, tx *pg.Tx, cond ActionUpdateCondition) (affected int, err error) {
	//update
	query := tx.ModelContext(ctx, action).
		Where("id = ?", cond.ID).
		Where("user_id=?", cond.UserID)
	//ruleId
	if cond.RuleID != "" {
		query.Where("rule_id=?", cond.RuleID)
	}
	setAction(query, cond)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (action *Action) UpdateByUser(ctx context.Context, tx *pg.Tx, cond ActionUpdateCondition) (affected int, err error) {
	//update
	query := tx.ModelContext(ctx, action).
		//Where("id = ?", cond.ID).
		Where("user_id=?", cond.UserID).
		Where("status != ?", constant.ActionStatusBan)

	if cond.RuleID != "" {
		query.Where("rule_id=?", cond.RuleID)
	}
	if cond.ID != "" {
		query.Where("id = ?", cond.ID)
	}
	setAction(query, cond)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (action *Action) Delete(ctx context.Context, tx *pg.Tx, req ActionDeleteCondition) (int, error) {
	query := tx.ModelContext(ctx, action).
		Where("user_id=?", req.UserID)

	if "" != req.ID {
		query.Where("id=?", req.ID)
	}
	if "" != req.RuleID {
		query.Where("rule_id=?", req.RuleID)
	}
	res, err := query.Delete()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (action *Action) Query(ctx context.Context, cond ActionQueryCondition) ([]*Action, error) {
	var queryBan bool
	actions := make([]*Action, 0)
	pgdb := db.GetPgInstance()
	//query relations.
	query := pgdb.ModelContext(ctx, &actions).
		Where("user_id=?", cond.UserID)

	//conditions.
	if !reflect.DeepEqual(cond, ActionQueryCondition{}) {
		if cond.ID != "" {
			query.Where("id=?", cond.ID)
		}
		if cond.IDs != nil && len(cond.IDs) > 0 {
			query.Where("id in (?)", pg.In(cond.IDs))
		}
		if cond.RuleID != "" {
			query.Where("rule_id=?", cond.RuleID)
		}
		if cond.ActionType != "" {
			query.Where("action_type=?", cond.ActionType)
		}
		if cond.Name != "" {
			query.Where("name=?", cond.Name)
		} else if cond.SearchKey != "" {
			query = query.Where("name like ?", "%"+cond.SearchKey+"%") //TODO: improve search key.
		}
		if nil != cond.ConfigStatus {
			query.Where("config_status=?", *cond.ConfigStatus)
		}
		queryBan = cond.FlagQueryBan
		//daoutils.WherePage(query, cond.Page)
	}
	//status not ban
	if !queryBan {
		query.Where("not status=?", "ban")
	}
	err := query.Select(&actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}

func (action *Action) Select(ctx context.Context) error {
	query := db.GetPgInstance().ModelContext(ctx, action).
		Where("id= ? ", action.ID).
		Where("user_id = ?", action.UserID)

	if "" != action.RuleID {
		query.Where("rule_id = ?", action.RuleID)
	}
	return query.Select()
}

func (action *Action) Exists(ctx context.Context) (bool, error) {
	query := db.GetPgInstance().ModelContext(ctx, action).Where("id=?", action.ID)
	return query.Exists()
}

func setAction(query *orm.Query, cond ActionUpdateCondition) {

	if cond.Name != "" {
		query.Set("name=?", cond.Name)
	}
	if cond.ActionType != "" {
		query.Set("action_type = ?", cond.ActionType)
	}
	if nil != cond.ErrorActionFlag {
		query.Set("error_action_flag = ?", *cond.ErrorActionFlag)
	}
	if cond.Configuration != nil {
		query.Set("configuration = ?", cond.Configuration)
	}
	if cond.Status != "" {
		query.Set("status=?", cond.Status)
	}
	if nil != cond.ConfigStatus {
		query.Set("config_status=?", *cond.ConfigStatus)
	}
	if cond.SinkID != "" {
		query.Set("sink_id=?", cond.SinkID)
	}
}

type ActionUpdateCondition struct {
	ID              string
	UserID          string
	RuleID          string
	SinkID          string
	Name            string
	Status          string
	ActionType      string
	Configuration   map[string]interface{}
	ConfigStatus    *bool
	ErrorActionFlag *bool
}

type ActionDeleteCondition struct {
	ID     string
	UserID string
	RuleID string
}

type ActionQueryCondition struct {
	ID              string
	IDs             []string
	UserID          string
	RuleID          string
	SinkID          string
	Name            string
	ActionType      string
	SearchKey       string
	ErrorActionFlag *bool
	ConfigStatus    *bool
	FlagQueryBan    bool
}
