package server

import (
	"errors"
	"fmt"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjmd5"
	"strings"
)

type ILogin interface {
	Login(req *sysuser.LoginReq,resp  *sysuser.LoginResp) error
}

func NewLogin(tp sysuser.LoginType) ILogin {
	switch tp {
	case sysuser.LoginType_NAME://通过用户名登录
		return &loginByName{}
	case sysuser.LoginType_PHONE://手机登录
		return &loginByPhone{}
	case sysuser.LoginType_EMAIL://邮箱登录
		return &loginByEmail{}
	default:
		panic("不支持该登录方式")
	}
}
type loginByName struct {}
func (*loginByName) Login(req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	req.UserNameOrPhoneOrEmail=strings.Trim(req.UserNameOrPhoneOrEmail,"")
	req.UserPasswordOrCode=strings.Trim(req.UserPasswordOrCode,"")
	if len(req.UserNameOrPhoneOrEmail)==0 || len(req.UserPasswordOrCode)==0 {
		return errors.New("用户名或密码不能为空")
	}
	db:=Conf.DbConfig.New()
	u:=models.SysUser{
	}
	err:=  db.Where(&models.SysUser{UserName: req.UserNameOrPhoneOrEmail, UserPassword:mzjmd5.MD5(req.UserPasswordOrCode)}).First(&u).Error
	if err != nil {
		// First 返回record not found表示没有数据， Find返回nil
		return errors.New("用户名或密码错误")
	}
	resp.UserName=u.UserName
	Conf.Jwt.Data=u
	resp.Token,err=Conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go Conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,Conf.Jwt.TimeOut) //添加到redis中
	return nil
}
type loginByEmail struct {}
func (*loginByEmail) Login(req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	req.UserNameOrPhoneOrEmail=strings.Trim(req.UserNameOrPhoneOrEmail,"")
	req.UserPasswordOrCode=strings.Trim(req.UserPasswordOrCode,"")
	if len(req.UserNameOrPhoneOrEmail)==0 || len(req.UserPasswordOrCode)==0 {
		return errors.New("邮箱和验证码不能为空")
	}
	db:=Conf.DbConfig.New()
	u:=models.SysUser{
		UserEmail: req.UserNameOrPhoneOrEmail,
	}
	err:=  db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}

	resp.UserName=u.UserName
	Conf.Jwt.Data=u
	resp.Token,err=Conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go Conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,Conf.Jwt.TimeOut) //添加到redis中
	return nil
}
type loginByPhone struct {}
func (*loginByPhone) Login(req *sysuser.LoginReq, resp *sysuser.LoginResp) error {
	req.UserNameOrPhoneOrEmail=strings.Trim(req.UserNameOrPhoneOrEmail,"")
	req.UserPasswordOrCode=strings.Trim(req.UserPasswordOrCode,"")
	if len(req.UserNameOrPhoneOrEmail)==0 || len(req.UserPasswordOrCode)==0 {
		return errors.New("电话或验证码不能为空")
	}
	db:=Conf.DbConfig.New()
	defer  db.Close()
	u:=models.SysUser{
		UserPhone: req.UserNameOrPhoneOrEmail,
	}
	err:=  db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}
	resp.UserName=u.UserName
	Conf.Jwt.Data=u
	resp.Token,err=Conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go Conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,Conf.Jwt.TimeOut) //添加到redis中
	return nil
}
