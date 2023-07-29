// Code generated by goctl. DO NOT EDIT.
// Source: stock.proto

package stocksvr

import (
	"context"

	"github.com/i-Things/things/src/stocksvr/internal/svc"
	"github.com/i-Things/things/src/stocksvr/types/pb/stock"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqStockLocationCreate = stock.ReqStockLocationCreate
	ReqStockLocationDelete = stock.ReqStockLocationDelete
	ReqStockLocationList   = stock.ReqStockLocationList
	ReqStockLocationUpdate = stock.ReqStockLocationUpdate
	ResStockLocationCreate = stock.ResStockLocationCreate
	ResStockLocationDelete = stock.ResStockLocationDelete
	ResStockLocationList   = stock.ResStockLocationList
	ResStockLocationUpdate = stock.ResStockLocationUpdate
	StockLocationItem      = stock.StockLocationItem

	Stocksvr interface {
		StockLocationCreate(ctx context.Context, in *ReqStockLocationCreate, opts ...grpc.CallOption) (*ResStockLocationCreate, error)
		StockLocationUpdate(ctx context.Context, in *ReqStockLocationUpdate, opts ...grpc.CallOption) (*ResStockLocationUpdate, error)
		StockLocationDelete(ctx context.Context, in *ReqStockLocationDelete, opts ...grpc.CallOption) (*ResStockLocationDelete, error)
		StockLocationList(ctx context.Context, in *ReqStockLocationList, opts ...grpc.CallOption) (*ResStockLocationList, error)
	}

	defaultStocksvr struct {
		cli zrpc.Client
	}

	directStocksvr struct {
		svcCtx *svc.ServiceContext
		svr    stock.StocksvrServer
	}
)

func NewStocksvr(cli zrpc.Client) Stocksvr {
	return &defaultStocksvr{
		cli: cli,
	}
}

func NewDirectStocksvr(svcCtx *svc.ServiceContext, svr stock.StocksvrServer) Stocksvr {
	return &directStocksvr{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultStocksvr) StockLocationCreate(ctx context.Context, in *ReqStockLocationCreate, opts ...grpc.CallOption) (*ResStockLocationCreate, error) {
	client := stock.NewStocksvrClient(m.cli.Conn())
	return client.StockLocationCreate(ctx, in, opts...)
}

func (d *directStocksvr) StockLocationCreate(ctx context.Context, in *ReqStockLocationCreate, opts ...grpc.CallOption) (*ResStockLocationCreate, error) {
	return d.svr.StockLocationCreate(ctx, in)
}

func (m *defaultStocksvr) StockLocationUpdate(ctx context.Context, in *ReqStockLocationUpdate, opts ...grpc.CallOption) (*ResStockLocationUpdate, error) {
	client := stock.NewStocksvrClient(m.cli.Conn())
	return client.StockLocationUpdate(ctx, in, opts...)
}

func (d *directStocksvr) StockLocationUpdate(ctx context.Context, in *ReqStockLocationUpdate, opts ...grpc.CallOption) (*ResStockLocationUpdate, error) {
	return d.svr.StockLocationUpdate(ctx, in)
}

func (m *defaultStocksvr) StockLocationDelete(ctx context.Context, in *ReqStockLocationDelete, opts ...grpc.CallOption) (*ResStockLocationDelete, error) {
	client := stock.NewStocksvrClient(m.cli.Conn())
	return client.StockLocationDelete(ctx, in, opts...)
}

func (d *directStocksvr) StockLocationDelete(ctx context.Context, in *ReqStockLocationDelete, opts ...grpc.CallOption) (*ResStockLocationDelete, error) {
	return d.svr.StockLocationDelete(ctx, in)
}

func (m *defaultStocksvr) StockLocationList(ctx context.Context, in *ReqStockLocationList, opts ...grpc.CallOption) (*ResStockLocationList, error) {
	client := stock.NewStocksvrClient(m.cli.Conn())
	return client.StockLocationList(ctx, in, opts...)
}

func (d *directStocksvr) StockLocationList(ctx context.Context, in *ReqStockLocationList, opts ...grpc.CallOption) (*ResStockLocationList, error) {
	return d.svr.StockLocationList(ctx, in)
}
