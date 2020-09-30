package server

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjuuid"
	"strings"
)
var (
	conf models.APIConfig
	)
type Server struct {

}
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
func (s Server) LoginByName(name,pwd string, resp *sysuser.LoginResp) error {
	name=strings.Trim(name,"")
	pwd=strings.Trim(pwd,"")
	if len(name)==0 || len(pwd)==0 {
		return errors.New("用户名或密码不能为空")
	}
	db:=conf.DbConfig.New()
	u:=models.SysUser{
	}
	err:=  db.Where(&models.SysUser{UserName: name, UserPassword:mzjmd5.MD5(pwd)}).First(&u).Error
	if err != nil {
		// First 返回record not found表示没有数据， Find返回nil
		return errors.New("用户名或密码错误")
	}
	 resp.UserName=u.UserName
	conf.Jwt.Data=u
	resp.Token,err=conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) LoginByEmail(email string, resp *sysuser.LoginResp) error {
	email=strings.Trim(email,"")
	if len(email)==0  {
		return errors.New("邮箱不能为空")
	}
	db:=conf.DbConfig.New()
	u:=models.SysUser{
		UserEmail: email,
	}
	err:=  db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}

	resp.UserName=u.UserName
	conf.Jwt.Data=u
	resp.Token,err=conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) LoginByPhone(phone string, resp *sysuser.LoginResp) error {
	phone=strings.Trim(phone,"")
	if len(phone)==0  {
		return errors.New("电话不能为空")
	}
	db:=conf.DbConfig.New()
	defer  db.Close()
	u:=models.SysUser{
		UserPhone: phone,
	}
	err:=  db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}
	resp.UserName=u.UserName
	conf.Jwt.Data=u
	resp.Token,err=conf.Jwt.CreateToken()
	if err != nil {
		return err
	}
	go conf.RedisConfig.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) Registry(req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {
	db:=conf.DbConfig.New()
	defer  db.Close()
	//判断是否存在用户名
	var count =0
	if err:=db.Model(&models.SysUser{}).Where(&models.SysUser{UserName: req.UserName}).Count(&count).Error;err != nil {
		return err
	}
	if count >0 {
		return errors.New("用户已存在")
	}
	if err:=db.Model(&models.SysUser{}).Where(&models.SysUser{UserPhone: req.UserPhone}).Count(&count).Error;err != nil {
		return err
	}
	if count >0 {
		return errors.New("该手机号已注册")
	}
	u:=models.SysUser{
		ID: mzjuuid.WorkerDefault(),
		UserName: req.UserName,
		UserPassword: mzjmd5.MD5(req.UserPassword),
		UserPhone: req.UserPhone,

	}
	fmt.Println(u.ID)
	return db.Create(&u).Error
}
func (s Server) VertifyCode()bool{//验证码验证
	return  false
}
