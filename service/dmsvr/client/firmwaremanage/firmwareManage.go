// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package firmwaremanage

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

	FirmwareManage interface {
		// 新增固件升级包
		FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error)
		FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error)
		FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error)
		FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error)
		FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error)
		// 附件信息更新
		OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error)
		// 附件列表搜索
		OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error)
		// 获取固件包对应设备版本列表
		OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error)
	}

	defaultFirmwareManage struct {
		cli zrpc.Client
	}

	directFirmwareManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.FirmwareManageServer
	}
)

func NewFirmwareManage(cli zrpc.Client) FirmwareManage {
	return &defaultFirmwareManage{
		cli: cli,
	}
}

func NewDirectFirmwareManage(svcCtx *svc.ServiceContext, svr dm.FirmwareManageServer) FirmwareManage {
	return &directFirmwareManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增固件升级包
func (m *defaultFirmwareManage) FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoCreate(ctx, in, opts...)
}

// 新增固件升级包
func (d *directFirmwareManage) FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error) {
	return d.svr.FirmwareInfoCreate(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoUpdate(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.FirmwareInfoUpdate(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoDelete(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error) {
	return d.svr.FirmwareInfoDelete(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoIndex(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error) {
	return d.svr.FirmwareInfoIndex(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoRead(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error) {
	return d.svr.FirmwareInfoRead(ctx, in)
}

// 附件信息更新
func (m *defaultFirmwareManage) OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareFileUpdate(ctx, in, opts...)
}

// 附件信息更新
func (d *directFirmwareManage) OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error) {
	return d.svr.OtaFirmwareFileUpdate(ctx, in)
}

// 附件列表搜索
func (m *defaultFirmwareManage) OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareFileIndex(ctx, in, opts...)
}

// 附件列表搜索
func (d *directFirmwareManage) OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error) {
	return d.svr.OtaFirmwareFileIndex(ctx, in)
}

// 获取固件包对应设备版本列表
func (m *defaultFirmwareManage) OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceInfo(ctx, in, opts...)
}

// 获取固件包对应设备版本列表
func (d *directFirmwareManage) OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error) {
	return d.svr.OtaFirmwareDeviceInfo(ctx, in)
}
