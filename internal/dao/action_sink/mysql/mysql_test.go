package mysql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/tkeel-io/rule-manager/internal/utils"
	"github.com/jmoiron/sqlx"
)

var ctx context.Context = context.Background()
var endpoints []string = []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai"}

func TestConnect(t *testing.T) {
	end := []string{"139.198.156.131:3306"}
	end = utils.GenerateUrlMysql(end, "root", "lpt316", "test")
	err := Connect(ctx, end, "test")
	if err != nil {
		t.Fatalf("connect to db: %s failed, error: %s\n", "test", err)
	}
}

func TestConnect2(t *testing.T) {
	err := Connect(ctx, endpoints, "test")
	if err != nil {
		t.Fatalf("connect to db: %s failed, error: %s\n", "test", err)
	}
}

func TestListTable(t *testing.T) {
	db := "test"
	tabs, err := ListTable(ctx, endpoints, db)
	if err != nil {
		t.Fatalf("list table failed, db: %s, error: %s \n", db, err)
	}
	fmt.Printf("tables:%v\n", tabs)
}

func TestGenSql(t *testing.T) {
	conf := &TableConfigs{
		Name:      "tb1",
		Database:  "test",
		Endpoints: []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"},
		Fields: []TableField{
			TableField{Name: "id", Type: "BigInt", IsPK: true},
			TableField{Name: "name", Type: "Varchar(20)", IsPK: true},
		},
	}
	sql, err := genSql(conf, false)
	if err != nil {
		t.Fatalf("gensql failed, error: %s\n", err)
	}
	fmt.Println(sql)
}

func TestCreateDb(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		sql, err := CreateTable(ctx, &TableConfigs{
			Name:      "tb2",
			Database:  "test",
			Endpoints: []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"},
			Fields: []TableField{
				TableField{Name: "id", Type: "BigInt", IsPK: true},
				TableField{Name: "name", Type: "Varchar(20)", IsPK: false},
			},
		})
		if err == nil {
			t.Fatalf("expect create db error, but got nil")
		}
		fmt.Println("sql", sql)
	})

	t.Run("t2", func(t *testing.T) {
		sql, err := CreateTable(ctx, &TableConfigs{
			Name:      "tb4",
			Database:  "test",
			Endpoints: []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"},
			Fields: []TableField{
				TableField{Name: "id", Type: "BigInt", IsPK: true},
				TableField{Name: "name", Type: "Varchar(20)", IsPK: false},
			},
		})
		if err != nil {
			t.Fatalf("expect create db succeed, bug got error: %s\n", err)
		}
		fmt.Println("sql", sql)
	})

	t.Run("t3", func(t *testing.T) {
		sql, err := CreateTable(ctx, &TableConfigs{
			Name:      "tb",
			Database:  "test5",
			Endpoints: []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"},
			Fields: []TableField{
				TableField{Name: "id", Type: "BigInt", IsPK: true},
				TableField{Name: "name", Type: "Varchar(20)", IsPK: false},
			},
		})
		if err == nil {
			t.Fatalf("expect create db error, but got nil")
		}
		fmt.Println("sql", sql)
	})
}

func TestGenerateSql(t *testing.T) {
	conf := &TableConfigs{
		Name:      "tb1",
		Database:  "test2",
		Endpoints: []string{"root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local"},
		Fields: []TableField{
			TableField{Name: "id", Type: "BigInt", IsPK: true},
			TableField{Name: "name", Type: "Varchar", IsPK: false},
		},
	}
	sql, err := GenerateSql(ctx, conf)
	if err != nil {
		t.Fatalf("generate sql failed, error: %s \n", err)
	}
	fmt.Println("sql", sql)
}

func TestGetTableFields(t *testing.T) {
	tableFields := GetTableFieldTypes()
	if len(tableFields) == 0 {
		t.Fatalf("get table fields failed")
	}
	fmt.Println("tableFields:", tableFields)
}

func TestMysqlDataTypeTinyInt(t *testing.T) {
	sql := "insert into val1 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	_, err = stmt.Exec(int64(7))
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypeFloat(t *testing.T) {
	sql := "insert into val2 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	_, err = stmt.Exec(float32(3.14))
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypChar(t *testing.T) {
	sql := "insert into val3 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	_, err = stmt.Exec("abc")
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypDate(t *testing.T) {
	sql := "insert into val4 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	d := time.Now()
	//"2021-12-22"
	_, err = stmt.Exec(d)
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypTime(t *testing.T) {
	sql := "insert into val5 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	//"15:00:00"
	d := time.Now()
	_, err = stmt.Exec(d)
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypDateTime(t *testing.T) {
	sql := "insert into val6 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	d := "2021-12-22 20:00:00"
	// d := time.Now()
	// d, err := time.Parse("2006-01-02 15:04:05", "2021-12-22 13:00:00")
	if err != nil {
		t.Fatalf("parse time failed, error: %s \n", err)
	}
	_, err = stmt.Exec(d)
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypTimestamp(t *testing.T) {
	sql := "insert into val7 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	// "2021-12-22 15:00:00"
	// d := time.Now()
	// d, err := time.Parse("2006-01-02 15:04:05", "2021-12-22 13:00:00")
	d, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-22 13:00:00", time.Local)
	fmt.Printf("d: %v\n", d)
	if err != nil {
		t.Fatalf("parse time failed, error: %s \n", err)
	}
	_, err = stmt.Exec(d)
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypMediumInt(t *testing.T) {
	sql := "insert into val8 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	_, err = stmt.Exec(int64(7))
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestMysqlDataTypBigInt(t *testing.T) {
	sql := "insert into val9 values(?)"
	db, err := sqlx.Open("mysql", "root:lpt316@tcp(139.198.156.131:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("open mysql failed, error: %s \n", err)
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatalf("prepare sql failed, error: %s \n", err)
	}
	_, err = stmt.Exec(int64(7))
	if err != nil {
		t.Fatalf("exec sql failed, error: %s \n", err)
	}
}

func TestTableInfo(t *testing.T) {
	tab, err := TableInfo(ctx, endpoints, "create_test1")
	if err != nil {
		t.Fatalf("get table info failed, error: %s\n", err)
	}
	fmt.Printf("table: %v", tab)
}
