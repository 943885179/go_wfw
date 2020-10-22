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
	initMenu()
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
		{Id: 1, Service: "com.weixiao.api.basic", Method: "Login", SrvExplain: ""},
		{Id: 2, Service: "com.weixiao.api.basic", Method: "Registry", SrvExplain: ""},
		{Id: 3, Service: "com.weixiao.api.basic", Method: "ChangePassword", SrvExplain: ""},
		{Id: 4, Service: "com.weixiao.api.basic", Method: "UserInfoList", SrvExplain: ""},
		{Id: 5, Service: "com.weixiao.api.basic", Method: "EditUser", SrvExplain: ""},

		{Id: 6, Service: "com.weixiao.api.basic", Method: "EditRole", SrvExplain: ""},
		{Id: 7, Service: "com.weixiao.api.basic", Method: "DelRole", SrvExplain: ""},
		{Id: 8, Service: "com.weixiao.api.basic", Method: "RoleList", SrvExplain: ""},

		{Id: 9, Service: "com.weixiao.api.basic", Method: "EditUserGroup", SrvExplain: ""},
		{Id: 10, Service: "com.weixiao.api.basic", Method: "DelUserGroup", SrvExplain: ""},
		{Id: 11, Service: "com.weixiao.api.basic", Method: "UserGroupList", SrvExplain: ""},

		{Id: 12, Service: "com.weixiao.api.basic", Method: "EditMenu", SrvExplain: ""},
		{Id: 13, Service: "com.weixiao.api.basic", Method: "DelMenu", SrvExplain: ""},
		{Id: 14, Service: "com.weixiao.api.basic", Method: "MenuList", SrvExplain: ""},

		{Id: 15, Service: "com.weixiao.api.basic", Method: "EditApi", SrvExplain: ""},
		{Id: 16, Service: "com.weixiao.api.basic", Method: "DelApi", SrvExplain: ""},
		{Id: 17, Service: "com.weixiao.api.basic", Method: "ApiList", SrvExplain: ""},

		{Id: 18, Service: "com.weixiao.api.basic", Method: "EditSrv", SrvExplain: ""},
		{Id: 19, Service: "com.weixiao.api.basic", Method: "DelSrv", SrvExplain: ""},
		{Id: 20, Service: "com.weixiao.api.basic", Method: "SrvList", SrvExplain: ""},

		{Id: 21, Service: "com.weixiao.api.basic", Method: "EditTree", SrvExplain: ""},
		{Id: 22, Service: "com.weixiao.api.basic", Method: "DelTree", SrvExplain: ""},
		{Id: 23, Service: "com.weixiao.api.basic", Method: "TreeList", SrvExplain: ""},

		{Id: 24, Service: "com.weixiao.api.file", Method: "UploadFile", SrvExplain: ""},
		{Id: 25, Service: "com.weixiao.api.file", Method: "GetFile", SrvExplain: ""},

		{Id: 27, Service: "com.weixiao.api.send", Method: "sendCode", SrvExplain: ""},
		{Id: 28, Service: "com.weixiao.api.send", Method: "send", SrvExplain: ""},
		{Id: 29, Service: "com.weixiao.api.send", Method: "sendAll", SrvExplain: ""},
		{Id: 30, Service: "com.weixiao.api.send", Method: "codeVerify", SrvExplain: ""},

		{Id: 31, Service: "com.weixiao.api.product", Method: "EditProduct", SrvExplain: ""},
		{Id: 32, Service: "com.weixiao.api.product", Method: "DelProduct", SrvExplain: ""},
		{Id: 33, Service: "com.weixiao.api.product", Method: "productList", SrvExplain: ""},
	}
	db.Create(srv)
}
func initApi() {
	db := conf.DbConfig.New()
	srv := []models.SysApi{
		{Id: 1, Service: "com.weixiao.web.basic", Method: "Login", ApiExplain: ""},
		{Id: 2, Service: "com.weixiao.web.basic", Method: "Registry", ApiExplain: ""},
		{Id: 3, Service: "com.weixiao.web.basic", Method: "ChangePassword", ApiExplain: ""},
		{Id: 4, Service: "com.weixiao.web.basic", Method: "UserInfoList", ApiExplain: ""},
		{Id: 5, Service: "com.weixiao.web.basic", Method: "EditUser", ApiExplain: ""},

		{Id: 6, Service: "com.weixiao.web.basic", Method: "EditRole", ApiExplain: ""},
		{Id: 7, Service: "com.weixiao.web.basic", Method: "DelRole", ApiExplain: ""},
		{Id: 8, Service: "com.weixiao.web.basic", Method: "RoleList", ApiExplain: ""},

		{Id: 9, Service: "com.weixiao.web.basic", Method: "EditUserGroup", ApiExplain: ""},
		{Id: 10, Service: "com.weixiao.web.basic", Method: "DelUserGroup", ApiExplain: ""},
		{Id: 11, Service: "com.weixiao.web.basic", Method: "UserGroupList", ApiExplain: ""},

		{Id: 12, Service: "com.weixiao.web.basic", Method: "EditMenu", ApiExplain: ""},
		{Id: 13, Service: "com.weixiao.web.basic", Method: "DelMenu", ApiExplain: ""},
		{Id: 14, Service: "com.weixiao.web.basic", Method: "MenuList", ApiExplain: ""},

		{Id: 15, Service: "com.weixiao.web.basic", Method: "EditApi", ApiExplain: ""},
		{Id: 16, Service: "com.weixiao.web.basic", Method: "DelApi", ApiExplain: ""},
		{Id: 17, Service: "com.weixiao.web.basic", Method: "ApiList", ApiExplain: ""},

		{Id: 18, Service: "com.weixiao.web.basic", Method: "EditSrv", ApiExplain: ""},
		{Id: 19, Service: "com.weixiao.web.basic", Method: "DelSrv", ApiExplain: ""},
		{Id: 20, Service: "com.weixiao.web.basic", Method: "SrvList", ApiExplain: ""},

		{Id: 21, Service: "com.weixiao.web.basic", Method: "EditTree", ApiExplain: ""},
		{Id: 22, Service: "com.weixiao.web.basic", Method: "DelTree", ApiExplain: ""},
		{Id: 23, Service: "com.weixiao.web.basic", Method: "TreeList", ApiExplain: ""},

		{Id: 24, Service: "com.weixiao.web.file", Method: "upload", ApiExplain: ""},
		{Id: 25, Service: "com.weixiao.web.file", Method: "uploadMutiple", ApiExplain: ""},
		{Id: 26, Service: "com.weixiao.web.file", Method: "showFile", ApiExplain: ""},
		{Id: 27, Service: "com.weixiao.web.file", Method: "fileById", ApiExplain: ""},

		{Id: 28, Service: "com.weixiao.web.send", Method: "sendCode", ApiExplain: ""},
		{Id: 29, Service: "com.weixiao.web.send", Method: "send", ApiExplain: ""},
		{Id: 30, Service: "com.weixiao.web.send", Method: "sendAll", ApiExplain: ""},
		{Id: 31, Service: "com.weixiao.web.send", Method: "codeVerify", ApiExplain: ""},

		{Id: 32, Service: "com.weixiao.web.product", Method: "EditProduct", ApiExplain: ""},
		{Id: 33, Service: "com.weixiao.web.product", Method: "DelProduct", ApiExplain: ""},
		{Id: 34, Service: "com.weixiao.web.product", Method: "productList", ApiExplain: ""},
	}
	db.Create(srv)
}

func initMenu() {
	db := conf.DbConfig.New()
	srv := []models.SysMenu{
		{
			Text:             "主导航",
			I18n:             "menu.main",
			Group:            true,
			HideInBreadcrumb: true,
			Children: []models.SysMenu{
				{
					Text: "仪表盘",
					I18n: "menu.dashboard",
					Icon: "anticon-dashboard",
					Children: []models.SysMenu{
						{
							Text: "仪表盘V1",
							I18n: "menu.dashboard.v1",
							Icon: "anticon-dashboard",
							Link: "/dashboard/v1",
						},
					},
				},
			},
		},
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
