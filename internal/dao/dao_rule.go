package dao

import (
	"context"
	"time"

	"github.com/tkeel-io/rule-manager/constant"
	daoutils "github.com/tkeel-io/rule-manager/internal/dao/utils"
	"git.internal.yunify.com/manage/common/db"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//https://cwiki.yunify.com/pages/viewpage.action?pageId=40438054
type Rule struct {
	Id           string                  `sql:"id,type:varchar(255), notnull, unique"`
	UserId       string                  `sql:"user_id,type:varchar(255)"`
	Name         string                  `sql:"name,type:varchar(255)"`
	Status       string                  `sql:"statustype:varchar(31)"`
	RuleDesc     string                  `sql:"rule_desc"`
	DataType     uint8                   `sql:"data_type"`
	SelectText   string                  `sql:"select_text"`
	SelectFields []*daoutils.SelectField `sql:"select_fields"`
	TopicType    string                  `sql:"topic_type,type:varchar(255)"`
	ShortTopic   string                  `sql:"short_topic,type:varchar(255)"`
	WhereText    string                  `sql:"where_text,type:varchar(2048)"`
	Ruleql       string                  `sql:"ruleql"`
	RawRequest   string                  `sql:"raw"`
	CreateTime   int64                   `sql:"create_time"`
	UpdateTime   int64                   `sql:"update_time"`
	TableName    struct{}                `sql:"mdmp_rules"`
}

//insert.
func (rule *Rule) Insert(ctx context.Context, tx *pg.Tx, rules []*Rule) (affected int, err error) {
	if nil == rules || len(rules) < 1 {
		return 0, nil
	}
	res, err := tx.ModelContext(ctx, &rules).Insert(&rules)
	if nil != err {
		return 0, err
	} else {
		return res.RowsAffected(), nil
	}
}

//update.
func (rule *Rule) Update(ctx context.Context, tx *pg.Tx, req *daoutils.RuleUpdateReq) (affected int, err error) {

	//update
	query := tx.ModelContext(ctx, rule).
		Where("user_id=?", req.UserId).
		Where("id = ?", req.Id).
		Where("status != ?", constant.RuleStatusBan)

	setRule(query, req)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

//delete
func (rule *Rule) Delete(ctx context.Context, tx *pg.Tx, req *daoutils.RuleDeleteReq) (int, error) {

	if "" == req.Id {
		return 0, nil
	}
	query := tx.ModelContext(ctx, rule).
		Where("id=?", req.Id).
		Where("user_id=?", req.UserId)

	res, err := query.Delete()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (rule *Rule) Select(ctx context.Context) error {
	return db.GetPgInstance().ModelContext(ctx, rule).
		Where("id= ? ", rule.Id).
		Where("user_id = ?", rule.UserId).
		Select()
}

//query
func (rule *Rule) Query(ctx context.Context, cond *daoutils.RuleQueryReq) ([]*Rule, error) {
	var queryBan bool
	rules := make([]*Rule, 0)
	pgdb := db.GetPgInstance()
	query := pgdb.ModelContext(ctx, &rules).
		Where("user_id=?", cond.UserId)

	if nil != cond {
		if nil != cond.Id {
			query.Where("id=?", *cond.Id)
		}
		if nil != cond.Ids {
			query.Where("id in (?)", pg.In(cond.Ids))
		}
		if nil != cond.DataType {
			query.Where("data_type=?", *cond.DataType)
		}
		if nil != cond.Name {
			query.Where("name=?", *cond.Name)
		} else if nil != cond.SearchKey {
			query = query.Where("name like ?", "%"+*cond.SearchKey+"%") //暂时这样写
		}
		//status not ban
		queryBan = cond.FlagQueryBan
		//...some fields...
		daoutils.WherePage(query, cond.Page)
	}
	if !queryBan {
		query.Where("not status=?", "ban")
	}
	err := query.Select(&rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (rule *Rule) Exists(ctx context.Context) (bool, error) {
	query := db.GetPgInstance().ModelContext(ctx, rule).Where("id = ?", rule.Id).Where("user_id=?", rule.UserId)
	return query.Exists()
}

func setRule(query *orm.Query, req *daoutils.RuleUpdateReq) {

	if req.Name != nil {
		query.Set("name = ?", *req.Name)
	}
	if nil != req.RuleDesc {
		query.Set("rule_desc = ?", *req.RuleDesc)
	}
	if nil != req.DataType {
		query.Set("data_type = ?", *req.DataType)
	}
	if nil != req.SelectText {
		query.Set("select_text = ?", *req.SelectText)
	}
	if nil != req.SelectFields {
		query.Set("select_fields= ?", req.SelectFields)
	}
	if nil != req.TopicType {
		query.Set("topic_type = ?", *req.TopicType)
	}
	if nil != req.ShortTopic {
		query.Set("short_topic = ?", *req.ShortTopic)
	}
	if nil != req.WhereText {
		query.Set("where_text = ?", *req.WhereText)
	}
	if nil != req.Status {
		query.Set("status = ?", *req.Status)
	}
	if nil != req.Raw {
		query.Set("raw = ?", *req.Raw)
	}
	//set ruleql.

	query.Set("ruleql = ?", req.Ruleql)
}
