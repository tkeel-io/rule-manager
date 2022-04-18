package grammar

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	//SQL := "CREATE TABLE iot_manage_dev.device_event_data (`device_id` String, `source_id` String, `event_id` String, `identifier` String, `type` String, `user_id` String, `thing_id` String, `metadata` String, `time` UInt64, `action_date` Date, `action_time` DateTime) ENGINE = MergeTree() PARTITION BY len(action_date+action_date2)  ORDER BY (user_id, device_id, event_id) TTL action_date + toIntervalDay(10) SETTINGS index_granularity = 8192"
	SQL := "CREATE TABLE iot_manage_dev.device_thing_data_distributed (`time` UInt64, `user_id` String, `device_id` String, `source_id` String, `thing_id` String, `identifier` String, `value_int32` Int32, `value_float` Float32, `value_double` Float64, `value_string` String, `value_enum` Enum8 ('0' = 0, '1' = 1, '2' = 2, '3' = 3, '4' = 4, '5' = 5, '6' = 6, '7' = 7, '8' = 8, '9' = 9, '10' = 10, '11' = 11, '12' = 12, '13' = 13, '14' = 14, '15' = 15, '16' = 16, '17' = 17, '18' = 18, '19' = 19, '20' = 20, '21' = 21, '22' = 22, '23' = 23, '24' = 24, '25' = 25, '26' = 26, '27' = 27, '28' = 28, '29' = 29, '30' = 30, '31' = 31, '32' = 32, '33' = 33, '34' = 34, '35' = 35, '36' = 36, '37' = 37, '38' = 38, '39' = 39, '40' = 40, '41' = 41, '42' = 42, '43' = 43, '44' = 44, '45' = 45, '46' = 46, '47' = 47, '48' = 48, '49' = 49, '50' = 50, '51' = 51, '53' = 53, '54' = 54, '55' = 55, '56' = 56, '57' = 57, '58' = 58, '59' = 59, '60' = 60, '61' = 61, '62' = 62, '63' = 63, '64' = 64, '65' = 65, '66' = 66, '67' = 67, '68' = 68, '69' = 69, '70' = 70, '71' = 71, '72' = 72, '73' = 73, '74' = 74, '75' = 75, '76' = 76, '77' = 77, '78' = 78, '79' = 79, '80' = 80, '81' = 81, '82' = 82, '83' = 83, '84' = 84, '85' = 85, '86' = 86, '87' = 87, '88' = 88, '89' = 89, '90' = 90, '91' = 91, '92' = 92, '93' = 93, '94' = 94, '95' = 95, '96' = 96, '97' = 97, '98' = 98, '99' = 99, '100' = 100), `value_bool` UInt8, `value_string_ex` String, `value_array_string` Array(String), `value_array_int32` Array(Int32), `value_array_float` Array(Float32), `value_array_double` Array(Float64), `action_date` Date, `action_time` DateTime, `tags` Nullable(String)) ENGINE = Distributed(logical_consistency_cluster, iot_manage_dev, device_thing_data, sipHash64(user_id))"
	//SQL = `show tables`
	epx := Parse(SQL)

	fmt.Printf("%v\n", epx)
	print("Fields", epx.Fields)
	//print("PartitionFields", epx.PartitionFields)
	//print("OrderByFields", epx.OrderByFields)
}

func print(name string, fields []*Field) {
	fmt.Println(name, ":")
	for idx, e := range fields {
		fmt.Println(idx, e)
	}

}
