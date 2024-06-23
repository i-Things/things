// Code generated by goctl. DO NOT EDIT.
// Source: ud.proto

package rule

import (
	"context"

	"github.com/i-Things/things/service/udsvr/internal/svc"
	"github.com/i-Things/things/service/udsvr/pb/ud"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlarmInfo                 = ud.AlarmInfo
	AlarmInfoIndexReq         = ud.AlarmInfoIndexReq
	AlarmInfoIndexResp        = ud.AlarmInfoIndexResp
	AlarmNotify               = ud.AlarmNotify
	AlarmRecord               = ud.AlarmRecord
	AlarmRecordCreateReq      = ud.AlarmRecordCreateReq
	AlarmRecordDealReq        = ud.AlarmRecordDealReq
	AlarmRecordIndexReq       = ud.AlarmRecordIndexReq
	AlarmRecordIndexResp      = ud.AlarmRecordIndexResp
	AlarmSceneDeleteReq       = ud.AlarmSceneDeleteReq
	AlarmSceneIndexReq        = ud.AlarmSceneIndexReq
	AlarmSceneIndexResp       = ud.AlarmSceneIndexResp
	AlarmSceneMultiSaveReq    = ud.AlarmSceneMultiSaveReq
	DeviceCore                = ud.DeviceCore
	Empty                     = ud.Empty
	PageInfo                  = ud.PageInfo
	SceneFlowInfo             = ud.SceneFlowInfo
	SceneInfo                 = ud.SceneInfo
	SceneInfoIndexReq         = ud.SceneInfoIndexReq
	SceneInfoIndexResp        = ud.SceneInfoIndexResp
	SceneLog                  = ud.SceneLog
	SceneLogAction            = ud.SceneLogAction
	SceneLogActionAlarm       = ud.SceneLogActionAlarm
	SceneLogActionDevice      = ud.SceneLogActionDevice
	SceneLogActionDeviceValue = ud.SceneLogActionDeviceValue
	SceneLogIndexReq          = ud.SceneLogIndexReq
	SceneLogIndexResp         = ud.SceneLogIndexResp
	SceneLogTrigger           = ud.SceneLogTrigger
	SceneLogTriggerDevice     = ud.SceneLogTriggerDevice
	TimeRange                 = ud.TimeRange
	WithID                    = ud.WithID

	Rule interface {
		// 场景
		SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*WithID, error)
		SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Empty, error)
		SceneInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error)
		SceneInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*SceneInfo, error)
		SceneManuallyTrigger(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		SceneLogIndex(ctx context.Context, in *SceneLogIndexReq, opts ...grpc.CallOption) (*SceneLogIndexResp, error)
		AlarmInfoCreate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*WithID, error)
		AlarmInfoUpdate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*Empty, error)
		AlarmInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		AlarmInfoIndex(ctx context.Context, in *AlarmInfoIndexReq, opts ...grpc.CallOption) (*AlarmInfoIndexResp, error)
		AlarmInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*AlarmInfo, error)
		// 告警关联场景联动
		AlarmSceneMultiCreate(ctx context.Context, in *AlarmSceneMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		AlarmSceneDelete(ctx context.Context, in *AlarmSceneDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		AlarmSceneIndex(ctx context.Context, in *AlarmSceneIndexReq, opts ...grpc.CallOption) (*AlarmSceneIndexResp, error)
		// 告警记录
		AlarmRecordIndex(ctx context.Context, in *AlarmRecordIndexReq, opts ...grpc.CallOption) (*AlarmRecordIndexResp, error)
		AlarmRecordCreate(ctx context.Context, in *AlarmRecordCreateReq, opts ...grpc.CallOption) (*Empty, error)
		AlarmRecordDeal(ctx context.Context, in *AlarmRecordDealReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultRule struct {
		cli zrpc.Client
	}

	directRule struct {
		svcCtx *svc.ServiceContext
		svr    ud.RuleServer
	}
)

func NewRule(cli zrpc.Client) Rule {
	return &defaultRule{
		cli: cli,
	}
}

func NewDirectRule(svcCtx *svc.ServiceContext, svr ud.RuleServer) Rule {
	return &directRule{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 场景
func (m *defaultRule) SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneInfoCreate(ctx, in, opts...)
}

// 场景
func (d *directRule) SceneInfoCreate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.SceneInfoCreate(ctx, in)
}

func (m *defaultRule) SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneInfoUpdate(ctx, in, opts...)
}

func (d *directRule) SceneInfoUpdate(ctx context.Context, in *SceneInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.SceneInfoUpdate(ctx, in)
}

func (m *defaultRule) SceneInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneInfoDelete(ctx, in, opts...)
}

func (d *directRule) SceneInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.SceneInfoDelete(ctx, in)
}

func (m *defaultRule) SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneInfoIndex(ctx, in, opts...)
}

