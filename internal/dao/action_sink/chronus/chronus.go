package chronus

import (
	"github.com/tkeel-io/rule-manager/internal/dao/action_sink"
)

const ChronusSinkId = "ChronusSinkId"

//format url: tcp://ip:port?database=dbname

const actionSinkType_ChronusDB = "chronus"
const SinkChronusHashTableKey = "rulemanager-sinks-chronus"
const SinkLogChronus = "[SinkChronus]"

type ModelField struct {
	Name string `json:"name" mapstructure:"name"`
	Type string `json:"type" mapstructure:"type"`
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

type MapField struct {
	TFieldName string     `json:"tfield_name" mapstructure:"tfield_name"`
	MField     ModelField `json:"mfield" mapstructure:"mfield"`
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

type ModelConfigs struct {
	DeviceName []string
	DeviceId   []string
	ThingId    string
}
type TableConfigs struct {
	Name       string
	Database   string
	Endpoints  []string
	Partitions []string
	Fields     []TableField
}

type ActionConfig struct {
	SinkId    string
	ActionId  string
	DbName    string
	TableName string
	Urls      []string
}
