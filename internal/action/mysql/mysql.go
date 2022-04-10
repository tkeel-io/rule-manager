package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tkeel-io/rule-util/pkg/log"
	logf "github.com/tkeel-io/rule-util/pkg/logfield"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	plugin "github.com/tkeel-io/rule-rulex/pkg/sink/plugin/mysql"
)

const DefaultUser = "default"
const DriverMysql = "mysql"

var notImpl = errors.New("method not implement")

type Transport struct{}

type Field struct {
	Name  string
	Type  string
	Is_PK bool
}

type Table struct {
	Name   string
	Fields []Field
}

func NewTransport() *Transport {
	return &Transport{}
}

func (t *Transport) CreateTable(ctx context.Context, endpoints []string, sql string) error {
	log.Info("create mysql table", logf.Any("endpoints", endpoints), logf.String("sql", sql))
	return t.ExceSQL(ctx, endpoints, sql)
}

func (t *Transport) ExceSQL(ctx context.Context, endpoints []string, sql string) error {
	if len(endpoints) == 0 {
		return errors.New("endpoints can not be e,pty")
	}
	servers := make([]*plugin.Server, len(endpoints))
	for idx, endpoint := range endpoints {
		log.Info("mysql init", logf.Any("endpoint", endpoint))
		db, err := sqlx.Open(DriverMysql, endpoint)
		if err != nil {
			log.Error("open mysql db failed", logf.Any("error", err))
			return err
		}
		if err := db.PingContext(ctx); err != nil {
			log.Error("ping mysql failed", logf.Any("error", err))
			return err
		}
		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)
		servers[idx] = &plugin.Server{
			Name:   endpoint,
			DB:     db,
			Weight: 1,
		}
	}

	balance := plugin.NewLoadBalanceRandom(servers)
	defer balance.Close()

	server := balance.Select([]*sqlx.DB{})
	_, err := server.DB.Exec(sql)
	if err != nil {
		log.Error("tx Commit error", logf.Error(err))
		return err
	}
	return nil
}

func (t *Transport) TableList(ctx context.Context, endpoints []string) (elem []string, err error) {
	log.Info("open db",
		logf.String("driver", DriverMysql),
		logf.Any("endpoints", endpoints),
		logf.Any("error", err))

	if len(endpoints) <= 0 {
		return nil, errors.New("endpoints can not be empty")
	}

	servers := make([]*plugin.Server, len(endpoints))

	for idx, endpoint := range endpoints {
		log.Info("mysql init ", logf.String("endpoint", endpoint))
		db, err := sqlx.Open(DriverMysql, endpoint)
		if err != nil {
			log.Error("open mysql error", logf.Any("error", err))
			return nil, err
		}

		if err := db.PingContext(ctx); err != nil {
			log.Error("ping mysql error", logf.Any("error", err))
			return nil, err
		}

		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)

		servers[idx] = &plugin.Server{
			DB:     db,
			Name:   endpoint,
			Weight: 1,
		}
	}

	balance := plugin.NewLoadBalanceRandom(servers)
	defer balance.Close()

	server := balance.Select([]*sqlx.DB{})

	preURL := "show tables;"
	rows, err := server.DB.Query(preURL)
	if err != nil {
		log.Error("tx Commit error", logf.Error(err))
		return nil, err
	}
	defer rows.Close()

	var tables = make([]string, 0)
	for rows.Next() {
		var value string
		rows.Scan(&value)
		tables = append(tables, value)
	}

	log.Debug("list mysql tables", logf.Any("tables", tables))
	return tables, nil
}

func (t *Transport) TableInfo(ctx context.Context, endpoints []string, tableName string) (*Table, error) {
	servers := make([]*plugin.Server, len(endpoints))
	for idx, endpoint := range endpoints {
		log.Info("mysql init", logf.Any("endpoint", endpoint))
		db, err := sqlx.Open(DriverMysql, endpoint)
		if err != nil {
			log.Error("open db", logf.String("driver", DriverMysql), logf.Any("error", err))
			return nil, err
		}

		if err = db.PingContext(ctx); err != nil {
			log.Error("ping mysql error", logf.Any("error", err))
			return nil, err
		}
		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)
		servers[idx] = &plugin.Server{
			DB:     db,
			Name:   endpoint,
			Weight: 1,
		}
	}
	balance := plugin.NewLoadBalanceRandom(servers)
	defer balance.Close()

	server := balance.Select([]*sqlx.DB{})

	// sql := fmt.Sprintf("show create table %s;", tableName)
	// rows, err := server.DB.QueryContext(ctx, sql)
	// if err != nil {
	// 	log.Error("tx commit error", logf.Error(err))
	// 	return nil, err
	// }
	// defer rows.Close()
	// var value string
	// if rows.Next() {
	// 	rows.Scan(&tableName, &value)
	// } else {
	// 	return nil, errors.New("no result found")
	// }
	// value = strings.ReplaceAll(value, "\n", "")
	// value = strings.ReplaceAll(value, "\t", " ")
	// fmt.Println("value：", value)
	// elem := grammar.Parse(value)
	// fmt.Printf("elem: %#v \n", elem)
	// return elem, nil

	db, err := getDbFromEndpoint(server.Name)
	if err != nil {
		log.Error("get db from endpoint error", logf.String("call", "TableInfo"), logf.Any("error", err))
		return nil, err
	}
	fields, err := getTableFields(ctx, server, db, tableName)
	if err != nil {
		log.Error("get table fields error", logf.String("db", db), logf.String("table", tableName), logf.Any("error", err))
		return nil, err
	}

	tab := &Table{
		Name:   tableName,
		Fields: make([]Field, 0),
	}
	for _, field := range fields {
		tab.Fields = append(tab.Fields, Field{
			Name:  field.Name,
			Type:  field.Type,
			Is_PK: field.Is_PK,
		})
	}
	return tab, nil
}

func getTableFields(ctx context.Context, server *plugin.Server, db, tableName string) ([]Field, error) {
	sql := "select column_name, data_type, column_key from information_schema.COLUMNS where table_schema = '%s' and table_name = '%s';"
	sql = fmt.Sprintf(sql, db, tableName)
	rows, err := server.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	var colName, dataType, colKey string
	fields := make([]Field, 0)
	for rows.Next() {
		rows.Scan(&colName, &dataType, &colKey)
		isPK := false
		if colKey == "PRI" {
			isPK = true
		}
		fields = append(fields, Field{
			Name:  colName,
			Type:  dataType,
			Is_PK: isPK,
		})
	}
	return fields, nil
}
