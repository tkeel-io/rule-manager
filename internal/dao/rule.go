package dao

import (
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
