package handler

import (
	"context"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
	
}
func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	return server.NewLogin(req.LoginType).Login(req,resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {
	return  server.NewRegistry().Registry(req,resp)
}