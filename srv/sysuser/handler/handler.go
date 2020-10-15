package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
}

func (h Handler) DelRole(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewRole().DelRole(req)
}

func (h Handler) DelUserGroup(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewUserGroup().DelUserGroup(req)
}

func (h Handler) DelMenu(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewMenu().DelMenu(req)
}

func (h Handler) DelApi(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewAPI().DelApi(req)
}

func (h Handler) DelSrv(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewSrv().DelSrv(req)
}

func (h Handler) DelTree(ctx context.Context, req *sysuser.DelReq, e *empty.Empty) error {
	return server.NewTree().DelTree(req)
}

func (h Handler) EditRole(ctx context.Context, req *sysuser.RoleReq, empty *empty.Empty) error {
	return server.NewRole().EditRole(req)
}

func (h Handler) EditUserGroup(ctx context.Context, req *sysuser.UserGroupReq, empty *empty.Empty) error {
	return server.NewUserGroup().EditUserGroup(req)
}

func (h Handler) EditMenu(ctx context.Context, req *sysuser.MenuReq, empty *empty.Empty) error {
	return server.NewMenu().EditMenu(req)
}

func (h Handler) EditApi(ctx context.Context, req *sysuser.ApiReq, empty *empty.Empty) error {
	return server.NewAPI().EditApi(req)
}

func (h Handler) EditSrv(ctx context.Context, req *sysuser.SrvReq, empty *empty.Empty) error {
	return server.NewSrv().EditSrv(req)
}

func (h Handler) EditTree(ctx context.Context, req *sysuser.TreeReq, empty *empty.Empty) error {
	return server.NewTree().EditTree(req)
}

func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	return server.NewLogin(req.LoginType).Login(req, resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, empty *empty.Empty) error {
	return server.NewRegistry().Registry(req)
}
