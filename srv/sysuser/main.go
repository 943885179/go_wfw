package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/handler"
	"qshapi/utils/mzjinit"
	v2 "qshapi/utils/mzjmicro/v2"
)
var (
	conf models.APIConfig
	sv *v2.Service
)
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
	log.Fatal(err)
	}
	sv=v2.NewService(conf.Services["userSrv"])

}
func main() {
	s:=sv.NewSrv()
	sysuser.RegisterUserSrvHandler(s.Server(),handler.Handler{})
	if err:=s.Run();err != nil {
		log.Fatal(err)
	}
}