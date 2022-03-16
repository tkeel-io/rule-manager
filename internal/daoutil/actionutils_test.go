package daoutil

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapCatRecursive(t *testing.T) {

	m1 := `
		{
			"name": 123,
			"maps":{
				"id":"id",
				"mapping":{
					"flag":"false"
				}
			}
		}
	`
	m2 := `
		{
			"name": "456",
			"maps":{
				"name":"name",
				"mapping":{
					"flag":false,
					"yun":"yun"
				}
			},
			"user":"dong"
		}
	`
	var mm1, mm2 = make(map[string]interface{}), make(map[string]interface{})
	json.Unmarshal([]byte(m1), &mm1)
	json.Unmarshal([]byte(m2), &mm2)
	m := MapCatRecursive(mm1, mm2)
	fmt.Println(m)
}
