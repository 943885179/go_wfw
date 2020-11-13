package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/srv/basic/server"
)

type Handler struct {
}

func (h Handler) EditQualifications(ctx context.Context, qualification *dbmodel.Qualification, id *dbmodel.Id) error {
	return server.NewQualifications().EditQualifications(qualification, id)
}

func (h Handler) DelQualifications(ctx context.Context, id *dbmodel.Id, id2 *dbmodel.Id) error {
	return server.NewQualifications().DelQualifications(id, id2)
}

func (h Handler) MenuListByUser(ctx context.Context, user *dbmodel.SysUser, menu *dbmodel.OnlyMenu) error {
	return server.NewMenu().MenuListByUser(user, menu)
}

func (h Handler) ApiListByUser(ctx context.Context, user *dbmodel.SysUser, api *dbmodel.OnlyApi) error {
	return server.NewAPI().ApiListByUser(user, api)
}

func (h Handler) SrvListByUser(ctx context.Context, user *dbmodel.SysUser, srv *dbmodel.OnlySrv) error {
	return server.NewSrv().SrvListByUser(user, srv)
}

func (h Handler) RoleTree(ctx context.Context, empty *empty.Empty, resp *dbmodel.TreeResp) error {
	return server.NewRole().RoleTree(empty, resp)
}

func (h Handler) MenuTree(ctx context.Context, empty *empty.Empty, resp *dbmodel.TreeResp) error {
	return server.NewMenu().MenuTree(empty, resp)
}

func (h Handler) TreeTree(ctx context.Context, empty *empty.Empty, resp *dbmodel.TreeResp) error {

	return server.NewTree().TreeTree(empty, resp)
}

func (h Handler) UserById(ctx context.Context, id *dbmodel.Id, user *dbmodel.SysUser) error {
	return server.NewUser().UserById(id, user)
}

func (h Handler) RoleById(ctx context.Context, id *dbmodel.Id, role *dbmodel.SysRole) error {
	return server.NewRole().RoleById(id, role)
}

func (h Handler) UserGroupById(ctx context.Context, id *dbmodel.Id, group *dbmodel.SysGroup) error {
	return server.NewUserGroup().UserGroupById(id, group)
}

func (h Handler) MenuById(ctx context.Context, id *dbmodel.Id, menu *dbmodel.SysMenu) error {
	return server.NewMenu().MenuById(id, menu)
}

func (h Handler) ApiById(ctx context.Context, id *dbmodel.Id, api *dbmodel.SysApi) error {
	return server.NewAPI().ApiById(id, api)
}

func (h Handler) SrvById(ctx context.Context, id *dbmodel.Id, srv *dbmodel.SysSrv) error {
	return server.NewSrv().SrvById(id, srv)
}

func (h Handler) TreeById(ctx context.Context, id *dbmodel.Id, tree *dbmodel.SysTree) error {
	return server.NewTree().TreeById(id, tree)
}

func (h Handler) ShopById(ctx context.Context, id *dbmodel.Id, shop *dbmodel.SysShop) error {
	return server.NewShop().ShopById(id, shop)
}

func (h Handler) EditUser(ctx context.Context, req *dbmodel.SysUser, resp *dbmodel.Id) error {
	return server.NewUser().EditUser(req, resp)
}

func (h Handler) UserGroupList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewUserGroup().UserGroupList(req, resp)
}
func (h Handler) MenuList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewMenu().MenuList(req, resp)
}

func (h Handler) ApiList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewAPI().ApiList(req, resp)
}

func (h Handler) SrvList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewSrv().SrvList(req, resp)
}

func (h Handler) TreeList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewTree().TreeList(req, resp)
}

func (h Handler) RoleList(ctx context.Context, req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	return server.NewRole().RoleList(req, resp)
}

func (h Handler) ChangePassword(ctx context.Context, req *basic.ChangePasswordReq, resp *dbmodel.Id) error {
	return server.NewUser().ChangePassword(req, resp)
}

func (h Handler) UserInfoList(ctx context.Context, req *basic.UserInfoListReq, resp *dbmodel.PageResp) error {
	return server.NewUser().UserInfoList(req, resp)
}

func (h Handler) DelRole(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewRole().DelRole(req, resp)
}

func (h Handler) DelUserGroup(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewUserGroup().DelUserGroup(req, resp)
}

func (h Handler) DelMenu(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewMenu().DelMenu(req, resp)
}

func (h Handler) DelApi(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewAPI().DelApi(req, resp)
}

func (h Handler) DelSrv(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewSrv().DelSrv(req, resp)
}

func (h Handler) DelTree(ctx context.Context, req *dbmodel.Id, resp *dbmodel.Id) error {
	return server.NewTree().DelTree(req, resp)
}

func (h Handler) EditRole(ctx context.Context, req *dbmodel.SysRole, resp *dbmodel.Id) error {
	return server.NewRole().EditRole(req, resp)
}

func (h Handler) EditUserGroup(ctx context.Context, req *dbmodel.SysGroup, resp *dbmodel.Id) error {
	return server.NewUserGroup().EditUserGroup(req, resp)
}

func (h Handler) EditMenu(ctx context.Context, req *dbmodel.SysMenu, resp *dbmodel.Id) error {
	return server.NewMenu().EditMenu(req, resp)
}

func (h Handler) EditApi(ctx context.Context, req *dbmodel.SysApi, resp *dbmodel.Id) error {
	return server.NewAPI().EditApi(req, resp)
}

func (h Handler) EditSrv(ctx context.Context, req *dbmodel.SysSrv, resp *dbmodel.Id) error {
	return server.NewSrv().EditSrv(req, resp)
}

func (h Handler) EditTree(ctx context.Context, req *dbmodel.SysTree, resp *dbmodel.Id) error {
	return server.NewTree().EditTree(req, resp)
}

func (h Handler) Login(ctx context.Context, req *basic.LoginReq, resp *basic.LoginResp) error {
	return server.NewLogin(req.LoginType).Login(req, resp)
}
func (h Handler) Registry(ctx context.Context, req *basic.RegistryReq, resp *dbmodel.Id) error {
	return server.NewRegistry().Registry(req, resp)
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
