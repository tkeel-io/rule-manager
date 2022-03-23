package dao

import (
	"fmt"
	"github.com/tkeel-io/rule-manager/constant"
	"gorm.io/gorm"
)

// Const for Rule's Type
const (
	MessageRule uint8 = iota + 1
	TimeseriesRule
)

// Const for Rule's Status
const (
	NotRunningStatus uint8 = iota
	RunningStatus
)

type Rule struct {
	gorm.Model
	UserID string `gorm:"index"`
	Name   string `gorm:"not null;size:255"`
	Status uint8  `gorm:"default:0;comment:'0:not_running,1:running'"`
	Desc   string
	Type   uint8 `gorm:"not null;index;comment:'1:message;2:timeseries'"`
}

func (r *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	if r.Status == 0 {
		r.Status = constant.RuleStatusStop
	}
	return
}

func (r *Rule) Select() *gorm.DB {
	return DB().Model(r).Where(r).First(r)
}

func (r *Rule) Exists() (bool, error) {
	var c string
	result := DB().Model(r).Select("1").Where(r).First(&c)
	if result.Error != nil || result.RowsAffected == 0 {
		return false, result.Error
	}

	return true, nil
}

func (r *Rule) SwitchStatus() error {
	result := DB().Model(r).Where(r).First(&r)
	if result.Error != nil {
		return result.Error
	}
	switch r.Status {
	case NotRunningStatus:
		r.Status = RunningStatus
	case RunningStatus:
		r.Status = NotRunningStatus
	}
	return DB().Model(r).Save(r).Error
}

type RuleEntities struct {
	UniqueKey string `gorm:"uniqueIndex;not null"`
	RuleID    uint
	EntityID  string

	Rule Rule
}

func (e *RuleEntities) BeforeCreate(tx *gorm.DB) error {
	if e.UniqueKey == "" {
		e.UniqueKey = GenUniqueKey(e.RuleID, e.EntityID)
	}

	return nil
}

func (e *RuleEntities) Find(tx *gorm.DB) []RuleEntities {
	if tx == nil {
		tx = DB().Model(e)
	}
	var entities []RuleEntities
	if e.RuleID != 0 {
		tx = tx.Preload("Rule")
	}
	tx.Where(e).Find(&entities)
	return entities
}

func (e *RuleEntities) FindEntityIDS() []string {
	type record struct {
		EntityID string
	}
	var records []record
	DB().Model(e).Where(e).Select("entities_id").Find(&records)
	var ids []string
	for _, r := range records {
		ids = append(ids, r.EntityID)
	}
	return ids
}

func (e *RuleEntities) Count(tx *gorm.DB) (c int64, err error) {
	if tx == nil {
		tx = DB().Model(e)
	}
	err = tx.Where(e).Count(&c).Error
	return
}

const separator = "-"

func GenUniqueKey(ruleID uint, entityID string) string {
	return fmt.Sprintf("%d%s%s", ruleID, separator, entityID)
}
