package mysql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	mysql "github.com/tkeel-io/rule-manager/internal/action/mysql"
	"github.com/tkeel-io/rule-manager/internal/dao/action_sink"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
	log "github.com/tkeel-io/rule-util/pkg/commonlog"
)

const MysqlSinkId = "MysqlSinkId"
const SinkLogMysql = "[SinkMysql]"

var formatCreateTable = "CREATE TABLE %s.%s(%s) ENGINE=InnoDB DEFAULT CHARSET=utf8"
var dilimiter, showDilimiter = ",", ",\n"

var EnumBaseFieldTypes []action_sink.BaseFieldType = []action_sink.BaseFieldType{
	{
		Name:        "TinyInt",
		LengthLimit: false,
	},
	{
		Name:        "SmallInt",
		LengthLimit: false,
	},
	{
		Name:        "MediumInt",
		LengthLimit: false,
	},
	{
		Name:        "Int",
		LengthLimit: false,
	},
	{
		Name:        "BigInt",
		LengthLimit: false,
	},
	{
		Name:        "Float",
		LengthLimit: false,
	},
	{
		Name:        "Double",
		LengthLimit: false,
	},
	{
		Name:        "Char",
		LengthLimit: true,
		MinLength:   0,
		MaxLength:   255,
	},
	{
		Name:        "Varchar",
		LengthLimit: true,
		MinLength:   0,
		MaxLength:   65535,
	},
	{
		Name:        "Text",
		LengthLimit: false,
	},
	{
		Name:        "Date",
		LengthLimit: false,
	},
	{
		Name:        "Time",
		LengthLimit: false,
	},
	{
		Name:        "DateTime",
		LengthLimit: false,
	},
	{
		Name:        "Timestamp",
		LengthLimit: false,
	},
}

type Table struct {
	Name   string
	Fields []action_sink.TableField
}

func (t Table) GetName() string {
	return t.Name
}

func (t Table) GetFields() []action_sink.TableField {
	return t.Fields
}

type TableField struct {
	Name string
	Type string
	IsPK bool
}

func (tf TableField) GetName() string {
	return tf.Name
}

func (tf TableField) GetType() string {
	return tf.Type
}

func (tf TableField) ISPK() bool {
	return tf.IsPK
}

type TableConfigs struct {
	Name      string
	Database  string
	Endpoints []string
	Fields    []TableField
}

func Connect(ctx context.Context, endpoints []string, database string) error {
	if len(endpoints) == 0 {
		return errors.New("endpoints can not be empty")
	}
	_, err := mysql.NewTransport().TableList(ctx, endpoints)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"desc":       "Connect mysql failed.",
			"endpointts": endpoints,
			"database":   database,
			"error":      err,
		})
	}
	return err
}

func GetTableFieldTypes() []string {
	types := []string{}
	for _, fieldType := range EnumBaseFieldTypes {
		data, _ := json.Marshal(&fieldType)
		types = append(types, string(data))
	}
	return types
}

func tableList(ctx context.Context, endpoints []string) ([]string, error) {
	tabs, err := mysql.NewTransport().TableList(ctx, endpoints)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "tableList",
			"endpoints": endpoints,
			"error":     err,
		})
		return nil, err
	}

	return tabs, nil
}

func ListTable(ctx context.Context, endpoints []string, db string) ([]action_sink.Table, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints error")
	}
	tabNames, err := tableList(ctx, endpoints)
	if err != nil {
		return nil, err
	}
	tables := make([]action_sink.Table, 0)
	for _, tabName := range tabNames {
		tab, err := mysql.NewTransport().TableInfo(ctx, endpoints, tabName)
		if err != nil {
			continue
		}
		t := Table{
			Name:   tab.Name,
			Fields: make([]action_sink.TableField, 0),
		}
		for _, field := range tab.Fields {
			t.Fields = append(t.Fields, TableField{
				Name: field.Name,
				Type: field.Type,
				IsPK: field.Is_PK,
			})
		}
		tables = append(tables, t)
	}
	return tables, nil
}

