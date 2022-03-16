package dao

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tkeel-io/rule-manager/constant"
	daoutils "github.com/tkeel-io/rule-manager/internal/dao/utils"
	"git.internal.yunify.com/manage/common/db"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//https://cwiki.yunify.com/pages/viewpage.action?pageId=40438054
type Action struct {
	Id              string                 `sql:"id,type:varchar(255), notnull, unique"`
	UserId          string                 `sql:"user_id,type:varchar(255)"`
	RuleId          string                 `sql:"rule_id,type:varchar(255), notnull"`
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

//insert.
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

//update.
func (action *Action) Update(ctx context.Context, tx *pg.Tx, req *daoutils.ActionUpdateReq) (affected int, err error) {
	//update
	query := tx.ModelContext(ctx, action).
		Where("id = ?", req.Id).
		Where("user_id=?", req.UserId)
	//ruleId
	if nil != req.RuleId {
		query.Where("rule_id=?", *req.RuleId)
	}
	setAction(query, req)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (action *Action) UpdateByUser(ctx context.Context, tx *pg.Tx, req *daoutils.ActionUpdateReq) (affected int, err error) {
	//update
	query := tx.ModelContext(ctx, action).
		//Where("id = ?", req.Id).
		Where("user_id=?", req.UserId).
		Where("status != ?", constant.ActionStatusBan)
	//ruleId
	if nil != req.RuleId {
		query.Where("rule_id=?", *req.RuleId)
	}
	if "" != req.Id {
		query.Where("id = ?", req.Id)
	}
	setAction(query, req)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

//delete
func (action *Action) Delete(ctx context.Context, tx *pg.Tx, req *daoutils.ActionDeleteReq) (int, error) {
	query := tx.ModelContext(ctx, action).
		Where("user_id=?", req.UserId)

	if "" != req.Id {
		query.Where("id=?", req.Id)
	}
	if "" != req.RuleId {
		query.Where("rule_id=?", req.RuleId)
	}
	res, err := query.Delete()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

//query
func (action *Action) Query(ctx context.Context, cond *daoutils.ActionQueryReq) ([]*Action, error) {
	var queryBan bool
	actions := make([]*Action, 0)
	pgdb := db.GetPgInstance()
	//query relations.
	query := pgdb.ModelContext(ctx, &actions).
		Where("user_id=?", cond.UserId)

	//conditions.
	if nil != cond {
		if nil != cond.Id {
			query.Where("id=?", *cond.Id)
		}
		if nil != cond.Ids {
			query.Where("id in (?)", pg.In(cond.Ids))
		}
		if nil != cond.RuleId {
			query.Where("rule_id=?", cond.RuleId)
		}
		if nil != cond.ActionType {
			query.Where("action_type=?", *cond.ActionType)
		}
		if nil != cond.Name {
			query.Where("name=?", *cond.Name)
		} else if nil != cond.SearchKey {
			query = query.Where("name like ?", "%"+*cond.SearchKey+"%") //暂时这样写
		}
		if nil != cond.ConfigStatus {
			query.Where("config_status=?", *cond.ConfigStatus)
		}
		queryBan = cond.FlagQueryBan
		daoutils.WherePage(query, cond.Page)
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
		Where("id= ? ", action.Id).
		Where("user_id = ?", action.UserId)

	if "" != action.RuleId {
		query.Where("rule_id = ?", action.RuleId)
	}
	return query.Select()
}

func (action *Action) Exists(ctx context.Context) (bool, error) {
	query := db.GetPgInstance().ModelContext(ctx, action).Where("id=?", action.Id)
	return query.Exists()
}

func setAction(query *orm.Query, req *daoutils.ActionUpdateReq) {

	if req.Name != nil {
		query.Set("name=?", *req.Name)
	}
	if nil != req.ActionType {
		query.Set("action_type = ?", *req.ActionType)
	}
	if nil != req.ErrorActionFlag {
		query.Set("error_action_flag = ?", *req.ErrorActionFlag)
	}
	if nil != req.Configuration {
		query.Set("configuration = ?", req.Configuration)
	}
	if nil != req.Status {
		query.Set("status=?", req.Status)
	}
	if nil != req.ConfigStatus {
		query.Set("config_status=?", req.ConfigStatus)
	}
	if nil != req.SinkId {
		query.Set("sink_id=?", req.SinkId)
	}
}
