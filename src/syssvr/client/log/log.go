// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package log

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq              = sys.ApiDeleteReq
	ApiInfo                   = sys.ApiInfo
	ApiInfoIndexReq           = sys.ApiInfoIndexReq
	ApiInfoIndexResp          = sys.ApiInfoIndexResp
	AppInfo                   = sys.AppInfo
	AppInfoIndexReq           = sys.AppInfoIndexReq
	AppInfoIndexResp          = sys.AppInfoIndexResp
	AppModuleIndexReq         = sys.AppModuleIndexReq
	AppModuleIndexResp        = sys.AppModuleIndexResp
	AppModuleMultiUpdateReq   = sys.AppModuleMultiUpdateReq
	AreaInfo                  = sys.AreaInfo
	AreaInfoIndexReq          = sys.AreaInfoIndexReq
	AreaInfoIndexResp         = sys.AreaInfoIndexResp
	AreaInfoTreeReq           = sys.AreaInfoTreeReq
	AreaInfoTreeResp          = sys.AreaInfoTreeResp
	AreaWithID                = sys.AreaWithID
	AuthApiInfo               = sys.AuthApiInfo
	ConfigResp                = sys.ConfigResp
	DateRange                 = sys.DateRange
	JwtToken                  = sys.JwtToken
	LoginLogCreateReq         = sys.LoginLogCreateReq
	LoginLogIndexReq          = sys.LoginLogIndexReq
	LoginLogIndexResp         = sys.LoginLogIndexResp
	LoginLogInfo              = sys.LoginLogInfo
	Map                       = sys.Map
	MenuInfo                  = sys.MenuInfo
	MenuInfoIndexReq          = sys.MenuInfoIndexReq
	MenuInfoIndexResp         = sys.MenuInfoIndexResp
	ModuleInfo                = sys.ModuleInfo
	ModuleInfoIndexReq        = sys.ModuleInfoIndexReq
	ModuleInfoIndexResp       = sys.ModuleInfoIndexResp
	OperLogCreateReq          = sys.OperLogCreateReq
	OperLogIndexReq           = sys.OperLogIndexReq
	OperLogIndexResp          = sys.OperLogIndexResp
	OperLogInfo               = sys.OperLogInfo
	PageInfo                  = sys.PageInfo
	PageInfo_OrderBy          = sys.PageInfo_OrderBy
	Point                     = sys.Point
	ProjectInfo               = sys.ProjectInfo
	ProjectInfoIndexReq       = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp      = sys.ProjectInfoIndexResp
	ProjectWithID             = sys.ProjectWithID
	Response                  = sys.Response
	RoleApiAuthReq            = sys.RoleApiAuthReq
	RoleApiIndexReq           = sys.RoleApiIndexReq
	RoleApiIndexResp          = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq     = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq           = sys.RoleAppIndexReq
	RoleAppIndexResp          = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq     = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq          = sys.RoleAppUpdateReq
	RoleInfo                  = sys.RoleInfo
	RoleInfoIndexReq          = sys.RoleInfoIndexReq
	RoleInfoIndexResp         = sys.RoleInfoIndexResp
	RoleMenuIndexReq          = sys.RoleMenuIndexReq
	RoleMenuIndexResp         = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq    = sys.RoleMenuMultiUpdateReq
	RoleModuleIndexReq        = sys.RoleModuleIndexReq
	RoleModuleIndexResp       = sys.RoleModuleIndexResp
	RoleModuleMultiUpdateReq  = sys.RoleModuleMultiUpdateReq
	TenantApiInfo             = sys.TenantApiInfo
	TenantAppApiIndexReq      = sys.TenantAppApiIndexReq
	TenantAppApiIndexResp     = sys.TenantAppApiIndexResp
	TenantAppCreateReq        = sys.TenantAppCreateReq
	TenantAppIndexReq         = sys.TenantAppIndexReq
	TenantAppIndexResp        = sys.TenantAppIndexResp
	TenantAppMenu             = sys.TenantAppMenu
	TenantAppMenuIndexReq     = sys.TenantAppMenuIndexReq
	TenantAppMenuIndexResp    = sys.TenantAppMenuIndexResp
	TenantAppModule           = sys.TenantAppModule
	TenantAppMultiUpdateReq   = sys.TenantAppMultiUpdateReq
	TenantAppWithIDOrCode     = sys.TenantAppWithIDOrCode
	TenantInfo                = sys.TenantInfo
	TenantInfoCreateReq       = sys.TenantInfoCreateReq
	TenantInfoIndexReq        = sys.TenantInfoIndexReq
	TenantInfoIndexResp       = sys.TenantInfoIndexResp
	TenantModuleCreateReq     = sys.TenantModuleCreateReq
	TenantModuleIndexReq      = sys.TenantModuleIndexReq
	TenantModuleIndexResp     = sys.TenantModuleIndexResp
	TenantModuleWithIDOrCode  = sys.TenantModuleWithIDOrCode
	UserArea                  = sys.UserArea
	UserAreaIndexReq          = sys.UserAreaIndexReq
	UserAreaIndexResp         = sys.UserAreaIndexResp
	UserAreaMultiUpdateReq    = sys.UserAreaMultiUpdateReq
	UserCaptchaReq            = sys.UserCaptchaReq
	UserCaptchaResp           = sys.UserCaptchaResp
	UserCheckTokenReq         = sys.UserCheckTokenReq
	UserCheckTokenResp        = sys.UserCheckTokenResp
	UserCreateResp            = sys.UserCreateResp
	UserForgetPwdReq          = sys.UserForgetPwdReq
	UserInfo                  = sys.UserInfo
	UserInfoCreateReq         = sys.UserInfoCreateReq
	UserInfoDeleteReq         = sys.UserInfoDeleteReq
	UserInfoIndexReq          = sys.UserInfoIndexReq
	UserInfoIndexResp         = sys.UserInfoIndexResp
	UserInfoReadReq           = sys.UserInfoReadReq
	UserLoginReq              = sys.UserLoginReq
	UserLoginResp             = sys.UserLoginResp
	UserProject               = sys.UserProject
	UserProjectIndexReq       = sys.UserProjectIndexReq
	UserProjectIndexResp      = sys.UserProjectIndexResp
	UserProjectMultiUpdateReq = sys.UserProjectMultiUpdateReq
	UserRegisterReq           = sys.UserRegisterReq
	UserRegisterResp          = sys.UserRegisterResp
	UserRoleIndexReq          = sys.UserRoleIndexReq
	UserRoleIndexResp         = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq    = sys.UserRoleMultiUpdateReq
	WithAppCodeID             = sys.WithAppCodeID
	WithID                    = sys.WithID
	WithIDCode                = sys.WithIDCode

	Log interface {
		LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error)
		OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error)
		LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error)
		OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultLog struct {
		cli zrpc.Client
	}

	directLog struct {
		svcCtx *svc.ServiceContext
		svr    sys.LogServer
	}
)

func NewLog(cli zrpc.Client) Log {
	return &defaultLog{
		cli: cli,
	}
}

func NewDirectLog(svcCtx *svc.ServiceContext, svr sys.LogServer) Log {
	return &directLog{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultLog) LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.LoginLogIndex(ctx, in, opts...)
}

func (d *directLog) LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error) {
	return d.svr.LoginLogIndex(ctx, in)
}

func (m *defaultLog) OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.OperLogIndex(ctx, in, opts...)
}

func (d *directLog) OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error) {
	return d.svr.OperLogIndex(ctx, in)
}

func (m *defaultLog) LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.LoginLogCreate(ctx, in, opts...)
}

func (d *directLog) LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.LoginLogCreate(ctx, in)
}

func (m *defaultLog) OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.OperLogCreate(ctx, in, opts...)
}

func (d *directLog) OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.OperLogCreate(ctx, in)
}
