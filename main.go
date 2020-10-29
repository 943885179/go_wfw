package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/dbmodel"
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
	//test()
}
func test() {
	var i = 0
	for {
		i++
		var u models.SysUser
		conf.DbConfig.New().First(&u)
		fmt.Println(i, u.UserName)
	}
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
		{Id: "1", Service: "com.weixiao.api.basic", Method: "Login", SrvExplain: ""},
		{Id: "2", Service: "com.weixiao.api.basic", Method: "Registry", SrvExplain: ""},
		{Id: "3", Service: "com.weixiao.api.basic", Method: "ChangePassword", SrvExplain: ""},
		{Id: "4", Service: "com.weixiao.api.basic", Method: "UserInfoList", SrvExplain: ""},
		{Id: "5", Service: "com.weixiao.api.basic", Method: "EditUser", SrvExplain: ""},
		{Id: "6", Service: "com.weixiao.api.basic", Method: "UserById", SrvExplain: ""},

		{Id: "7", Service: "com.weixiao.api.basic", Method: "EditRole", SrvExplain: ""},
		{Id: "8", Service: "com.weixiao.api.basic", Method: "DelRole", SrvExplain: ""},
		{Id: "9", Service: "com.weixiao.api.basic", Method: "RoleList", SrvExplain: ""},
		{Id: "10", Service: "com.weixiao.api.basic", Method: "RoleById", SrvExplain: ""},

		{Id: "11", Service: "com.weixiao.api.basic", Method: "EditUserGroup", SrvExplain: ""},
		{Id: "12", Service: "com.weixiao.api.basic", Method: "DelUserGroup", SrvExplain: ""},
		{Id: "13", Service: "com.weixiao.api.basic", Method: "UserGroupList", SrvExplain: ""},
		{Id: "14", Service: "com.weixiao.api.basic", Method: "UserGroupById", SrvExplain: ""},

		{Id: "15", Service: "com.weixiao.api.basic", Method: "EditMenu", SrvExplain: ""},
		{Id: "16", Service: "com.weixiao.api.basic", Method: "DelMenu", SrvExplain: ""},
		{Id: "17", Service: "com.weixiao.api.basic", Method: "MenuList", SrvExplain: ""},
		{Id: "18", Service: "com.weixiao.api.basic", Method: "MenuById", SrvExplain: ""},

		{Id: "19", Service: "com.weixiao.api.basic", Method: "EditApi", SrvExplain: ""},
		{Id: "20", Service: "com.weixiao.api.basic", Method: "DelApi", SrvExplain: ""},
		{Id: "21", Service: "com.weixiao.api.basic", Method: "ApiList", SrvExplain: ""},
		{Id: "22", Service: "com.weixiao.api.basic", Method: "ApiById", SrvExplain: ""},

		{Id: "23", Service: "com.weixiao.api.basic", Method: "EditSrv", SrvExplain: ""},
		{Id: "24", Service: "com.weixiao.api.basic", Method: "DelSrv", SrvExplain: ""},
		{Id: "25", Service: "com.weixiao.api.basic", Method: "SrvList", SrvExplain: ""},
		{Id: "26", Service: "com.weixiao.api.basic", Method: "SrvById", SrvExplain: ""},

		{Id: "27", Service: "com.weixiao.api.basic", Method: "EditTree", SrvExplain: ""},
		{Id: "28", Service: "com.weixiao.api.basic", Method: "DelTree", SrvExplain: ""},
		{Id: "29", Service: "com.weixiao.api.basic", Method: "TreeList", SrvExplain: ""},
		{Id: "30", Service: "com.weixiao.api.basic", Method: "TreeById", SrvExplain: ""},

		{Id: "31", Service: "com.weixiao.api.basic", Method: "EditShop", SrvExplain: ""},
		{Id: "32", Service: "com.weixiao.api.basic", Method: "DelShop", SrvExplain: ""},
		{Id: "33", Service: "com.weixiao.api.basic", Method: "ShopList", SrvExplain: ""},
		{Id: "34", Service: "com.weixiao.api.basic", Method: "ShopById", SrvExplain: ""},

		{Id: "35", Service: "com.weixiao.api.file", Method: "UploadFile", SrvExplain: ""},
		{Id: "36", Service: "com.weixiao.api.file", Method: "GetFile", SrvExplain: ""},

		{Id: "37", Service: "com.weixiao.api.send", Method: "sendCode", SrvExplain: ""},
		{Id: "38", Service: "com.weixiao.api.send", Method: "send", SrvExplain: ""},
		{Id: "39", Service: "com.weixiao.api.send", Method: "sendAll", SrvExplain: ""},
		{Id: "40", Service: "com.weixiao.api.send", Method: "codeVerify", SrvExplain: ""},

		{Id: "41", Service: "com.weixiao.api.product", Method: "EditProduct", SrvExplain: ""},
		{Id: "42", Service: "com.weixiao.api.product", Method: "DelProduct", SrvExplain: ""},
		{Id: "43", Service: "com.weixiao.api.product", Method: "productList", SrvExplain: ""},

		{Id: "44", Service: "com.weixiao.api.basic", Method: "TreeTree", SrvExplain: ""},
		{Id: "45", Service: "com.weixiao.api.basic", Method: "MenuTree", SrvExplain: ""},
		{Id: "46", Service: "com.weixiao.api.basic", Method: "RoleTree", SrvExplain: ""},
	}
	db.Create(srv)
}
func initApi() {
	db := conf.DbConfig.New()
	srv := []models.SysApi{
		{Id: "1", Service: "com.weixiao.web.basic", Method: "Login", ApiExplain: ""},
		{Id: "2", Service: "com.weixiao.web.basic", Method: "Registry", ApiExplain: ""},
		{Id: "3", Service: "com.weixiao.web.basic", Method: "ChangePassword", ApiExplain: ""},
		{Id: "4", Service: "com.weixiao.web.basic", Method: "UserInfoList", ApiExplain: ""},
		{Id: "5", Service: "com.weixiao.web.basic", Method: "EditUser", ApiExplain: ""},
		{Id: "6", Service: "com.weixiao.web.basic", Method: "UserById", ApiExplain: ""},

		{Id: "7", Service: "com.weixiao.web.basic", Method: "EditRole", ApiExplain: ""},
		{Id: "8", Service: "com.weixiao.web.basic", Method: "DelRole", ApiExplain: ""},
		{Id: "9", Service: "com.weixiao.web.basic", Method: "RoleList", ApiExplain: ""},
		{Id: "10", Service: "com.weixiao.web.basic", Method: "RoleById", ApiExplain: ""},

		{Id: "11", Service: "com.weixiao.web.basic", Method: "EditUserGroup", ApiExplain: ""},
		{Id: "12", Service: "com.weixiao.web.basic", Method: "DelUserGroup", ApiExplain: ""},
		{Id: "13", Service: "com.weixiao.web.basic", Method: "UserGroupList", ApiExplain: ""},
		{Id: "14", Service: "com.weixiao.web.basic", Method: "UserGroupById", ApiExplain: ""},

		{Id: "15", Service: "com.weixiao.web.basic", Method: "EditMenu", ApiExplain: ""},
		{Id: "16", Service: "com.weixiao.web.basic", Method: "DelMenu", ApiExplain: ""},
		{Id: "17", Service: "com.weixiao.web.basic", Method: "MenuList", ApiExplain: ""},
		{Id: "18", Service: "com.weixiao.web.basic", Method: "MenuById", ApiExplain: ""},

		{Id: "19", Service: "com.weixiao.web.basic", Method: "EditApi", ApiExplain: ""},
		{Id: "20", Service: "com.weixiao.web.basic", Method: "DelApi", ApiExplain: ""},
		{Id: "21", Service: "com.weixiao.web.basic", Method: "ApiList", ApiExplain: ""},
		{Id: "22", Service: "com.weixiao.web.basic", Method: "ApiById", ApiExplain: ""},

		{Id: "23", Service: "com.weixiao.web.basic", Method: "EditSrv", ApiExplain: ""},
		{Id: "24", Service: "com.weixiao.web.basic", Method: "DelSrv", ApiExplain: ""},
		{Id: "25", Service: "com.weixiao.web.basic", Method: "SrvList", ApiExplain: ""},
		{Id: "26", Service: "com.weixiao.web.basic", Method: "SrvById", ApiExplain: ""},

		{Id: "27", Service: "com.weixiao.web.basic", Method: "EditTree", ApiExplain: ""},
		{Id: "28", Service: "com.weixiao.web.basic", Method: "DelTree", ApiExplain: ""},
		{Id: "29", Service: "com.weixiao.web.basic", Method: "TreeList", ApiExplain: ""},
		{Id: "30", Service: "com.weixiao.web.basic", Method: "TreeById", ApiExplain: ""},

		{Id: "31", Service: "com.weixiao.web.basic", Method: "EditShop", ApiExplain: ""},
		{Id: "32", Service: "com.weixiao.web.basic", Method: "DelShop", ApiExplain: ""},
		{Id: "33", Service: "com.weixiao.web.basic", Method: "ShopList", ApiExplain: ""},
		{Id: "34", Service: "com.weixiao.web.basic", Method: "ShopById", ApiExplain: ""},

		{Id: "37", Service: "com.weixiao.web.send", Method: "sendCode", ApiExplain: ""},
		{Id: "38", Service: "com.weixiao.web.send", Method: "send", ApiExplain: ""},
		{Id: "39", Service: "com.weixiao.web.send", Method: "sendAll", ApiExplain: ""},
		{Id: "40", Service: "com.weixiao.web.send", Method: "codeVerify", ApiExplain: ""},

		{Id: "41", Service: "com.weixiao.web.product", Method: "EditProduct", ApiExplain: ""},
		{Id: "42", Service: "com.weixiao.web.product", Method: "DelProduct", ApiExplain: ""},
		{Id: "43", Service: "com.weixiao.web.product", Method: "productList", ApiExplain: ""},

		{Id: "44", Service: "com.weixiao.web.file", Method: "upload", ApiExplain: ""},
		{Id: "45", Service: "com.weixiao.web.file", Method: "uploadMutiple", ApiExplain: ""},
		{Id: "46", Service: "com.weixiao.web.file", Method: "showFile", ApiExplain: ""},
		{Id: "47", Service: "com.weixiao.web.file", Method: "fileById", ApiExplain: ""},

		{Id: "48", Service: "com.weixiao.web.send", Method: "sendCode", ApiExplain: ""},
		{Id: "49", Service: "com.weixiao.web.send", Method: "send", ApiExplain: ""},
		{Id: "50", Service: "com.weixiao.web.send", Method: "sendAll", ApiExplain: ""},
		{Id: "51", Service: "com.weixiao.web.send", Method: "codeVerify", ApiExplain: ""},

		{Id: "52", Service: "com.weixiao.web.product", Method: "EditProduct", ApiExplain: ""},
		{Id: "53", Service: "com.weixiao.web.product", Method: "DelProduct", ApiExplain: ""},
		{Id: "54", Service: "com.weixiao.web.product", Method: "productList", ApiExplain: ""},

		{Id: "55", Service: "com.weixiao.web.basic", Method: "TreeTree", ApiExplain: ""},
		{Id: "56", Service: "com.weixiao.web.basic", Method: "MenuTree", ApiExplain: ""},
		{Id: "57", Service: "com.weixiao.web.basic", Method: "RoleTree", ApiExplain: ""},
		{Id: "58", Service: "com.weixiao.web.basic", Method: "Token", ApiExplain: ""},
	}
	db.Create(srv)
}

