package utils

import (
	"strconv"
	"time"

	types "git.internal.yunify.com/manage/common/constant"
	"git.internal.yunify.com/manage/common/log"
	thing "git.internal.yunify.com/manage/common/proto/thing"
)

const SlimulateThingDataLogTitle = "[SimulateData]"

func GenMessage(fields map[string]*ThingTypeDesc) (res map[string]interface{}) {
	res = genMessage(fields)
	return res
}

func genMessage(fields map[string]*ThingTypeDesc) map[string]interface{} {

	msg := make(map[string]interface{})
	for name, field := range fields {
		msg[name] = genValue(field)
	}
	return msg
}

type ThingTypeDesc struct {
	Type   string                 `json:"type"`
	Define map[string]interface{} `json:"define"`
}

type Value struct {
	Value interface{} `json:"value"`
	Time  int64       `json:"time"`
}

type Metadata struct {
	EntityId  string   `json:"entityId"`
	ModelId   string   `json:"modelId"`
	SourceId  []string `json:"sourceId"`
	EpochTime int64    `json:"epochTime"`
}

type ThingMessage struct {
	Id       string                 `json:"id"`
	Version  string                 `json:"version"`
	Type     string                 `json:"type"`
	Metadata Metadata               `json:"metadata"`
	Params   map[string]interface{} `json:"params"`
}

func genValue(field *ThingTypeDesc) interface{} {

	defer func() {
		if err := recover(); nil != err {
			log.WarnWithFields(SlimulateThingDataLogTitle, log.Fields{
				"desc":  "generate value failed.",
				"error": err,
			})
		}
	}()
	t := time.Now().Unix()
	switch field.Type {
	case types.THING_PROPERTY_TYPE_INT32:
		return &Value{
			Value: int32(32),
			Time:  t,
		}
	case types.THING_PROPERTY_TYPE_FLOAT32:
		return &Value{
			Value: float32(32),
			Time:  t,
		}
	case types.THING_PROPERTY_TYPE_FLOAT64:
		return &Value{
			Value: float64(64),
			Time:  t,
		}
	case types.THING_PROPERTY_TYPE_ARRAY:
		return genArray(field)
	case types.THING_PROPERTY_TYPE_BOOL:
		return &Value{
			Value: true,
			Time:  t,
		}
	case types.THING_PROPERTY_TYPE_ENUM:
		return genEnum(field)
	case types.THING_PROPERTY_TYPE_DATE:
		return &Value{
			Value: "Mon, 02 Jan 2006 15:04:05 GMT",
			Time:  t,
		}
	case types.THING_PROPERTY_TYPE_STRUCT:
		return genStruct(field)
	case types.THING_PROPERTY_TYPE_STRING:
		return &Value{
			Value: "value",
			Time:  t,
		}
	}
	return nil
}

func GetType(t thing.EnumMetaType) string {
	var tp string
	switch t {
	case thing.EnumMetaType_INT32:
		tp = types.THING_PROPERTY_TYPE_INT32
	case thing.EnumMetaType_FLOAT:
		tp = types.THING_PROPERTY_TYPE_FLOAT32
	case thing.EnumMetaType_DOUBLE:
		tp = types.THING_PROPERTY_TYPE_FLOAT64
	case thing.EnumMetaType_STRING:
		tp = types.THING_PROPERTY_TYPE_STRING
	case thing.EnumMetaType_ENUM:
		tp = types.THING_PROPERTY_TYPE_ENUM
	case thing.EnumMetaType_ARRAY:
		tp = types.THING_PROPERTY_TYPE_ARRAY
	case thing.EnumMetaType_BOOL:
		tp = types.THING_PROPERTY_TYPE_BOOL
	case thing.EnumMetaType_STRUCT:
		tp = types.THING_PROPERTY_TYPE_STRUCT
	case thing.EnumMetaType_DATE:
		tp = types.THING_PROPERTY_TYPE_DATE
	default:
	}
	return tp
}

func genArray(field *ThingTypeDesc) interface{} {

	values := make([]interface{}, 0)

	sizeS, _ := field.Define["size"].(string)
	if size, err := strconv.Atoi(sizeS); nil != err {
		panic(err)
	} else {
		item, _ := field.Define["item"].(map[string]interface{})
		typ, _ := item["type"].(string)
		var define map[string]interface{}
		switch typ {
		case types.THING_PROPERTY_TYPE_ARRAY:
			fallthrough
		case types.THING_PROPERTY_TYPE_BOOL:
			fallthrough
		case types.THING_PROPERTY_TYPE_ENUM:
			fallthrough
		case types.THING_PROPERTY_TYPE_STRUCT:
			define, _ = item["define"].(map[string]interface{})
		default:
			//do nothing...
		}
		//generate value.
		v := genValue(&ThingTypeDesc{
			Type:   typ,
			Define: define,
		})
		for i := 0; i < size; i++ {
			if nil != v {
				values = append(values, v)
			}
		}
	}
	return values
}

func genStruct(field *ThingTypeDesc) interface{} {

	value := make(map[string]interface{})
	items, _ := field.Define["items"].([]interface{})
	if nil == items {
		items = field.Define["item"].([]interface{})
	}
	for _, itemI := range items {
		var def map[string]interface{}
		item := itemI.(map[string]interface{})
		typ, _ := item["type"].(string)
		name, _ := item["identifier"].(string)
		define, ok := item["define"]
		if !ok {
			//define = nil
		} else {
			def, _ = define.(map[string]interface{})
		}
		value[name] = genValue(&ThingTypeDesc{
			Type:   typ,
			Define: def,
		})
	}
	return value
}

func genEnum(field *ThingTypeDesc) interface{} {

	var value interface{}
	for _, val := range field.Define {
		value = val
	}
	return &Value{
		Value: value,
		Time:  time.Now().Unix(),
	}
}
