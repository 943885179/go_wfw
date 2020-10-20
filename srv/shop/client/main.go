package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/proto/shop"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "shopCli"
	svName  = "shopSrv"
	conf    models.APIConfig
	client  shop.ShopSrvService
	resp    = mzjgin.Resp{}
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[cliName]
	cliName = service.Name
	svName = conf.Services[svName].Name
	fmt.Println(svName)
	client = shop.NewShopSrvService(svName, service.NewRoundSrv().Options().Client)
	s := service.NewGinWeb(SrvGin())
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
func SrvGin() *gin.Engine {
	g := mzjgin.NewGin().Default(cliName)
	r := g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c, "商家webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c, gin.H{
				"webconfig": conf.Services[cliName],
				"service":   conf.Services[svName].Name,
			})
		})
		r.POST("EditShop", EditShop)
		r.POST("DelShop", DelShop)
		r.POST("ShopList", ShopList)
	}
	return g
}

func EditShop(c *gin.Context) {
	req := dbmodel.SysShop{}
	c.Bind(&req)
	result, err := client.EditShop(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelShop(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelShop(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ShopList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.ShopList(context.TODO(), &req)
	var rs []dbmodel.SysShop
	for _, any := range result.Data {
		var r dbmodel.SysShop
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
}
