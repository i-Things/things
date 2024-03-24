// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/logic/productmanage"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"
)

type ProductManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedProductManageServer
}

func NewProductManageServer(svcCtx *svc.ServiceContext) *ProductManageServer {
	return &ProductManageServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductManageServer) ProductInit(ctx context.Context, in *dm.ProductInitReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductInitLogic(ctx, s.svcCtx)
	return l.ProductInit(in)
}

// 新增产品
func (s *ProductManageServer) ProductInfoCreate(ctx context.Context, in *dm.ProductInfo) (*dm.Empty, error) {
	l := productmanagelogic.NewProductInfoCreateLogic(ctx, s.svcCtx)
	return l.ProductInfoCreate(in)
}

// 更新产品
func (s *ProductManageServer) ProductInfoUpdate(ctx context.Context, in *dm.ProductInfo) (*dm.Empty, error) {
	l := productmanagelogic.NewProductInfoUpdateLogic(ctx, s.svcCtx)
	return l.ProductInfoUpdate(in)
}

// 删除产品
func (s *ProductManageServer) ProductInfoDelete(ctx context.Context, in *dm.ProductInfoDeleteReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductInfoDeleteLogic(ctx, s.svcCtx)
	return l.ProductInfoDelete(in)
}

// 获取产品信息列表
func (s *ProductManageServer) ProductInfoIndex(ctx context.Context, in *dm.ProductInfoIndexReq) (*dm.ProductInfoIndexResp, error) {
	l := productmanagelogic.NewProductInfoIndexLogic(ctx, s.svcCtx)
	return l.ProductInfoIndex(in)
}

// 获取产品信息详情
func (s *ProductManageServer) ProductInfoRead(ctx context.Context, in *dm.ProductInfoReadReq) (*dm.ProductInfo, error) {
	l := productmanagelogic.NewProductInfoReadLogic(ctx, s.svcCtx)
	return l.ProductInfoRead(in)
}

// 更新产品物模型
func (s *ProductManageServer) ProductSchemaUpdate(ctx context.Context, in *dm.ProductSchemaUpdateReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductSchemaUpdateLogic(ctx, s.svcCtx)
	return l.ProductSchemaUpdate(in)
}

// 新增产品
func (s *ProductManageServer) ProductSchemaCreate(ctx context.Context, in *dm.ProductSchemaCreateReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductSchemaCreateLogic(ctx, s.svcCtx)
	return l.ProductSchemaCreate(in)
}

// 批量新增物模型,只新增没有的,已有的不处理
func (s *ProductManageServer) ProductSchemaMultiCreate(ctx context.Context, in *dm.ProductSchemaMultiCreateReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductSchemaMultiCreateLogic(ctx, s.svcCtx)
	return l.ProductSchemaMultiCreate(in)
}

// 删除产品
func (s *ProductManageServer) ProductSchemaDelete(ctx context.Context, in *dm.ProductSchemaDeleteReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductSchemaDeleteLogic(ctx, s.svcCtx)
	return l.ProductSchemaDelete(in)
}

// 获取产品信息列表
func (s *ProductManageServer) ProductSchemaIndex(ctx context.Context, in *dm.ProductSchemaIndexReq) (*dm.ProductSchemaIndexResp, error) {
	l := productmanagelogic.NewProductSchemaIndexLogic(ctx, s.svcCtx)
	return l.ProductSchemaIndex(in)
}

// 删除产品
func (s *ProductManageServer) ProductSchemaTslImport(ctx context.Context, in *dm.ProductSchemaTslImportReq) (*dm.Empty, error) {
	l := productmanagelogic.NewProductSchemaTslImportLogic(ctx, s.svcCtx)
	return l.ProductSchemaTslImport(in)
}

// 获取产品信息列表
func (s *ProductManageServer) ProductSchemaTslRead(ctx context.Context, in *dm.ProductSchemaTslReadReq) (*dm.ProductSchemaTslReadResp, error) {
	l := productmanagelogic.NewProductSchemaTslReadLogic(ctx, s.svcCtx)
	return l.ProductSchemaTslRead(in)
}

// 脚本管理
func (s *ProductManageServer) ProductCustomRead(ctx context.Context, in *dm.ProductCustomReadReq) (*dm.ProductCustom, error) {
	l := productmanagelogic.NewProductCustomReadLogic(ctx, s.svcCtx)
	return l.ProductCustomRead(in)
}

func (s *ProductManageServer) ProductCustomUpdate(ctx context.Context, in *dm.ProductCustom) (*dm.Empty, error) {
	l := productmanagelogic.NewProductCustomUpdateLogic(ctx, s.svcCtx)
	return l.ProductCustomUpdate(in)
}

// 新增产品
func (s *ProductManageServer) ProductCategoryCreate(ctx context.Context, in *dm.ProductCategory) (*dm.WithID, error) {
	l := productmanagelogic.NewProductCategoryCreateLogic(ctx, s.svcCtx)
	return l.ProductCategoryCreate(in)
}

// 更新产品
func (s *ProductManageServer) ProductCategoryUpdate(ctx context.Context, in *dm.ProductCategory) (*dm.Empty, error) {
	l := productmanagelogic.NewProductCategoryUpdateLogic(ctx, s.svcCtx)
	return l.ProductCategoryUpdate(in)
}

// 删除产品
func (s *ProductManageServer) ProductCategoryDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := productmanagelogic.NewProductCategoryDeleteLogic(ctx, s.svcCtx)
	return l.ProductCategoryDelete(in)
}

// 获取产品信息列表
func (s *ProductManageServer) ProductCategoryIndex(ctx context.Context, in *dm.ProductCategoryIndexReq) (*dm.ProductCategoryIndexResp, error) {
	l := productmanagelogic.NewProductCategoryIndexLogic(ctx, s.svcCtx)
	return l.ProductCategoryIndex(in)
}

// 获取产品信息详情
func (s *ProductManageServer) ProductCategoryRead(ctx context.Context, in *dm.ProductCategoryReadReq) (*dm.ProductCategory, error) {
	l := productmanagelogic.NewProductCategoryReadLogic(ctx, s.svcCtx)
	return l.ProductCategoryRead(in)
}