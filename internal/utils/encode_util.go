package utils

import "encoding/json"

func Encode2String(v interface{}) string {
	buf, err := json.Marshal(v)
	if nil != err {
		return err.Error()
	}
	return string(buf)
}
