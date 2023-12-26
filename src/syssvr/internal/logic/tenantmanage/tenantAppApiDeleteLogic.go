package tenantmanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type TenantAppApiDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTenantAppApiDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TenantAppApiDeleteLogic {
	return &TenantAppApiDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TenantAppApiDeleteLogic) TenantAppApiDelete(in *sys.WithAppCodeID) (*sys.Response, error) {
	if err := ctxs.IsRoot(l.ctx); err != nil {
		return nil, err
	}
	ctxs.GetUserCtx(l.ctx).AllTenant = true
	defer func() {
		ctxs.GetUserCtx(l.ctx).AllTenant = false
	}()
	err := relationDB.NewTenantAppApiRepo(l.ctx).Delete(l.ctx, in.Id)
	return &sys.Response{}, err
}
