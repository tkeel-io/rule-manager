package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestTableList(t *testing.T) {
	trans := NewTransport()
	ctx := context.Background()
	endpoints := []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"}
	tables, err := trans.TableList(ctx, endpoints)
	if err != nil {
		t.Fatalf("mysql table list fail, error: %s \n", err)
	}
	fmt.Printf("mysql table list succeed, tbales: %#v \n", tables)
	t.Logf("mysql table list succeed, tbales: %#v \n", tables)
}

func TestTableinfo(t *testing.T) {
	trans := NewTransport()
	ctx := context.Background()
	endpoints := []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"}
	tableDef, err := trans.TableInfo(ctx, endpoints, "temp")

	if err != nil {
		t.Fatalf("get mysql table info failed, error: %s \n", err)
	}
	fmt.Println("tableDef:", tableDef)
}
