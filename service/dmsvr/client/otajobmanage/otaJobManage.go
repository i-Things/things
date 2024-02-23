// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package otajobmanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                      = dm.ActionRespReq
	ActionSendReq                      = dm.ActionSendReq
	ActionSendResp                     = dm.ActionSendResp
	CommonSchemaCreateReq              = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq               = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp              = dm.CommonSchemaIndexResp
	CommonSchemaInfo                   = dm.CommonSchemaInfo
	CommonSchemaUpdateReq              = dm.CommonSchemaUpdateReq
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceCountInfo                    = dm.DeviceCountInfo
	DeviceCountReq                     = dm.DeviceCountReq
	DeviceCountResp                    = dm.DeviceCountResp
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCount                    = dm.DeviceInfoCount
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq           = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                  = dm.DeviceInfoReadReq
	DeviceTypeCountReq                 = dm.DeviceTypeCountReq
	DeviceTypeCountResp                = dm.DeviceTypeCountResp
	DynamicUpgradeJobReq               = dm.DynamicUpgradeJobReq
	Empty                              = dm.Empty
	EventIndex                         = dm.EventIndex
	EventIndexResp                     = dm.EventIndexResp
	EventLogIndexReq                   = dm.EventLogIndexReq
	Firmware                           = dm.Firmware
	FirmwareFile                       = dm.FirmwareFile
	FirmwareInfo                       = dm.FirmwareInfo
	FirmwareInfoDeleteReq              = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp             = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq               = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp              = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp               = dm.FirmwareInfoReadResp
	FirmwareResp                       = dm.FirmwareResp
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq            = dm.GroupDeviceMultiSaveReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndex                        = dm.HubLogIndex
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	JobReq                             = dm.JobReq
	OTAModuleDeleteReq                 = dm.OTAModuleDeleteReq
	OTAModuleDetail                    = dm.OTAModuleDetail
	OTAModuleIndexReq                  = dm.OTAModuleIndexReq
	OTAModuleIndexResp                 = dm.OTAModuleIndexResp
	OTAModuleReq                       = dm.OTAModuleReq
	OTAModuleVersionsIndexResp         = dm.OTAModuleVersionsIndexResp
	OTATaskByDeviceCancelReq           = dm.OTATaskByDeviceCancelReq
	OTATaskByDeviceNameReq             = dm.OTATaskByDeviceNameReq
	OTATaskByJobCancelReq              = dm.OTATaskByJobCancelReq
	OTATaskByJobIndexReq               = dm.OTATaskByJobIndexReq
	OTATaskConfirmReq                  = dm.OTATaskConfirmReq
	OTATaskReUpgradeReq                = dm.OTATaskReUpgradeReq
	OTAUnfinishedTaskByDeviceIndexReq  = dm.OTAUnfinishedTaskByDeviceIndexReq
	OTAUnfinishedTaskByDeviceIndexResp = dm.OTAUnfinishedTaskByDeviceIndexResp
	OtaCommonResp                      = dm.OtaCommonResp
	OtaFirmwareCreateReq               = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq               = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq           = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp          = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile                    = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq            = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp           = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                 = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq                = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp               = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo                    = dm.OtaFirmwareInfo
	OtaFirmwareReadReq                 = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp                = dm.OtaFirmwareReadResp
	OtaFirmwareResp                    = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq               = dm.OtaFirmwareUpdateReq
	OtaFirmwareVerifyReq               = dm.OtaFirmwareVerifyReq
	OtaJobByDeviceIndexReq             = dm.OtaJobByDeviceIndexReq
	OtaJobByFirmwareIndexReq           = dm.OtaJobByFirmwareIndexReq
	OtaJobInfo                         = dm.OtaJobInfo
	OtaJobInfoIndexResp                = dm.OtaJobInfoIndexResp
	OtaModuleInfo                      = dm.OtaModuleInfo
	OtaPageInfo                        = dm.OtaPageInfo
	OtaPromptIndexReq                  = dm.OtaPromptIndexReq
	OtaPromptIndexResp                 = dm.OtaPromptIndexResp
	OtaTaskBatchReq                    = dm.OtaTaskBatchReq
	OtaTaskBatchResp                   = dm.OtaTaskBatchResp
	OtaTaskByJobIndexResp              = dm.OtaTaskByJobIndexResp
	OtaTaskCancleReq                   = dm.OtaTaskCancleReq
	OtaTaskCreatResp                   = dm.OtaTaskCreatResp
	OtaTaskCreateReq                   = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq             = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq              = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp             = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo                  = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq            = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq               = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq                    = dm.OtaTaskIndexReq
	OtaTaskIndexResp                   = dm.OtaTaskIndexResp
	OtaTaskInfo                        = dm.OtaTaskInfo
	OtaTaskReadReq                     = dm.OtaTaskReadReq
	OtaTaskReadResp                    = dm.OtaTaskReadResp
	OtaUpTaskInfo                      = dm.OtaUpTaskInfo
	PageInfo                           = dm.PageInfo
	PageInfo_OrderBy                   = dm.PageInfo_OrderBy
	Point                              = dm.Point
	ProductCategory                    = dm.ProductCategory
	ProductCategoryIndexReq            = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp           = dm.ProductCategoryIndexResp
	ProductCategoryReadReq             = dm.ProductCategoryReadReq
	ProductCustom                      = dm.ProductCustom
	ProductCustomReadReq               = dm.ProductCustomReadReq
	ProductInfo                        = dm.ProductInfo
	ProductInfoDeleteReq               = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                = dm.ProductInfoIndexReq
	ProductInfoIndexResp               = dm.ProductInfoIndexResp
	ProductInfoReadReq                 = dm.ProductInfoReadReq
	ProductRemoteConfig                = dm.ProductRemoteConfig
	ProductSchemaCreateReq             = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq             = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq              = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp             = dm.ProductSchemaIndexResp
	ProductSchemaInfo                  = dm.ProductSchemaInfo
	ProductSchemaTslImportReq          = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq            = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp           = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq             = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq        = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp       = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg             = dm.PropertyControlSendMsg
	PropertyControlSendReq             = dm.PropertyControlSendReq
	PropertyControlSendResp            = dm.PropertyControlSendResp
	PropertyGetReportSendReq           = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp          = dm.PropertyGetReportSendResp
	PropertyIndex                      = dm.PropertyIndex
	PropertyIndexResp                  = dm.PropertyIndexResp
	PropertyLatestIndexReq             = dm.PropertyLatestIndexReq
	PropertyLogIndexReq                = dm.PropertyLogIndexReq
	ProtocolConfigField                = dm.ProtocolConfigField
	ProtocolConfigInfo                 = dm.ProtocolConfigInfo
	ProtocolInfo                       = dm.ProtocolInfo
	ProtocolInfoIndexReq               = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp              = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq              = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq               = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp              = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq            = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp           = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq             = dm.RemoteConfigPushAllReq
	RespReadReq                        = dm.RespReadReq
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndex                        = dm.SdkLogIndex
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	Tag                                = dm.Tag
	TimeRange                          = dm.TimeRange
	UpgradeJobResp                     = dm.UpgradeJobResp
	UserDeviceCollectSave              = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq            = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp           = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo                = dm.UserDeviceShareInfo
	UserDeviceShareReadReq             = dm.UserDeviceShareReadReq
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq
	WithID                             = dm.WithID
	WithIDCode                         = dm.WithIDCode

	OTAJobManage interface {
		// 验证升级包
		OtaVerifyJobCreate(ctx context.Context, in *OtaFirmwareVerifyReq, opts ...grpc.CallOption) (*UpgradeJobResp, error)
		// 创建静态升级批次
		OtaStaticUpgradeJobCreate(ctx context.Context, in *StaticUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error)
		// 创建动态升级批次
		OtaDynamicUpgradeJobCreate(ctx context.Context, in *DynamicUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error)
		// 获取升级包下的升级任务批次列表
		OtaJobByFirmwareIndex(ctx context.Context, in *OtaJobByFirmwareIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error)
		// 获取设备所在的升级包升级批次列表
		OtaJobByDeviceIndex(ctx context.Context, in *OtaJobByDeviceIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error)
		// 查询指定升级批次的详情
		OtaJobRead(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*OtaJobInfo, error)
		// 取消动态升级策略
		CancelOTAStrategyByJob(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultOTAJobManage struct {
		cli zrpc.Client
	}

	directOTAJobManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.OTAJobManageServer
	}
)

