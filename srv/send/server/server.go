package server

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/send"
	"qshapi/utils/mzjcode"
	"qshapi/utils/mzjinit"
)

var conf models.APIConfig
const sendHeard="code_"

type Server struct {

}

func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
func (s *Server)SendCodeEmail(email string,resp *send.SendCodeResp) error {
	if len(email)==0 {
		return errors.New("邮箱不能为空")
	}
	resp.Code=mzjcode.GetRandCode(conf.Yzm.Width)//获取六位验证码随机数

	e:=conf.EmailConfig.Copy()
	e.To=[]string{email}
	e.HTML=fmt.Sprintf("<h3>验证码：%s</h3>",resp.Code)
	go e.Send()
	go  conf.RedisConfig.Set(fmt.Sprintf("%s%s",sendHeard,email),resp.Code,conf.Yzm.TimeOut) //添加到redis中
	return nil
}
func (s *Server)SendCodePhone(phone string,resp *send.SendCodeResp)  error{
	if len(phone)==0 {
		return errors.New("电话不能为空")
	}
	resp.Code=mzjcode.GetRandCode(conf.Yzm.Width)//获取六位验证码随机数
	//TODO:发送手机验证码
	go  conf.RedisConfig.Set(fmt.Sprintf("%s%s",sendHeard,phone),resp.Code,conf.Yzm.TimeOut) //添加到redis中
	return errors.New("手机验证码发正在开发中...")
}
func (s *Server)SendEmail(msg string,to ...string)error  {
	if len(to)==0  ||len(msg)==0 {
		return errors.New("邮箱或消息不能为空")
	}
	e:=conf.EmailConfig.Copy()
	e.To=to
	e.HTML=msg
	go e.Send()
	return nil
}
func (s *Server)SendPhone(msg string,to ...string)error  {
	if len(to)==0 ||len(msg)==0 {
		return errors.New("电话或消息不能为空")
	}
	//TODO:发送手机消息
	return errors.New("手机短信群发正在开发中...")
}
func (s *Server)CodeVerify(req *send.CodeVerifyReq, resp *send.CodeVerifyResp) error {
	if len(req.EmailOrPhone)==0 || len(req.Code)==0 {
		return errors.New("电话/邮箱/验证码不能为空")
	}
	v,err:=conf.RedisConfig.Get(fmt.Sprintf("%s%s",sendHeard,req.EmailOrPhone))
	if err != nil {
		return err
	}
	resp.Verify=v==req.Code
	return nil
}