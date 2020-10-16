package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "userCli"
	svName  = "userSrv"
	conf    models.APIConfig
	client  sysuser.UserSrvService
	resp    = mzjgin.Resp{}
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[cliName]
	client = sysuser.NewUserSrvService(conf.Services[svName].Name, service.NewRoundSrv().Options().Client)
	s := service.NewGinWeb(SrvGin())
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
func SrvGin() *gin.Engine {
	g := mzjgin.NewGin().Default()
	r := g.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			resp.APIOK(c, "用户webapi")
		})
		r.POST("/", func(c *gin.Context) {
			resp.APIOK(c, gin.H{
				"webconfig": conf.Services[cliName],
				"service":   conf.Services[svName].Name,
			})
		})
		r.POST("Login", Login)
		r.POST("Registry", Registry)
		r.POST("UserInfoList", UserInfoList)

		r.POST("EditApi", EditApi)
		r.POST("DelApi", DelApi)
		r.POST("EditSrv", EditSrv)
		r.POST("DelSrv", DelSrv)
		r.POST("EditRole", EditRole)
		r.POST("DelRole", DelRole)
		r.POST("EditUserGroup", EditUserGroup)
		r.POST("DelUserGroup", DelUserGroup)
		r.POST("EditMenu", EditMenu)
		r.POST("DelMenu", DelMenu)
		r.POST("EditTree", EditTree)
		r.POST("DelTree", DelTree)
	}
	return g
}
func Login(c *gin.Context) {
	req := sysuser.LoginReq{}
	c.Bind(&req)
	result, err := client.Login(context.TODO(), &req)
	resp.MicroResp(c, result, err)
	//if err != nil {
	//	resp.APIError(c,err.Error())
	//	return
	//}
	// resp.APIOK(c,result)
}
func Registry(c *gin.Context) {
	req := sysuser.RegistryReq{}
	c.Bind(&req)
	result, err := client.Registry(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func EditApi(c *gin.Context) {
	req := sysuser.ApiReq{} //😩 sysuser.ApiReq{}直接定义成这个然后使用中文会转化报错，还是麻烦点吧
	c.Bind(&req)
	fmt.Println("参数", &req)
	result, err := client.EditApi(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditSrv(c *gin.Context) {
	req := sysuser.SrvReq{}
	c.Bind(&req)
	result, err := client.EditSrv(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditTree(c *gin.Context) {
	req := sysuser.TreeReq{}
	c.Bind(&req)
	result, err := client.EditTree(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditMenu(c *gin.Context) {
	req := sysuser.MenuReq{}
	c.Bind(&req)
	result, err := client.EditMenu(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditUserGroup(c *gin.Context) {
	req := sysuser.UserGroupReq{}
	c.Bind(&req)
	result, err := client.EditUserGroup(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditRole(c *gin.Context) {
	req := sysuser.RoleReq{}
	c.Bind(&req)
	result, err := client.EditRole(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func DelTree(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelTree(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelSrv(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelSrv(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelApi(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelApi(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelMenu(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelMenu(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelUserGroup(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelUserGroup(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelRole(c *gin.Context) {
	req := sysuser.DelReq{}
	c.Bind(&req)
	result, err := client.DelRole(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func UserInfoList(c *gin.Context) {
	req := sysuser.UserInfoListReq{}
	c.Bind(&req)
	result, err := client.UserInfoList(context.TODO(), &req)
	if err != nil {
		resp.MicroResp(c, result, err)
		return
	}
	var users models.SysUser
	json.Unmarshal(result.Data.Value, &users)
	resp.MicroResp(c, users, err)
}
