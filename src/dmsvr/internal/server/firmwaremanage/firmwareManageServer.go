// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/firmwaremanage"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type FirmwareManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedFirmwareManageServer
}

func NewFirmwareManageServer(svcCtx *svc.ServiceContext) *FirmwareManageServer {
	return &FirmwareManageServer{
		svcCtx: svcCtx,
	}
}

// 新增固件升级包
func (s *FirmwareManageServer) FirmwareInfoCreate(ctx context.Context, in *dm.Firmware) (*dm.FirmwareResp, error) {
	l := firmwaremanagelogic.NewFirmwareInfoCreateLogic(ctx, s.svcCtx)
	return l.FirmwareInfoCreate(in)
}

func (s *FirmwareManageServer) FirmwareInfoUpdate(ctx context.Context, in *dm.FirmwareInfo) (*dm.OtaCommonResp, error) {
	l := firmwaremanagelogic.NewFirmwareInfoUpdateLogic(ctx, s.svcCtx)
	return l.FirmwareInfoUpdate(in)
}

func (s *FirmwareManageServer) FirmwareInfoDelete(ctx context.Context, in *dm.FirmwareInfoDeleteReq) (*dm.FirmwareInfoDeleteResp, error) {
	l := firmwaremanagelogic.NewFirmwareInfoDeleteLogic(ctx, s.svcCtx)
	return l.FirmwareInfoDelete(in)
}

func (s *FirmwareManageServer) FirmwareInfoIndex(ctx context.Context, in *dm.FirmwareInfoIndexReq) (*dm.FirmwareInfoIndexResp, error) {
	l := firmwaremanagelogic.NewFirmwareInfoIndexLogic(ctx, s.svcCtx)
	return l.FirmwareInfoIndex(in)
}

func (s *FirmwareManageServer) FirmwareInfoRead(ctx context.Context, in *dm.FirmwareInfoReadReq) (*dm.FirmwareInfoReadResp, error) {
	l := firmwaremanagelogic.NewFirmwareInfoReadLogic(ctx, s.svcCtx)
	return l.FirmwareInfoRead(in)
}

// 附件信息更新
func (s *FirmwareManageServer) OtaFirmwareFileUpdate(ctx context.Context, in *dm.OtaFirmwareFileReq) (*dm.OtaFirmwareFileResp, error) {
	l := firmwaremanagelogic.NewOtaFirmwareFileUpdateLogic(ctx, s.svcCtx)
	return l.OtaFirmwareFileUpdate(in)
}

// 附件列表搜索
func (s *FirmwareManageServer) OtaFirmwareFileIndex(ctx context.Context, in *dm.OtaFirmwareFileIndexReq) (*dm.OtaFirmwareFileIndexResp, error) {
	l := firmwaremanagelogic.NewOtaFirmwareFileIndexLogic(ctx, s.svcCtx)
	return l.OtaFirmwareFileIndex(in)
}

// 获取固件包对应设备版本列表
func (s *FirmwareManageServer) OtaFirmwareDeviceInfo(ctx context.Context, in *dm.OtaFirmwareDeviceInfoReq) (*dm.OtaFirmwareDeviceInfoResp, error) {
	l := firmwaremanagelogic.NewOtaFirmwareDeviceInfoLogic(ctx, s.svcCtx)
	return l.OtaFirmwareDeviceInfo(in)
}