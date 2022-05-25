package xgrpc

import (
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/tkeel-io/rule-manager/config"
	_ "github.com/tkeel-io/rule-manager/internal/types"
	"github.com/tkeel-io/rule-util/metadata"
	pb "github.com/tkeel-io/rule-util/metadata/v1"
	"github.com/tkeel-io/rule-util/pkg/registry"
	etcdv3 "github.com/tkeel-io/rule-util/pkg/registry/etcd3"
	"github.com/tkeel-io/rule-util/rulex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var opts = []grpc.DialOption{
	grpc.WithBalancerName("round_robin"),
	grpc.WithInsecure(),
	grpc.WithUnaryInterceptor(
		grpc_opentracing.UnaryClientInterceptor(),
	),
	// grpc.WithTimeout(timeout),
	grpc.WithBackoffMaxDelay(time.Second * 3),
	grpc.WithInitialWindowSize(1 << 30),
	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(64 * 1024 * 1024)),
}

type Client struct {
	rulex pb.RulexNodeActionClient
	rule  pb.RuleActionClient
}

func NewClient() *Client {
	discovery, err0 := etcdv3.New(&registry.Config{
		Endpoints: config.GetConfig().Etcd.Address,
	})
	if err0 != nil {
		panic(err0)
	}

	return &Client{
		rulex: newRulexNodeActionClient(discovery),
		rule:  newRuleActionClient(discovery),
	}
}

func newRuleActionClient(discovery *etcdv3.Discovery) pb.RuleActionClient {
	// timeout := time.Duration(config.RPCClient.Timeout)
	resolverBuilder := discovery.GRPCResolver()
	resolver.Register(resolverBuilder)
	conn, err := grpc.Dial(
		resolverBuilder.Scheme()+":///"+metadata.APPID,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	return pb.NewRuleActionClient(conn)
}

func newRulexNodeActionClient(discovery *etcdv3.Discovery) pb.RulexNodeActionClient {
	// timeout := time.Duration(config.RPCClient.Timeout)
	resolverBuilder := discovery.GRPCResolver()
	resolver.Register(resolverBuilder)
	conn, err1 := grpc.Dial(
		resolverBuilder.Scheme()+":///"+rulex.APPID,
		opts...,
	)
	if err1 != nil {
		panic(err1)
	}
	return pb.NewRulexNodeActionClient(conn)
}

func (c *Client) Rulex() pb.RulexNodeActionClient {
	return c.rulex
}

func (c *Client) Rule() pb.RuleActionClient {
	return c.rule
}
