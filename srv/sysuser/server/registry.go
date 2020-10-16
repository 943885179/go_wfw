package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
	"strings"
)

type IRegistry interface {
	Registry(req *sysuser.RegistryReq, resp *sysuser.EditResp) error
}

func NewRegistry() IRegistry {
	return &Registry{}
}

type Registry struct{}

func (*Registry) Registry(req *sysuser.RegistryReq, resp *sysuser.EditResp) error {
	req.UserName = strings.Trim(req.UserName, "")
	req.UserPassword = strings.Trim(req.UserPassword, "")
	req.UserPasswordAgain = strings.Trim(req.UserPasswordAgain, "")
	req.UserPhone = strings.Trim(req.UserPhone, "")
	req.UserPhoneCode = strings.Trim(req.UserPhoneCode, "")
	if len(req.UserName) == 0 || len(req.UserPassword) == 0 || len(req.UserPasswordAgain) == 0 {
		return errors.New("用户名或密码不能为空")
	}
	if strings.Trim(req.UserPassword, "") != strings.Trim(req.UserPasswordAgain, "") {
		return errors.New("两次密码不一致")
	}
	if len(req.UserPhone) == 0 {
		return errors.New("手机号不能为空")
	}
	if len(req.UserPhoneCode) == 0 {
		return errors.New("请输入验证码")
	}
	if v, err := CodeVerify(req.UserPhone, req.UserPhoneCode); err != nil || !v {
		//return errors.New("验证码错误")
	}
	db := Conf.DbConfig.New()
	//判断是否存在用户名
	var count int64 = 0
	if err := db.Model(&models.SysUser{}).Where(&models.SysUser{UserName: req.UserName}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	if err := db.Model(&models.SysUser{}).Where(&models.SysUser{UserPhone: req.UserPhone}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该手机号已注册")
	}
	req.UserPassword = mzjmd5.MD5(req.UserPassword)
	u := &models.SysUser{}
	mzjstruct.CopyStruct(req, u)
	u.Id = mzjuuid.WorkerDefault()
	resp.Id = u.Id
	return db.Create(u).Error
}
