package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/proto/dbmodel"
	"qshapi/proto/product"
	"qshapi/srv/product/server"
)

type Handler struct {
}

func (h Handler) ProductList(ctx context.Context, req *product.ProductListReq, resp *dbmodel.PageResp) error {

	return server.NewProduct().ProductList(req, resp)
}

func (h Handler) EditProductByIds(ctx context.Context, ids *dbmodel.Ids, empty *empty.Empty) error {

	return server.NewProduct().EditProductByIds(ids)
}

func (h Handler) EditProductSku(ctx context.Context, sku *dbmodel.ProductSku, id *dbmodel.Id) error {
	return server.NewProductSku().EditProductSku(sku, id)
}

func (h Handler) DelProductSku(ctx context.Context, id *dbmodel.Id, id2 *dbmodel.Id) error {
	return server.NewProductSku().DelProductSku(id, id2)
}

func (h Handler) ProductSkuById(ctx context.Context, id *dbmodel.Id, sku *dbmodel.ProductSku) error {
	return server.NewProductSku().ProductSkuById(id, sku)
}

func (h Handler) ProductById(ctx context.Context, id *dbmodel.Id, product *dbmodel.Product) error {
	return server.NewProduct().ProductById(id, product)
}

func (h Handler) EditProduct(ctx context.Context, req *dbmodel.Product, resp *dbmodel.Id) error {
	return server.NewProduct().EditProduct(req, resp)
}

func (h Handler) DelProduct(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewProduct().DelProduct(req, resp)
}
