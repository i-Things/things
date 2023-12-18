// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package tenantmanage

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq                  = sys.ApiDeleteReq
	ApiInfo                       = sys.ApiInfo
	ApiInfoIndexReq               = sys.ApiInfoIndexReq
	ApiInfoIndexResp              = sys.ApiInfoIndexResp
	AppInfo                       = sys.AppInfo
	AppInfoIndexReq               = sys.AppInfoIndexReq
	AppInfoIndexResp              = sys.AppInfoIndexResp
	AreaInfo                      = sys.AreaInfo
	AreaInfoDeleteReq             = sys.AreaInfoDeleteReq
	AreaInfoIndexReq              = sys.AreaInfoIndexReq
	AreaInfoIndexResp             = sys.AreaInfoIndexResp
	AreaInfoReadReq               = sys.AreaInfoReadReq
	AreaInfoTreeReq               = sys.AreaInfoTreeReq
	AreaInfoTreeResp              = sys.AreaInfoTreeResp
	AuthApiInfo                   = sys.AuthApiInfo
	ConfigResp                    = sys.ConfigResp
	DateRange                     = sys.DateRange
	JwtToken                      = sys.JwtToken
	LoginLogCreateReq             = sys.LoginLogCreateReq
	LoginLogIndexReq              = sys.LoginLogIndexReq
	LoginLogIndexResp             = sys.LoginLogIndexResp
	LoginLogInfo                  = sys.LoginLogInfo
	Map                           = sys.Map
	MenuInfo                      = sys.MenuInfo
	MenuInfoIndexReq              = sys.MenuInfoIndexReq
	MenuInfoIndexResp             = sys.MenuInfoIndexResp
	OperLogCreateReq              = sys.OperLogCreateReq
	OperLogIndexReq               = sys.OperLogIndexReq
	OperLogIndexResp              = sys.OperLogIndexResp
	OperLogInfo                   = sys.OperLogInfo
	PageInfo                      = sys.PageInfo
	PageInfo_OrderBy              = sys.PageInfo_OrderBy
	Point                         = sys.Point
	ProjectInfo                   = sys.ProjectInfo
	ProjectInfoDeleteReq          = sys.ProjectInfoDeleteReq
	ProjectInfoIndexReq           = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp          = sys.ProjectInfoIndexResp
	ProjectInfoReadReq            = sys.ProjectInfoReadReq
	ReqWithID                     = sys.ReqWithID
	ReqWithIDCode                 = sys.ReqWithIDCode
	Response                      = sys.Response
	RoleApiAuthReq                = sys.RoleApiAuthReq
	RoleApiIndexReq               = sys.RoleApiIndexReq
	RoleApiIndexResp              = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq         = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq               = sys.RoleAppIndexReq
	RoleAppIndexResp              = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq         = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq              = sys.RoleAppUpdateReq
	RoleInfo                      = sys.RoleInfo
	RoleInfoIndexReq              = sys.RoleInfoIndexReq
	RoleInfoIndexResp             = sys.RoleInfoIndexResp
	RoleMenuIndexReq              = sys.RoleMenuIndexReq
	RoleMenuIndexResp             = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq        = sys.RoleMenuMultiUpdateReq
	TenantAppIndexReq             = sys.TenantAppIndexReq
	TenantAppIndexResp            = sys.TenantAppIndexResp
	TenantAppMultiUpdateReq       = sys.TenantAppMultiUpdateReq
	TenantInfo                    = sys.TenantInfo
	TenantInfoCreateReq           = sys.TenantInfoCreateReq
	TenantInfoIndexReq            = sys.TenantInfoIndexReq
	TenantInfoIndexResp           = sys.TenantInfoIndexResp
	UserAuthArea                  = sys.UserAuthArea
	UserAuthAreaIndexReq          = sys.UserAuthAreaIndexReq
	UserAuthAreaIndexResp         = sys.UserAuthAreaIndexResp
	UserAuthAreaMultiUpdateReq    = sys.UserAuthAreaMultiUpdateReq
	UserAuthProject               = sys.UserAuthProject
	UserAuthProjectIndexReq       = sys.UserAuthProjectIndexReq
	UserAuthProjectIndexResp      = sys.UserAuthProjectIndexResp
	UserAuthProjectMultiUpdateReq = sys.UserAuthProjectMultiUpdateReq
	UserCheckTokenReq             = sys.UserCheckTokenReq
	UserCheckTokenResp            = sys.UserCheckTokenResp
	UserCreateResp                = sys.UserCreateResp
	UserInfo                      = sys.UserInfo
	UserInfoCreateReq             = sys.UserInfoCreateReq
	UserInfoDeleteReq             = sys.UserInfoDeleteReq
	UserInfoIndexReq              = sys.UserInfoIndexReq
	UserInfoIndexResp             = sys.UserInfoIndexResp
	UserInfoReadReq               = sys.UserInfoReadReq
	UserLoginReq                  = sys.UserLoginReq
	UserLoginResp                 = sys.UserLoginResp
	UserRegister1Req              = sys.UserRegister1Req
	UserRegister1Resp             = sys.UserRegister1Resp
	UserRegister2Req              = sys.UserRegister2Req
	UserRoleIndexReq              = sys.UserRoleIndexReq
	UserRoleIndexResp             = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq        = sys.UserRoleMultiUpdateReq

	TenantManage interface {
		// 新增区域
		TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 更新区域
		TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除区域
		TenantInfoDelete(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*Response, error)
		// 获取租户信息详情
		TenantInfoRead(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*TenantInfo, error)
		// 获取租户信息列表
		TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error)
		TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error)
		TenantAppMultiUpdate(ctx context.Context, in *TenantAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultTenantManage struct {
		cli zrpc.Client
	}

	directTenantManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.TenantManageServer
	}
)

func NewTenantManage(cli zrpc.Client) TenantManage {
	return &defaultTenantManage{
		cli: cli,
	}
}

func NewDirectTenantManage(svcCtx *svc.ServiceContext, svr sys.TenantManageServer) TenantManage {
	return &directTenantManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增区域
func (m *defaultTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoCreate(ctx, in, opts...)
}

// 新增区域
func (d *directTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantInfoCreate(ctx, in)
}

// 更新区域
func (m *defaultTenantManage) TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoUpdate(ctx, in, opts...)
}

// 更新区域
func (d *directTenantManage) TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantInfoUpdate(ctx, in)
}

// 删除区域
func (m *defaultTenantManage) TenantInfoDelete(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoDelete(ctx, in, opts...)
}

// 删除区域
func (d *directTenantManage) TenantInfoDelete(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantInfoDelete(ctx, in)
}

// 获取租户信息详情
func (m *defaultTenantManage) TenantInfoRead(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoRead(ctx, in, opts...)
}

// 获取租户信息详情
func (d *directTenantManage) TenantInfoRead(ctx context.Context, in *ReqWithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
	return d.svr.TenantInfoRead(ctx, in)
}

// 获取租户信息列表
func (m *defaultTenantManage) TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoIndex(ctx, in, opts...)
}

// 获取租户信息列表
func (d *directTenantManage) TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error) {
	return d.svr.TenantInfoIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error) {
	return d.svr.TenantAppIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppMultiUpdate(ctx context.Context, in *TenantAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMultiUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMultiUpdate(ctx context.Context, in *TenantAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppMultiUpdate(ctx, in)
}