package server

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjgorm"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjjwt"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjredis"
)
var (
	conf models.APIConfig
	dbConfig *mzjgorm.DbConfig
	)
type Server struct {

}
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
	dbConfig=mzjgorm.NewDbConfig(conf.DbConfig)
	if dbConfig == nil {
		log.Fatal("数据库连接初始化失败",conf.DbConfig)
	}
}
func (s Server) LoginByName(req *sysuser.LoginByNameReq, resp *sysuser.LoginResp) error {
	req.UserPassword=mzjmd5.MD5(req.UserPassword)//数据加密
	db:=dbConfig.New()
	u:=models.SysUser{
		UserName: req.UserName,
		UserPassword: req.UserPassword,
	}
	err:=  db.First(&u).Error
	if err != nil {
		return err
	}
	fmt.Println("读取到的客户为",u)
	if u.ID==0 {
		return errors.New("用户名或密码错误")
	}
	 resp.UserName=u.UserName
	 resp.Token,err=mzjjwt.NewToken(conf.Jwt).CreateToken()
	if err != nil {
		return err
	}
	go mzjredis.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) LoginByEmail( req *sysuser.LoginByEmailReq, resp *sysuser.LoginResp) error {
	db:=dbConfig.New()
	u:=models.SysUser{
		UserEmail: req.UserEmail,
	}
	err:=  db.First(&u).Error
	if err != nil {
		return err
	}
	resp.UserName=u.UserName
	resp.Token,err=mzjjwt.NewToken(conf.Jwt).CreateToken()
	if err != nil {
		return err
	}
	go mzjredis.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) LoginByPhone(req *sysuser.LoginByPhoneReq, resp *sysuser.LoginResp) error {
	db:=dbConfig.New()
	u:=models.SysUser{
		UserPhone: req.UserPhone,
	}
	err:=  db.First(&u).Error
	if err != nil {
		return err
	}
	resp.UserName=u.UserName
	resp.Token,err=mzjjwt.NewToken(conf.Jwt).CreateToken()
	if err != nil {
		return err
	}
	go mzjredis.Set(fmt.Sprintf("LoginByName_%s",u.UserName),resp.Token,conf.Jwt.TimeOut) //添加到redis中
	return nil
}

func (s Server) Registry(req *sysuser.RegistryReq, resp *sysuser.RegistryResp) error {
	panic("implement me")
}

