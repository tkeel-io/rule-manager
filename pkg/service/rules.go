package service

import (
	"context"
	"time"

	"git.internal.yunify.com/manage/common/db"
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/rule-manager/internal/dao"
	daorequest "github.com/tkeel-io/rule-manager/internal/dao/utils"
	"github.com/tkeel-io/rule-manager/internal/daoutils"
	erro "github.com/tkeel-io/rule-manager/internal/errors"
	xutils "github.com/tkeel-io/rule-manager/internal/utils"
)

//Log
const (
	ServiceLogRuleCreate = "[RuleCreate]"
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
	//print request [debug]
	printInputDebug(ServiceLogRuleCreate, req)

	//coonstraint total(rule) by user, : queryTotalRule()

	var tx *pg.Tx
	var ruleId string
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	tx, err = db.GetTransaction()
	if nil != err {
		log.ErrorWithFields(ServiceLogRuleCreate, log.Fields{
			"error": err,
		})
		return nil, erro.New(erro.InternalError, err)
	}

	if ruleId, err = xutils.GenerateID(ctx, req.UserId, xutils.RuleIdPrefix); nil != err {
		log.ErrorWithFields(ServiceLogRuleCreate, log.Fields{
			"error": err,
		})
		return nil, erro.New(erro.InternalError, err)
	}
	defer func() {
		err = daorequest.OnTransacation(tx, err, ServiceLogRuleCreate, log.Fields{
			"desc": "rule create successful.",
		})
	}()
	rule := dao.Rule{
		Id:           ruleId,
		UserId:       req.UserId,
		Name:         req.Name,
		Status:       constant.RuleStatusStop,
		RuleDesc:     req.RuleDesc,
		DataType:     uint8(req.DataType),
		SelectText:   req.SelectText,
		SelectFields: make([]*daorequest.SelectField, 0),
		TopicType:    req.TopicType,
		ShortTopic:   req.ShortTopic,
		WhereText:    req.Where,
		RawRequest:   req.Raw,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	}
	for _, field := range req.SelectFields {
		rule.SelectFields = append(rule.SelectFields, &daorequest.SelectField{
			Expr:  field.Expr,
			Type:  field.Type,
			Alias: field.Alias,
		})
	}
	log.Error(xutils.Encode2String(rule.SelectFields))
	//generate select_text.
	rule.SelectText = daoutils.GenerateSelectText(rule.TopicType, rule.SelectFields)

	//*: all fields.
	//generate ruleql
	rule.Ruleql = xutils.GenerateRuleql(&rule)
	//insert table.
	_, err = rule.Insert(ctx, tx, []*dao.Rule{&rule})
	if nil != err {
		return nil, erro.New(erro.InternalError, err)
	}
	return &pb.RuleCreateResp{
		Id: rule.Id,
	}, nil
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
