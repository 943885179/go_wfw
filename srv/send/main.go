package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/send"
	"qshapi/srv/send/handler"
	"qshapi/utils/mzjinit"
)

var (
	svName = "sendSrv"
	conf   models.APIConfig
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[svName]
	s := service.NewSrv()
	send.RegisterSendSrvHandler(s.Server(), handler.Handler{})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
	/*addrs := make([]string, 1)
	addrs[0] = "127.0.0.1:8848"
	registry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = addrs
	})
	service := micro.NewService(
		// Set service name
		micro.Name("my.micro.service"),
		// Set service registry
		micro.Registry(registry),
	)
	service.Run()*/
}
