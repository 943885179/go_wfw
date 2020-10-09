package server

import (
	"qshapi/proto/send"
)

//调用其他服务
var (
	webName="sendWeb"
	svName="sendSrv"
	client send.SendSrvService
)
func main() {
	service := Conf.Services[webName]
	client=send.NewSendSrvService(Conf.Services[svName].Name,service.NewRoundSrv().Options().Client)
}