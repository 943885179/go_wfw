package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjmd5"
)

var (
	conf models.APIConfig
)

type Server struct {
}

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	dbInit()
	initApi()
	initSrv()
	initAdmin()
}

func dbInit() {
	db := conf.DbConfig.New()
	/*menus := []models.SysTree{
		{
			Text: "地址管理",
			Code: "00000000",
		},
		{
			Text: "商品类别管理",
			Code: "00000001",
		},
	}
	db.Save(&menus)*/
	db.AutoMigrate(
		&models.SysMenu{},
		&models.SysApi{},
		&models.SysSrv{},
		&models.SysRole{},
		&models.SysGroup{},
		&models.SysTree{},
		&models.SysUser{},
		&models.SysFile{},
		&models.SysShop{},
		&models.SysShopCustomer{},
		&models.LogisticsAddress{},
		&models.Product{},
		&models.ProductSku{},
		&models.ProductLog{},
		&models.PartServant{},
		&models.Qualification{},
		&models.QualificationsRange{},
		&models.Express{},
		&models.Freight{},
		&models.Orders{},
		&models.OrderItem{},
		&models.OrderLog{},
		&models.OrderEvaluate{},
		&models.OrderItemPartServant{},
		&models.Cart{},
		&models.Prop{},
		&models.PropLog{},
	)
	fmt.Println("数据库初始化成功")
}

func initSrv() {
	db := conf.DbConfig.New()
	srv := []models.SysSrv{
		{Id: 1, Service: "com.weixiao.api.user", Method: "Login", SrvExplain: ""},
		{Id: 2, Service: "com.weixiao.api.user", Method: "Registry", SrvExplain: ""},
		{Id: 3, Service: "com.weixiao.api.user", Method: "ChangePassword", SrvExplain: ""},
		{Id: 4, Service: "com.weixiao.api.user", Method: "UserInfoList", SrvExplain: ""},
		{Id: 5, Service: "com.weixiao.api.user", Method: "EditUser", SrvExplain: ""},

		{Id: 6, Service: "com.weixiao.api.user", Method: "EditRole", SrvExplain: ""},
		{Id: 7, Service: "com.weixiao.api.user", Method: "DelRole", SrvExplain: ""},
		{Id: 8, Service: "com.weixiao.api.user", Method: "RoleList", SrvExplain: ""},

		{Id: 9, Service: "com.weixiao.api.user", Method: "EditUserGroup", SrvExplain: ""},
		{Id: 10, Service: "com.weixiao.api.user", Method: "DelUserGroup", SrvExplain: ""},
		{Id: 11, Service: "com.weixiao.api.user", Method: "UserGroupList", SrvExplain: ""},

		{Id: 12, Service: "com.weixiao.api.user", Method: "EditMenu", SrvExplain: ""},
		{Id: 13, Service: "com.weixiao.api.user", Method: "DelMenu", SrvExplain: ""},
		{Id: 14, Service: "com.weixiao.api.user", Method: "MenuList", SrvExplain: ""},

		{Id: 15, Service: "com.weixiao.api.user", Method: "EditApi", SrvExplain: ""},
		{Id: 16, Service: "com.weixiao.api.user", Method: "DelApi", SrvExplain: ""},
		{Id: 17, Service: "com.weixiao.api.user", Method: "ApiList", SrvExplain: ""},

		{Id: 18, Service: "com.weixiao.api.user", Method: "EditSrv", SrvExplain: ""},
		{Id: 19, Service: "com.weixiao.api.user", Method: "DelSrv", SrvExplain: ""},
		{Id: 20, Service: "com.weixiao.api.user", Method: "SrvList", SrvExplain: ""},

		{Id: 21, Service: "com.weixiao.api.user", Method: "EditTree", SrvExplain: ""},
		{Id: 22, Service: "com.weixiao.api.user", Method: "DelTree", SrvExplain: ""},
		{Id: 23, Service: "com.weixiao.api.user", Method: "TreeList", SrvExplain: ""},

		{Id: 24, Service: "com.weixiao.api.file", Method: "UploadFile", SrvExplain: ""},
		{Id: 25, Service: "com.weixiao.api.file", Method: "GetFile", SrvExplain: ""},

		{Id: 27, Service: "com.weixiao.api.send", Method: "sendCode", SrvExplain: ""},
		{Id: 28, Service: "com.weixiao.api.send", Method: "send", SrvExplain: ""},
		{Id: 29, Service: "com.weixiao.api.send", Method: "sendAll", SrvExplain: ""},
		{Id: 30, Service: "com.weixiao.api.send", Method: "codeVerify", SrvExplain: ""},

		{Id: 31, Service: "com.weixiao.api.shop", Method: "Editshop", SrvExplain: ""},
		{Id: 32, Service: "com.weixiao.api.shop", Method: "Delshop", SrvExplain: ""},
		{Id: 33, Service: "com.weixiao.api.shop", Method: "shopList", SrvExplain: ""},
	}
	db.Create(srv)
}
func initApi() {
	db := conf.DbConfig.New()
	srv := []models.SysApi{
		{Id: 1, Service: "com.weixiao.web.user", Method: "Login", ApiExplain: ""},
		{Id: 2, Service: "com.weixiao.web.user", Method: "Registry", ApiExplain: ""},
		{Id: 3, Service: "com.weixiao.web.user", Method: "ChangePassword", ApiExplain: ""},
		{Id: 4, Service: "com.weixiao.web.user", Method: "UserInfoList", ApiExplain: ""},
		{Id: 5, Service: "com.weixiao.web.user", Method: "EditUser", ApiExplain: ""},

		{Id: 6, Service: "com.weixiao.web.user", Method: "EditRole", ApiExplain: ""},
		{Id: 7, Service: "com.weixiao.web.user", Method: "DelRole", ApiExplain: ""},
		{Id: 8, Service: "com.weixiao.web.user", Method: "RoleList", ApiExplain: ""},

		{Id: 9, Service: "com.weixiao.web.user", Method: "EditUserGroup", ApiExplain: ""},
		{Id: 10, Service: "com.weixiao.web.user", Method: "DelUserGroup", ApiExplain: ""},
		{Id: 11, Service: "com.weixiao.web.user", Method: "UserGroupList", ApiExplain: ""},

		{Id: 12, Service: "com.weixiao.web.user", Method: "EditMenu", ApiExplain: ""},
		{Id: 13, Service: "com.weixiao.web.user", Method: "DelMenu", ApiExplain: ""},
		{Id: 14, Service: "com.weixiao.web.user", Method: "MenuList", ApiExplain: ""},

		{Id: 15, Service: "com.weixiao.web.user", Method: "EditApi", ApiExplain: ""},
		{Id: 16, Service: "com.weixiao.web.user", Method: "DelApi", ApiExplain: ""},
		{Id: 17, Service: "com.weixiao.web.user", Method: "ApiList", ApiExplain: ""},

		{Id: 18, Service: "com.weixiao.web.user", Method: "EditSrv", ApiExplain: ""},
		{Id: 19, Service: "com.weixiao.web.user", Method: "DelSrv", ApiExplain: ""},
		{Id: 20, Service: "com.weixiao.web.user", Method: "SrvList", ApiExplain: ""},

		{Id: 21, Service: "com.weixiao.web.user", Method: "EditTree", ApiExplain: ""},
		{Id: 22, Service: "com.weixiao.web.user", Method: "DelTree", ApiExplain: ""},
		{Id: 23, Service: "com.weixiao.web.user", Method: "TreeList", ApiExplain: ""},

		{Id: 24, Service: "com.weixiao.web.file", Method: "upload", ApiExplain: ""},
		{Id: 25, Service: "com.weixiao.web.file", Method: "uploadMutiple", ApiExplain: ""},
		{Id: 26, Service: "com.weixiao.web.file", Method: "showFile", ApiExplain: ""},
		{Id: 27, Service: "com.weixiao.web.file", Method: "fileById", ApiExplain: ""},

		{Id: 28, Service: "com.weixiao.web.send", Method: "sendCode", ApiExplain: ""},
		{Id: 29, Service: "com.weixiao.web.send", Method: "send", ApiExplain: ""},
		{Id: 30, Service: "com.weixiao.web.send", Method: "sendAll", ApiExplain: ""},
		{Id: 31, Service: "com.weixiao.web.send", Method: "codeVerify", ApiExplain: ""},

		{Id: 32, Service: "com.weixiao.web.shop", Method: "Editshop", ApiExplain: ""},
		{Id: 33, Service: "com.weixiao.web.shop", Method: "Delshop", ApiExplain: ""},
		{Id: 34, Service: "com.weixiao.web.shop", Method: "shopList", ApiExplain: ""},
	}
	db.Create(srv)
}

