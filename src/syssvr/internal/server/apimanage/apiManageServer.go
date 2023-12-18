// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/logic/apimanage"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
)

type ApiManageServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedApiManageServer
}

func NewApiManageServer(svcCtx *svc.ServiceContext) *ApiManageServer {
	return &ApiManageServer{
		svcCtx: svcCtx,
	}
}

func (s *ApiManageServer) ApiInfoCreate(ctx context.Context, in *sys.ApiInfo) (*sys.Response, error) {
	l := apimanagelogic.NewApiInfoCreateLogic(ctx, s.svcCtx)
	return l.ApiInfoCreate(in)
}

func (s *ApiManageServer) ApiInfoIndex(ctx context.Context, in *sys.ApiInfoIndexReq) (*sys.ApiInfoIndexResp, error) {
	l := apimanagelogic.NewApiInfoIndexLogic(ctx, s.svcCtx)
	return l.ApiInfoIndex(in)
}

func (s *ApiManageServer) ApiInfoUpdate(ctx context.Context, in *sys.ApiInfo) (*sys.Response, error) {
	l := apimanagelogic.NewApiInfoUpdateLogic(ctx, s.svcCtx)
	return l.ApiInfoUpdate(in)
}

func (s *ApiManageServer) ApiInfoDelete(ctx context.Context, in *sys.ReqWithID) (*sys.Response, error) {
	l := apimanagelogic.NewApiInfoDeleteLogic(ctx, s.svcCtx)
	return l.ApiInfoDelete(in)
}