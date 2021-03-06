// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/basic/basic.proto

package basic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
	dbmodel "qshapi/proto/dbmodel"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for BasicSrv service

type BasicSrvService interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error)
	Registry(ctx context.Context, in *RegistryReq, opts ...client.CallOption) (*dbmodel.Id, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...client.CallOption) (*dbmodel.Id, error)
	UserInfoList(ctx context.Context, in *UserInfoListReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	EditUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.Id, error)
	UserById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysUser, error)
	EditRole(ctx context.Context, in *dbmodel.SysRole, opts ...client.CallOption) (*dbmodel.Id, error)
	DelRole(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	RoleList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	RoleById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysRole, error)
	RoleTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error)
	EditUserGroup(ctx context.Context, in *dbmodel.SysGroup, opts ...client.CallOption) (*dbmodel.Id, error)
	DelUserGroup(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	UserGroupList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	UserGroupById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysGroup, error)
	EditMenu(ctx context.Context, in *dbmodel.SysMenu, opts ...client.CallOption) (*dbmodel.Id, error)
	DelMenu(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	MenuList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	MenuTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error)
	EditArea(ctx context.Context, in *dbmodel.SysArea, opts ...client.CallOption) (*dbmodel.Id, error)
	DelArea(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	AreaList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	AreaTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error)
	AreaById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysArea, error)
	MenuListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlyMenu, error)
	MenuById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysMenu, error)
	EditApi(ctx context.Context, in *dbmodel.SysApi, opts ...client.CallOption) (*dbmodel.Id, error)
	DelApi(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	ApiList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	ApiById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysApi, error)
	ApiListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlyApi, error)
	EditSrv(ctx context.Context, in *dbmodel.SysSrv, opts ...client.CallOption) (*dbmodel.Id, error)
	DelSrv(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	SrvList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	SrvById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysSrv, error)
	SrvListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlySrv, error)
	EditTree(ctx context.Context, in *dbmodel.SysTree, opts ...client.CallOption) (*dbmodel.Id, error)
	DelTree(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	TreeList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	TreeById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysTree, error)
	TreeTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error)
	TreeByType(ctx context.Context, in *TreeType, opts ...client.CallOption) (*dbmodel.TreeResp, error)
	EditShop(ctx context.Context, in *dbmodel.SysShop, opts ...client.CallOption) (*dbmodel.Id, error)
	DelShop(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	ShopList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error)
	ShopById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysShop, error)
	EditQualification(ctx context.Context, in *dbmodel.Qualification, opts ...client.CallOption) (*dbmodel.Id, error)
	EditQualifications(ctx context.Context, in *Qualifications, opts ...client.CallOption) (*empty.Empty, error)
	DelQualification(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error)
	QualificationByForeignId(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*Qualifications, error)
}

type basicSrvService struct {
	c    client.Client
	name string
}

func NewBasicSrvService(name string, c client.Client) BasicSrvService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "basic"
	}
	return &basicSrvService{
		c:    c,
		name: name,
	}
}