func NewOTAJobManage(cli zrpc.Client) OTAJobManage {
	return &defaultOTAJobManage{
		cli: cli,
	}
}

func NewDirectOTAJobManage(svcCtx *svc.ServiceContext, svr dm.OTAJobManageServer) OTAJobManage {
	return &directOTAJobManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 验证升级包
func (m *defaultOTAJobManage) OtaVerifyJobCreate(ctx context.Context, in *OtaFirmwareVerifyReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaVerifyJobCreate(ctx, in, opts...)
}

// 验证升级包
func (d *directOTAJobManage) OtaVerifyJobCreate(ctx context.Context, in *OtaFirmwareVerifyReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	return d.svr.OtaVerifyJobCreate(ctx, in)
}

// 创建静态升级批次
func (m *defaultOTAJobManage) OtaStaticUpgradeJobCreate(ctx context.Context, in *StaticUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaStaticUpgradeJobCreate(ctx, in, opts...)
}

// 创建静态升级批次
func (d *directOTAJobManage) OtaStaticUpgradeJobCreate(ctx context.Context, in *StaticUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	return d.svr.OtaStaticUpgradeJobCreate(ctx, in)
}

// 创建动态升级批次
func (m *defaultOTAJobManage) OtaDynamicUpgradeJobCreate(ctx context.Context, in *DynamicUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaDynamicUpgradeJobCreate(ctx, in, opts...)
}

// 创建动态升级批次
func (d *directOTAJobManage) OtaDynamicUpgradeJobCreate(ctx context.Context, in *DynamicUpgradeJobReq, opts ...grpc.CallOption) (*UpgradeJobResp, error) {
	return d.svr.OtaDynamicUpgradeJobCreate(ctx, in)
}

// 获取升级包下的升级任务批次列表
func (m *defaultOTAJobManage) OtaJobByFirmwareIndex(ctx context.Context, in *OtaJobByFirmwareIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaJobByFirmwareIndex(ctx, in, opts...)
}

// 获取升级包下的升级任务批次列表
func (d *directOTAJobManage) OtaJobByFirmwareIndex(ctx context.Context, in *OtaJobByFirmwareIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error) {
	return d.svr.OtaJobByFirmwareIndex(ctx, in)
}

// 获取设备所在的升级包升级批次列表
func (m *defaultOTAJobManage) OtaJobByDeviceIndex(ctx context.Context, in *OtaJobByDeviceIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaJobByDeviceIndex(ctx, in, opts...)
}

// 获取设备所在的升级包升级批次列表
func (d *directOTAJobManage) OtaJobByDeviceIndex(ctx context.Context, in *OtaJobByDeviceIndexReq, opts ...grpc.CallOption) (*OtaJobInfoIndexResp, error) {
	return d.svr.OtaJobByDeviceIndex(ctx, in)
}

// 查询指定升级批次的详情
func (m *defaultOTAJobManage) OtaJobRead(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*OtaJobInfo, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.OtaJobRead(ctx, in, opts...)
}

// 查询指定升级批次的详情
func (d *directOTAJobManage) OtaJobRead(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*OtaJobInfo, error) {
	return d.svr.OtaJobRead(ctx, in)
}

// 取消动态升级策略
func (m *defaultOTAJobManage) CancelOTAStrategyByJob(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewOTAJobManageClient(m.cli.Conn())
	return client.CancelOTAStrategyByJob(ctx, in, opts...)
}

// 取消动态升级策略
func (d *directOTAJobManage) CancelOTAStrategyByJob(ctx context.Context, in *JobReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.CancelOTAStrategyByJob(ctx, in)
}
