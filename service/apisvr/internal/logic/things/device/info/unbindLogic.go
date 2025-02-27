package info

import (
	"context"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnbindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindLogic {
	return &UnbindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnbindLogic) Unbind(req *types.DeviceInfoUnbindReq) error {
	_, err := l.svcCtx.DeviceM.DeviceInfoUnbind(l.ctx, utils.Copy[dm.DeviceInfoUnbindReq](req))

	return err
}
