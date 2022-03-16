package types

import pb "git.internal.yunify.com/MDMP2/api/metadata/v1"

type GrpcClient interface {
	Rule() pb.RuleActionClient
	RuleNode() pb.RulexNodeActionClient
}
