package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
	v2 "qshapi/utils/mzjmicro/v2"
)

var (
	conf models.APIConfig
	sv *v2.Service
	client sysuser.UserSrvService
	resp mzjgin.Resp
)
func init(){
	if err:=mzjinit.InitByMicroConfig("config.conf",&conf);err != nil {
		log.Fatal(err)
	}
	resp= mzjgin.Resp{}
	sv=v2.NewService(conf.Services["userWeb"])
	client=sysuser.NewUserSrvService(conf.Services["userSrv"].Name,sv.NewRoundSrv().Options().Client)
}
func main() {
	s:=sv.NewGinWeb(userGin())
	if err:=s.Run();err!= nil {
		log.Fatal(err)
	}
}

func userGin() *gin.Engine {
	g:=gin.Default()
	r:=g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c,"用户webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c,"用户webapi")
		})
		r.POST("LoginByName",LoginByName)
		r.POST("LoginByEmail",LoginByEmail)
		r.POST("LoginByPhone",LoginByPhone)
		r.POST("Registry",Registry)
	}
	return g
}

func LoginByName(c *gin.Context)  {
	req:=&sysuser.LoginByNameReq{}
	c.Bind(req)
	result, err := client.LoginByName(context.TODO(), req)
	if err != nil {
		resp.APIError(c,err.Error())
	}
	 resp.APIOK(c,result)
}
func LoginByEmail(c *gin.Context)  {

}
func LoginByPhone(c *gin.Context)  {

}
func Registry(c *gin.Context)  {

}