func (c *basicSrvService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*LoginResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.Login", in)
	out := new(LoginResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) Registry(ctx context.Context, in *RegistryReq, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.Registry", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ChangePassword", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) UserInfoList(ctx context.Context, in *UserInfoListReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.UserInfoList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditUser", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) UserById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysUser, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.UserById", in)
	out := new(dbmodel.SysUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditRole(ctx context.Context, in *dbmodel.SysRole, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditRole", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelRole(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelRole", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) RoleList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.RoleList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) RoleById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysRole, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.RoleById", in)
	out := new(dbmodel.SysRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) RoleTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.RoleTree", in)
	out := new(dbmodel.TreeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditUserGroup(ctx context.Context, in *dbmodel.SysGroup, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditUserGroup", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelUserGroup(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelUserGroup", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) UserGroupList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.UserGroupList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) UserGroupById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysGroup, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.UserGroupById", in)
	out := new(dbmodel.SysGroup)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditMenu(ctx context.Context, in *dbmodel.SysMenu, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditMenu", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelMenu(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelMenu", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) MenuList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.MenuList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) MenuTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.MenuTree", in)
	out := new(dbmodel.TreeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditArea(ctx context.Context, in *dbmodel.SysArea, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditArea", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelArea(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelArea", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) AreaList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.AreaList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) AreaTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.AreaTree", in)
	out := new(dbmodel.TreeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) AreaById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysArea, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.AreaById", in)
	out := new(dbmodel.SysArea)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) MenuListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlyMenu, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.MenuListByUser", in)
	out := new(dbmodel.OnlyMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) MenuById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysMenu, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.MenuById", in)
	out := new(dbmodel.SysMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditApi(ctx context.Context, in *dbmodel.SysApi, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditApi", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelApi(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelApi", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ApiList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ApiList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ApiById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysApi, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ApiById", in)
	out := new(dbmodel.SysApi)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ApiListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlyApi, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ApiListByUser", in)
	out := new(dbmodel.OnlyApi)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditSrv(ctx context.Context, in *dbmodel.SysSrv, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditSrv", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelSrv(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelSrv", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) SrvList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.SrvList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) SrvById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysSrv, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.SrvById", in)
	out := new(dbmodel.SysSrv)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) SrvListByUser(ctx context.Context, in *dbmodel.SysUser, opts ...client.CallOption) (*dbmodel.OnlySrv, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.SrvListByUser", in)
	out := new(dbmodel.OnlySrv)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditTree(ctx context.Context, in *dbmodel.SysTree, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditTree", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelTree(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelTree", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) TreeList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.TreeList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) TreeById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysTree, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.TreeById", in)
	out := new(dbmodel.SysTree)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) TreeTree(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*dbmodel.TreeResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.TreeTree", in)
	out := new(dbmodel.TreeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) TreeByType(ctx context.Context, in *TreeType, opts ...client.CallOption) (*dbmodel.TreeResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.TreeByType", in)
	out := new(dbmodel.TreeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditShop(ctx context.Context, in *dbmodel.SysShop, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditShop", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelShop(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelShop", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ShopList(ctx context.Context, in *dbmodel.PageReq, opts ...client.CallOption) (*dbmodel.PageResp, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ShopList", in)
	out := new(dbmodel.PageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) ShopById(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.SysShop, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.ShopById", in)
	out := new(dbmodel.SysShop)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditQualification(ctx context.Context, in *dbmodel.Qualification, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditQualification", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) EditQualifications(ctx context.Context, in *Qualifications, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.EditQualifications", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) DelQualification(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*dbmodel.Id, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.DelQualification", in)
	out := new(dbmodel.Id)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicSrvService) QualificationByForeignId(ctx context.Context, in *dbmodel.Id, opts ...client.CallOption) (*Qualifications, error) {
	req := c.c.NewRequest(c.name, "BasicSrv.QualificationByForeignId", in)
	out := new(Qualifications)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BasicSrv service

type BasicSrvHandler interface {
	Login(context.Context, *LoginReq, *LoginResp) error
	Registry(context.Context, *RegistryReq, *dbmodel.Id) error
	ChangePassword(context.Context, *ChangePasswordReq, *dbmodel.Id) error
	UserInfoList(context.Context, *UserInfoListReq, *dbmodel.PageResp) error
	EditUser(context.Context, *dbmodel.SysUser, *dbmodel.Id) error
	UserById(context.Context, *dbmodel.Id, *dbmodel.SysUser) error
	EditRole(context.Context, *dbmodel.SysRole, *dbmodel.Id) error
	DelRole(context.Context, *dbmodel.Id, *dbmodel.Id) error
	RoleList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	RoleById(context.Context, *dbmodel.Id, *dbmodel.SysRole) error
	RoleTree(context.Context, *empty.Empty, *dbmodel.TreeResp) error
	EditUserGroup(context.Context, *dbmodel.SysGroup, *dbmodel.Id) error
	DelUserGroup(context.Context, *dbmodel.Id, *dbmodel.Id) error
	UserGroupList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	UserGroupById(context.Context, *dbmodel.Id, *dbmodel.SysGroup) error
	EditMenu(context.Context, *dbmodel.SysMenu, *dbmodel.Id) error
	DelMenu(context.Context, *dbmodel.Id, *dbmodel.Id) error
	MenuList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	MenuTree(context.Context, *empty.Empty, *dbmodel.TreeResp) error
	EditArea(context.Context, *dbmodel.SysArea, *dbmodel.Id) error
	DelArea(context.Context, *dbmodel.Id, *dbmodel.Id) error
	AreaList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	AreaTree(context.Context, *empty.Empty, *dbmodel.TreeResp) error
	AreaById(context.Context, *dbmodel.Id, *dbmodel.SysArea) error
	MenuListByUser(context.Context, *dbmodel.SysUser, *dbmodel.OnlyMenu) error
	MenuById(context.Context, *dbmodel.Id, *dbmodel.SysMenu) error
	EditApi(context.Context, *dbmodel.SysApi, *dbmodel.Id) error
	DelApi(context.Context, *dbmodel.Id, *dbmodel.Id) error
	ApiList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	ApiById(context.Context, *dbmodel.Id, *dbmodel.SysApi) error
	ApiListByUser(context.Context, *dbmodel.SysUser, *dbmodel.OnlyApi) error
	EditSrv(context.Context, *dbmodel.SysSrv, *dbmodel.Id) error
	DelSrv(context.Context, *dbmodel.Id, *dbmodel.Id) error
	SrvList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	SrvById(context.Context, *dbmodel.Id, *dbmodel.SysSrv) error
	SrvListByUser(context.Context, *dbmodel.SysUser, *dbmodel.OnlySrv) error
	EditTree(context.Context, *dbmodel.SysTree, *dbmodel.Id) error
	DelTree(context.Context, *dbmodel.Id, *dbmodel.Id) error
	TreeList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	TreeById(context.Context, *dbmodel.Id, *dbmodel.SysTree) error
	TreeTree(context.Context, *empty.Empty, *dbmodel.TreeResp) error
	TreeByType(context.Context, *TreeType, *dbmodel.TreeResp) error
	EditShop(context.Context, *dbmodel.SysShop, *dbmodel.Id) error
	DelShop(context.Context, *dbmodel.Id, *dbmodel.Id) error
	ShopList(context.Context, *dbmodel.PageReq, *dbmodel.PageResp) error
	ShopById(context.Context, *dbmodel.Id, *dbmodel.SysShop) error
	EditQualification(context.Context, *dbmodel.Qualification, *dbmodel.Id) error
	EditQualifications(context.Context, *Qualifications, *empty.Empty) error
	DelQualification(context.Context, *dbmodel.Id, *dbmodel.Id) error
	QualificationByForeignId(context.Context, *dbmodel.Id, *Qualifications) error
}

func RegisterBasicSrvHandler(s server.Server, hdlr BasicSrvHandler, opts ...server.HandlerOption) error {
	type basicSrv interface {
		Login(ctx context.Context, in *LoginReq, out *LoginResp) error
		Registry(ctx context.Context, in *RegistryReq, out *dbmodel.Id) error
		ChangePassword(ctx context.Context, in *ChangePasswordReq, out *dbmodel.Id) error
		UserInfoList(ctx context.Context, in *UserInfoListReq, out *dbmodel.PageResp) error
		EditUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.Id) error
		UserById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysUser) error
		EditRole(ctx context.Context, in *dbmodel.SysRole, out *dbmodel.Id) error
		DelRole(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		RoleList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		RoleById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysRole) error
		RoleTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error
		EditUserGroup(ctx context.Context, in *dbmodel.SysGroup, out *dbmodel.Id) error
		DelUserGroup(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		UserGroupList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		UserGroupById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysGroup) error
		EditMenu(ctx context.Context, in *dbmodel.SysMenu, out *dbmodel.Id) error
		DelMenu(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		MenuList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		MenuTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error
		EditArea(ctx context.Context, in *dbmodel.SysArea, out *dbmodel.Id) error
		DelArea(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		AreaList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		AreaTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error
		AreaById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysArea) error
		MenuListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlyMenu) error
		MenuById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysMenu) error
		EditApi(ctx context.Context, in *dbmodel.SysApi, out *dbmodel.Id) error
		DelApi(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		ApiList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		ApiById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysApi) error
		ApiListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlyApi) error
		EditSrv(ctx context.Context, in *dbmodel.SysSrv, out *dbmodel.Id) error
		DelSrv(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		SrvList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		SrvById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysSrv) error
		SrvListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlySrv) error
		EditTree(ctx context.Context, in *dbmodel.SysTree, out *dbmodel.Id) error
		DelTree(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		TreeList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		TreeById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysTree) error
		TreeTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error
		TreeByType(ctx context.Context, in *TreeType, out *dbmodel.TreeResp) error
		EditShop(ctx context.Context, in *dbmodel.SysShop, out *dbmodel.Id) error
		DelShop(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		ShopList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error
		ShopById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysShop) error
		EditQualification(ctx context.Context, in *dbmodel.Qualification, out *dbmodel.Id) error
		EditQualifications(ctx context.Context, in *Qualifications, out *empty.Empty) error
		DelQualification(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error
		QualificationByForeignId(ctx context.Context, in *dbmodel.Id, out *Qualifications) error
	}
	type BasicSrv struct {
		basicSrv
	}
	h := &basicSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&BasicSrv{h}, opts...))
}

type basicSrvHandler struct {
	BasicSrvHandler
}

func (h *basicSrvHandler) Login(ctx context.Context, in *LoginReq, out *LoginResp) error {
	return h.BasicSrvHandler.Login(ctx, in, out)
}

func (h *basicSrvHandler) Registry(ctx context.Context, in *RegistryReq, out *dbmodel.Id) error {
	return h.BasicSrvHandler.Registry(ctx, in, out)
}

func (h *basicSrvHandler) ChangePassword(ctx context.Context, in *ChangePasswordReq, out *dbmodel.Id) error {
	return h.BasicSrvHandler.ChangePassword(ctx, in, out)
}

func (h *basicSrvHandler) UserInfoList(ctx context.Context, in *UserInfoListReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.UserInfoList(ctx, in, out)
}

func (h *basicSrvHandler) EditUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditUser(ctx, in, out)
}