func initMenu() {
	db := conf.DbConfig.New()
	srv := []models.SysMenu{
		{
			Id:               "1",
			Key:              "1",
			Text:             "主导航",
			Title:            "主导航",
			I18N:             "menu.main",
			Group:            true,
			HideInBreadcrumb: true,
			Children: []models.SysMenu{
				{
					Id:    "2",
					Key:   "2",
					Text:  "仪表盘",
					Title: "仪表盘",
					I18N:  "menu.dashboard",
					Icon:  "anticon-dashboard",
					Children: []models.SysMenu{
						{
							Id:    "4",
							Key:   "4",
							Text:  "仪表盘V1",
							Title: "仪表盘V1",
							I18N:  "menu.dashboard.v1",
							Icon:  "anticon-dashboard",
							Link:  "/dashboard/v1",
						},
					}},
				{
					Id:    "3",
					Key:   "3",
					Text:  "基础信息管理",
					Title: "基础信息管理",
					I18N:  "menu.basic",
					Icon:  "anticon-setting",
					Children: []models.SysMenu{
						{
							Id:    "5",
							Key:   "5",
							Text:  "用户管理",
							Title: "用户管理",
							I18N:  "menu.basic.user",
							Icon:  "anticon-user",
							Link:  "/basic/user",
						},
						{
							Id:    "6",
							Key:   "6",
							Text:  "菜单管理",
							Title: "菜单管理",
							I18N:  "menu.basic.menu",
							Icon:  "anticon-dashboard",
							Link:  "/basic/menu",
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
		Title:       "超级管理员",
		RoleExplain: "超级管理员有所有权限",
		Menus:       m,
		Apis:        a,
		Srvs:        s,
		Id:          "1", //mzjuuid.WorkerDefault(),
	}
	//db.Create(r)
	//创建一个用户组
	gr := models.SysGroup{
		GroupName:    "超级用户组",
		GroupExplain: "超级用户组",
		Id:           "1", // mzjuuid.WorkerDefault(),
	}
	gr.Roles = append(gr.Roles, r)
	u := models.SysUser{
		Id:           "1", // mzjuuid.WorkerDefault(),
		UserName:     "admin",
		UserPhone:    "18206840781",
		UserEmail:    "943885179@qq.com",
		UserPassword: mzjmd5.MD5("123"),
		UserQq:       "943885179",
		UserWx:       "18206840781",
		UserType:     dbmodel.UserType_ADMIN,
	}
	u.Groups = append(u.Groups, gr) //admin可以不设置用户组权限了
	//u.Roles = append(u.Roles, r) //这个可以不用
	db.Create(&u)
	//db.Updates(&u)
}
