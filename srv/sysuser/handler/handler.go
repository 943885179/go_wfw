package handler

import (
	"context"
	"fmt"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
	
}
func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	fmt.Println("登录")
	return server.NewLogin(req.LoginType).Login(req,resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {
	return  server.NewRegistry().Registry(req,resp)
}