func (h *basicSrvHandler) UserById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysUser) error {
	return h.BasicSrvHandler.UserById(ctx, in, out)
}

func (h *basicSrvHandler) EditRole(ctx context.Context, in *dbmodel.SysRole, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditRole(ctx, in, out)
}

func (h *basicSrvHandler) DelRole(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelRole(ctx, in, out)
}

func (h *basicSrvHandler) RoleList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.RoleList(ctx, in, out)
}

func (h *basicSrvHandler) RoleById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysRole) error {
	return h.BasicSrvHandler.RoleById(ctx, in, out)
}

func (h *basicSrvHandler) RoleTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error {
	return h.BasicSrvHandler.RoleTree(ctx, in, out)
}

func (h *basicSrvHandler) EditUserGroup(ctx context.Context, in *dbmodel.SysGroup, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditUserGroup(ctx, in, out)
}

func (h *basicSrvHandler) DelUserGroup(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelUserGroup(ctx, in, out)
}

func (h *basicSrvHandler) UserGroupList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.UserGroupList(ctx, in, out)
}

func (h *basicSrvHandler) UserGroupById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysGroup) error {
	return h.BasicSrvHandler.UserGroupById(ctx, in, out)
}

func (h *basicSrvHandler) EditMenu(ctx context.Context, in *dbmodel.SysMenu, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditMenu(ctx, in, out)
}

