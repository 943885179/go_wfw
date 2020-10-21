package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/product"
	"qshapi/srv/product/handler"
	"qshapi/utils/mzjinit"
)

var (
	svName = "shopSrv"
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
	product.RegisterProductSrvHandler(s.Server(), handler.Handler{})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
