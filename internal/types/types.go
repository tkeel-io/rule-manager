package types

import pb "github.com/tkeel-io/rule-util/metadata/v1"

type GrpcClient interface {
	Rule() pb.RuleActionClient
	RuleNode() pb.RulexNodeActionClient
}