func (d *directRule) SceneInfoIndex(ctx context.Context, in *SceneInfoIndexReq, opts ...grpc.CallOption) (*SceneInfoIndexResp, error) {
	return d.svr.SceneInfoIndex(ctx, in)
}

func (m *defaultRule) SceneInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*SceneInfo, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneInfoRead(ctx, in, opts...)
}

func (d *directRule) SceneInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*SceneInfo, error) {
	return d.svr.SceneInfoRead(ctx, in)
}

func (m *defaultRule) SceneManuallyTrigger(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneManuallyTrigger(ctx, in, opts...)
}

func (d *directRule) SceneManuallyTrigger(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.SceneManuallyTrigger(ctx, in)
}

func (m *defaultRule) SceneLogIndex(ctx context.Context, in *SceneLogIndexReq, opts ...grpc.CallOption) (*SceneLogIndexResp, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.SceneLogIndex(ctx, in, opts...)
}

func (d *directRule) SceneLogIndex(ctx context.Context, in *SceneLogIndexReq, opts ...grpc.CallOption) (*SceneLogIndexResp, error) {
	return d.svr.SceneLogIndex(ctx, in)
}

func (m *defaultRule) AlarmInfoCreate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmInfoCreate(ctx, in, opts...)
}

func (d *directRule) AlarmInfoCreate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.AlarmInfoCreate(ctx, in)
}

func (m *defaultRule) AlarmInfoUpdate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmInfoUpdate(ctx, in, opts...)
}

func (d *directRule) AlarmInfoUpdate(ctx context.Context, in *AlarmInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmInfoUpdate(ctx, in)
}

func (m *defaultRule) AlarmInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmInfoDelete(ctx, in, opts...)
}

func (d *directRule) AlarmInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmInfoDelete(ctx, in)
}

func (m *defaultRule) AlarmInfoIndex(ctx context.Context, in *AlarmInfoIndexReq, opts ...grpc.CallOption) (*AlarmInfoIndexResp, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmInfoIndex(ctx, in, opts...)
}

func (d *directRule) AlarmInfoIndex(ctx context.Context, in *AlarmInfoIndexReq, opts ...grpc.CallOption) (*AlarmInfoIndexResp, error) {
	return d.svr.AlarmInfoIndex(ctx, in)
}

func (m *defaultRule) AlarmInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*AlarmInfo, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmInfoRead(ctx, in, opts...)
}

func (d *directRule) AlarmInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*AlarmInfo, error) {
	return d.svr.AlarmInfoRead(ctx, in)
}

// 告警关联场景联动
func (m *defaultRule) AlarmSceneMultiCreate(ctx context.Context, in *AlarmSceneMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmSceneMultiCreate(ctx, in, opts...)
}

// 告警关联场景联动
func (d *directRule) AlarmSceneMultiCreate(ctx context.Context, in *AlarmSceneMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmSceneMultiCreate(ctx, in)
}

func (m *defaultRule) AlarmSceneDelete(ctx context.Context, in *AlarmSceneDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmSceneDelete(ctx, in, opts...)
}

func (d *directRule) AlarmSceneDelete(ctx context.Context, in *AlarmSceneDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmSceneDelete(ctx, in)
}

func (m *defaultRule) AlarmSceneIndex(ctx context.Context, in *AlarmSceneIndexReq, opts ...grpc.CallOption) (*AlarmSceneIndexResp, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmSceneIndex(ctx, in, opts...)
}

func (d *directRule) AlarmSceneIndex(ctx context.Context, in *AlarmSceneIndexReq, opts ...grpc.CallOption) (*AlarmSceneIndexResp, error) {
	return d.svr.AlarmSceneIndex(ctx, in)
}

// 告警记录
func (m *defaultRule) AlarmRecordIndex(ctx context.Context, in *AlarmRecordIndexReq, opts ...grpc.CallOption) (*AlarmRecordIndexResp, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmRecordIndex(ctx, in, opts...)
}

// 告警记录
func (d *directRule) AlarmRecordIndex(ctx context.Context, in *AlarmRecordIndexReq, opts ...grpc.CallOption) (*AlarmRecordIndexResp, error) {
	return d.svr.AlarmRecordIndex(ctx, in)
}

func (m *defaultRule) AlarmRecordCreate(ctx context.Context, in *AlarmRecordCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmRecordCreate(ctx, in, opts...)
}

func (d *directRule) AlarmRecordCreate(ctx context.Context, in *AlarmRecordCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmRecordCreate(ctx, in)
}

func (m *defaultRule) AlarmRecordDeal(ctx context.Context, in *AlarmRecordDealReq, opts ...grpc.CallOption) (*Empty, error) {
	client := ud.NewRuleClient(m.cli.Conn())
	return client.AlarmRecordDeal(ctx, in, opts...)
}

func (d *directRule) AlarmRecordDeal(ctx context.Context, in *AlarmRecordDealReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.AlarmRecordDeal(ctx, in)
}
