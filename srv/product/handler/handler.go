package handler

import (
	"context"
	"qshapi/proto/dbmodel"
	"qshapi/srv/product/server"
)

type Handler struct {
}

func (h Handler) EditProduct(ctx context.Context, req *dbmodel.Product, resp *dbmodel.Id) error {
	return server.NewProduct().EditProduct(req, resp)
}

func (h Handler) DelProduct(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewProduct().DelProduct(req, resp)
}

func (h Handler) ProductList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewProduct().ProductList(req, resp)
}
