package chronus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	chronus "github.com/tkeel-io/rule-manager/internal/action/clickhouse"
	"github.com/tkeel-io/rule-manager/internal/dao/action_sink"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	log "github.com/tkeel-io/rule-util/pkg/commonlog"
)

var aa []string = []string{
	//integer.
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"UInt8",
	"UInt16",
	"UInt32",
	"UInt64",

	//float...
	"Float32",
	"Float64",

	//string
	"String",

	//date.
	"Date",
}

var EnumBaseFieldTypes []action_sink.BaseFieldType = []action_sink.BaseFieldType{
	{
		Name:        "Int8",
		LengthLimit: false,
	},
	{
		Name:        "Int16",
		LengthLimit: false,
	},
	{
		Name:        "Int32",
		LengthLimit: false,
	},
	{
		Name:        "Int64",
		LengthLimit: false,
	},
	{
		Name:        "UInt8",
		LengthLimit: false,
	},
	{
		Name:        "UInt16",
		LengthLimit: false,
	},
	{
		Name:        "UInt32",
		LengthLimit: false,
	},
	{
		Name:        "UInt64",
		LengthLimit: false,
	},
	{
		Name:        "Float32",
		LengthLimit: false,
	},
	{
		Name:        "Float64",
		LengthLimit: false,
	},
	{
		Name:        "String",
		LengthLimit: false,
	},
	{
		Name:        "Date",
		LengthLimit: false,
	},
}

//dbname, table, fields, engine, partition, orders, ttl, index
var formatCreateTable = "CREATE TABLE %s.%s(%s) ENGINE = MergeTree() "

var defaultConf = "ENGINE = MergeTree() PARTITION BY action_date ORDER BY (user_id, device_id, event_id) TTL action_date + toIntervalDay(10) SETTINGS index_granularity = 8192"

const (
	index_granularity = "SETTINGS index_granularity = 8192 "
)

var dilimiter, showDilimiter = ",", ",\n"

var EmptySQL = ""

// "ENGINE = %s PARTITION BY %s  ORDER BY %s TTL %s SETTINGS index_granularity = %d"

func genSql(dbname string, conf *TableConfigs, showFlag bool) (string, error) {
	//gen select.
	var err error
	selects := []string{}
	orders := []string{}
	for _, field := range conf.Fields {
		if field.IsPK {
			orders = append(orders, field.Name)
		}
		if !xutils.ContainFieldType(EnumBaseFieldTypes, field.Type) {
			err = errors.New(fmt.Sprintf("unkown type(%s)", field.Type))
			log.ErrorWithFields(SinkLogChronus, log.Fields{
				"desc":  "generate sql failed.",
				"error": err,
			})
			return EmptySQL, err
		}
		selects = append(selects, fmt.Sprintf("%s %s", field.Name, field.Type))
	}
	if len(selects) == 0 {
		err = errors.New("table fields is empty.")
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":  "generate sql failed.",
			"error": err,
		})
		return EmptySQL, err
	}
	dim := dilimiter
	if showFlag {
		dim = showDilimiter
	}
	selectText := strings.Join(selects, dim)
	sql := fmt.Sprintf(formatCreateTable, dbname, conf.Name, selectText)

	//Partitions       []string
	if len(conf.Partitions) > 0 {
		sql = fmt.Sprintf("%s PARTITION BY (%s) ", sql, strings.Join(conf.Partitions, ","))
	}
	//Orders           []string
	if len(orders) > 0 {
		sql = fmt.Sprintf("%s ORDER BY (%s) ", sql, strings.Join(orders, ","))
	}
	return sql + index_granularity, nil
}

func GetTableFieldTypes() []string {
	types := []string{}
	for _, fieldTypes := range EnumBaseFieldTypes {
		data, _ := json.Marshal(&fieldTypes)
		types = append(types, string(data))
	}
	return types
}

