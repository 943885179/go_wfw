package handler

import (
	"context"
	"errors"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/server"
	"strings"
)

type Handler struct {
	
}


var sv server.Server
func init(){
	sv=server.Server{}
}

func (h Handler) Login(ctx context.Context, req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	switch req.LoginType {
	case sysuser.LoginType_NAME://通过用户名登录
		return sv.LoginByName(req.UserNameOrPhoneOrEmail,req.UserPasswordOrCode,resp)
	case sysuser.LoginType_PHONE://手机登录
		return sv.LoginByPhone(req.UserNameOrPhoneOrEmail,resp)
	case sysuser.LoginType_EMAIL://邮箱登录
		return sv.LoginByEmail(req.UserNameOrPhoneOrEmail,resp)
	default:
		return errors.New("暂时不支持该登录方式")
	}
}
func (h Handler) Registry(ctx context.Context, req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {
	req.UserName=strings.Trim(req.UserName,"")
	req.UserPassword=strings.Trim(req.UserPassword,"")
	req.UserPasswordAgain=strings.Trim(req.UserPasswordAgain,"")
	req.UserPhone=strings.Trim(req.UserPhone,"")
	req.UserPhoneCode=strings.Trim(req.UserPhoneCode,"")
	if  len(req.UserName)==0 ||len( req.UserPassword)==0 ||len(req.UserPasswordAgain)==0 {
		return errors.New("用户名或密码不能为空")
	}
	if   strings.Trim( req.UserPassword,"") != strings.Trim( req.UserPasswordAgain,""){
		return errors.New("两次密码不一致")
	}
	if  len(req.UserPhone)==0 {
		return errors.New("手机号不能为空")
	}
	if len(req.UserPhoneCode)==0 {
		return errors.New("请输入验证码")
	}
	return  sv.Registry(req,resp)
}

