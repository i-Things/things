package rulelogic

import (
	"context"
	"github.com/i-Things/things/src/udsvr/internal/logic"
	"github.com/i-Things/things/src/udsvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/udsvr/internal/svc"
	"github.com/i-Things/things/src/udsvr/pb/ud"

	"github.com/zeromicro/go-zero/core/logx"
)

type SceneInfoIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSceneInfoIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SceneInfoIndexLogic {
	return &SceneInfoIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SceneInfoIndexLogic) SceneInfoIndex(in *ud.SceneInfoIndexReq) (*ud.SceneInfoIndexResp, error) {
	f := relationDB.SceneInfoFilter{AreaID: in.AreaID, Status: in.Status, Name: in.Name}
	list, err := relationDB.NewSceneInfoRepo(l.ctx).FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	total, err := relationDB.NewSceneInfoRepo(l.ctx).CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}

	return &ud.SceneInfoIndexResp{List: PoToSceneInfoPbs(list), Total: total}, nil
}