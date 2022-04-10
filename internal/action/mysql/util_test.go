package mysql

import (
	"fmt"
	"testing"
)

func TestGetDB(t *testing.T) {
	endpoint := "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, err := getDbFromEndpoint(endpoint)
	if err != nil {
		t.Fatalf("findDbFromEndpoint failed, error: %s \n", err)

	}
	if db != "test" {
		t.Fatalf("expect db test, but get %s \n", db)
	}
	fmt.Println("db:", db)
}
