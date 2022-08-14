package device

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoUpdateLogic {
	return &InfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoUpdateLogic) InfoUpdate(req *types.DeviceInfoUpdateReq) error {
	dmReq := &dm.DeviceInfo{
		ProductID:  req.ProductID,  //产品id 只读
		DeviceName: req.DeviceName, //设备名称 读写
		LogLevel:   req.LogLevel,   // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试  读写
		Tags:       toTagsMap(req.Tags),
	}
	_, err := l.svcCtx.DmRpc.DeviceInfoUpdate(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.ManageDevice|req=%v|err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}