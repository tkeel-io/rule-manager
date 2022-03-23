package utils

import (
	"fmt"
	"testing"
)

func DC_Test(t *testing.T) {
	a := map[string]interface{}{
		"a": 1,
		"b": struct {
			Name string
		}{
			Name: "john",
		},
	}
	var b map[string]interface{}
	DeepCopy(a, b)
	fmt.Println(b)
}

func TestGenerateUrlMysql(t *testing.T) {
	endpoints := []string{"139.198.156.131:3306"}
	user := "root"
	pass := "lpt316"
	db := "test"
	urls := GenerateUrlMysql(endpoints, user, pass, db)
	fmt.Println("urls:", urls)
}
