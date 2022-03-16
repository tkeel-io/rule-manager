package service

import (
	"context"
	"github.com/tkeel-io/rule-manager/internal/daoutil"
	"time"

	"git.internal.yunify.com/manage/common/db"
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/rule-manager/internal/dao"
	"github.com/tkeel-io/rule-manager/internal/utils"
)

// Log prefix
const (
	CreatePrefixTag      = "[RuleCreate]"
	ServiceLogRuleUpdate = "[RuleUpdate]"
	ServiceLogRuleDelete = "[RuleDelete]"
	ServiceLogRuleQuery  = "[RuleQuery]"
	ServiceLogRuleStatus = "[RuleStatus]"
	ServiceLogRuleDebug  = "[RuleDebug]"
	ServiceLogRuleError  = "[RuleError]"
)

type RulesService struct {
	pb.UnimplementedRulesServer
}

func NewRulesService() *RulesService {
	return &RulesService{}
}

func (s *RulesService) RuleCreate(ctx context.Context, req *pb.RuleCreateReq) (res *pb.RuleCreateResp, err error) {
	printInputDebug(CreatePrefixTag, req)

	//coonstraint total(rule) by user, : queryTotalRule()
	var tx *pg.Tx
	var ruleID string
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	tx, err = db.GetTransaction()
	if nil != err {
		log.ErrorWithFields(CreatePrefixTag, log.Fields{
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}

	if ruleID, err = utils.GenerateID(ctx, req.UserId, utils.RuleID); nil != err {
		log.ErrorWithFields(CreatePrefixTag, log.Fields{
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}
	rule := dao.Rule{
		ID:           ruleID,
		UserID:       req.UserId,
		Name:         req.Name,
		Status:       constant.RuleStatusStop,
		RuleDesc:     req.RuleDesc,
		DataType:     uint8(req.DataType),
		SelectText:   req.SelectText,
		SelectFields: make([]*dao.SelectField, 0),
		TopicType:    req.TopicType,
		ShortTopic:   req.ShortTopic,
		WhereText:    req.Where,
		RawRequest:   req.Raw,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	}
	for _, field := range req.SelectFields {
		rule.SelectFields = append(rule.SelectFields, &dao.SelectField{
			Expr:  field.Expr,
			Type:  field.Type,
			Alias: field.Alias,
		})
	}
	log.Debugf(utils.Encode2String(rule.SelectFields))
	rule.SelectText = daoutil.GenerateSelectText(rule.TopicType, rule.SelectFields)

	rule.Ruleql = utils.GenerateRuleql(&rule)
	//insert table.
	//generate ruleql
	//*: all fields.
	//generate select_text.
	if _, err = rule.Insert(ctx, tx, &rule); nil != err {
		defer func() {
			err = daoutil.CommitTransaction(tx, err, CreatePrefixTag, log.Fields{
				"desc": "rule create successful.",
			})
		}()
		return nil, pb.ErrInternalError()
	}
	return &pb.RuleCreateResp{
		Id: rule.ID,
	}, nil
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
