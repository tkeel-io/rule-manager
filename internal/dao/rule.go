package dao

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/pkg/errors"
	"github.com/tkeel-io/core-broker/pkg/util"
	"github.com/tkeel-io/kit/log"
	"gorm.io/gorm"
	"reflect"
	"strings"
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

// The field for Target type
const (
	TargetTypeKafka = iota + 1
	TargetTypeObjectStorage
)

type Rule struct {
	gorm.Model
	UserID      string `gorm:"index"`
	SubID       uint
	Name        string `gorm:"not null;size:255"`
	Status      uint8  `gorm:"default:0;comment:'0:not_running,1:running'"`
	Desc        string
	Type        uint8  `gorm:"not null;index;comment:'1:message;2:timeseries'"`
	ErrEntityID string `gorm:"size:255;null"`
}

func (r *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ErrEntityID == "" {
		randomStr := util.GenerateRandString(8)
		r.ErrEntityID = fmt.Sprintf("ERROR-%s", randomStr)
	}
	_, err = CoreClient.CreateEntity(r.ErrEntityID, r.UserID, "RULE-MANAGER")
	return
}

func (r *Rule) BeforeDelete(tx *gorm.DB) (err error) {
	// 清理相关的数据
	if err = tx.Session(&gorm.Session{NewDB: true}).Model(&RuleEntities{}).
		Where("rule_id = ?", r.ID).Delete(&RuleEntities{}).Error; err != nil {
		return err
	}

	if err = tx.Session(&gorm.Session{NewDB: true}).Model(&Target{}).
		Where("rule_id = ?", r.ID).Delete(&Target{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Rule) Select() *gorm.DB {
	return DB().Model(r).Where(r).First(r)
}

func (r *Rule) Exists() (bool, error) {
	var c string
	result := DB().Model(r).Select("1").Where(r).Limit(1).First(&c)
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

func (r *Rule) Subscribe(id uint) error {
	r.SubID = id
	methodName := fmt.Sprintf("v1/core-broker/subscribe/%d/entities", id)
	type request struct {
		entities []string
	}
	contentData, err := json.Marshal(request{entities: []string{r.ErrEntityID}})
	if err != nil {
		return errors.Wrap(err, "subscriptionRequestData marshal error")
	}

	content := &dapr.DataContent{
		Data:        contentData,
		ContentType: "application/json",
	}
	_, err = d.InvokeMethodWithContent(context.Background(), "core-broker", "POST", methodName, content)
	if err != nil {
		log.Error("subscribe error", err)
		return errors.Wrap(err, "subscribe error")
	}

	return DB().Model(r).Save(r).Error
}

type RuleEntities struct {
	UniqueKey string `gorm:"uniqueIndex;size:255"`
	RuleID    uint
	EntityID  string

	Rule Rule
}

func (e *RuleEntities) BeforeCreate(tx *gorm.DB) error {
	if e.UniqueKey == "" {
		e.UniqueKey = GenUniqueKey(e.RuleID, e.EntityID)
	}

	if err := UpdateEntityRuleInfo(e.EntityID, e.UniqueKey, add); err != nil {
		log.Error("call update entity error", err)
		return err
	}

	return nil
}

func (e *RuleEntities) BeforeDelete(tx *gorm.DB) error {
	if err := UpdateEntityRuleInfo(e.EntityID, e.UniqueKey, reduce); err != nil {
		log.Error("call update entity error", err)
		return err
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
	DB().Model(e).Where(e).Select("entity_id").Find(&records)
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

type Target struct {
	ID     uint  `gorm:"primaryKey"`
	Type   uint8 `gorm:"not null;index;comment:'1:kafka;2:object-storage'"`
	Host   string
	Value  string
	Ext    *string `gorm:"type:json,null"`
	RuleID uint

	Rule Rule
}

func (t *Target) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (t *Target) Create() error {
	return DB().Model(t).Create(t).Error
}

func (t *Target) Find() error {
	return DB().Model(t).Preload("Rule").Where(t).First(t).Error
}

func (t *Target) FindAndAuth(userID string) error {
	if err := DB().Model(t).Preload("Rule").Where(t).First(t).Error; err != nil {
		return err
	}

	if reflect.DeepEqual(t.Rule, Rule{}) && t.RuleID != 0 {
		var rule Rule
		if DB().Model(&Rule{}).
			Where("id=?", t.RuleID).
			Where("user_id=?", userID).
			First(&rule).Error != nil {
			return fmt.Errorf("rule not found")
		}
		t.Rule = rule
	}

	if t.Rule.UserID != userID {
		return fmt.Errorf("rule %d is not yours", t.RuleID)
	}
	return nil
}

func (t *Target) Delete() error {
	return DB().Model(t).Where(t).Delete(t).Error
}

const separator = "-"

func GenUniqueKey(ruleID uint, entityID string) string {
	return fmt.Sprintf("%d%s%s", ruleID, separator, entityID)
}

type choice uint8

const (
	add choice = iota + 1
	reduce
)

func UpdateEntityRuleInfo(entityID, ruleinfo string, c choice) error {
	separator := ","
	patchData := make([]map[string]interface{}, 0)

	device, err := CoreClient.GetDeviceEntity(entityID)
	log.Debug("get device entity:", device)
	if err != nil {
		log.Error("get entity err:", err)
		return err
	}
	val := ruleinfo
	switch c {
	case add:
		if strings.Contains(device.Properties.SysField.RuleInfo, ruleinfo) {
			return nil
		}
		if device.Properties.SysField.RuleInfo != "" {
			val = strings.Join([]string{device.Properties.SysField.RuleInfo, ruleinfo}, separator)
			val = strings.Join([]string{device.Properties.SysField.RuleInfo, ruleinfo}, separator)
		}
	case reduce:
		info := strings.Split(device.Properties.SysField.RuleInfo, separator)
		validAddresses := make([]string, 0, len(info))
		for i := range info {
			if info[i] != ruleinfo {
				validAddresses = append(validAddresses, info[i])
			}
		}
		if len(validAddresses) != 0 {
			val = strings.Join(validAddresses, separator)
		} else {
			val = ""
		}
		log.Debugf("generated val: %s", val)
	}

	patchData = append(patchData, map[string]interface{}{
		"operator": "replace",
		"path":     "sysField._ruleInfo",
		"value":    val,
	})

	log.Debug("patchData:", patchData)
	log.Debug("call patch on choice (add 1, reduce 2):", c)

	if err = CoreClient.PatchEntity(entityID, patchData); err != nil {
		err = errors.Wrap(err, "patch entity err")
		return err
	}

	return nil
}
