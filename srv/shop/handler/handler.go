package handler

import (
	"context"
	"qshapi/proto/dbmodel"
	"qshapi/srv/shop/server"
)

type Handler struct {
}

func (h Handler) EditShop(ctx context.Context, req *dbmodel.SysShop, resp *dbmodel.Id) error {
	return server.NewShop().EditShop(req, resp)
}

func (h Handler) DelShop(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewShop().DelShop(req, resp)
}

func (h Handler) ShopList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewShop().ShopList(req, resp)
}
