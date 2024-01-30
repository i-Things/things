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
	AreaInfoReadReq           = sys.AreaInfoReadReq
	AreaWithID                = sys.AreaWithID
	AuthApiInfo               = sys.AuthApiInfo
	ConfigResp                = sys.ConfigResp
	DataArea                  = sys.DataArea
	DataAreaIndexReq          = sys.DataAreaIndexReq
	DataAreaIndexResp         = sys.DataAreaIndexResp
	DataAreaMultiDeleteReq    = sys.DataAreaMultiDeleteReq
	DataAreaMultiUpdateReq    = sys.DataAreaMultiUpdateReq
	DataProject               = sys.DataProject
	DataProjectIndexReq       = sys.DataProjectIndexReq
	DataProjectIndexResp      = sys.DataProjectIndexResp
	DataProjectMultiUpdateReq = sys.DataProjectMultiUpdateReq
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
	UserAreaApplyCreateReq    = sys.UserAreaApplyCreateReq
	UserAreaApplyDealReq      = sys.UserAreaApplyDealReq
	UserAreaApplyIndexReq     = sys.UserAreaApplyIndexReq
	UserAreaApplyIndexResp    = sys.UserAreaApplyIndexResp
	UserAreaApplyInfo         = sys.UserAreaApplyInfo
	UserCaptchaReq            = sys.UserCaptchaReq
	UserCaptchaResp           = sys.UserCaptchaResp
	UserChangePwdReq          = sys.UserChangePwdReq
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
	UserRegisterReq           = sys.UserRegisterReq
	UserRegisterResp          = sys.UserRegisterResp
	UserRoleIndexReq          = sys.UserRoleIndexReq
	UserRoleIndexResp         = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq    = sys.UserRoleMultiUpdateReq
	WithAppCodeID             = sys.WithAppCodeID
	WithID                    = sys.WithID
	WithIDCode                = sys.WithIDCode

	TenantManage interface {
		// 新增区域
		TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error)
		// 更新区域
		TenantInfoUpdate(ctx context.Context, in *TenantInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除区域
		TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error)
		// 获取租户信息详情
		TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error)
		// 获取租户信息列表
		TenantInfoIndex(ctx context.Context, in *TenantInfoIndexReq, opts ...grpc.CallOption) (*TenantInfoIndexResp, error)
		TenantAppIndex(ctx context.Context, in *TenantAppIndexReq, opts ...grpc.CallOption) (*TenantAppIndexResp, error)
		TenantAppCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error)
		TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Response, error)
		TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error)
		TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Response, error)
		TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error)
		TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Response, error)
		TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error)
		TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error)
		TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Response, error)
		TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error)
		TenantAppApiCreate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*WithID, error)
		TenantAppApiIndex(ctx context.Context, in *TenantAppApiIndexReq, opts ...grpc.CallOption) (*TenantAppApiIndexResp, error)
		TenantAppApiUpdate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*Response, error)
		TenantAppApiDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error)
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
func (m *defaultTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoCreate(ctx, in, opts...)
}

// 新增区域
func (d *directTenantManage) TenantInfoCreate(ctx context.Context, in *TenantInfoCreateReq, opts ...grpc.CallOption) (*WithID, error) {
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
func (m *defaultTenantManage) TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoDelete(ctx, in, opts...)
}

// 删除区域
func (d *directTenantManage) TenantInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantInfoDelete(ctx, in)
}

// 获取租户信息详情
func (m *defaultTenantManage) TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantInfoRead(ctx, in, opts...)
}

// 获取租户信息详情
func (d *directTenantManage) TenantInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*TenantInfo, error) {
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

func (m *defaultTenantManage) TenantAppCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppDelete(ctx context.Context, in *TenantAppWithIDOrCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleMultiCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleMultiCreate(ctx context.Context, in *TenantAppCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppModuleMultiCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleCreate(ctx context.Context, in *TenantModuleCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppModuleCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleIndex(ctx context.Context, in *TenantModuleIndexReq, opts ...grpc.CallOption) (*TenantModuleIndexResp, error) {
	return d.svr.TenantAppModuleIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppModuleDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppModuleDelete(ctx context.Context, in *TenantModuleWithIDOrCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppModuleDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuCreate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.TenantAppMenuCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuIndex(ctx context.Context, in *TenantAppMenuIndexReq, opts ...grpc.CallOption) (*TenantAppMenuIndexResp, error) {
	return d.svr.TenantAppMenuIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuUpdate(ctx context.Context, in *TenantAppMenu, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppMenuUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppMenuDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppMenuDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppMenuDelete(ctx, in)
}

func (m *defaultTenantManage) TenantAppApiCreate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppApiCreate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppApiCreate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.TenantAppApiCreate(ctx, in)
}

func (m *defaultTenantManage) TenantAppApiIndex(ctx context.Context, in *TenantAppApiIndexReq, opts ...grpc.CallOption) (*TenantAppApiIndexResp, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppApiIndex(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppApiIndex(ctx context.Context, in *TenantAppApiIndexReq, opts ...grpc.CallOption) (*TenantAppApiIndexResp, error) {
	return d.svr.TenantAppApiIndex(ctx, in)
}

func (m *defaultTenantManage) TenantAppApiUpdate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppApiUpdate(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppApiUpdate(ctx context.Context, in *TenantApiInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppApiUpdate(ctx, in)
}

func (m *defaultTenantManage) TenantAppApiDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewTenantManageClient(m.cli.Conn())
	return client.TenantAppApiDelete(ctx, in, opts...)
}

func (d *directTenantManage) TenantAppApiDelete(ctx context.Context, in *WithAppCodeID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.TenantAppApiDelete(ctx, in)
}
