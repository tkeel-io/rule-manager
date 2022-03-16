package service

import (
	"context"

	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
)

type RulesService struct {
	pb.UnimplementedRulesServer
}

func NewRulesService() *RulesService {
	return &RulesService{}
}

func (s *RulesService) RuleCreate(ctx context.Context, req *pb.RuleCreateReq) (*pb.RuleCreateResp, error) {
	return &pb.RuleCreateResp{}, nil
}
func (s *RulesService) RuleUpdate(ctx context.Context, req *pb.RuleUpdateReq) (*pb.RuleUpdateResp, error) {
	return &pb.RuleUpdateResp{}, nil
}
func (s *RulesService) RuleDelete(ctx context.Context, req *pb.RuleDeleteReq) (*pb.RuleDeleteResp, error) {
	return &pb.RuleDeleteResp{}, nil
}
func (s *RulesService) RuleGet(ctx context.Context, req *pb.RuleGetReq) (*pb.Rule, error) {
	return &pb.Rule{}, nil
}
func (s *RulesService) RuleQuery(ctx context.Context, req *pb.RuleQueryReq) (*pb.RuleQueryResp, error) {
	return &pb.RuleQueryResp{}, nil
}
func (s *RulesService) RuleStatus(ctx context.Context, req *pb.RuleStatusReq) (*pb.RuleStatusResp, error) {
	return &pb.RuleStatusResp{}, nil
}
func (s *RulesService) RuleStart(ctx context.Context, req *pb.RuleStartReq) (*pb.RuleStartResp, error) {
	return &pb.RuleStartResp{}, nil
}
func (s *RulesService) RuleStop(ctx context.Context, req *pb.RuleStopReq) (*pb.RuleStopResp, error) {
	return &pb.RuleStopResp{}, nil
}
func (s *RulesService) RuleDebug(ctx context.Context, req *pb.RuleDebugReq) (*pb.RuleDebugResp, error) {
	return &pb.RuleDebugResp{}, nil
}
func (s *RulesService) RuleDebugMessage(ctx context.Context, req *pb.RuleDebugMsgReq) (*pb.RuleDebugMsgResp, error) {
	return &pb.RuleDebugMsgResp{}, nil
}
func (s *RulesService) RuleError(ctx context.Context, req *pb.RuleErrorReq) (*pb.RuleErrorResp, error) {
	return &pb.RuleErrorResp{}, nil
}
