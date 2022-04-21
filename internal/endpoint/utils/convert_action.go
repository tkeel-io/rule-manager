package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/rule-manager/constant"
	dao "github.com/tkeel-io/rule-manager/internal/dao"
	"github.com/tkeel-io/rule-manager/internal/dao/action_sink"
	sink_chronus "github.com/tkeel-io/rule-manager/internal/dao/action_sink/chronus"
	sink_mysql "github.com/tkeel-io/rule-manager/internal/dao/action_sink/mysql"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	util "github.com/tkeel-io/rule-manager/pkg/util"
	commonlog "github.com/tkeel-io/rule-util/pkg/commonlog"
	"gorm.io/gorm"
)

const convertLogTitle = "[ActionConvert]"

type Action struct {
	Id         string
	UserId     string
	ActionType string
	Sink       string
	ErrorFlag  bool
	Metadata   map[string]string
	Body       []byte
}

type Field struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ActionOption struct {
	Version string                 `json:"version"`
	Option  map[string]interface{} `json:"option"`
}

//---------------------------------convert action
func ConvertAction(ctx context.Context, actionId, ruleId uint) *Action {
	targetConnd := &dao.Target{RuleID: ruleId}

	targets := make([]*dao.Target, 0)
	tx := dao.DB().Model(targetConnd).Where(targetConnd)
	result := tx.Find(&targets)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("query error", result.Error)
		return nil
	}

	for _, ac := range targets {
		return ConvertActionZ(ac)
	}
	return nil
}

func ConvertActionZ(ac *dao.Target) *Action {

	const version = "v0.0.1"
	var options = make(map[string]interface{})

	configuration := make(map[string]interface{})
	switch ac.SinkType {
	case constant.Action1Type_Kafka:
		sink := ac.Host
		topic := ac.Value
		options["sink"] = xutils.GenerateUrlKafka(sink, "user", "password", topic)
		options["topic"] = topic
		configuration["sink"] = options["sink"]
		configuration["topic"] = options["topic"]
	case constant.Action1Type_Chronus, constant.Action1Type_MYSQL, constant.Action1Type_POSTGRESQL:

		if ac.Ext == nil {
			log.L().Error("target config error")
			return nil
		}
		err := json.Unmarshal([]byte(*ac.Ext), &configuration)
		if err != nil {
			log.Error(err)
			return nil
		}
		o, ok := configuration[constant.MappingInfoKey]
		if !ok {
			commonlog.ErrorWithFields(convertLogTitle, commonlog.Fields{
				"call":      "ConvertActionZ",
				"adtion_id": ac.ID,
				"rule_id":   ac.RuleID,
				"err":       "configuration has no mapping.",
			})
			return nil
		}
		mapping := util.NewConnectInfoFromJson(o)
		if nil == mapping {
			commonlog.ErrorWithFields(convertLogTitle, commonlog.Fields{
				"call":      "ConvertActionZ",
				"adtion_id": ac.ID,
				"rule_id":   ac.RuleID,
				"err":       "configuration has no mapping.",
			})
			return nil
		}

		//fields mapping..
		//get table informations.
		var tableInfo action_sink.Table
		switch ac.SinkType {
		case constant.Action1Type_Chronus:
			tableInfo, err = sink_chronus.TableInfo(context.Background(), mapping.GetEndpoints(), mapping.TableName)
		case constant.Action1Type_MYSQL:
			tableInfo, err = sink_mysql.TableInfo(context.Background(), mapping.GetEndpoints(), mapping.TableName)
		}

		if nil != err {
			commonlog.ErrorWithFields(convertLogTitle, commonlog.Fields{
				"call":      "ConvertActionZ",
				"table":     mapping.TableName,
				"database":  mapping.GetDatabase(),
				"endpoints": mapping.GetEndpoints(),
				"error":     err,
			})
			return nil
		}
		ftypes := make(map[string]string)
		for _, field := range tableInfo.GetFields() {
			ftypes[field.GetName()] = field.GetType()
		}
		mfields := make(map[string]*Field)
		timestampValue := ""
		for _, field := range mapping.Maps {
			tfname := field.TField.Name
			if _, ok := ftypes[tfname]; ok {
				if field.MField.Name == "" {
					continue
				}
				mfields[tfname] = &Field{
					Type:  ftypes[tfname],
					Value: fmt.Sprintf("properties.telemetry.%s.value", field.MField.Name),
				}
				timestampValue = fmt.Sprintf("properties.telemetry.%s.ts", field.MField.Name)
			}
		}
		if len(mfields) > 0 {
			mfields["id"] = &Field{
				Type:  "String",
				Value: "id",
			}
			mfields["timestamp"] = &Field{
				Type:  "Int64",
				Value: timestampValue,
			}
		}
		//construct
		if len(mapping.GetEndpoints()) > 0 {
			configuration["sink"] = mapping.GetEndpoints()[0]
		} else {
			configuration["sink"] = ""

		}
		options["urls"] = mapping.GetEndpoints()
		options["dbName"] = mapping.GetDatabase()
		options["table"] = mapping.TableName
		options["fields"] = mfields
		commonlog.InfoWithFields(convertLogTitle, commonlog.Fields{
			"call":      "ConvertActionZ",
			"table":     mapping.TableName,
			"database":  mapping.GetDatabase(),
			"endpoints": mapping.GetEndpoints(),
			"tableInfo": tableInfo,
			"option":    options,
			"mapping":   mapping,
			"error":     err,
		})
	default:
	}

	act := &Action{
		Id:         actionId2ActionID(ac.ID),
		UserId:     ac.Rule.UserID,
		ActionType: ac.SinkType,
		Sink:       getString(configuration, "sink"),
		Metadata: map[string]string{
			"version": version,
			"option":  xutils.Encode2String(options),
		},
		ErrorFlag: false,
	}
	return act
}

func getString(m map[string]interface{}, key string) string {
	if value, ok := m[key]; ok {
		if v, ok := value.(string); ok {
			return v
		}
	}
	return ""
}

func ruleId2RulexID(id uint) string {
	return fmt.Sprintf("rule-%d", id)
}

func actionId2ActionID(id uint) string {
	return fmt.Sprintf("action-%d", id)
}
