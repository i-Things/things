package share

import (
	"context"
	"github.com/i-Things/things/service/apisvr/internal/logic/things"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.UserDeviceShareReadReq) error {
	_, err := l.svcCtx.UserDevice.UserDeviceShareDelete(l.ctx, &dm.UserDeviceShareReadReq{
		Id:     req.ID,
		Device: things.ToDmDeviceCorePb(req.Device),
	})

	return err
}
