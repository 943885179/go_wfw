package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
)

var (
	cliName = "basicCli"
	svName  = "basicSrv"
	conf    models.APIConfig
	client  basic.UserSrvService
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
	client = basic.NewUserSrvService(svName, service.NewRoundSrv().Options().Client)
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

		r.POST("EditUser", EditUser)
		r.POST("EditApi", EditApi)
		r.POST("EditSrv", EditSrv)
		r.POST("EditRole", EditRole)
		r.POST("EditUserGroup", EditUserGroup)
		r.POST("EditMenu", EditMenu)
		r.POST("EditTree", EditTree)

		r.POST("DelApi", DelApi)
		r.POST("DelSrv", DelSrv)
		r.POST("DelRole", DelRole)
		r.POST("DelUserGroup", DelUserGroup)
		r.POST("DelMenu", DelMenu)
		r.POST("DelTree", DelTree)
		r.POST("ChangePassword", ChangePassword)

		r.POST("UserInfoList", UserInfoList)
		r.POST("RoleList", RoleList)
		r.POST("TreeList", TreeList)
		r.POST("ApiList", ApiList)
		r.POST("SrvList", SrvList)
		r.POST("MenuList", MenuList)
		r.POST("UserGroupList", UserGroupList)

		r.POST("EditShop", EditShop)
		r.POST("DelShop", DelShop)
		r.POST("ShopList", ShopList)
	}
	return g
}
func Login(c *gin.Context) {
	//mzjgin.SrvRole(c, svName, "Login") //服务权限判断
	req := basic.LoginReq{}
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
	req := basic.RegistryReq{}
	c.Bind(&req)
	result, err := client.Registry(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditUser(c *gin.Context) {
	req := dbmodel.SysUser{}
	c.Bind(&req)
	result, err := client.EditUser(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditApi(c *gin.Context) {
	req := dbmodel.SysApi{} //😩 basic.Api{}直接定义成这个然后使用中文会转化报错，还是麻烦点吧
	c.Bind(&req)
	fmt.Println("参数", &req)
	result, err := client.EditApi(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditSrv(c *gin.Context) {
	req := dbmodel.SysSrv{}
	c.Bind(&req)
	result, err := client.EditSrv(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditTree(c *gin.Context) {
	req := dbmodel.SysTree{}
	c.Bind(&req)
	result, err := client.EditTree(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditMenu(c *gin.Context) {
	req := dbmodel.SysMenu{}
	c.Bind(&req)
	result, err := client.EditMenu(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditUserGroup(c *gin.Context) {
	req := dbmodel.SysGroup{}
	c.Bind(&req)
	result, err := client.EditUserGroup(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditRole(c *gin.Context) {
	req := dbmodel.SysRole{}
	c.Bind(&req)
	result, err := client.EditRole(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func DelTree(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelTree(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelSrv(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelSrv(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelApi(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelApi(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelMenu(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelMenu(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelUserGroup(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelUserGroup(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelRole(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelRole(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ApiList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.ApiList(context.TODO(), &req)
	var rs []dbmodel.SysApi
	for _, any := range result.Data {
		var r dbmodel.SysApi
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}
func SrvList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.SrvList(context.TODO(), &req)
	var rs []dbmodel.SysSrv
	for _, any := range result.Data {
		var r dbmodel.SysSrv
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}
func UserGroupList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.UserGroupList(context.TODO(), &req)
	var rs []dbmodel.SysGroup
	for _, any := range result.Data {
		var r dbmodel.SysGroup
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}
func TreeList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.TreeList(context.TODO(), &req)
	var rs []dbmodel.SysTree
	for _, any := range result.Data {
		var r dbmodel.SysTree
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}
func MenuList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.MenuList(context.TODO(), &req)
	var rs []dbmodel.SysMenu
	for _, any := range result.Data {
		var r dbmodel.SysMenu
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}
func RoleList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.RoleList(context.TODO(), &req)
	var rs []dbmodel.SysRole
	for _, any := range result.Data {
		var r dbmodel.SysRole
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}

func UserInfoList(c *gin.Context) {
	req := basic.UserInfoListReq{}
	c.Bind(&req)
	result, err := client.UserInfoList(context.TODO(), &req)
	var us []dbmodel.SysUser
	for _, any := range result.Data {
		var u dbmodel.SysUser
		ptypes.UnmarshalAny(any, &u)
		us = append(us, u)
	}
	resp.MicroTotalResp(c, result.Total, us, err)
	/*if err != nil {
		resp.MicroResp(c, result, err)
		return
	}
	var users []models.SysUser
	fmt.Println(string(result.Data.Value))
	json.Unmarshal(result.Data.Value, &users)
	resp.MicroTotalResp(c, result.Total, users, err)*/
}
func ChangePassword(c *gin.Context) {
	req := basic.ChangePasswordReq{}
	c.Bind(&req)
	result, err := client.ChangePassword(context.TODO(), &req)
	resp.MicroResp(c, result, err)

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