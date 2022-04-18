package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tkeel-io/rule-manager/internal/action/clickhouse/grammar"
	"github.com/tkeel-io/rule-util/pkg/log"

	plugin "github.com/tkeel-io/rule-rulex/pkg/sink/plugin/clickhouse"
	logf "github.com/tkeel-io/rule-util/pkg/logfield"
)

//desc device_event_data;

var DriverChronus = "chronus"

type Option struct {
	//Addrs     string `json:"addrs,omitempty"`
	Urls      []string `json:"urls"`
	DbName    string   `json:"dbName,omitempty"`
	Table     string   `json:"table,omitempty"`
	TimeField string   `json:"timestamp_field,omitempty"`
}

//clickhouse plugin
type clickhouse struct {
	entityType, entityID string
	balance              plugin.LoadBalance
	//db     *sqlx.DB
	option *Option
}

//func Connect(ctx context.Context, Endpoints []string) (err error) {
//	servers := make([]*plugin.Server, len(Endpoints))
//	for k, v := range Endpoints {
//		log.Info("clickhouse init " + v)
//		db, err := sqlx.Open("clickhouse", v)
//		if err != nil {
//			log.Error("open clickhouse", logf.Any("error", err))
//			return err
//		}
//		if err = db.PingContext(ctx); err != nil {
//			log.Error("ping clickhouse", logf.Any("error", err))
//			return err
//		}
//		db.SetConnMaxLifetime(30 * time.Second)
//		db.SetMaxOpenConns(5)
//		servers[k] = &plugin.Server{db, v, 1}
//	}
//
//	balance := plugin.NewLoadBalanceRandom(servers)
//
//	server := balance.Select([]*sqlx.DB{})
//
//	preURL := "show tables;"
//	rows, err := server.DB.Query(preURL)
//	if err != nil {
//		log.GlobalLogger().Bg().Error("tx Commit error",
//			logf.String("error", err.Error()))
//		return err
//	}
//	defer rows.Close()
//	var value string
//	if rows.Next() {
//		rows.Scan(&value)
//	}
//	fmt.Println(value)
//
//	return nil
//}

func (this *Transport) ExceSQL(ctx context.Context, Endpoints []string, SQL string) error {
	servers := make([]*plugin.Server, len(Endpoints))
	for k, v := range Endpoints {
		log.Info("clickhouse init ", logf.Any("endpoint", v))
		db, err := sqlx.Open(DriverChronus, v)
		if err != nil {
			log.Error("open db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return err
		}
		if err = db.PingContext(ctx); err != nil {
			log.Error("ping db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return err
		}
		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)
		servers[k] = &plugin.Server{db, v, 1}
	}

	balance := plugin.NewLoadBalanceRandom(servers)
	defer balance.Close()

	server := balance.Select([]*sqlx.DB{})

	preURL := fmt.Sprintf("%s;", SQL)
	_, err := server.DB.Exec(preURL)
	if err != nil {
		log.Error("tx Commit error", logf.Error(err))
		return err
	}
	return nil
}

func (this *Transport) TableInfo(ctx context.Context, Endpoints []string, tableName string) (elem *grammar.ClickHouseListener, err error) {
	servers := make([]*plugin.Server, len(Endpoints))
	for k, v := range Endpoints {
		log.Info("clickhouse init ", logf.Any("endpoint", v))
		db, err := sqlx.Open(DriverChronus, v)
		if err != nil {
			log.Error("open db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return nil, err
		}
		if err = db.PingContext(ctx); err != nil {
			log.Error("ping db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return nil, err
		}
		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)
		servers[k] = &plugin.Server{db, v, 1}
	}

	balance := plugin.NewLoadBalanceRandom(servers)
	defer balance.Close()

	server := balance.Select([]*sqlx.DB{})

	preURL := fmt.Sprintf("show create table %s;", tableName)
	rows, err := server.DB.Query(preURL)
	if err != nil {
		log.Error("tx Commit error", logf.Error(err))
		return nil, err
	}
	defer rows.Close()
	var value string
	if rows.Next() {
		rows.Scan(&value)
	}

	//fmt.Println(value)
	elem = grammar.Parse(value)

	return elem, nil
}

func (this *Transport) TableList(ctx context.Context, Endpoints []string) (elem []string, err error) {
	log.Info("open db",
		logf.String("driver", DriverChronus),
		logf.Any("endpoints", Endpoints),
		logf.Any("error", err))

	if len(Endpoints) <= 0 {
		return nil, errors.New("endpoint can not be empty")
	}

	servers := make([]*plugin.Server, len(Endpoints))
	for k, v := range Endpoints {
		log.Info("clickhouse init ", logf.Any("endpoint", v))
		db, err := sqlx.Open(DriverChronus, v)
		if err != nil {
			log.Error("open db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return nil, err
		}
		if err = db.PingContext(ctx); err != nil {
			log.Error("ping db",
				logf.String("driver", DriverChronus),
				logf.Any("error", err))
			return nil, err
		}
		db.SetConnMaxLifetime(30 * time.Second)
		db.SetMaxOpenConns(5)
		servers[k] = &plugin.Server{db, v, 1}
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
		fmt.Println("-", value)
		tables = append(tables, value)
	}

	return tables, nil
}
