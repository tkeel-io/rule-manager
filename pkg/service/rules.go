package service

import (
	"context"
	"encoding/json"
	"github.com/tkeel-io/core-broker/pkg/core"
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
	CreatePrefixTag     = "[RuleCreate]"
	UpdatePrefixTag     = "[RuleUpdate]"
	DeletePrefixTag     = "[RuleDelete]"
	QueryPrefixTag      = "[RuleQuery]"
	RuleStatusPrefixTag = "[RuleStatus]"
	DebugPrefixTag      = "[RuleDebug]"
	ErrorPrefixTag      = "[RuleError]"
)

type RulesService struct {
	pb.UnimplementedRulesServer
	Core *core.Client
}

func NewRulesService() *RulesService {
	client, err := core.NewCoreClient()
	if err != nil {
		tkeelLog.Fatal("init core sdk failed", err)
	}
	return &RulesService{
		Core: client,
	}
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

	return &pb.RuleUpdateResp{
		Id:           ruleUpdate.ID,
		UserId:       ruleUpdate.UserID,
		Name:         ruleUpdate.Name,
		RuleDesc:     ruleUpdate.RuleDesc,
		DataType:     uint32(ruleUpdate.DataType),
		SelectText:   ruleUpdate.SelectText,
		SelectFields: req.SelectFields,
		TopicType:    ruleUpdate.TopicType,
		ShortTopic:   ruleUpdate.ShortTopic,
		Where:        ruleUpdate.WhereText,
		Raw:          ruleUpdate.Raw,
	}, nil
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
	r, err := s.RuleQuery(ctx, &pb.RuleQueryReq{
		Id:     &pb.String{Value: req.Id},
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if r.Rules == nil || len(r.Rules) == 0 {
		return nil, pb.ErrRuleNotFound()
	}
	rule := r.Rules[0]
	return &pb.Rule{
		Id:           rule.Id,
		UserId:       rule.UserId,
		Name:         rule.Name,
		RuleDesc:     rule.RuleDesc,
		DataType:     rule.DataType,
		SelectText:   rule.SelectText,
		TopicType:    rule.TopicType,
		ShortTopic:   rule.ShortTopic,
		Where:        rule.Where,
		Status:       rule.Status,
		Ruleql:       rule.Ruleql,
		Raw:          rule.Raw,
		Topic:        rule.Topic,
		ConfigStatus: rule.ConfigStatus,
		SelectFields: rule.SelectFields,
		LastError:    rule.LastError,
		CreateTime:   rule.CreateTime,
		UpdateTime:   rule.UpdateTime,
	}, nil
}

func (s *RulesService) RuleQuery(ctx context.Context, req *pb.RuleQueryReq) (*pb.RuleQueryResp, error) {
	//print request [debug]
	printInputDebug(QueryPrefixTag, req)

	var rule = &dao.Rule{}
	var action = &dao.Action{}
	var actions []*dao.Action
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()

	cond := dao.RuleQueryCondition{
		IDs:    req.Ids,
		UserID: req.UserId,
	}
	{
		if nil != req.Id {
			cond.ID = req.Id.Value
		}
		if nil != req.Name {
			cond.Name = req.Name.Value
		}
		if nil != req.DataType {
			value := uint8(req.DataType.Value)
			cond.DataType = value
		}
		if nil != req.TopicType {
			cond.TopicType = req.TopicType.Value
		}
		if nil != req.ShortTopic {
			cond.ShortTopic = req.ShortTopic.Value
		}
		if nil != req.SearchKey {
			cond.SearchKey = req.SearchKey.Value
		}
	}
	//query.
	rules, err := rule.Query(ctx, cond)
	if nil != err {
		log.ErrorWithFields(QueryPrefixTag, log.Fields{
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}
	total, err := queryTotalRule(cond.UserID)
	if nil != err {
		log.ErrorWithFields(QueryPrefixTag, log.Fields{
			"error": err,
		})
		return nil, pb.ErrInternalError()
	}
	res := &pb.RuleQueryResp{
		Total: total,
		Rules: make([]*pb.Rule, 0),
	}
	for _, ru := range rules {
		//过滤逻辑删除的rule。
		if constant.RuleStatusBan == ru.Status {
			continue
		}
		//field: SelectFeilds:
		fields := make([]*pb.SelectField, 0)
		for _, field := range ru.SelectFields {
			fields = append(fields, &pb.SelectField{
				Expr:  field.Expr,
				Type:  field.Type,
				Alias: field.Alias,
			})
		}
		r := &pb.Rule{
			Id:           ru.ID,
			UserId:       ru.UserID,
			Name:         ru.Name,
			Status:       ru.Status,
			RuleDesc:     ru.RuleDesc,
			DataType:     int32(ru.DataType),
			SelectText:   ru.SelectText,
			SelectFields: fields,
			TopicType:    ru.TopicType,
			Topic:        utils.GenerateTopic(ru),
			ShortTopic:   ru.ShortTopic,
			Where:        ru.WhereText,
			Ruleql:       ru.Ruleql,
			Raw:          ru.RawRequest,
			CreateTime:   ru.CreateTime,
			UpdateTime:   ru.UpdateTime,
		}
		//check rule.
		status, checkRule := true, false
		checkRule, _ = daoutil.ValidateRule(ctx, ru)
		//query action by rule.Id
		checkAction, checkActionErr := false, false
		actions, err = action.Query(ctx, dao.ActionQueryCondition{
			RuleID:       r.Id,
			UserID:       ru.UserID,
			ConfigStatus: &status,
		})
		if nil != err {
			log.Error(err)
			return nil, pb.ErrInternalError()
		}
		for _, ac := range actions {
			if ac.ErrorActionFlag {
				checkActionErr = true
			} else {
				checkAction = true
			}
		}
		r.ConfigStatus = &pb.ConfigStatus{
			DataSelectFlag:   checkRule,
			DataDispatchFlag: checkAction,
			DataErrorFlag:    checkActionErr,
		}
		res.Rules = append(res.Rules, r)
	}

	return res, nil
}

func (s *RulesService) RuleStatus(ctx context.Context, req *pb.RuleStatusReq) (*pb.RuleStatusResp, error) {
	// callback error.
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	var err error
	defer cancelHandler()
	defer func() {
		fields := log.Fields{
			"call":     "RuleStatus",
			"rule_id":  req.Id,
			"operator": req.Operator,
			"status":   req.Status,
			"userId":   req.UserId,
			"error":    err,
		}
		if nil == err {
			log.InfoWithFields(RuleStatusPrefixTag, fields)
		} else {
			log.ErrorWithFields(RuleStatusPrefixTag, fields)
		}
	}()

	switch req.Operator {
	case pb.RuleStatusReq_STATUS_READ:
		var status string
		if status, err = s.getRuleStatus(ctx, req.Id, req.UserId); nil == err {
			return &pb.RuleStatusResp{
				Id:     req.Id,
				Status: status,
			}, nil
		}
		return nil, pb.ErrInternalError()

	case pb.RuleStatusReq_STATUS_WRITE:
		//var ruleReq *endpointutil.Rule
		//if ruleReq, err = endpointutil.ConvertRule(ctx, req.Id, req.UserId); nil != err {
		//	return nil, err
		//}
		switch req.Status {
		case constant.CommandStatusRuleStart:
			//if err = endpoint.GetMetadataEndpointIns().AddRule(ctx, ruleReq); nil != err {
			//}
			//update rule status pg.
			if err = daoutil.UpdateStatusPG(ctx, req.Id, req.UserId, constant.RuleStatusStarting); nil == err {
				// update action status pg.
				if err = daoutil.UpdateStatusPGA(ctx, req.Id, "", req.UserId, constant.RuleStatusStop); nil == err {
					return &pb.RuleStatusResp{
						Id:     req.Id,
						Status: constant.RuleStatusStarting,
					}, nil
				}
			}

		case constant.CommandStatusRuleStop:
			//if err = endpoint.GetMetadataEndpointIns().DelRule(ctx, ruleReq); nil != err {
			//}
			//update pg.
			if err = daoutil.UpdateStatusPG(ctx, req.Id, req.UserId, constant.RuleStatusStopping); nil == err {
				return &pb.RuleStatusResp{
					Id:     req.Id,
					Status: constant.RuleStatusStopping,
				}, nil
			}

		default:
			err = errors.New("rule status invalid")
		}
	default:
		err = errors.New("rule operator invalid")
	}
	return nil, err
}

func (s *RulesService) RuleStart(ctx context.Context, req *pb.RuleStartReq) (*pb.RuleStartResp, error) {
	r, err := s.RuleStatus(ctx, &pb.RuleStatusReq{
		Id:       req.Id,
		UserId:   req.UserId,
		Operator: pb.RuleStatusReq_STATUS_WRITE,
		Status:   constant.CommandStatusRuleStart,
	})
	if err != nil {
		tkeelLog.Error(err)
		return nil, err
	}
	return &pb.RuleStartResp{
		Id:     r.Id,
		Status: r.Status,
	}, nil
}

func (s *RulesService) RuleStop(ctx context.Context, req *pb.RuleStopReq) (*pb.RuleStopResp, error) {
	r, err := s.RuleStatus(ctx, &pb.RuleStatusReq{
		Id:       req.Id,
		UserId:   req.UserId,
		Operator: pb.RuleStatusReq_STATUS_WRITE,
		Status:   constant.CommandStatusRuleStop,
	})
	if err != nil {
		tkeelLog.Error(err)
		return nil, err
	}
	return &pb.RuleStopResp{
		Id:     r.Id,
		Status: r.Status,
	}, nil
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
		var r *metapb.RuleExecResponse
		//if r, err = endpoint.GetRuleNodeEndpointIns().ExecRule(ctx, req.RuleId, req.UserId, msg); nil != err {
		//}
		if "" != r.ErrMsg {
			return nil, errors.New(r.ErrMsg)
		}
		log.ErrorWithFields(DebugPrefixTag, log.Fields{
			"rule_id": req.RuleId,
			"user_id": req.UserId,
			"error":   err,
		})
		return &pb.RuleDebugResp{
			RuleId: req.RuleId,
			Topic:  r.Message.Topic(),
			Data:   r.Message.Data(),
		}, nil

	} else if nil == err {
		return nil, pb.ErrRuleNotFound()
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
	return nil, err
}

func (s *RulesService) RuleDebugMessage(ctx context.Context, req *pb.RuleDebugMsgReq) (*pb.RuleDebugMsgResp, error) {
	var (
		buf       []byte
		postTypes string
		rule      = &dao.Rule{
			ID:     req.RuleId,
			UserID: req.UserId,
		}

		err error
	)
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	defer func() {
		if nil != err {
			log.ErrorWithFields(DebugPrefixTag, log.Fields{
				"rule_id": req.RuleId,
				"user_id": req.UserId,
				"error":   err.Error(),
			})
		}
	}()

	if err = rule.Select(ctx); nil != err {
		return nil, pb.ErrInternalError()
	}

	//var endp = endpoint.NewThingEndpoint()
	//var me *endpoint.EventPropertyInfo
	thingID := utils.GetThingId(rule)
	e, err := s.Core.GetDeviceEntity(req.DeviceId)
	if err != nil {
		log.ErrorWithFields(DebugPrefixTag, log.Fields{
			"rule_id": req.RuleId,
			"user_id": req.UserId,
			"message": "call Core GET Entity ERR",
			"error":   err,
		})
		return nil, pb.ErrInternalError()
	}
	message := &utils.ThingMessage{
		Id:      "模拟数据输入...",
		Version: "0.0.0",
		Type:    postTypes,
		Metadata: utils.Metadata{
			EntityId:  req.DeviceId,
			ModelId:   thingID,
			SourceId:  []string{req.DeviceId},
			EpochTime: time.Now().Unix(),
		},
		Params: map[string]interface{}{"device_info": e},
	}

	if buf, err = json.Marshal(message); err != nil {
		return nil, pb.ErrInternalError()
	}

	return &pb.RuleDebugMsgResp{
		RuleId:  req.RuleId,
		Message: buf,
	}, nil
}

func (s *RulesService) RuleError(ctx context.Context, req *pb.RuleErrorReq) (*pb.RuleErrorResp, error) {
	/*
		in.RuleId
	*/
	log.ErrorWithFields(ErrorPrefixTag, log.Fields{
		"rule_id": req.RuleId,
	})
	return &pb.RuleErrorResp{
		RuleId: req.RuleId,
		Error:  nil,
	}, nil
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

func (s RulesService) getRuleStatus(ctx context.Context, id, userID string) (status string, err error) {
	//print request [debug]
	ctx, cancelHandler := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancelHandler()
	rule := dao.Rule{}
	res, err := rule.Query(ctx, dao.RuleQueryCondition{
		ID:     id,
		UserID: userID,
	})
	if nil != err || nil == res || len(res) != 1 {
		log.ErrorWithFields("[RuleStatus]", log.Fields{
			"query response": res,
			"count":          len(res),
			"error:":         err,
		})
		if len(res) == 0 {
			err = errors.New("rule not existed")
		}
		return "", err
	}
	return res[0].Status, nil
}

func queryTotalRule(id string) (int32, error) {
	rule := &dao.Rule{}
	rules, err := rule.Query(context.TODO(), dao.RuleQueryCondition{
		UserID: id,
	})
	return int32(len(rules)), err
}
