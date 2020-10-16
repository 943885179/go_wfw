package handler

import (
	"context"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
}

func (h Handler) ChangePassword(ctx context.Context, req *sysuser.ChangePasswordReq, resp *sysuser.EditResp) error {
	return server.NewUser().ChangePassword(req, resp)
}

func (h Handler) UserInfoList(ctx context.Context, req *sysuser.UserInfoListReq, resp *sysuser.UserInfoListResp) error {
	return server.NewUser().UserInfoList(req, resp)
}

func (h Handler) DelRole(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewRole().DelRole(req, resp)
}

func (h Handler) DelUserGroup(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewUserGroup().DelUserGroup(req, resp)
}

func (h Handler) DelMenu(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewMenu().DelMenu(req, resp)
}

func (h Handler) DelApi(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewAPI().DelApi(req, resp)
}

func (h Handler) DelSrv(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewSrv().DelSrv(req, resp)
}

func (h Handler) DelTree(ctx context.Context, req *sysuser.DelReq, resp *sysuser.EditResp) error {
	return server.NewTree().DelTree(req, resp)
}

func (h Handler) EditRole(ctx context.Context, req *sysuser.RoleReq, resp *sysuser.EditResp) error {
	return server.NewRole().EditRole(req, resp)
}

func (h Handler) EditUserGroup(ctx context.Context, req *sysuser.UserGroupReq, resp *sysuser.EditResp) error {
	return server.NewUserGroup().EditUserGroup(req, resp)
}

func (h Handler) EditMenu(ctx context.Context, req *sysuser.MenuReq, resp *sysuser.EditResp) error {
	return server.NewMenu().EditMenu(req, resp)
}

func (h Handler) EditApi(ctx context.Context, req *sysuser.ApiReq, resp *sysuser.EditResp) error {
	return server.NewAPI().EditApi(req, resp)
}

func (h Handler) EditSrv(ctx context.Context, req *sysuser.SrvReq, resp *sysuser.EditResp) error {
	return server.NewSrv().EditSrv(req, resp)
}

func (h Handler) EditTree(ctx context.Context, req *sysuser.TreeReq, resp *sysuser.EditResp) error {
	return server.NewTree().EditTree(req, resp)
}

func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	return server.NewLogin(req.LoginType).Login(req, resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.EditResp) error {
	return server.NewRegistry().Registry(req, resp)
}
