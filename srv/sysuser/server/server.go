package server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/send"
	"qshapi/utils/mzjinit"
)
var (
	cliName="midCli"
	svName="sendSrv"
	sendClient send.SendSrvService
	Conf models.APIConfig
	)
func init(){
	if err:=mzjinit.Default(&Conf);err != nil {
		log.Fatal(err)
	}
	service := Conf.Services[cliName]
	sendClient=send.NewSendSrvService(Conf.Services[svName].Name,service.NewRoundSrv().Options().Client)
}
func CodeVerify(emailOrPhone,code string)(bool,error){
	req:=&send.CodeVerifyReq{
		Code: code,
		EmailOrPhone: emailOrPhone,
	}
	resp, err := sendClient.CodeVerify(context.Background(), req)
	if err !=nil {
		fmt.Println("读取服务失败",err)
		return false,err
	}
	return resp.Verify,nil
}