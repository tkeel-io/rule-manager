package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"git.internal.yunify.com/manage/common/db"
	interLog "git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	"github.com/tkeel-io/rule-manager/config"
	"github.com/tkeel-io/rule-manager/internal/dao"
	"github.com/tkeel-io/rule-manager/pkg/server"
	"github.com/tkeel-io/rule-manager/pkg/service"

	openapi "github.com/tkeel-io/rule-manager/api/openapi/v1"
	rule "github.com/tkeel-io/rule-manager/api/rule/v1"
)

var (
	// Name app.
	Name string
	// HTTPAddr string.
	HTTPAddr string
	// GRPCAddr string.
	GRPCAddr string
)

func main() {
	parseFlags()
	initDB()
	initInternalLog()

	httpSrv := server.NewHTTPServer(HTTPAddr)
	grpcSrv := server.NewGRPCServer(GRPCAddr)
	serverList := []transport.Server{httpSrv, grpcSrv}

	s := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		serverList...,
	)

	{
		OpenapiSrv := service.NewOpenapiService()
		openapi.RegisterOpenapiHTTPServer(httpSrv.Container, OpenapiSrv)
		openapi.RegisterOpenapiServer(grpcSrv.GetServe(), OpenapiSrv)

		RulesSrv := service.NewRulesService()
		rule.RegisterRulesHTTPServer(httpSrv.Container, RulesSrv)
		rule.RegisterRulesServer(grpcSrv.GetServe(), RulesSrv)
	}

	if err := s.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := s.Stop(context.TODO()); err != nil {
		panic(err)
	}
}

func parseFlags() {
	flag.StringVar(&Name, "name", "rule-manager", "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", ":31234", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":31233", "grpc listen address.")
	flag.StringVar(&config.PGHost, "pg_host", "localhost", "postgres host.")
	flag.StringVar(&config.PGPort, "pg_port", "5432", "postgres port.")
	flag.StringVar(&config.PGUser, "pg_user", "username", "postgres user.")
	flag.StringVar(&config.PGPassword, "pg_password", "password", "postgres password.")

	flag.Parse()
}

func initDB() {
	db.InitPG(
		config.PGHost,
		config.PGPort,
		config.PGUser,
		config.PGPassword,
		config.PGDatabase,
	)

	if err := db.CheckTables([]interface{}{
		(*dao.Action)(nil),
		(*dao.Rule)(nil),
	}); err != nil {
		log.Fatal("check database table failed", err)
	}
}

func initInternalLog() {
	interLog.Init(nil)
	interLog.SetLogLevelFromStr("debug")
}
