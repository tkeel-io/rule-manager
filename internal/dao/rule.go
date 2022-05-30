package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/rule-manager/config"
	"github.com/tkeel-io/rule-manager/pkg/metrics"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Const for Rule's Type
const (
	TypeMessageRule uint8 = iota + 1
	TypeTimeseriesRule
)

// Const for Rule's Status
const (
	StatusNotRunning uint8 = iota
	StatusRunning
)

// The field for Target type
const (
	TargetTypeKafka = iota + 1
	TargetTypeObjectStorage
)

// The key of the entity properties
const (
	KeyRawData   = "rawData"
	KeyTeleMetry = "telemetry"
)

// rule type.
const (
	RuleTypeMessage    = 1
	RuleTypeTimeseries = 2
)

const SubscribeService string = "http://localhost:3500/v1.0/invoke/keel/method/apis/core-broker/v1/subscribe/%d"

const SubscriptionIDFormat = "%s_%d_%s"

type Rule struct {
	gorm.Model
	UserID      string `gorm:"index;uniqueIndex:user_rule"`
	TenantID    string
	ModelID     string
	ModelName   string
	SubID       uint
	SubEndpoint string
	Name        string `gorm:"not null;size:255;uniqueIndex:user_rule"`
	Status      uint8  `gorm:"default:0;comment:'0:not_running,1:running'"`
	Desc        string
	Type        uint8 `gorm:"not null;index;comment:'1:message;2:timeseries'"`
}

func (s *Rule) UpdateTenantID(userID, tenantID string) {
	DB().Model(&Rule{}).Where(&Rule{UserID: userID}).Update("tenant_id", tenantID)
}

func (r *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (r *Rule) BeforeDelete(tx *gorm.DB) (err error) {
	// 清理相关的数据
	var res []RuleEntities
	newtx := tx.Session(&gorm.Session{NewDB: true}).Model(&RuleEntities{})
	if err = newtx.Where("rule_id = ?", r.ID).Find(&res).Error; err != nil {
		return err
	}

	for _, re := range res {
		if err = newtx.Where(&re).Delete(&re).Error; err != nil {
			return err
		}
	}
	newtx = tx.Session(&gorm.Session{NewDB: true}).Model(&Target{})
	if err = newtx.Where("rule_id = ?", r.ID).Delete(&Target{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Rule) SelectFirst() *gorm.DB {
	return DB().Model(r).Where(r).First(r)
}

func (r *Rule) Exists() (bool, error) {
	result := DB().Model(r).Where(r).Limit(1).First(r)
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
	case StatusNotRunning:
		r.Status = StatusRunning
	case StatusRunning:
		r.Status = StatusNotRunning
	}
	return DB().Model(r).Save(r).Error
}

func (r *Rule) Subscribe(id uint, auth string) error {
	r.SubID = id

	url := fmt.Sprintf(SubscribeService, r.SubID)

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}
	//	req.Header.Add("Authorization", c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Tkeel-Auth", auth)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	response := make(map[string]interface{})
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respData, &response); err != nil {
		log.Errorf("unmarshal response content: %s \n err:%v", string(respData), err)
		return err
	}
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return errors.New("no data")
	}

	dataEndpoint, ok := data["endpoint"].(string)
	if !ok {
		return errors.New("no endpoint")
	}

	endpointSplit := strings.Split(dataEndpoint, "/")
	r.SubEndpoint = endpointSplit[len(endpointSplit)-1]

	return DB().Model(r).Save(r).Error
}

func (r *Rule) Unsubscribe() error {
	r.SubID = 0
	return DB().Model(r).Save(r).Error
}

func (r *Rule) InitMetrics() {
	type res struct {
		Tenant string
		Count  int
	}
	ress := make([]res, 0)
	DB().Model(r).Select("tenant_id as tenant, count(1) as count").Group("tenant").Find(&ress)
	for _, v := range ress {
		metrics.CollectorRuleNumber.WithLabelValues(v.Tenant).Set(float64(v.Count))
	}
}

