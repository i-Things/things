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

	ProductManage interface {
		// 新增设备
		ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 更新设备
		ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除设备
		ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取设备信息列表
		ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
		// 获取设备信息详情
		ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
		// 更新产品物模型
		ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品物模型
		ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error)
	}

	defaultProductManage struct {
		cli zrpc.Client
	}

	directProductManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.ProductManageServer
	}
)

func NewProductManage(cli zrpc.Client) ProductManage {
	return &defaultProductManage{
		cli: cli,
	}
}

func NewDirectProductManage(svcCtx *svc.ServiceContext, svr dm.ProductManageServer) ProductManage {
	return &directProductManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增设备
func (m *defaultProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoCreate(ctx, in, opts...)
}

// 新增设备
func (d *directProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoCreate(ctx, in)
}

// 更新设备
func (m *defaultProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoUpdate(ctx, in, opts...)
}

// 更新设备
func (d *directProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoUpdate(ctx, in)
}

// 删除设备
func (m *defaultProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoDelete(ctx, in, opts...)
}

// 删除设备
func (d *directProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoDelete(ctx, in)
}

// 获取设备信息列表
func (m *defaultProductManage) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoIndex(ctx, in, opts...)
}

// 获取设备信息列表
func (d *directProductManage) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	return d.svr.ProductInfoIndex(ctx, in)
}

// 获取设备信息详情
func (m *defaultProductManage) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoRead(ctx, in, opts...)
}

// 获取设备信息详情
func (d *directProductManage) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	return d.svr.ProductInfoRead(ctx, in)
}

// 更新产品物模型
func (m *defaultProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductSchemaUpdate(ctx, in)
}

// 获取产品物模型
func (m *defaultProductManage) ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaRead(ctx, in, opts...)
}

// 获取产品物模型
func (d *directProductManage) ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error) {
	return d.svr.ProductSchemaRead(ctx, in)
}
