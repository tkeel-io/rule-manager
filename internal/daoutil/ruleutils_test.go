package daoutils

import (
	"fmt"
	"testing"

	daorequest "github.com/tkeel-io/rule-manager/internal/dao/utils"
)

func TestGenText(t *testing.T) {

	fields := []*daorequest.SelectField{

		&daorequest.SelectField{
			Expr: "*",
		},
		&daorequest.SelectField{
			Expr: "deviceId()",
		},
		&daorequest.SelectField{
			Expr:  "upTime()",
			Alias: "upT",
		},
		&daorequest.SelectField{
			Expr:  "user",
			Alias: "user",
		},
		&daorequest.SelectField{
			Expr: "value",
		},
	}

	text := GenerateSelectText(fields)
	fmt.Println(text)

}
