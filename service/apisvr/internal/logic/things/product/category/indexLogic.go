package category

import (
	"context"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/apisvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.ProductCategoryIndexReq) (resp *types.ProductCategoryIndexResp, err error) {
	dmReq := &dm.ProductCategoryIndexReq{
		Name:     req.Name,
		Page:     logic.ToDmPageRpc(req.Page),
		ParentID: req.ParentID,
	}
	dmResp, err := l.svcCtx.ProductM.ProductCategoryIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceInfo req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.ProductCategory, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		pi := ProductCategoryToApi(v)
		pis = append(pis, pi)
	}
	return &types.ProductCategoryIndexResp{
		Total: dmResp.Total,
		List:  pis,
		Num:   int64(len(pis)),
	}, nil
}