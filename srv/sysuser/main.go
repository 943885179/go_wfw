package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/srv/sysuser/handler"
	"qshapi/utils/mzjinit"
)

var (
	svName = "userSrv"
	conf   models.APIConfig
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)

}
func main() {
	service := conf.Services[svName]
	s := service.NewSrv()
	sysuser.RegisterUserSrvHandler(s.Server(), handler.Handler{})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