func (h *basicSrvHandler) DelMenu(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelMenu(ctx, in, out)
}

func (h *basicSrvHandler) MenuList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.MenuList(ctx, in, out)
}

func (h *basicSrvHandler) MenuTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error {
	return h.BasicSrvHandler.MenuTree(ctx, in, out)
}

func (h *basicSrvHandler) EditArea(ctx context.Context, in *dbmodel.SysArea, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditArea(ctx, in, out)
}

func (h *basicSrvHandler) DelArea(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelArea(ctx, in, out)
}

func (h *basicSrvHandler) AreaList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.AreaList(ctx, in, out)
}

func (h *basicSrvHandler) AreaTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error {
	return h.BasicSrvHandler.AreaTree(ctx, in, out)
}

func (h *basicSrvHandler) AreaById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysArea) error {
	return h.BasicSrvHandler.AreaById(ctx, in, out)
}

func (h *basicSrvHandler) MenuListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlyMenu) error {
	return h.BasicSrvHandler.MenuListByUser(ctx, in, out)
}

func (h *basicSrvHandler) MenuById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysMenu) error {
	return h.BasicSrvHandler.MenuById(ctx, in, out)
}

func (h *basicSrvHandler) EditApi(ctx context.Context, in *dbmodel.SysApi, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditApi(ctx, in, out)
}

