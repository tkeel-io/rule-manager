package service

import (
	"context"

	v1 "github.com/tkeel-io/rule-manager/api/openapi/v1"
	"github.com/tkeel-io/rule-manager/pkg/util"
	openapi_v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// OpenapiService is a openapi service.
type OpenapiService struct {
	v1.UnimplementedOpenapiServer
}

// NewOpenapiService new a openapi service.
func NewOpenapiService() *OpenapiService {
	return &OpenapiService{
		UnimplementedOpenapiServer: v1.UnimplementedOpenapiServer{},
	}
}

// AddonsIdentify implements AddonsIdentify.OpenapiServer.
func (s *OpenapiService) AddonsIdentify(ctx context.Context, in *openapi_v1.AddonsIdentifyRequest) (*openapi_v1.AddonsIdentifyResponse, error) {
	return &openapi_v1.AddonsIdentifyResponse{
		Res: util.GetV1ResultBadRequest("not declare addons"),
	}, nil
}

// Identify implements Identify.OpenapiServer.
func (s *OpenapiService) Identify(ctx context.Context, in *emptypb.Empty) (*openapi_v1.IdentifyResponse, error) {
	profiles := map[string]*openapi_v1.ProfileSchema{
		"rule_max":  &openapi_v1.ProfileSchema{Type: "number", Title: "规则最大数", Default: 5, MultipleOf: 1, Maximum: 20, Minimum: 0},
	}
	return &openapi_v1.IdentifyResponse{
		Res:                     util.GetV1ResultOK(),
		PluginId:                "rule-manager",
		Version:                 "0.4.1",
		TkeelVersion:            "v0.4.0",
		DisableManualActivation: true,
		Profiles: profiles,
	}, nil
}

// Status implements Status.OpenapiServer.
func (s *OpenapiService) Status(ctx context.Context, in *emptypb.Empty) (*openapi_v1.StatusResponse, error) {
	return &openapi_v1.StatusResponse{
		Res:    util.GetV1ResultOK(),
		Status: openapi_v1.PluginStatus_RUNNING,
	}, nil
}

// TenantEnable implements TenantEnable.OpenapiServer.
func (s *OpenapiService) TenantEnable(ctx context.Context, in *openapi_v1.TenantEnableRequest) (*openapi_v1.TenantEnableResponse, error) {
	return &openapi_v1.TenantEnableResponse{
		Res: util.GetV1ResultOK(),
	}, nil
}

// TenantDisable implements TenantDisable.OpenapiServer.
func (s *OpenapiService) TenantDisable(ctx context.Context, in *openapi_v1.TenantDisableRequest) (*openapi_v1.TenantDisableResponse, error) {
	return &openapi_v1.TenantDisableResponse{
		Res: util.GetV1ResultOK(),
	}, nil
}
