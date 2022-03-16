package dao

import (
	"context"
	"errors"
	"reflect"
	"time"

	"git.internal.yunify.com/manage/common/db"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Rule struct {
	ID           string         `sql:"id,type:varchar(255), notnull, unique"`
	UserID       string         `sql:"user_id,type:varchar(255)"`
	Name         string         `sql:"name,type:varchar(255)"`
	Status       string         `sql:"statustype:varchar(31)"`
	RuleDesc     string         `sql:"rule_desc"`
	DataType     uint8          `sql:"data_type"`
	SelectText   string         `sql:"select_text"`
	SelectFields []*SelectField `sql:"select_fields"`
	TopicType    string         `sql:"topic_type,type:varchar(255)"`
	ShortTopic   string         `sql:"short_topic,type:varchar(255)"`
	WhereText    string         `sql:"where_text,type:varchar(2048)"`
	Ruleql       string         `sql:"ruleql"`
	RawRequest   string         `sql:"raw"`
	CreateTime   int64          `sql:"create_time"`
	UpdateTime   int64          `sql:"update_time"`
	TableName    struct{}       `sql:"rules"`
}

func (rule *Rule) Insert(ctx context.Context, tx *pg.Tx, rules ...*Rule) (affected int, err error) {
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

func (rule *Rule) Update(ctx context.Context, tx *pg.Tx, cond RuleUpdateCondition) (affected int, err error) {

	//update
	query := tx.ModelContext(ctx, rule).
		Where("user_id=?", cond.UserID).
		Where("id = ?", cond.ID)

	setRule(query, cond)
	query.Set("update_time = ?", time.Now().Unix())
	res, err := query.Update()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (rule *Rule) Delete(ctx context.Context, tx *pg.Tx, cond RuleDeleteCondition) (int, error) {

	if "" == cond.Id {
		return 0, nil
	}
	query := tx.ModelContext(ctx, rule).
		Where("id=?", cond.Id).
		Where("user_id=?", cond.UserId)

	res, err := query.Delete()
	if nil != err {
		return 0, err
	}
	return res.RowsAffected(), nil
}

func (rule *Rule) Select(ctx context.Context) error {
	return db.GetPgInstance().ModelContext(ctx, rule).
		Where("id= ? ", rule.ID).
		Where("user_id = ?", rule.UserID).
		Select()
}

// Query By conditions
func (rule *Rule) Query(ctx context.Context, cond RuleQueryCondition) ([]*Rule, error) {
	if cond.UserID == "" {
		return nil, errors.New("user id is empty")
	}
	var queryBan bool
	rules := make([]*Rule, 0)
	dbConn := db.GetPgInstance()
	query := dbConn.ModelContext(ctx, &rules).
		Where("user_id=?", cond.UserID)

	if !reflect.DeepEqual(cond, RuleQueryCondition{}) {
		if cond.ID != "" {
			query.Where("id=?", cond.ID)
		}
		if nil != cond.IDs {
			query.Where("id in (?)", pg.In(cond.IDs))
		}
		if 0 != cond.DataType {
			query.Where("data_type=?", cond.DataType)
		}
		if cond.Name != "" {
			query.Where("name=?", cond.Name)
		} else if cond.SearchKey != "" {
			query = query.Where("name like ?", "%"+cond.SearchKey+"%") // TODO: improve this
		}
		//status not ban
		queryBan = cond.FlagQueryBan
		//...some fields...
		//daoutils.WherePage(query, cond.Page)
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
	query := db.GetPgInstance().ModelContext(ctx, rule).Where("id = ?", rule.ID).Where("user_id=?", rule.UserID)
	return query.Exists()
}

func setRule(query *orm.Query, req RuleUpdateCondition) {

	if req.Name != "" {
		query.Set("name = ?", req.Name)
	}
	if req.RuleDesc != "" {
		query.Set("rule_desc = ?", req.RuleDesc)
	}
	if req.DataType != 0 {
		query.Set("data_type = ?", req.DataType)
	}
	if req.SelectText != "" {
		query.Set("select_text = ?", req.SelectText)
	}
	if nil != req.SelectFields {
		query.Set("select_fields= ?", req.SelectFields)
	}
	if req.TopicType != "" {
		query.Set("topic_type = ?", req.TopicType)
	}
	if req.ShortTopic != "" {
		query.Set("short_topic = ?", req.ShortTopic)
	}
	if req.WhereText != "" {
		query.Set("where_text = ?", req.WhereText)
	}
	if req.Status != "" {
		query.Set("status = ?", req.Status)
	}
	if req.Raw != "" {
		query.Set("raw = ?", req.Raw)
	}
	query.Set("ruleql = ?", req.Ruleql)
}

// SelectField 用于对 rules 进行存储
type SelectField struct {
	Expr  string
	Alias string
	Type  string
}

type RuleDeleteCondition struct {
	Id     string
	UserId string
}

type RuleQueryCondition struct {
	ID           string
	IDs          []string
	UserID       string
	Name         string
	DataType     uint8
	TopicType    string
	ShortTopic   string
	SearchKey    string
	FlagQueryBan bool
}

type RuleUpdateCondition struct {
	ID           string
	UserID       string
	Name         string
	Status       string
	RuleDesc     string
	DataType     uint8
	SelectText   string
	SelectFields []*SelectField
	TopicType    string
	ShortTopic   string
	WhereText    string
	Ruleql       string
	Raw          string
}
