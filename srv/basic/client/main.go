package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjgin"
	"qshapi/utils/mzjinit"
	"strconv"
)

var (
	cliName = "basicCli"
	svName  = "basicSrv"
	conf    models.APIConfig
	client  basic.BasicSrvService
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
	client = basic.NewBasicSrvService(svName, service.NewRoundSrv().Options().Client)
	s := service.NewGinWeb(SrvGin())
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
func SrvGin() *gin.Engine {
	g := mzjgin.NewGin().Default(cliName)
	//g.Use(mzjgin.TokenAuthMiddleware(cliName)) 权限认证
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
		//r.POST("Login", mzjgin.APITokenMiddleware, mzjgin.TokenAuthMiddleware(""), Login)
		r.POST("Login", Login)
		r.POST("Registry", Registry)

		r.POST("EditUser", EditUser)
		r.POST("EditApi", EditApi)
		r.POST("EditSrv", EditSrv)
		r.POST("EditRole", EditRole)
		r.POST("EditUserGroup", EditUserGroup)
		r.POST("EditMenu", EditMenu)
		r.POST("EditTree", EditTree)
		r.POST("EditArea", EditArea)

		r.POST("DelApi", DelApi)
		r.POST("DelSrv", DelSrv)
		r.POST("DelRole", DelRole)
		r.POST("DelUserGroup", DelUserGroup)
		r.POST("DelMenu", DelMenu)
		r.POST("DelTree", DelTree)
		r.POST("DelArea", DelArea)

		r.GET("/UserById/:id", UserById)
		r.GET("/ApiById/:id", ApiById)
		r.GET("/SrvById/:id", SrvById)
		r.GET("/RoleById/:id", RoleById)
		r.GET("/UserGroupById/:id", UserGroupById)
		r.GET("/MenuById/:id", MenuById)
		r.GET("/TreeById/:id", TreeById)
		r.GET("/AreaById/:id", AreaById)

		r.GET("RoleTree", RoleTree)
		r.GET("MenuTree", MenuTree)
		r.GET("TreeTree", TreeTree)
		r.GET("AreaTree", AreaTree)

		r.GET("/TreeByType/:type", TreeByType)

		r.POST("ChangePassword", ChangePassword)

		r.POST("UserInfoList", UserInfoList)
		r.POST("RoleList", RoleList)
		r.POST("TreeList", TreeList)
		r.POST("ApiList", ApiList)
		r.POST("SrvList", SrvList)
		r.POST("MenuList", MenuList)
		r.POST("UserGroupList", UserGroupList)
		r.POST("AreaList", AreaList)

		r.POST("EditShop", EditShop)
		r.POST("DelShop", DelShop)
		r.POST("ShopList", ShopList)
		r.GET("/ShopById/:id", ShopById)

		r.POST("EditQualification", EditQualification)
		r.POST("EditQualifications", EditQualifications)
		r.POST("DelQualification", DelQualification)
		r.GET("/QualificationByForeignId/:id", QualificationByForeignId)

		r.GET("Token", Token)
	}
	return g
}

func Token(c *gin.Context) {
	//user, err := mzjgin.TokenResp(c)
	token := c.Request.Header.Get("token")
	user, err := mzjgin.TokenResp(token)
	result, err := client.MenuListByUser(context.TODO(), user.User)
	resp.MicroResp(c, result, err)
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
	//添加当前用户
	req.UserId = mzjgin.UserId
	result, err := client.EditShop(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelShop(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	//添加当前用户
	req.UserId = mzjgin.UserId
	result, err := client.DelShop(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ShopList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	req.UserId = mzjgin.UserId
	req.UserType = mzjgin.UserType
	result, err := client.ShopList(context.TODO(), &req)
	var rs []dbmodel.SysShop
	for _, any := range result.Data {
		var r dbmodel.SysShop
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
}

func UserById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.UserById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func UserGroupById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.UserGroupById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func RoleById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.RoleById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func MenuById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.MenuById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func TreeById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.TreeById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func AreaById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.AreaById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ApiById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.ApiById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func SrvById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.SrvById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func ShopById(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.ShopById(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func TreeTree(c *gin.Context) {
	result, err := client.TreeTree(context.TODO(), &empty.Empty{})
	resp.MicroResp(c, result, err)
}

func MenuTree(c *gin.Context) {
	result, err := client.MenuTree(context.TODO(), &empty.Empty{})
	resp.MicroResp(c, result, err)
}

func RoleTree(c *gin.Context) {
	result, err := client.RoleTree(context.TODO(), &empty.Empty{})
	resp.MicroResp(c, result, err)
}
func EditQualification(c *gin.Context) {
	req := dbmodel.Qualification{}
	c.Bind(&req)
	result, err := client.EditQualification(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditQualifications(c *gin.Context) {
	req := basic.Qualifications{}
	c.Bind(&req)
	result, err := client.EditQualifications(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelQualification(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelQualification(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func QualificationByForeignId(c *gin.Context) {
	req := dbmodel.Id{
		Id: c.Param("id"),
	}
	result, err := client.QualificationByForeignId(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}

func TreeByType(c *gin.Context) {
	//req := basic.TreeType{}
	//c.Bind(&req)
	treetype, _ := strconv.Atoi(c.Param("type"))
	req := basic.TreeType{}
	switch treetype {
	case int(dbmodel.TreeType_AREA):
		req.TreeType = dbmodel.TreeType_AREA
	case int(dbmodel.TreeType_PRD):
		req.TreeType = dbmodel.TreeType_PRD
	case int(dbmodel.TreeType_QUA):
		req.TreeType = dbmodel.TreeType_QUA
	case int(dbmodel.TreeType_USER):
		req.TreeType = dbmodel.TreeType_USER
	case int(dbmodel.TreeType_FILE):
		req.TreeType = dbmodel.TreeType_FILE
	case int(dbmodel.TreeType_DEP):
		req.TreeType = dbmodel.TreeType_DEP
	case int(dbmodel.TreeType_MENU):
		req.TreeType = dbmodel.TreeType_MENU
	case int(dbmodel.TreeType_LOGIN):
		req.TreeType = dbmodel.TreeType_LOGIN
	case int(dbmodel.TreeType_OTHER):
		req.TreeType = dbmodel.TreeType_OTHER
	default:
		resp.APIError(c, "暂时不支持该类型")
		return
	}
	result, err := client.TreeByType(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func EditArea(c *gin.Context) {
	req := dbmodel.SysArea{}
	c.Bind(&req)
	result, err := client.EditArea(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func DelArea(c *gin.Context) {
	req := dbmodel.Id{}
	c.Bind(&req)
	result, err := client.DelArea(context.TODO(), &req)
	resp.MicroResp(c, result, err)
}
func AreaList(c *gin.Context) {
	req := dbmodel.PageReq{}
	c.Bind(&req)
	result, err := client.AreaList(context.TODO(), &req)
	var rs []dbmodel.SysArea
	for _, any := range result.Data {
		var r dbmodel.SysArea
		ptypes.UnmarshalAny(any, &r)
		rs = append(rs, r)
	}
	resp.MicroTotalResp(c, result.Total, rs, err)
	//resp.MicroResp(c, result, err)
}

func AreaTree(c *gin.Context) {
	result, err := client.AreaTree(context.TODO(), &empty.Empty{})
	resp.MicroResp(c, result, err)
}
