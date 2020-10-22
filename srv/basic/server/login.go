package server

import (
	"errors"
	"fmt"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjstruct"
	"strings"
	"time"
)

type ILogin interface {
	Login(req *basic.LoginReq, resp *basic.LoginResp) error
}

func NewLogin(tp dbmodel.LoginType) ILogin {
	switch tp {
	case dbmodel.LoginType_NAME: //通过用户名登录
		return &loginByName{}
	case dbmodel.LoginType_PHONE: //手机登录
		return &loginByPhone{}
	case dbmodel.LoginType_EMAIL: //邮箱登录
		return &loginByEmail{}
	default:
		panic("不支持该登录方式")
	}
}

type loginByName struct{}

func (*loginByName) Login(req *basic.LoginReq, resp *basic.LoginResp) error {
	req.UserNameOrPhoneOrEmail = strings.Trim(req.UserNameOrPhoneOrEmail, "")
	req.UserPasswordOrCode = strings.Trim(req.UserPasswordOrCode, "")
	if len(req.UserNameOrPhoneOrEmail) == 0 || len(req.UserPasswordOrCode) == 0 {
		return errors.New("用户名或密码不能为空")
	}
	db := Conf.DbConfig.New().Model(&models.SysUser{})
	db = db.Preload("Roles").Preload("Groups").Preload("Groups.Roles") //注意大小写
	db = db.Preload("Roles.Srvs").Preload("Roles.Apis").Preload("Roles.Menus").Preload("Roles.Menus.Children").Preload("Roles.Menus.Children.Children")
	db = db.Preload("Groups.Roles.Srvs").Preload("Groups.Roles.Apis").Preload("Groups.Roles.Menus").Preload("Groups.Roles.Menus.Children").Preload("Groups.Roles.Menus.Children.Children")
	db = db.Preload("Province").Preload("City").Preload("Area") //地址
	db = db.Preload("Icon")                                     //头像
	u := models.SysUser{}
	err := db.Where(&models.SysUser{UserName: req.UserNameOrPhoneOrEmail, UserPassword: mzjmd5.MD5(req.UserPasswordOrCode)}).First(&u).Error
	if err != nil {
		// First 返回record not found表示没有数据， Find返回nil
		return errors.New("用户名或密码错误")
	}
	addToken(u, resp)
	return nil
}

type loginByEmail struct{}

func (*loginByEmail) Login(req *basic.LoginReq, resp *basic.LoginResp) error {
	req.UserNameOrPhoneOrEmail = strings.Trim(req.UserNameOrPhoneOrEmail, "")
	req.UserPasswordOrCode = strings.Trim(req.UserPasswordOrCode, "")
	if len(req.UserNameOrPhoneOrEmail) == 0 || len(req.UserPasswordOrCode) == 0 {
		return errors.New("邮箱和验证码不能为空")
	}
	if v, err := CodeVerify(req.UserNameOrPhoneOrEmail, req.UserPasswordOrCode); err != nil || !v {
		return errors.New("验证码错误")
	}
	db := Conf.DbConfig.New().Model(&models.SysUser{})
	db = db.Preload("Roles").Preload("Groups").Preload("Groups.Roles") //注意大小写
	db = db.Preload("Roles.Srvs").Preload("Roles.Apis").Preload("Roles.Menus").Preload("Roles.Menus.Children").Preload("Roles.Menus.Children.Children")
	db = db.Preload("Groups.Roles.Srvs").Preload("Groups.Roles.Apis").Preload("Groups.Roles.Menus").Preload("Groups.Roles.Menus.Children").Preload("Groups.Roles.Menus.Children.Children")
	db = db.Preload("Province").Preload("City").Preload("Area") //地址
	db = db.Preload("Icon")
	u := models.SysUser{
		UserEmail: req.UserNameOrPhoneOrEmail,
	}
	err := db.Where(&u).First(&u).Error
	if err != nil {
		return errors.New("用户未注册")
	}
	addToken(u, resp)
	return nil
}

type loginByPhone struct{}

func (*loginByPhone) Login(req *basic.LoginReq, resp *basic.LoginResp) error {
	req.UserNameOrPhoneOrEmail = strings.Trim(req.UserNameOrPhoneOrEmail, "")
	req.UserPasswordOrCode = strings.Trim(req.UserPasswordOrCode, "")
	if len(req.UserNameOrPhoneOrEmail) == 0 || len(req.UserPasswordOrCode) == 0 {
		return errors.New("电话或验证码不能为空")
	}
	if v, err := CodeVerify(req.UserNameOrPhoneOrEmail, req.UserPasswordOrCode); err != nil || !v {
		return errors.New("验证码错误")
	}
	db := Conf.DbConfig.New().Model(&models.SysUser{})
	db = db.Preload("Roles").Preload("Groups").Preload("Groups.Roles") //注意大小写
	db = db.Preload("Roles.Srvs").Preload("Roles.Apis").Preload("Roles.Menus").Preload("Roles.Menus.Children").Preload("Roles.Menus.Children.Children")
	db = db.Preload("Groups.Roles.Srvs").Preload("Groups.Roles.Apis").Preload("Groups.Roles.Menus").Preload("Groups.Roles.Menus.Children").Preload("Groups.Roles.Menus.Children.Children")
	db = db.Preload("Province").Preload("City").Preload("Area") //地址
	db = db.Preload("Icon")
	u := models.SysUser{
		UserPhone: req.UserNameOrPhoneOrEmail,
	}
	err := db.Where(&u).First(&u).Error
	if err != nil {
		return errors.New("用户未注册")
	}
	addToken(u, resp)
	return nil
}
func addToken(u models.SysUser, resp *basic.LoginResp) {

	Conf.Jwt.Data = fmt.Sprintf("Login_%s", u.UserName) // u
	tk, _ := Conf.Jwt.CreateToken()
	resp.Token = &basic.TokenResp{
		Token:   tk,
		Expired: int64(Conf.Jwt.TimeOut * time.Second),
	}
	mzjstruct.CopyStruct(&u, &resp.User)
	go Conf.RedisConfig.Set(fmt.Sprintf("Login_%s", u.UserName), u, Conf.Jwt.TimeOut*time.Second) //添加到redis中
}
