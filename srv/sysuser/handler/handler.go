package handler

import (
	"context"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
}

func (h Handler) EditUser(ctx context.Context, req *sysuser.SysUser, resp *sysuser.EditResp) error {
	return server.NewUser().EditUser(req, resp)
}

func (h Handler) UserGroupList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewUserGroup().UserGroupList(req, resp)
}
func (h Handler) MenuList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewMenu().MenuList(req, resp)
}

func (h Handler) ApiList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewAPI().ApiList(req, resp)
}

func (h Handler) SrvList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewSrv().SrvList(req, resp)
}

func (h Handler) TreeList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewTree().TreeList(req, resp)
}

func (h Handler) RoleList(ctx context.Context, req *sysuser.PageReq, resp *sysuser.PageResp) error {
	return server.NewRole().RoleList(req, resp)
}

func (h Handler) ChangePassword(ctx context.Context, req *sysuser.ChangePasswordReq, resp *sysuser.EditResp) error {
	return server.NewUser().ChangePassword(req, resp)
}

func (h Handler) UserInfoList(ctx context.Context, req *sysuser.UserInfoListReq, resp *sysuser.PageResp) error {
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

func (h Handler) EditRole(ctx context.Context, req *sysuser.SysRole, resp *sysuser.EditResp) error {
	return server.NewRole().EditRole(req, resp)
}

func (h Handler) EditUserGroup(ctx context.Context, req *sysuser.SysGroup, resp *sysuser.EditResp) error {
	return server.NewUserGroup().EditUserGroup(req, resp)
}

func (h Handler) EditMenu(ctx context.Context, req *sysuser.SysMenu, resp *sysuser.EditResp) error {
	return server.NewMenu().EditMenu(req, resp)
}

func (h Handler) EditApi(ctx context.Context, req *sysuser.SysApi, resp *sysuser.EditResp) error {
	return server.NewAPI().EditApi(req, resp)
}

func (h Handler) EditSrv(ctx context.Context, req *sysuser.SysSrv, resp *sysuser.EditResp) error {
	return server.NewSrv().EditSrv(req, resp)
}

func (h Handler) EditTree(ctx context.Context, req *sysuser.SysTree, resp *sysuser.EditResp) error {
	return server.NewTree().EditTree(req, resp)
}

func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	return server.NewLogin(req.LoginType).Login(req, resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.EditResp) error {
	return server.NewRegistry().Registry(req, resp)
}