func initAdmin() { //超级管理员
	db := conf.DbConfig.New()
	var a []models.SysApi
	var s []models.SysSrv
	var m []models.SysMenu
	db.Model(&models.SysApi{}).Find(&a)
	db.Model(&models.SysSrv{}).Find(&s)
	db.Model(&models.SysMenu{}).Find(&m)
	//创建一个超级管理员角色
	r := models.SysRole{
		RoleName:    "超级管理员",
		RoleExplain: "超级管理员有所有权限",
		Menus:       m,
		Apis:        a,
		Srvs:        s,
		Id:          1, //mzjuuid.WorkerDefault(),
	}
	//db.Create(r)
	//创建一个用户组
	gr := models.SysGroup{
		GroupName:    "超级用户组",
		GroupExplain: "超级用户组",
		Id:           1, // mzjuuid.WorkerDefault(),
	}
	gr.Roles = append(gr.Roles, r)
	u := models.SysUser{
		Id:           1, // mzjuuid.WorkerDefault(),
		UserName:     "admin",
		UserPhone:    "18206840781",
		UserEmail:    "943885179@qq.com",
		UserPassword: mzjmd5.MD5("123"),
		UserQq:       "943885179",
		UserWx:       "18206840781",
	}
	u.Groups = append(u.Groups, gr)
	//u.Roles = append(u.Roles, r) //这个可以不用
	db.Create(&u)
}
