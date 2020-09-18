package handler

import (
	"context"
	"errors"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
)

type Handler struct {
	
}
var sv server.Server
func init(){
	sv=server.Server{}
}
func (h Handler) LoginByName(ctx context.Context, req *sysuser.LoginByNameReq, resp *sysuser.LoginResp) error {
	if len(req.UserName)==0 || len(req.UserPassword)==0 {
		return errors.New("用户名或密码不能为空")
	}
	return  sv.LoginByName(req,resp)
}

func (h Handler) LoginByEmail(ctx context.Context, req *sysuser.LoginByEmailReq, resp *sysuser.LoginResp) error {
	if len(req.UserEmail)==0 || len(req.Code)==0 {
		return errors.New("邮箱或验证码不能为空")
	}
	return  sv.LoginByEmail(req,resp)
}

func (h Handler) LoginByPhone(ctx context.Context, req *sysuser.LoginByPhoneReq, resp *sysuser.LoginResp) error {
	if len(req.UserPhone)==0 || len(req.Code)==0 {
		return errors.New("电话或验证码不能为空")
	}
	return sv.LoginByPhone(req, resp)
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {

	return  sv.Registry(req,resp)
}

