package handler

import (
	"context"
	"qshapi/proto/dbmodel"
	"qshapi/srv/orders/server"
)

type Handler struct {
}

func (h Handler) EditOrders(ctx context.Context, req *dbmodel.Orders, resp *dbmodel.Id) error {
	return server.NewOrder().EditOrder(req, resp)
}

func (h Handler) DelOrders(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewOrder().DelOrder(req, resp)
}

func (h Handler) OrderLists(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewOrder().OrderList(req, resp)
}
