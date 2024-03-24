package startup

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/i-Things/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/i-Things/share/caches"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/eventBus"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/event/deviceMsgEvent"
	"github.com/i-Things/things/service/dmsvr/internal/event/otaEvent"
	"github.com/i-Things/things/service/dmsvr/internal/event/serverEvent"
	"github.com/i-Things/things/service/dmsvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/internal/repo/event/subscribe/server"
	"github.com/i-Things/things/service/dmsvr/internal/repo/event/subscribe/subDev"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
)

func Init(svcCtx *svc.ServiceContext) {
	InitCache(svcCtx)
	TimerInit(svcCtx)
	InitSubscribe(svcCtx)
	InitEventBus(svcCtx)
}
func InitSubscribe(svcCtx *svc.ServiceContext) {
	{
		cli, err := subDev.NewSubDev(svcCtx.Config.Event)
		logx.Must(err)
		err = cli.Subscribe(func(ctx context.Context) subDev.InnerSubEvent {
			return deviceMsgEvent.NewDeviceMsgHandle(ctx, svcCtx)
		})
		logx.Must(err)
	}
	{
		cli, err := server.NewServer(svcCtx.Config.Event)
		logx.Must(err)
		err = cli.Subscribe(func(ctx context.Context) server.ServerHandle {
			return serverEvent.NewServerHandle(ctx, svcCtx)
		})
		logx.Must(err)
	}
}

func InitCache(svcCtx *svc.ServiceContext) {
	productCache, err := caches.NewCache(caches.CacheConfig[dm.ProductInfo]{
		KeyType:   eventBus.ServerCacheKeyDmProduct,
		FastEvent: svcCtx.FastEvent,
		GetData: func(ctx context.Context, key string) (*dm.ProductInfo, error) {
			db := relationDB.NewProductInfoRepo(ctx)
			pi, err := db.FindOneByFilter(ctx, relationDB.ProductFilter{
				ProductIDs: []string{key}, WithProtocol: true, WithCategory: true})
			pb := logic.ToProductInfo(ctx, svcCtx, pi)
			return pb, err
		},
		ExpireTime: 10 * time.Minute,
	})
	logx.Must(err)
	svcCtx.ProductCache = productCache
	deviceCache, err := caches.NewCache(caches.CacheConfig[dm.DeviceInfo]{
		KeyType:   eventBus.ServerCacheKeyDmDevice,
		FastEvent: svcCtx.FastEvent,
		GetData: func(ctx context.Context, key string) (*dm.DeviceInfo, error) {
			db := relationDB.NewDeviceInfoRepo(ctx)
			productID, deviceName, _ := strings.Cut(key, ":")
			di, err := db.FindOneByFilter(ctx, relationDB.DeviceFilter{
				ProductID: productID, DeviceName: deviceName})
			pb := logic.ToDeviceInfo(di)
			return pb, err
		},
		ExpireTime: 10 * time.Minute,
	})
	logx.Must(err)
	svcCtx.DeviceCache = deviceCache
}

func InitEventBus(svcCtx *svc.ServiceContext) {
	err := svcCtx.FastEvent.Subscribe(eventBus.DmDeviceInfoDelete, func(ctx context.Context, t time.Time, body []byte) error {
		var value devices.Core
		err := json.Unmarshal(body, &value)
		if err != nil {
			return err
		}
		err = relationDB.NewGroupDeviceRepo(ctx).DeleteByFilter(ctx, relationDB.GroupDeviceFilter{
			ProductID:  value.ProductID,
			DeviceName: value.DeviceName,
		})
		logx.WithContext(ctx).Infof("DeviceGroupHandle value:%v err:%v", utils.Fmt(value), err)
		return err
	})
	logx.Must(err)
	err = svcCtx.FastEvent.Subscribe(eventBus.DmOtaJobDelayRun, func(ctx context.Context, t time.Time, body []byte) error {
		return otaEvent.NewOtaEvent(svcCtx, ctx).JobDelayRun(cast.ToInt64(string(body)))
	})
	logx.Must(err)
	err = svcCtx.FastEvent.Subscribe(eventBus.DmOtaDeviceUpgradePush, func(ctx context.Context, t time.Time, body []byte) error {
		return otaEvent.NewOtaEvent(svcCtx, ctx).DeviceUpgradePush()
	})
	logx.Must(err)
	err = svcCtx.FastEvent.Start()
	logx.Must(err)
}

func TimerInit(svcCtx *svc.ServiceContext) {
	ctx := context.Background()
	_, err := svcCtx.TimedM.TaskInfoCreate(ctx, &timedmanage.TaskInfo{
		GroupCode: def.TimedIThingsQueueGroupCode,                                              //组编码
		Type:      1,                                                                           //任务类型 1 定时任务 2 延时任务
		Name:      "iThings ota升级定时任务",                                                         // 任务名称
		Code:      "iThingsOtaDeviceUpgradePush",                                               //任务编码
		Params:    fmt.Sprintf(`{"topic":"%s","payload":""}`, eventBus.DmOtaDeviceUpgradePush), // 任务参数,延时任务如果没有传任务参数会拿数据库的参数来执行
		CronExpr:  "@every 5s",                                                                 // cron执行表达式
		Status:    def.StatusWaitRun,                                                           // 状态
		Priority:  3,                                                                           //优先级: 10:critical 最高优先级  3: default 普通优先级 1:low 低优先级
	})
	if err != nil && !errors.Cmp(errors.Fmt(err), errors.Duplicate) {
		logx.Must(err)
	}
}