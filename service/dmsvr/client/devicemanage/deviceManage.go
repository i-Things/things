// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package devicemanage

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

	DeviceManage interface {
		// 鉴定是否是root账号(提供给mqtt broker)
		RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Empty, error)
		// 新增设备
		DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error)
		// 更新设备
		DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error)
		// 删除设备
		DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取设备信息列表
		DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error)
		// 批量更新设备状态
		DeviceInfoMultiUpdate(ctx context.Context, in *DeviceInfoMultiUpdateReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取设备信息详情
		DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error)
		// 绑定网关下子设备设备
		DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取绑定信息的设备信息列表
		DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error)
		// 删除网关下子设备
		DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 设备计数
		DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCount, error)
		// 设备类型
		DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error)
		DeviceCount(ctx context.Context, in *DeviceCountReq, opts ...grpc.CallOption) (*DeviceCountResp, error)
	}

	defaultDeviceManage struct {
		cli zrpc.Client
	}

	directDeviceManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceManageServer
	}
)

func NewDeviceManage(cli zrpc.Client) DeviceManage {
	return &defaultDeviceManage{
		cli: cli,
	}
}

func NewDirectDeviceManage(svcCtx *svc.ServiceContext, svr dm.DeviceManageServer) DeviceManage {
	return &directDeviceManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 鉴定是否是root账号(提供给mqtt broker)
func (m *defaultDeviceManage) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.RootCheck(ctx, in, opts...)
}

// 鉴定是否是root账号(提供给mqtt broker)
func (d *directDeviceManage) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.RootCheck(ctx, in)
}

// 新增设备
func (m *defaultDeviceManage) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoCreate(ctx, in, opts...)
}

// 新增设备
func (d *directDeviceManage) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceInfoCreate(ctx, in)
}

// 更新设备
func (m *defaultDeviceManage) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoUpdate(ctx, in, opts...)
}

// 更新设备
func (d *directDeviceManage) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceInfoUpdate(ctx, in)
}

// 删除设备
func (m *defaultDeviceManage) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoDelete(ctx, in, opts...)
}

// 删除设备
func (d *directDeviceManage) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceInfoDelete(ctx, in)
}

// 获取设备信息列表
func (m *defaultDeviceManage) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoIndex(ctx, in, opts...)
}

// 获取设备信息列表
func (d *directDeviceManage) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	return d.svr.DeviceInfoIndex(ctx, in)
}

// 批量更新设备状态
func (m *defaultDeviceManage) DeviceInfoMultiUpdate(ctx context.Context, in *DeviceInfoMultiUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoMultiUpdate(ctx, in, opts...)
}

// 批量更新设备状态
func (d *directDeviceManage) DeviceInfoMultiUpdate(ctx context.Context, in *DeviceInfoMultiUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceInfoMultiUpdate(ctx, in)
}

// 获取设备信息详情
func (m *defaultDeviceManage) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoRead(ctx, in, opts...)
}

// 获取设备信息详情
func (d *directDeviceManage) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	return d.svr.DeviceInfoRead(ctx, in)
}

// 绑定网关下子设备设备
func (m *defaultDeviceManage) DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayMultiCreate(ctx, in, opts...)
}

// 绑定网关下子设备设备
func (d *directDeviceManage) DeviceGatewayMultiCreate(ctx context.Context, in *DeviceGatewayMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceGatewayMultiCreate(ctx, in)
}

// 获取绑定信息的设备信息列表
func (m *defaultDeviceManage) DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayIndex(ctx, in, opts...)
}

// 获取绑定信息的设备信息列表
func (d *directDeviceManage) DeviceGatewayIndex(ctx context.Context, in *DeviceGatewayIndexReq, opts ...grpc.CallOption) (*DeviceGatewayIndexResp, error) {
	return d.svr.DeviceGatewayIndex(ctx, in)
}

// 删除网关下子设备
func (m *defaultDeviceManage) DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceGatewayMultiDelete(ctx, in, opts...)
}

// 删除网关下子设备
func (d *directDeviceManage) DeviceGatewayMultiDelete(ctx context.Context, in *DeviceGatewayMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.DeviceGatewayMultiDelete(ctx, in)
}

// 设备计数
func (m *defaultDeviceManage) DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCount, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceInfoCount(ctx, in, opts...)
}

// 设备计数
func (d *directDeviceManage) DeviceInfoCount(ctx context.Context, in *DeviceInfoCountReq, opts ...grpc.CallOption) (*DeviceInfoCount, error) {
	return d.svr.DeviceInfoCount(ctx, in)
}

// 设备类型
func (m *defaultDeviceManage) DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceTypeCount(ctx, in, opts...)
}

// 设备类型
func (d *directDeviceManage) DeviceTypeCount(ctx context.Context, in *DeviceTypeCountReq, opts ...grpc.CallOption) (*DeviceTypeCountResp, error) {
	return d.svr.DeviceTypeCount(ctx, in)
}

func (m *defaultDeviceManage) DeviceCount(ctx context.Context, in *DeviceCountReq, opts ...grpc.CallOption) (*DeviceCountResp, error) {
	client := dm.NewDeviceManageClient(m.cli.Conn())
	return client.DeviceCount(ctx, in, opts...)
}

func (d *directDeviceManage) DeviceCount(ctx context.Context, in *DeviceCountReq, opts ...grpc.CallOption) (*DeviceCountResp, error) {
	return d.svr.DeviceCount(ctx, in)
}
