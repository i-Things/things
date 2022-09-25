// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq          = dm.AccessAuthReq
	DeviceInfo             = dm.DeviceInfo
	DeviceInfoDeleteReq    = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq     = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp    = dm.DeviceInfoIndexResp
	DeviceInfoReadReq      = dm.DeviceInfoReadReq
	GroupDeviceCreateReq   = dm.GroupDeviceCreateReq
	GroupDeviceDeleteReq   = dm.GroupDeviceDeleteReq
	GroupDeviceIndexReq    = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp   = dm.GroupDeviceIndexResp
	GroupInfo              = dm.GroupInfo
	GroupInfoCreateReq     = dm.GroupInfoCreateReq
	GroupInfoDeleteReq     = dm.GroupInfoDeleteReq
	GroupInfoIndexReq      = dm.GroupInfoIndexReq
	GroupInfoIndexResp     = dm.GroupInfoIndexResp
	GroupInfoReadReq       = dm.GroupInfoReadReq
	GroupInfoUpdateReq     = dm.GroupInfoUpdateReq
	LoginAuthReq           = dm.LoginAuthReq
	PageInfo               = dm.PageInfo
	ProductInfo            = dm.ProductInfo
	ProductInfoDeleteReq   = dm.ProductInfoDeleteReq
	ProductInfoIndexReq    = dm.ProductInfoIndexReq
	ProductInfoIndexResp   = dm.ProductInfoIndexResp
	ProductInfoReadReq     = dm.ProductInfoReadReq
	ProductSchema          = dm.ProductSchema
	ProductSchemaReadReq   = dm.ProductSchemaReadReq
	ProductSchemaUpdateReq = dm.ProductSchemaUpdateReq
	Response               = dm.Response
	RootCheckReq           = dm.RootCheckReq

	DeviceGroup interface {
		// 创建分组
		GroupInfoCreate(ctx context.Context, in *GroupInfoCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 获取分组信息列表
		GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error)
		// 获取分组信息详情
		GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error)
		// 更新分组
		GroupInfoUpdate(ctx context.Context, in *GroupInfoUpdateReq, opts ...grpc.CallOption) (*Response, error)
		// 删除分组
		GroupInfoDelete(ctx context.Context, in *GroupInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 创建分组设备
		GroupDeviceCreate(ctx context.Context, in *GroupDeviceCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 获取分组设备信息列表
		GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error)
		// 删除分组设备
		GroupDeviceDelete(ctx context.Context, in *GroupDeviceDeleteReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDeviceGroup struct {
		cli zrpc.Client
	}

	directDeviceGroup struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceGroupServer
	}
)

func NewDeviceGroup(cli zrpc.Client) DeviceGroup {
	return &defaultDeviceGroup{
		cli: cli,
	}
}

func NewDirectDeviceGroup(svcCtx *svc.ServiceContext, svr dm.DeviceGroupServer) DeviceGroup {
	return &directDeviceGroup{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 创建分组
func (m *defaultDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfoCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoCreate(ctx, in, opts...)
}

// 创建分组
func (d *directDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfoCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.GroupInfoCreate(ctx, in)
}

// 获取分组信息列表
func (m *defaultDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoIndex(ctx, in, opts...)
}

// 获取分组信息列表
func (d *directDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	return d.svr.GroupInfoIndex(ctx, in)
}

// 获取分组信息详情
func (m *defaultDeviceGroup) GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoRead(ctx, in, opts...)
}

// 获取分组信息详情
func (d *directDeviceGroup) GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error) {
	return d.svr.GroupInfoRead(ctx, in)
}

// 更新分组
func (m *defaultDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfoUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoUpdate(ctx, in, opts...)
}

// 更新分组
func (d *directDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfoUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.GroupInfoUpdate(ctx, in)
}

// 删除分组
func (m *defaultDeviceGroup) GroupInfoDelete(ctx context.Context, in *GroupInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoDelete(ctx, in, opts...)
}

// 删除分组
func (d *directDeviceGroup) GroupInfoDelete(ctx context.Context, in *GroupInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.GroupInfoDelete(ctx, in)
}

// 创建分组设备
func (m *defaultDeviceGroup) GroupDeviceCreate(ctx context.Context, in *GroupDeviceCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceCreate(ctx, in, opts...)
}

// 创建分组设备
func (d *directDeviceGroup) GroupDeviceCreate(ctx context.Context, in *GroupDeviceCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.GroupDeviceCreate(ctx, in)
}

// 获取分组设备信息列表
func (m *defaultDeviceGroup) GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceIndex(ctx, in, opts...)
}

// 获取分组设备信息列表
func (d *directDeviceGroup) GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error) {
	return d.svr.GroupDeviceIndex(ctx, in)
}

// 删除分组设备
func (m *defaultDeviceGroup) GroupDeviceDelete(ctx context.Context, in *GroupDeviceDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceDelete(ctx, in, opts...)
}

// 删除分组设备
func (d *directDeviceGroup) GroupDeviceDelete(ctx context.Context, in *GroupDeviceDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.GroupDeviceDelete(ctx, in)
}
