package server

import (
	"qshapi/proto/send"
)

//调用其他服务
var (
	webName="sendWeb"
	svName="sendSrv"
	SendClient send.SendSrvService
)
func init() {
	service := Conf.Services[webName]
	SendClient=send.NewSendSrvService(Conf.Services[svName].Name,service.NewRoundSrv().Options().Client)
}