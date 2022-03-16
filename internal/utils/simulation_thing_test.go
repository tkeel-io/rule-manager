package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestThing(t *testing.T) {

	fields := map[string]*ThingTypeDesc{
		"appControlStatus": &ThingTypeDesc{
			Type: "string",
			Define: map[string]interface{}{
				"max": "1000",
			},
		},
	}

	message := GenMessage(fields)

	buf, _ := json.MarshalIndent(message, "	", "	")

	fmt.Printf("%v\n", string(buf))

}

func TestThingStruct(t *testing.T) {

	fields := map[string]*ThingTypeDesc{
		"appControlStatus": &ThingTypeDesc{
			Type: "struct",
			Define: map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"identifier": "name",
						"name":       "网卡名称",
						"type":       "string",
					},
					{
						"identifier": "netIn",
						"name":       "网卡接收流量",
						"type":       "int32",
					},
					{
						"identifier": "netOut",
						"name":       "网卡发送流量",
						"type":       "int32",
					},
				},
			},
		},
	}

	message := GenMessage(fields)

	buf, _ := json.MarshalIndent(message, "	", "	")

	fmt.Printf("%v\n", string(buf))

}

func TestThingArray(t *testing.T) {

	fields := map[string]*ThingTypeDesc{
		"appControlStatus": &ThingTypeDesc{
			Type: "array",
			Define: map[string]interface{}{
				"item": map[string]interface{}{
					"define": map[string]interface{}{
						"items": []map[string]interface{}{
							{
								"identifier": "name",
								"name":       "网卡名称",
								"type":       "string",
							},
							{
								"identifier": "netIn",
								"name":       "网卡接收流量",
								"type":       "int32",
							},
							{
								"identifier": "netOut",
								"name":       "网卡发送流量",
								"type":       "int32",
							},
						},
					},
					"type": "struct",
				},
				"size": "2",
			},
		},
	}

	message := GenMessage(fields)

	buf, _ := json.MarshalIndent(message, "	", "	")

	fmt.Printf("%v\n", string(buf))

}
