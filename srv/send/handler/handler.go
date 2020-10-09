package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/proto/send"
	"qshapi/srv/send/server"
)
type Handler struct {
	
}
func (h Handler) SendCode(ctx context.Context, req *send.SendCodeReq, resp *send.SendCodeResp) error {
	return server.NewServer(req.SendType).SendCode(req.EmailOrPhone,resp)
}
func (h Handler) CodeVerify(ctx context.Context, req *send.CodeVerifyReq, resp *send.CodeVerifyResp) error {
	return server.NewServer(req.SendType).CodeVerify(req,resp)
}
func (h Handler) Send(ctx context.Context, req *send.SendReq, empty *empty.Empty) error {
	return server.NewServer(req.SendType).Send(req.Msg,req.EmailOrPhone)
}
func (h Handler) SendAll(ctx context.Context, req *send.SendAllReq, empty *empty.Empty) error {
	return server.NewServer(req.SendType).Send(req.Msg,req.EmailOrPhone...)
}
