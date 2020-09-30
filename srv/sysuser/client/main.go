package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	webName="userWeb"
	svName="userSrv"
	conf models.APIConfig
	client sysuser.UserSrvService
	resp= mzjgin.Resp{}
)
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[webName]
	client=sysuser.NewUserSrvService(conf.Services[svName].Name, service.NewRoundSrv().Options().Client)
	s:= service.NewGinWeb(SrvGin())
	if err:=s.Run();err!= nil {
		log.Fatal(err)
	}
}
func SrvGin() *gin.Engine {
	g:=mzjgin.NewGin().Default()
	r:=g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c,"用户webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c,gin.H{
				"webconfig":conf.Services[webName],
				"service":conf.Services[svName].Name,
			})
		})
		r.POST("Login",Login)
		r.POST("Registry",Registry)
	}
	return g
}

func Login(c *gin.Context)  {
	req:=&sysuser.LoginReq{}
	c.Bind(req)
	result, err := client.Login(context.TODO(), req)
	resp.MicroResp(c,result,err)
	//if err != nil {
	//	resp.APIError(c,err.Error())
	//	return
	//}
	// resp.APIOK(c,result)
}
func Registry(c *gin.Context)  {
	req:=&sysuser.RegistryReq{}
	c.Bind(req)
	result, err := client.Registry(context.TODO(), req)
	resp.MicroResp(c,result,err)
}