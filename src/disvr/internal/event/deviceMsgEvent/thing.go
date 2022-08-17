package deviceMsgEvent

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/disvr/internal/domain/service/deviceSend"
	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type ThingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	schema *schema.Model
	topics []string
	dreq   deviceSend.DeviceReq
	dd     deviceMsg.SchemaDataRepo
}

func NewThingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThingLogic {
	return &ThingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThingLogic) initMsg(msg *deviceMsg.PublishMsg) error {
	var err error
	l.schema, err = l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, msg.ProductID)
	if err != nil {
		return errors.Database.AddDetail(err)
	}
	err = utils.Unmarshal(msg.Payload, &l.dreq)
	if err != nil {
		return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
	}
	l.dd = l.svcCtx.SchemaMsgRepo
	l.topics = strings.Split(msg.Topic, "/")
	return nil
}

func (l *ThingLogic) DeviceResp(msg *deviceMsg.PublishMsg, err error, data map[string]any) {
	topic, payload := deviceSend.GenThingDeviceRespData(l.dreq.Method, l.dreq.ClientToken, l.topics, err, data)
	er := l.svcCtx.PubDev.PublishToDev(l.ctx, topic, payload)
	if er != nil {
		l.Errorf("DeviceResp|PublishToDev failure err:%v", er)
		return
	}
	l.Infof("ThingLogic|DeviceResp|topic:%v payload:%v err:%v", topic, string(payload), err)
}

func (l *ThingLogic) HandlePropertyReport(msg *deviceMsg.PublishMsg) error {
	tp, err := l.dreq.VerifyReqParam(l.schema, schema.PROPERTY)
	if err != nil {
		l.DeviceResp(msg, err, nil)
		return err
	} else if len(tp) == 0 {
		err := errors.Parameter.AddDetail("need right param")
		l.DeviceResp(msg, err, nil)
		return err
	}
	params := deviceSend.ToVal(tp)
	timeStamp := l.dreq.GetTimeStamp(msg.Timestamp)
	err = l.dd.InsertPropertiesData(l.ctx, l.schema, msg.ProductID, msg.DeviceName, params, timeStamp)
	if err != nil {
		l.DeviceResp(msg, errors.Database, nil)
		l.Errorf("HandlePropertyReport|InsertPropertyData|err=%+v", err)
		return err
	}
	l.DeviceResp(msg, errors.OK, nil)
	return nil
}

func (l *ThingLogic) HandlePropertyGetStatus(msg *deviceMsg.PublishMsg) error {
	respData := make(map[string]any, len(l.schema.Properties))
	switch l.dreq.Type {
	case deviceSend.Report:
		for id, _ := range l.schema.Property {
			data, err := l.dd.GetPropertyDataByID(l.ctx,
				deviceMsg.FilterOpt{
					Page:       def.PageInfo2{Size: 1},
					ProductID:  msg.ProductID,
					DeviceName: []string{msg.DeviceName},
					DataID:     id})
			if err != nil {
				l.Errorf("HandlePropertyGetStatus|GetPropertyDataByID|get id:%s|err:%s",
					id, err.Error())
				return err
			}
			if len(data) == 0 {
				l.Infof("HandlePropertyGetStatus|GetPropertyDataByID|not find id:%s", id)
				continue
			}
			respData[id] = data[0].Param
		}
	default:
		err := errors.Parameter.AddDetailf("not suppot type :%s", l.dreq.Type)
		l.DeviceResp(msg, err, nil)
		return err
	}
	l.DeviceResp(msg, errors.OK, respData)
	return nil
}

func (l *ThingLogic) HandleProperty(msg *deviceMsg.PublishMsg) error {
	l.Infof("ThingLogic|HandleProperty")
	switch l.dreq.Method {
	case deviceSend.Report, deviceSend.ReportInfo:
		return l.HandlePropertyReport(msg)
	case deviceSend.GetStatus:
		return l.HandlePropertyGetStatus(msg)
	case deviceSend.ControlReply:
		return l.HandleResp(msg)
	default:
		return errors.Method
	}
	return nil
}

func (l *ThingLogic) HandleEvent(msg *deviceMsg.PublishMsg) error {
	l.Infof("ThingLogic|HandleEvent")
	dbData := deviceMsg.EventData{}
	dbData.ID = l.dreq.EventID
	dbData.Type = l.dreq.Type
	if l.dreq.Method != deviceSend.EventPost {
		return errors.Method
	}
	tp, err := l.dreq.VerifyReqParam(l.schema, schema.EVENT)
	if err != nil {
		l.DeviceResp(msg, err, nil)
		return err
	}
	dbData.Params = deviceSend.ToVal(tp)
	dbData.TimeStamp = l.dreq.GetTimeStamp(msg.Timestamp)

	err = l.dd.InsertEventData(l.ctx, msg.ProductID, msg.DeviceName, &dbData)
	if err != nil {
		l.DeviceResp(msg, errors.Database, nil)
		l.Errorf("InsertEventData|err=%+v", err)
		return errors.Database.AddDetail(err)
	}
	l.DeviceResp(msg, errors.OK, nil)

	return nil
}
func (l *ThingLogic) HandleResp(msg *deviceMsg.PublishMsg) error {
	l.Infof("ThingLogic|HandleResp")
	//todo 这里后续需要处理异步获取消息的情况
	return nil
}

func (l *ThingLogic) Handle(msg *deviceMsg.PublishMsg) (err error) {
	l.Infof("ThingLogic|req=%v", msg)
	err = l.initMsg(msg)
	if err != nil {
		return err
	}
	err = func() error {
		if len(l.topics) < 5 || l.topics[1] != "up" {
			return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
		}
		switch l.topics[2] {
		case devices.PropertyMethod: //属性上报
			return l.HandleProperty(msg)
		case devices.EventMethod: //事件上报
			return l.HandleEvent(msg)
		case devices.ActionMethod: //设备响应行为执行结果
			return l.HandleResp(msg)
		default:
			return errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
		}
		return nil
	}()
	l.svcCtx.HubLogRepo.Insert(l.ctx, &deviceMsg.HubLog{
		ProductID:  msg.ProductID,
		Action:     "publish",
		Timestamp:  l.dreq.GetTimeStamp(msg.Timestamp), // 操作时间
		DeviceName: msg.DeviceName,
		TranceID:   utils.TraceIdFromContext(l.ctx),
		RequestID:  l.dreq.ClientToken,
		Content:    string(msg.Payload),
		Topic:      msg.Topic,
		ResultType: errors.Fmt(err).GetCode(),
	})
	return err
}