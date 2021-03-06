package server

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/send"
	"qshapi/utils/mzjcode"
	"qshapi/utils/mzjinit"
	"time"
)

var conf models.APIConfig
const sendHeard="code_"
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}

type ISend interface {
	SendCode(phoneOrEmail string,resp *send.SendCodeResp)error
	Send(msg string,to ...string)error
	CodeVerify(req *send.CodeVerifyReq, resp *send.CodeVerifyResp)error
}

func NewServer(tp send.SendType)ISend{
	switch  tp {
	case send.SendType_PHONE:
		return &phoneServer{}
	case send.SendType_EMAIL:
		return &emailServer{}
	default:
		panic("暂不支持")
	}
}

type sendServer struct {}
func (s *sendServer)CodeVerify(req *send.CodeVerifyReq, resp *send.CodeVerifyResp) error {
	if len(req.EmailOrPhone)==0 || len(req.Code)==0 {
		return errors.New("电话/邮箱/验证码不能为空")
	}
	key:=fmt.Sprintf("%s%s",sendHeard,req.EmailOrPhone)
	i, err:= conf.RedisConfig.Exists(key)
	if err != nil {
		return err
	}
	if i==0 {
		return errors.New("验证码不存在或已过期")
	}
	v,err:=conf.RedisConfig.Get(key)
	if err != nil {
		return err
	}
	resp.Verify=v==req.Code
	if !resp.Verify {
		return errors.New("验证码错误")
	}
	return nil
}

type emailServer struct {
	sendServer
}
func (s *emailServer)SendCode(email string,resp *send.SendCodeResp) error {
	if len(email)==0 {
		return errors.New("邮箱不能为空")
	}
	code:=mzjcode.GetRandCode(conf.Yzm.Width)//获取六位验证码随机数
	e:=conf.EmailConfig.Copy()
	e.To=[]string{email}
	e.HTML=fmt.Sprintf("<h3>验证码：%s</h3>",code)
	//go e.Send()
	go  conf.RedisConfig.Set(fmt.Sprintf("%s%s",sendHeard,email),code,conf.Yzm.TimeOut*time.Second) //添加到redis中
	if conf.RedisConfig.IsDebug{
		resp.Code=code
	}
	resp.Code=code
	return nil
}
func (s *emailServer)Send(msg string,to ...string)error  {
	if len(to)==0  ||len(msg)==0 {
		return errors.New("邮箱或消息不能为空")
	}
	e:=conf.EmailConfig.Copy()
	e.To=to
	e.HTML=msg
	go e.Send()
	return nil
}

type phoneServer struct {
	sendServer
}
func (s *phoneServer)SendCode(phone string,resp *send.SendCodeResp)  error{
	if len(phone)==0 {
		return errors.New("电话不能为空")
	}
	code:=mzjcode.GetRandCode(conf.Yzm.Width)//获取六位验证码随机数
	//TODO:发送手机验证码
	go  conf.RedisConfig.Set(fmt.Sprintf("%s%s",sendHeard,phone),code,conf.Yzm.TimeOut*time.Second) //添加到redis中
	if conf.RedisConfig.IsDebug{
		resp.Code=code
	}
	return errors.New("手机验证码发正在开发中...")
}
func (s *phoneServer)Send(msg string,to ...string)error  {
	if len(to)==0 ||len(msg)==0 {
		return errors.New("电话或消息不能为空")
	}
	//TODO:发送手机消息
	return errors.New("手机短信群发正在开发中...")
}