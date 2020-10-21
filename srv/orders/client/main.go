package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/proto/orders"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "OrdersCli"
	svName  = "OrdersSrv"
	conf    models.APIConfig
	client  orders.OrderSrvService
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
	client = orders.NewOrderSrvService(svName, service.NewRoundSrv().Options().Client)
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
		r.POST("EditOrder", EditOrder)
		r.POST("DelOrder", DelOrder)
		r.POST("OrderList", OrderList)
	}
	return g
}

func EditOrder(c *gin.Context) {
	req := dbmodel.Orders{}
	c.Bind(&req)
	result, err := client.EditOrder(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelOrder(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelOrder(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func OrderList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.OrderList(context.TODO(), &req)
	var rs []dbmodel.Orders
	for _, any := range result.Data {
		var r dbmodel.Orders
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
}