type RuleEntities struct {
	UniqueKey string `gorm:"uniqueIndex;size:255"`
	RuleID    uint
	EntityID  string

	Rule Rule
}

func (e *RuleEntities) BeforeCreate(tx *gorm.DB) error {
	if e.EntityID == "" || e.RuleID == 0 {
		return errors.New("RuleEntities.EntityID or RuleEntities.RuleID is empty")
	}
	if e.UniqueKey == "" {
		e.UniqueKey = GenUniqueKey(e.RuleID, e.EntityID)
	}
	subscribeID := fmt.Sprintf(SubscriptionIDFormat, e.EntityID, e.RuleID, config.RuleTopic)
	field := "*"
	switch int(e.Rule.Type) {
	case RuleTypeMessage:
		field = fmt.Sprintf("%s as %s", KeyRawData, KeyRawData)
	case RuleTypeTimeseries:
		field = fmt.Sprintf("%s as %s", KeyTeleMetry, KeyTeleMetry)
	default:
		log.Error("rule type ", e)
	}
	if err := CoreClient().Subscribe(subscribeID, e.EntityID, config.RuleTopic, e.Rule.TenantID, field); err != nil {
		log.Error("Subscribe entity failed", "entity", e.EntityID, "topic", config.RuleTopic, "error", err)
		return err
	}

	if err := UpdateEntityRuleInfo(e.EntityID, "RULE:"+e.UniqueKey, add); err != nil {
		log.Error("call update entity error", err)
		return err
	}

	return nil
}

func (e *RuleEntities) BeforeDelete(tx *gorm.DB) error {
	if e.EntityID == "" || e.RuleID == 0 {
		return errors.New("RuleEntities.EntityID or RuleEntities.RuleID is empty")
	}
	if e.UniqueKey == "" {
		e.UniqueKey = GenUniqueKey(e.RuleID, e.EntityID)
	}
	subscribeID := fmt.Sprintf(SubscriptionIDFormat, e.EntityID, e.RuleID, config.RuleTopic)
	var rule Rule
	if DB().Model(&Rule{}).Where("id=?", e.RuleID).First(&rule).Error != nil {
		return fmt.Errorf("rule not found")
	}
	if err := CoreClient().Unsubscribe(subscribeID, rule.TenantID); err != nil {
		log.Error("call unsubscribe error", err)
		//	return err
	}

	if err := UpdateEntityRuleInfo(e.EntityID, "RULE:"+e.UniqueKey, reduce); err != nil {
		log.Error("call update entity error", err)
		//	return err
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
	Ext    *string `gorm:"type:json;null"`
	RuleID uint

	Status       string `gorm:"column:status"`
	ConfigStatus bool   `gorm:"column:config_status"`
	//	Configuration   map[string]interface{} `gorm:"column:configuration"`
	SinkType        string `gorm:"column:sink_type"`
	SinkId          string `gorm:"column:sink_id"`
	ErrorActionFlag bool   `gorm:"column:error_action_flag"`
	CreateTime      int64  `gorm:"column:create_time"`
	UpdateTime      int64  `gorm:"column:update_time"`
	Rule            Rule
}

func (t *Target) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}

func (t *Target) Update() (err error) {
	return DB().Model(t).Save(t).Error
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
	tx := DB().Begin()
	result := tx.Delete(t)
	if result.Error != nil {
		tx.Rollback()
		log.Error("delete target", result.Error)
	}
	tx.Commit()

	return result.Error
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

	device, err := CoreClient().GetDeviceEntity(entityID)
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
		}
	case reduce:
		info := strings.Split(device.Properties.SysField.RuleInfo, separator)
		splinted := make([]string, 0, len(info))
		for i := range info {
			if info[i] != ruleinfo {
				splinted = append(splinted, info[i])
			}
		}
		if len(splinted) != 0 {
			val = strings.Join(splinted, separator)
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

	if err = CoreClient().PatchEntity(entityID, patchData); err != nil {
		err = errors.Wrap(err, "patch entity err")
		return err
	}

	return nil
}