func (h *basicSrvHandler) DelApi(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelApi(ctx, in, out)
}

func (h *basicSrvHandler) ApiList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.ApiList(ctx, in, out)
}

func (h *basicSrvHandler) ApiById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysApi) error {
	return h.BasicSrvHandler.ApiById(ctx, in, out)
}

func (h *basicSrvHandler) ApiListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlyApi) error {
	return h.BasicSrvHandler.ApiListByUser(ctx, in, out)
}

func (h *basicSrvHandler) EditSrv(ctx context.Context, in *dbmodel.SysSrv, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditSrv(ctx, in, out)
}

func (h *basicSrvHandler) DelSrv(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelSrv(ctx, in, out)
}

func (h *basicSrvHandler) SrvList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.SrvList(ctx, in, out)
}

func (h *basicSrvHandler) SrvById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysSrv) error {
	return h.BasicSrvHandler.SrvById(ctx, in, out)
}

func (h *basicSrvHandler) SrvListByUser(ctx context.Context, in *dbmodel.SysUser, out *dbmodel.OnlySrv) error {
	return h.BasicSrvHandler.SrvListByUser(ctx, in, out)
}

func (h *basicSrvHandler) EditTree(ctx context.Context, in *dbmodel.SysTree, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditTree(ctx, in, out)
}

func (h *basicSrvHandler) DelTree(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelTree(ctx, in, out)
}

func (h *basicSrvHandler) TreeList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.TreeList(ctx, in, out)
}

func (h *basicSrvHandler) TreeById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysTree) error {
	return h.BasicSrvHandler.TreeById(ctx, in, out)
}

func (h *basicSrvHandler) TreeTree(ctx context.Context, in *empty.Empty, out *dbmodel.TreeResp) error {
	return h.BasicSrvHandler.TreeTree(ctx, in, out)
}

func (h *basicSrvHandler) TreeByType(ctx context.Context, in *TreeType, out *dbmodel.TreeResp) error {
	return h.BasicSrvHandler.TreeByType(ctx, in, out)
}

func (h *basicSrvHandler) EditShop(ctx context.Context, in *dbmodel.SysShop, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditShop(ctx, in, out)
}

func (h *basicSrvHandler) DelShop(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelShop(ctx, in, out)
}

func (h *basicSrvHandler) ShopList(ctx context.Context, in *dbmodel.PageReq, out *dbmodel.PageResp) error {
	return h.BasicSrvHandler.ShopList(ctx, in, out)
}

func (h *basicSrvHandler) ShopById(ctx context.Context, in *dbmodel.Id, out *dbmodel.SysShop) error {
	return h.BasicSrvHandler.ShopById(ctx, in, out)
}

func (h *basicSrvHandler) EditQualification(ctx context.Context, in *dbmodel.Qualification, out *dbmodel.Id) error {
	return h.BasicSrvHandler.EditQualification(ctx, in, out)
}

func (h *basicSrvHandler) EditQualifications(ctx context.Context, in *Qualifications, out *empty.Empty) error {
	return h.BasicSrvHandler.EditQualifications(ctx, in, out)
}

func (h *basicSrvHandler) DelQualification(ctx context.Context, in *dbmodel.Id, out *dbmodel.Id) error {
	return h.BasicSrvHandler.DelQualification(ctx, in, out)
}

func (h *basicSrvHandler) QualificationByForeignId(ctx context.Context, in *dbmodel.Id, out *Qualifications) error {
	return h.BasicSrvHandler.QualificationByForeignId(ctx, in, out)
}
