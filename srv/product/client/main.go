package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/proto/product"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "productCli"
	svName  = "productSrv"
	conf    models.APIConfig
	client  product.ProductSrvService
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
	client = product.NewProductSrvService(svName, service.NewRoundSrv().Options().Client)
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
		r.POST("EditProduct", EditProduct)
		r.POST("DelProduct", DelProduct)
		r.POST("DelProductSku", DelProductSku)
		r.POST("ProductList", ProductList)
		r.POST("EditProductByIds", EditProductByIds)
		r.GET("/ProductById/:id", ProductById)
	}
	return g
}

func EditProduct(c *gin.Context) {
	req := dbmodel.Product{}
	c.Bind(&req)
	req.ShopId = mzjgin.ShopId
	result, err := client.EditProduct(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelProduct(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelProduct(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ProductList(c *gin.Context) {
	req := product.ProductListReq{}
	c.Bind(&req)
	req.UserId = mzjgin.UserId
	req.ShopId = mzjgin.ShopId
	req.Token = mzjgin.LoginToken
	result, err := client.ProductList(context.TODO(), &req)
	var rs []dbmodel.Product
	for _, any := range result.Data {
		var r dbmodel.Product
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
}
func ProductById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.ProductById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func DelProductSku(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelProductSku(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditProductByIds(c *gin.Context) {
	req := dbmodel.Ids{}
	c.Bind(&req)
	result, err := client.EditProductByIds(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
