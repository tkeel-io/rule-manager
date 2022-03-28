package service

import (
	"context"
	"github.com/Shopify/sarama"
	"strings"

	"github.com/tkeel-io/core-broker/pkg/auth"
	"github.com/tkeel-io/core-broker/pkg/core"
	"github.com/tkeel-io/core-broker/pkg/deviceutil"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	"github.com/tkeel-io/kit/log"
	tkeelLog "github.com/tkeel-io/kit/log"
	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	"github.com/tkeel-io/rule-manager/internal/dao"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// Log prefix
const (
	CreatePrefixTag = "[RuleCreate]"
	UpdatePrefixTag = "[RuleUpdate]"
	DeletePrefixTag = "[RuleDelete]"
	QueryPrefixTag  = "[RuleQuery]"
)

var (
	ErrUnmatched      = errors.New("delete records are not matched whit user")
	ErrDeviceNotFound = errors.New("device not found")
)

type RulesService struct {
	pb.UnimplementedRulesServer
	Core *core.Client
}

func NewRulesService() *RulesService {
	if dao.CoreClient == nil {
		if err := dao.SetCoreClientUp(); err != nil {
			tkeelLog.Fatal("setup core client failed", err)
		}
	}
	return &RulesService{
		Core: dao.CoreClient,
	}
}

func (s *RulesService) RuleCreate(ctx context.Context, req *pb.RuleCreateReq) (res *pb.RuleCreateResp, err error) {
	printInputDebug(CreatePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := dao.Rule{
		UserID: user.ID,
		Name:   req.Name,
		Status: dao.NotRunningStatus,
		Desc:   req.Desc,
		Type:   uint8(req.Type),
	}

	result := dao.DB().Model(&rule).Create(&rule)
	if result.Error != nil {
		log.Error(CreatePrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}
	return &pb.RuleCreateResp{
		Id:        uint64(rule.ID),
		Name:      rule.Name,
		Desc:      rule.Desc,
		Status:    uint32(rule.Status),
		Type:      uint32(rule.Type),
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	}, nil
}

func (s *RulesService) RuleUpdate(ctx context.Context, req *pb.RuleUpdateReq) (*pb.RuleUpdateResp, error) {
	printInputDebug(UpdatePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}

	var c int
	result := dao.DB().Model(&rule).Select("1").
		Where(&rule).
		First(&c)
	if errors.Is(
		result.Error,
		gorm.ErrRecordNotFound,
	) || result.RowsAffected == 0 {
		return nil, pb.ErrForbidden()
	}

	result = dao.DB().Model(&rule).First(&rule)
	if result.Error != nil {
		tkeelLog.Error(UpdatePrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	rule.Name = req.Name
	rule.Desc = req.Desc

	result = dao.DB().Save(&rule)
	if result.Error != nil {
		return nil, pb.ErrInternalError()
	}

	return &pb.RuleUpdateResp{
		Id:        uint64(rule.ID),
		Name:      rule.Name,
		Desc:      rule.Desc,
		Status:    uint32(rule.Status),
		Type:      uint32(rule.Type),
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	}, nil
}

func (s *RulesService) RuleDelete(ctx context.Context, req *pb.RuleDeleteReq) (*pb.RuleDeleteResp, error) {
	//print request [debug]
	printInputDebug(DeletePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	result := dao.DB().Model(&rule).Where(&rule).First(&rule)
	if result.Error != nil {
		tkeelLog.Error(DeletePrefixTag, result.Error)
		return nil, pb.ErrForbidden()
	}

	result = dao.DB().Delete(&rule)
	if result.Error != nil {
		tkeelLog.Error(DeletePrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	return &pb.RuleDeleteResp{}, nil
}

func (s *RulesService) RuleGet(ctx context.Context, req *pb.RuleGetReq) (*pb.Rule, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	if result := rule.Select(); result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	var ds, ts int64
	if err = dao.DB().Model(&dao.RuleEntities{}).Where("rule_id = ?", rule.ID).Limit(1).Count(&ds).Error; err != nil {
		log.Error("query rule entities count error", err)
	}
	if err = dao.DB().Model(&dao.Target{}).Where("rule_id = ?", rule.ID).Limit(1).Count(&ts).Error; err != nil {
		log.Error("query rule target count error", err)
	}

	return &pb.Rule{
		Id:            uint64(rule.ID),
		Name:          rule.Name,
		Desc:          rule.Desc,
		Status:        uint32(rule.Status),
		Type:          uint32(rule.Type),
		CreatedAt:     rule.CreatedAt.Unix(),
		UpdatedAt:     rule.UpdatedAt.Unix(),
		SubId:         uint32(rule.SubID),
		DevicesStatus: uint32(ds),
		TargetsStatus: uint32(ts),
	}, nil
}

func (s *RulesService) RuleQuery(ctx context.Context, req *pb.RuleQueryReq) (*pb.RuleQueryResp, error) {
	//print request [debug]
	printInputDebug(QueryPrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}

	tkeelLog.Debug("query rule", req)
	page, err := pagination.Parse(req)
	if err != nil {
		tkeelLog.Error(QueryPrefixTag, err)
		return nil, pb.ErrInternalError()
	}

	rule := &dao.Rule{UserID: user.ID}
	tx := dao.DB().Model(&rule).Where(&rule)

	fillPagination(tx, page)

	if req.Id != nil && req.Ids != nil && len(req.Ids) > 0 {
		return nil, pb.ErrInvalidArgument()
	}

	if req.Id != nil {
		tx.Where("id = ?", req.Id.Value)
	}

	if req.Ids != nil && len(req.Ids) > 0 {
		tx.Where("id in (?)", req.Ids)
	}

	if req.Name != nil {
		tx.Where("name = ?", req.Name.Value)
	}

	if req.Type != nil {
		tx.Where("type = ?", req.Type.Value)
	}

	if req.Status != nil {
		tx.Where("status = ?", req.Status.Value)
	}

	rules := make([]*dao.Rule, 0)
	result := tx.Find(&rules)
	if result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}
	var count int64
	result = tx.Count(&count)
	if result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	resp := &pb.RuleQueryResp{}

	page.SetTotal(uint(count))
	if err = page.FillResponse(resp); err != nil {
		tkeelLog.Error(QueryPrefixTag, err)
		return nil, err
	}
	resp.Data = make([]*pb.Rule, 0, len(rules))

	for _, r := range rules {
		var ds, ts int64
		if err = dao.DB().Model(&dao.RuleEntities{}).Where("rule_id = ?", r.ID).Limit(1).Count(&ds).Error; err != nil {
			log.Error("query rule entities count error", err)
		}
		if err = dao.DB().Model(&dao.Target{}).Where("rule_id = ?", r.ID).Limit(1).Count(&ds).Error; err != nil {
			log.Error("query rule target count error", err)
		}
		resp.Data = append(resp.Data, &pb.Rule{
			Id:            uint64(r.ID),
			Name:          r.Name,
			Desc:          r.Desc,
			Status:        uint32(r.Status),
			Type:          uint32(r.Type),
			CreatedAt:     r.CreatedAt.Unix(),
			UpdatedAt:     r.UpdatedAt.Unix(),
			DevicesStatus: uint32(ds),
			TargetsStatus: uint32(ts),
			SubId:         uint32(r.SubID),
		})
	}
	return resp, nil
}

func (s *RulesService) RuleStatusSwitch(ctx context.Context, req *pb.RuleStatusSwitchReq) (*pb.RuleStatusSwitchResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	result := dao.DB().Model(&rule).Where(&rule).First(&rule)
	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrForbidden()
	}
	rule.Status = uint8(req.Status)
	result = dao.DB().Save(&rule)
	if result.Error != nil {
		return nil, pb.ErrInternalError()
	}
	return &pb.RuleStatusSwitchResp{Status: uint32(rule.Status), Id: uint64(rule.ID)}, nil
}

func (s *RulesService) GetRuleDevicesID(ctx context.Context, req *pb.RuleDevicesIDReq) (*pb.RuleDevicesIDResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	resp := &pb.RuleDevicesIDResp{}
	reModel := dao.RuleEntities{RuleID: uint(req.Id)}
	ids := reModel.FindEntityIDS()
	resp.DevicesIds = ids

	return resp, nil
}

func (s *RulesService) AddDevicesToRule(ctx context.Context, req *pb.AddDevicesToRuleReq) (*emptypb.Empty, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	if len(req.DevicesIds) == 0 {
		return &emptypb.Empty{}, nil
	}
	if err := saveDevicesToRule(rule, req.DevicesIds); err != nil {
		tkeelLog.Error("save rule_entities records err", err)
		return nil, pb.ErrInternalError()
	}

	return &emptypb.Empty{}, nil
}

func (s *RulesService) RemoveDevicesFromRule(ctx context.Context, req *pb.RemoveDevicesFromRuleReq) (*emptypb.Empty, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}
	ids := strings.Split(req.DevicesIds, ",")
	tx := dao.DB().Begin()
	if err = removeDevicesFromRule(tx, rule, ids); err != nil {
		defer func() {
			tx.Rollback()
		}()
		if errors.Is(err, ErrUnmatched) {
			return nil, pb.ErrForbidden()
		}
		return nil, pb.ErrInternalError()
	}
	defer func() {
		tx.Commit()
	}()

	return &emptypb.Empty{}, nil
}

func (s *RulesService) GetRuleDevices(ctx context.Context, req *pb.RuleDevicesReq) (*pb.RuleDevicesResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	page, err := pagination.Parse(req)
	if err != nil {
		tkeelLog.Error("parse pagination error", err)
		return nil, pb.ErrInvalidArgument()
	}

	resp := &pb.RuleDevicesResp{}

	ruleEntitiesCondition := dao.RuleEntities{
		RuleID: uint(req.Id),
	}
	total, err := ruleEntitiesCondition.Count(nil)
	if err != nil {
		tkeelLog.Error("query total error", err)
		return nil, pb.ErrInternalError()
	}
	page.SetTotal(uint(total))

	tx := dao.DB().Model(&ruleEntitiesCondition)
	if page.Required() {
		tx = tx.Limit(int(page.Limit())).Offset(int(page.Offset()))
	}
	ress := ruleEntitiesCondition.Find(tx)
	if len(ress) == 0 {
		return resp, nil
	}

	resp.Data, err = s.getDevicesFromCore(user.Token, ress)
	if err != nil {
		tkeelLog.Error("get devices from core error", err)
		return nil, pb.ErrInternalError()
	}
	if err = page.FillResponse(resp); err != nil {
		tkeelLog.Error("fill response error", err)
		return nil, pb.ErrInternalError()
	}
	return resp, nil
}

func (s *RulesService) CreateRuleTarget(ctx context.Context, req *pb.CreateRuleTargetReq) (*pb.CreateRuleTargetResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	if req.Type == 0 || req.Host == "" || req.Value == "" {
		return nil, pb.ErrInvalidArgument()
	}

	ruleTarget := &dao.Target{
		RuleID: uint(req.Id),
		Type:   uint8(req.Type),
		Host:   req.Host,
		Value:  req.Value,
		Ext:    req.Ext,
	}
	if err = ruleTarget.Create(); err != nil {
		tkeelLog.Error("save target record error", err)
		return nil, pb.ErrInternalError()
	}

	return &pb.CreateRuleTargetResp{
		Id:    uint64(ruleTarget.ID),
		Type:  uint32(ruleTarget.Type),
		Host:  ruleTarget.Host,
		Value: ruleTarget.Value,
		Ext:   ruleTarget.Ext,
	}, nil
}

func (s *RulesService) UpdateRuleTarget(ctx context.Context, req *pb.UpdateRuleTargetReq) (*pb.UpdateRuleTargetResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}
	target := &dao.Target{RuleID: rule.ID, ID: uint(req.TargetId)}
	err = target.Find()
	if err != nil {
		tkeelLog.Error("target not found", err)
		return nil, pb.ErrForbidden()
	}

	target.Ext = req.Ext
	target.Value = req.Value
	target.Host = req.Host

	if err = dao.DB().Save(target).Error; err != nil {
		tkeelLog.Error("save target record error", err)
		return nil, pb.ErrInternalError()
	}
	return &pb.UpdateRuleTargetResp{
		Id:    uint64(target.ID),
		Type:  uint32(target.Type),
		Host:  target.Host,
		Value: target.Value,
		Ext:   target.Ext,
	}, nil
}

func (s *RulesService) ListRuleTarget(ctx context.Context, req *pb.ListRuleTargetReq) (*pb.ListRuleTargetResp, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	page, err := pagination.Parse(req)
	if err != nil {
		tkeelLog.Error(QueryPrefixTag, err)
		return nil, pb.ErrInternalError()
	}

	targetConnd := &dao.Target{RuleID: rule.ID}
	var total int64

	targets := make([]*dao.Target, 0)
	tx := dao.DB().Model(targetConnd).Where(targetConnd)
	result := tx.Count(&total)
	if result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}
	page.SetTotal(uint(total))

	if page.Required() {
		fillPagination(tx, page)
	}
	result = tx.Find(&targets)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	data := make([]*pb.CreateRuleTargetResp, 0, len(targets))
	for _, target := range targets {
		t := &pb.CreateRuleTargetResp{
			Id:    uint64(target.ID),
			Type:  uint32(target.Type),
			Host:  target.Host,
			Value: target.Value,
			Ext:   target.Ext,
		}
		data = append(data, t)
	}

	resp := &pb.ListRuleTargetResp{Data: data}
	if err = page.FillResponse(resp); err != nil {
		tkeelLog.Error("fill response error", err)
		return nil, pb.ErrInternalError()
	}

	return resp, nil
}

func (s RulesService) DeleteRuleTarget(ctx context.Context, req *pb.DeleteRuleTargetReq) (*emptypb.Empty, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	_, err = rule.Exists()
	if err != nil {
		tkeelLog.Error("user and rule are not match", err)
		return nil, pb.ErrForbidden()
	}

	target := &dao.Target{RuleID: rule.ID, ID: uint(req.TargetId)}
	err = target.FindAndAuth(user.ID)
	if err != nil {
		tkeelLog.Error("target not found", err)
		return nil, pb.ErrForbidden()
	}

	if err = target.Delete(); err != nil {
		tkeelLog.Error("delete target record error", err)
		return nil, pb.ErrInternalError()
	}
	return &emptypb.Empty{}, nil
}

func (s *RulesService) TestConnectToKafka(ctx context.Context, req *pb.TestConnectToKafkaReq) (*emptypb.Empty, error) {

	endpoints := strings.Split(req.Host, ",")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	client, err := sarama.NewSyncProducer(endpoints, config)
	if err != nil {
		//log.Error(err)
		log.Errorf(`
			"call":      "KafkaVerify",
			"desc":      "failed open consumer.",
			"endpoints": endpoints,
			"sinkType":  "kafka",
			"error":     %s,
		`, err)
		return nil, pb.ErrInvalidArgument()
	}
	client.Close()

	return &emptypb.Empty{}, nil
}

func (s *RulesService) getDevicesFromCore(token string, ress []dao.RuleEntities) ([]*pb.Device, error) {
	dc := deviceutil.NewClient(token)
	devices := make([]*pb.Device, 0, len(ress))
	for _, re := range ress {
		bytes, err := dc.Search(deviceutil.EntitySearch, deviceutil.Conditions{deviceutil.DeviceQuery(re.EntityID)})
		if err != nil {
			log.Error("query device by device id err:", err)
			return nil, err
		}
		resp, err := deviceutil.ParseSearchEntityResponse(bytes)
		if err != nil {
			log.Error("parse device search response err:", err)
			return nil, err
		}
		if len(resp.Data.Items) == 0 {
			log.Error("device not found:", re.EntityID)
			return nil, ErrDeviceNotFound
		}
		d := &pb.Device{
			Id:        re.EntityID,
			Name:      resp.Data.Items[0].Properties.BasicInfo.Name,
			Template:  resp.Data.Items[0].Properties.BasicInfo.TemplateName,
			GroupName: resp.Data.Items[0].Properties.BasicInfo.ParentName,
			Status:    "offline",
		}
		if resp.Data.Items[0].Properties.ConnectionInfo.IsOnline {
			d.Status = "online"
		}
		devices = append(devices, d)
	}
	return devices, nil
}

func removeDevicesFromRule(tx *gorm.DB, rule *dao.Rule, ids []string) error {
	for _, id := range ids {
		e := dao.RuleEntities{RuleID: rule.ID, EntityID: id, UniqueKey: dao.GenUniqueKey(rule.ID, id)}
		result := tx.
			Where(&e).
			Delete(&e)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func saveDevicesToRule(rule *dao.Rule, ids []string) error {
	if rule == nil || rule.ID == 0 {
		return errors.New("rule is nil or unusable")
	}
	if len(ids) == 0 {
		return nil
	}

	ress := make([]dao.RuleEntities, 0, len(ids))
	for _, id := range ids {
		ress = append(ress, dao.RuleEntities{
			RuleID:   rule.ID,
			EntityID: id,
			Rule:     *rule,
		})
	}

	return dao.DB().Create(&ress).Error
}

func fillPagination(tx *gorm.DB, p pagination.Page) {
	if p.Required() {
		tx.Limit(int(p.Limit())).Offset(int(p.Offset()))
	}
	if p.IsDescending {
		if p.SearchKey != "" && !strings.Contains(p.SearchKey, ",") {
			tx.Order(p.SearchKey + " desc")
		}
	}
}
