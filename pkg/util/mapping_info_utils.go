package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/goinggo/mapstructure"
	commonlog "github.com/tkeel-io/rule-util/pkg/commonlog"
)

// 将[]string定义为MyStringList类型
type MyStringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

//sort slice.
func SortStringSlice(s []string) []string {
	ss := MyStringList(s)
	sort.Sort(ss)
	return []string(ss)
}

type ConnectInfo struct {
	User      string `json:"user" mapstructure:"user"`
	password  string
	Database  string   `json:"database" mapstructure:"database"`
	Endpoints []string `json:"endpoints" mapstructure:"endpoints"`
	SinkType  string   `json:"sink_type"`
}

func (this *ConnectInfo) GetPassword() string {
	return this.password
}
func (this *ConnectInfo) SetPassword(pass string) {
	this.password = pass
}

type ModelField struct {
	Name string `json:"name" mapstructure:"name"`
	Type string `json:"type" mapstructure:"type"`
}

type MapField struct {
	TFieldName string     `json:"tfield_name" mapstructure:"tfield_name"`
	MField     ModelField `json:"mfield" mapstructure:"mfield"`
}

type MappingInfo struct {
	ConnInfo  ConnectInfo `json:"conn" mapstructure:"conn"`
	TableName string      `json:"table" mapstructure:"table"`
	Maps      []MapField  `json:"maps" mapstructure:"maps"`
}

func (this *MappingInfo) SetConnInfo(conn ConnectInfo) {
	this.ConnInfo = conn
}

func (this *MappingInfo) GetConnInfo() ConnectInfo {
	return this.ConnInfo
}

func (this *MappingInfo) GetUser() string {
	return this.ConnInfo.User
}
func (this *MappingInfo) GetPassword() string {
	return this.ConnInfo.password
}
func (this *MappingInfo) GetDatabase() string {
	return this.ConnInfo.Database
}
func (this *MappingInfo) GetEndpoints() []string {
	return this.ConnInfo.Endpoints
}

func (this *ConnectInfo) Key() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("[%s]:[%v]", "sink-", SortStringSlice(this.Endpoints)))))
}

func (this *ConnectInfo) Value() []byte {

	buf, err := json.Marshal(this)
	if nil != err {
		commonlog.ErrorWithFields("[MappingInfo]", commonlog.Fields{
			"call":      "Marshal ConnectInfo failed.",
			"user":      this.User,
			"password":  this.password,
			"database":  this.Database,
			"endpoints": this.Endpoints,
			"error":     err,
		})
	}
	return buf
}

func NewConnectInfoFromJson(v interface{}) *MappingInfo {
	if nil == v {
		return nil
	}
	m, ok := v.(map[string]interface{})
	if !ok {
		commonlog.ErrorWithFields("[MappingInfo]", commonlog.Fields{
			"value": v,
			"error": "convert failed.",
		})
		return nil
	}
	mapInfo := &MappingInfo{}
	err := mapstructure.Decode(m, mapInfo)
	if err != nil {
		commonlog.ErrorWithFields("[MappingInfo]", commonlog.Fields{
			"value": v,
			"error": err,
		})
		return nil
	}
	return mapInfo
}

func MergeMapping(m1, m2 *MappingInfo) *MappingInfo {

	if nil == m1 {
		return m2
	}
	if nil != m2 {
		//update maps just.
		mapsm := make(map[string]MapField)
		for _, m := range m1.Maps {
			mapsm[m.TFieldName] = m
		}
		for _, m := range m2.Maps {
			mapsm[m.TFieldName] = m
		}

		//construct...
		m1.Maps = make([]MapField, 0)
		for _, m := range mapsm {
			m1.Maps = append(m1.Maps, m)
		}
	}
	return m1
}
