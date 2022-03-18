package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	interLog "git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/kit/transport"
	"github.com/tkeel-io/kit/transport/grpc"
	transportHTTP "github.com/tkeel-io/kit/transport/http"
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
	servers := []transport.Server{httpSrv, grpcSrv}

	s := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		servers...,
	)

	register(httpSrv, grpcSrv)

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

func register(httpSrv *transportHTTP.Server, grpcSrv *grpc.Server) {
	{
		OpenapiSrv := service.NewOpenapiService()
		openapi.RegisterOpenapiHTTPServer(httpSrv.Container, OpenapiSrv)
		openapi.RegisterOpenapiServer(grpcSrv.GetServe(), OpenapiSrv)

		RulesSrv := service.NewRulesService()
		rule.RegisterRulesHTTPServer(httpSrv.Container, RulesSrv)
		rule.RegisterRulesServer(grpcSrv.GetServe(), RulesSrv)
	}
}

func parseFlags() {
	flag.StringVar(&Name, "name", "rule-manager", "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", ":31234", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":31233", "grpc listen address.")

	flag.StringVar(&config.DSN, "dsn", "root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local", "database dsn")

	flag.Parse()
}

func initDB() {
	if err := dao.Setup(); err != nil {
		log.Fatal("setup database failed", err)
	}
}

func initInternalLog() {
	interLog.Init(nil)
	interLog.SetLogLevelFromStr("debug")
}
