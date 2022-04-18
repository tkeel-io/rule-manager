package clickhouse

import (
	"context"

	log "github.com/tkeel-io/rule-util/pkg/commonlog"
)

type Transport struct {
}

var default_user = "default"

//url, user, password, db
var formatUrl = "%s?username=%s&password=%s&database=%s"

/*
	对于数据库功能封装：
		1. 测试连接
		2. 获取db的table列表
		3. 获取table的详细信息：fields，configs
*/

func NewTransport() *Transport {
	return &Transport{}
}

func GetTransport() *Transport {
	return NewTransport()
}

func (this *Transport) CreateTable(ctx context.Context, endpoints []string, sql string) error {
	//dbname, table, fields, engine, partition, orders, ttl, index
	log.InfoWithFields("[<chronus>table create] ", log.Fields{
		"endpoints": endpoints,
		"sql":       sql,
	})
	return this.ExceSQL(ctx, endpoints, sql)
}
