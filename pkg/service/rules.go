package service

import (
	"context"
	"time"

	metapb "git.internal.yunify.com/MDMP2/api/metadata/v1"
	"git.internal.yunify.com/manage/common/db"
	"git.internal.yunify.com/manage/common/log"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	tkeelLog "github.com/tkeel-io/kit/log"
	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/rule-manager/internal/dao"
	"github.com/tkeel-io/rule-manager/internal/daoutil"
	"github.com/tkeel-io/rule-manager/internal/event"
	"github.com/tkeel-io/rule-manager/internal/utils"
)

// Log prefix
const (
	CreatePrefixTag      = "[RuleCreate]"
	UpdatePrefixTag      = "[RuleUpdate]"
	DeletePrefixTag      = "[RuleDelete]"
	ServiceLogRuleQuery  = "[RuleQuery]"
	ServiceLogRuleStatus = "[RuleStatus]"
	DebugPrefixTag       = "[RuleDebug]"
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
	defer func() {
		err = daoutil.CommitTransaction(tx, err, CreatePrefixTag, log.Fields{
			"desc": "rule create successful.",
		})
	}()

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
		return nil, pb.ErrInternalError()
	}
	return &pb.RuleCreateResp{
		Id: rule.ID,
	}, nil
}
func (s *RulesService) RuleUpdate(ctx context.Context, req *pb.RuleUpdateReq) (*pb.RuleUpdateResp, error) {
	//print request [debug]
	printInputDebug(UpdatePrefixTag, req)

	var rule = &dao.Rule{
		ID:     req.Id,
		UserID: req.UserId,
	}
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	tx, err := db.GetTransaction()
	if nil != err {
		log.InfoWithFields(UpdatePrefixTag, log.Fields{
			"desc":  "create trasaction failed.",
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}
	defer func() {
		err = daoutil.CommitTransaction(tx, err, UpdatePrefixTag, log.Fields{
			"desc":    "update rule successful.",
			"rule_id": rule.ID,
		})
	}()

	// update的transaction尚未commit，所以需要手动更新rule
	if err = rule.Select(ctx); nil != err {
		return nil, pb.ErrInternalError()
	}
	//fill fields.
	ruleUpdate := dao.RuleUpdateCondition{
		ID:     req.Id,
		UserID: req.UserId,
	}
	{
		var flagGenSelectText bool

		// Update name
		rule.Name, ruleUpdate.Name = req.Name, req.Name

		// Update ruleDesc
		rule.RuleDesc, ruleUpdate.RuleDesc = req.RuleDesc, req.RuleDesc

		// update dataType
		rule.DataType, ruleUpdate.DataType = uint8(req.DataType), uint8(req.DataType)

		// Update SelectField
		rule.SelectText, ruleUpdate.SelectText = req.SelectText, req.SelectText

		// Update TopicType
		rule.TopicType, ruleUpdate.TopicType = req.TopicType, req.TopicType
		if constant.TopicTypeRaw == rule.TopicType {
			flagGenSelectText = true
		}

		// Update ShortTopic
		rule.ShortTopic, ruleUpdate.ShortTopic = req.ShortTopic, req.ShortTopic

		// Update Where
		rule.WhereText, ruleUpdate.WhereText = req.Where, req.Where

		// Update Raw
		ruleUpdate.Raw = req.Raw

		if 0 < len(req.SelectFields) || flagGenSelectText {
			ruleUpdate.SelectFields = make([]*dao.SelectField, 0)
			for _, field := range req.SelectFields {
				ruleUpdate.SelectFields = append(ruleUpdate.SelectFields, &dao.SelectField{
					Expr:  field.Expr,
					Type:  field.Type,
					Alias: field.Alias,
				})
			}
			//generate select_text.
			rule.SelectText = daoutil.GenerateSelectText(rule.TopicType, ruleUpdate.SelectFields)
			log.DebugWithFields(UpdatePrefixTag, log.Fields{
				"rule_id":       ruleUpdate.ID,
				"user_id":       ruleUpdate.UserID,
				"select_fields": ruleUpdate.SelectFields,
				"select_text":   rule.SelectText,
			})
			ruleUpdate.SelectText = rule.SelectText
		}
	}
	//generate ruleql, then update
	ruleUpdate.Ruleql = utils.GenerateRuleql(rule)
	_, err = rule.Update(ctx, tx, ruleUpdate)
	if nil != err {
		return nil, pb.ErrInternalError()
	}
	s.Report(constant.EVENT_RULE_UPDATE, rule.ID, rule.UserID)

	return &pb.RuleUpdateResp{}, nil
}
func (s *RulesService) RuleDelete(ctx context.Context, req *pb.RuleDeleteReq) (*pb.RuleDeleteResp, error) {
	//print request [debug]
	printInputDebug(DeletePrefixTag, req)

	var (
		tx     *pg.Tx
		err    error
		rule   dao.Rule
		action dao.Action
	)

	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	tx, err = db.GetTransaction()
	if nil != err {
		log.ErrorWithFields(DeletePrefixTag, log.Fields{
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}
	defer func() {
		err = daoutil.CommitTransaction(tx, err, DeletePrefixTag, log.Fields{
			"desc":    "delete rule successful.",
			"rule_id": rule.ID,
		})
	}()

	// delete rule and actions.
	if _, err = action.Delete(ctx, tx, dao.ActionDeleteCondition{
		RuleID: req.Id,
		UserID: req.UserId,
	}); nil != err {
		tkeelLog.Error("delete actions failed.", err)
		return nil, pb.ErrInternalError()
	}

	if _, err = rule.Delete(ctx, tx, dao.RuleDeleteCondition{
		ID:     req.Id,
		UserID: req.UserId,
	}); nil != err {
		tkeelLog.Error("delete rule failed.", err)
		return nil, pb.ErrInternalError()
	}
	s.Report(constant.EVENT_RULE_DELETE, rule.ID, rule.UserID)
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
	//print request [debug]
	var err error
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	log.DebugWithFields(DebugPrefixTag, log.Fields{
		"rule_id":    req.RuleId,
		"user_id":    req.UserId,
		"thing_id":   req.ThingId,
		"device_id":  req.DeviceId,
		"topic_type": req.TopicType,
		"message":    string(req.Data),
	})
	ru := &dao.Rule{}
	var rules []*dao.Rule
	rules, err = ru.Query(ctx, dao.RuleQueryCondition{
		ID:     req.RuleId,
		UserID: req.UserId,
	})
	if nil == err && 1 == len(rules) {
		//construct message.
		msg := &metapb.ProtoMessage{
			Id:      req.RuleId,
			RawData: req.Data,
		}
		//construct topic
		msg.SetTopic(utils.GenerateTopicMsg(req.ThingId, req.DeviceId, req.TopicType))
		msg.SetDomain(req.UserId)
		msg.SetEntity(req.DeviceId)
		//var r *metapb.RuleExecResponse
		//if r, err = endpoint.GetRuleNodeEndpointIns().ExecRule(ctx, req.RuleId, req.UserId, msg); nil == err {
		//	if "" != r.ErrMsg {
		//		return nil, errors.New(erro.RuleStatusMessage + r.ErrMsg)
		//	}
		//	return &dto.RuleDebugResp{
		//		RuleId: in.RuleId,
		//		Topic:  r.Message.Topic(),
		//		Data:   r.Message.Data(),
		//	}, nil
		//}
		log.ErrorWithFields(DebugPrefixTag, log.Fields{
			"rule_id": req.RuleId,
			"user_id": req.UserId,
			"error":   err,
		})
	} else {
		if nil == err {
			err = errors.New("rule not exists")
		}
	}
	defer func(err error) {
		if nil != err {
			log.ErrorWithFields(DebugPrefixTag, log.Fields{
				"rule_id": req.RuleId,
				"user_id": req.UserId,
				"error":   err,
			})
		}
	}(err)
	return &pb.RuleDebugResp{}, nil
}
func (s *RulesService) RuleDebugMessage(ctx context.Context, req *pb.RuleDebugMsgReq) (*pb.RuleDebugMsgResp, error) {
	return &pb.RuleDebugMsgResp{}, nil
}
func (s *RulesService) RuleError(ctx context.Context, req *pb.RuleErrorReq) (*pb.RuleErrorResp, error) {
	return &pb.RuleErrorResp{}, nil
}

func (s RulesService) Report(et constant.EventType, ruleID, userID string) {
	log.DebugWithFields("[RuleReport] ", log.Fields{
		"type":    et,
		"rule_id": ruleID,
	})
	event.Collect(event.Message{
		Type: event.MessageTypeRuleActive,
		Data: event.NewRuleActiveEvent(et, ruleID, userID, nil),
	})

}