func Connect(ctx context.Context, endpoints []string, database string) error {
	_, err := chronus.GetTransport().TableList(ctx, endpoints)
	if nil != err {
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "Connect chronus failed.",
			"endpointts": endpoints,
			"database":   database,
			"error":      err,
		})
	}
	return err
}

//generate sql.
func GenerateSql(ctx context.Context, conf *TableConfigs) (string, error) {
	if conf.Name == "" {
		conf.Name = "tableName"
	}
	return genSql(conf.Database, conf, true)
}

//create table.
func CreateTable(ctx context.Context, conf *TableConfigs) (string, error) {

	sql, err := genSql(conf.Database, conf, false)
	if nil != err {
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "Connect chronus failed.",
			"endpointts": conf.Endpoints,
			"database":   conf.Database,
			"table":      conf.Name,
			"partitions": conf.Partitions,
			"fields":     conf.Fields,
			"sql":        sql,
			"error":      err,
		})
		return "", err
	}
	err = chronus.GetTransport().CreateTable(ctx, conf.Endpoints, sql)
	if nil != err {
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "Connect chronus failed.",
			"endpointts": conf.Endpoints,
			"database":   conf.Database,
			"table":      conf.Name,
			"partitions": conf.Partitions,
			"fields":     conf.Fields,
			"sql":        sql,
			"error":      err,
		})
		return "", err
	}
	return sql, err
}

//table list
func tableList(ctx context.Context, endpoints []string) ([]string, error) {
	ts, err := chronus.GetTransport().TableList(ctx, endpoints)
	if nil != err {
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "get table list failed.",
			"endpointts": endpoints,
			"error":      err,
		})
		return nil, err
	}
	return ts, nil
}

func TableInfo(ctx context.Context, endpoints []string, tableName string) (action_sink.Table, error) {

	var (
		err    error
		exists bool
		ts     []string
	)
	ts, err = tableList(ctx, endpoints)
	for _, t := range ts {
		if t == tableName {
			exists = true
		}
	}
	if err != nil {
		return nil, pb.ErrFailedClickhouseConnection()
	}
	if !exists {
		err = errors.New("table not exists.")
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "get table information failed.",
			"endpointts": endpoints,
			"table":      tableName,
			"error":      err,
		})
		return nil, pb.ErrFailedTableInfo()
	}
	tableinfo, err := chronus.GetTransport().TableInfo(ctx, endpoints, tableName)
	if nil != err {
		log.ErrorWithFields(SinkLogChronus, log.Fields{
			"desc":       "get table information failed.",
			"endpointts": endpoints,
			"table":      tableName,
			"error":      err,
		})
		return nil, pb.ErrFailedTableInfo()
	}
	//translate table.
	fields := []action_sink.TableField{}
	for _, field := range tableinfo.Fields {
		fields = append(fields, TableField{
			Name: field.Name,
			Type: field.Type,
		})
	}

	return &Table{
		Name:   tableName,
		Fields: fields,
	}, err
}

func ListTable(ctx context.Context, endpoints []string) ([]action_sink.Table, error) {

	var (
		err    error
		ts     []string
		tables []action_sink.Table = make([]action_sink.Table, 0)
	)
	ts, err = tableList(ctx, endpoints)
	if err != nil {
		return nil, err
	}

	for _, t := range ts {

		tableinfo, err := chronus.GetTransport().TableInfo(ctx, endpoints, t)
		if nil != err {
			log.ErrorWithFields(SinkLogChronus, log.Fields{
				"desc":       "get table information failed.",
				"endpointts": endpoints,
				"table":      t,
				"error":      err,
			})
			return nil, err
		}
		//translate table.
		fields := []action_sink.TableField{}
		for _, field := range tableinfo.Fields {
			fields = append(fields, TableField{
				Name: field.Name,
				Type: field.Type,
			})
		}

		tables = append(tables, &Table{
			Name:   t,
			Fields: fields,
		})
	}
	return tables, err
}
