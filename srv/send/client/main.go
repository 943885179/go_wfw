package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/send"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "sendCli"
	svName  = "sendSrv"
	conf    models.APIConfig
	client  send.SendSrvService
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
	client = send.NewSendSrvService(conf.Services[svName].Name, service.NewRoundSrv().Options().Client)
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
			resp.APIOK(c, "消息webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c, gin.H{
				"webconfig": conf.Services[cliName],
				"service":   conf.Services[svName].Name,
			})
		})
		r.POST("sendCode", sendCode)
		r.POST("send", sendMsg)
		r.POST("sendAll", sendAll)
		r.POST("codeVerify", codeVerify)
	}
	return g
}

func sendCode(c *gin.Context) {
	req := &send.SendCodeReq{}
	c.Bind(req)
	result, err := client.SendCode(context.TODO(), req)
	resp.MicroResp(c, result, err)
}
func sendMsg(c *gin.Context) {
	req := &send.SendReq{}
	c.Bind(req)
	result, err := client.Send(context.TODO(), req)
	resp.MicroResp(c, result, err)
}
func sendAll(c *gin.Context) {
	req := &send.SendAllReq{}
	c.Bind(req)
	result, err := client.SendAll(context.TODO(), req)
	resp.MicroResp(c, result, err)
}
func codeVerify(c *gin.Context) {
	req := &send.CodeVerifyReq{}
	c.Bind(req)
	result, err := client.CodeVerify(context.TODO(), req)
	resp.MicroResp(c, result, err)
}