func TableInfo(ctx context.Context, endpoints []string, tableName string) (action_sink.Table, error) {
	if len(endpoints) == 0 {
		return nil, errors.New("endpoints can not be empty")
	}
	tableNames, err := tableList(ctx, endpoints)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "TableInfo",
			"endpoints": endpoints,
			"tableName": tableName,
			"error":     err,
		})
		return nil, err
	}

	exist := false
	for _, tabName := range tableNames {
		if tabName == tableName {
			exist = true
			break
		}
	}

	if !exist {
		err := errors.New(fmt.Sprintf("table %s not exist", tableName))
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "TableInfo",
			"endpoints": endpoints,
			"tableName": tableName,
			"error":     err,
		})
		return nil, err
	}

	tab, err := mysql.NewTransport().TableInfo(ctx, endpoints, tableName)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "TableInfo",
			"endpoints": endpoints,
			"tableName": tableName,
			"error":     err,
		})
		return nil, err
	}
	table := Table{
		Name:   tableName,
		Fields: []action_sink.TableField{},
	}
	for _, field := range tab.Fields {
		table.Fields = append(table.Fields, TableField{
			Name: field.Name,
			Type: field.Type,
			IsPK: field.Is_PK,
		})
	}
	return table, nil
}

func CreateTable(ctx context.Context, conf *TableConfigs) (string, error) {
	if len(conf.Endpoints) == 0 {
		return "", errors.New("endpoints can not be empty")
	}
	sql, err := genSql(conf, false)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "CreateTable",
			"endpoints": conf.Endpoints,
			"database":  conf.Database,
			"name":      conf.Name,
			"error":     err,
		})
		return "", err
	}

	err = mysql.NewTransport().CreateTable(ctx, conf.Endpoints, sql)
	if err != nil {
		log.ErrorWithFields(SinkLogMysql, log.Fields{
			"call":      "CreateTable",
			"endpoints": conf.Endpoints,
			"database":  conf.Database,
			"name":      conf.Name,
			"error":     err,
		})
		return "", err
	}
	return sql, nil
}

func GenerateSql(ctx context.Context, conf *TableConfigs) (string, error) {
	if len(conf.Endpoints) == 0 {
		return "", errors.New("endpoints can not be empty")
	}
	return genSql(conf, true)
}

//CREATE TABLE `temp` (`id` bigint(20) PRIMARY KEY,`name` varchar(100)) ENGINE=InnoDB DEFAULT CHARSET=utf8;
func genSql(conf *TableConfigs, showFlag bool) (string, error) {
	if conf.Name == "" || conf.Database == "" || len(conf.Fields) == 0 {
		return "", errors.New("database/tableName/fields can not be empty")
	}
	selects := []string{}
	primaryKeys := []string{}
	for _, field := range conf.Fields {
		if field.Name == "" || field.Type == "" {
			return "", errors.New("field name/dateType can not be empty")
		}
		if !xutils.ContainFieldType(EnumBaseFieldTypes, field.Type) {
			err := fmt.Errorf("unknow datatype %s for mysql", field.Type)
			log.ErrorWithFields(SinkLogMysql, log.Fields{
				"call":  "genSql",
				"conf":  conf,
				"error": err,
			})
			return "", err
		}
		if field.IsPK {
			primaryKeys = append(primaryKeys, field.Name)
		}
		selects = append(selects, fmt.Sprintf("%s %s", field.Name, field.Type))
	}
	if len(selects) == 0 {
		err := errors.New("table fields is empty")
		log.DebugWithFields(SinkLogMysql, log.Fields{
			"call":  "genSql",
			"conf":  conf,
			"error": err,
		})
		return "", err
	}
	if len(primaryKeys) != 0 {
		selects = append(selects, fmt.Sprintf("primary key (%s)", strings.Join(primaryKeys, dilimiter)))
	}
	dim := dilimiter
	if showFlag {
		dim = showDilimiter
	}
	selectText := strings.Join(selects, dim)
	sql := fmt.Sprintf(formatCreateTable, conf.Database, conf.Name, selectText)
	return sql, nil
}
