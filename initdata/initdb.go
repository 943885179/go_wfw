package main

import (
	"fmt"
	"qshapi/models"
)

func dbInit() {
	db := conf.DbConfig.New()
	db.AutoMigrate(
		//***********服务一：基础信息服务***************
		&models.SysValue{},     //key_value存储
		&models.SysTree{},      //tree存储
		&models.SysMenu{},      //菜单
		&models.SysArea{},      //地址
		&models.SysWebApi{},    //webApi
		&models.SysSrv{},       //rpc 服务
		&models.SysRole{},      //角色
		&models.SysUserGroup{}, //用户组
		&models.SysUser{},      //用户
		//***********服务二：文件服务***************
		&models.SysFile{}, // 文件
		//***********服务三：店铺服务***************
		&models.SysShop{}, //商铺
		&models.SysShopCustomer{},
		&models.LogisticsAddress{},
		&models.Express{},
		&models.Freight{},
		//***********服务四：商品服务***************
		&models.Product{},
		&models.ProductSku{},
		&models.ProductLog{},

		//***********服务五：分佣服务***************
		&models.PartServant{},
		//***********服务六：资质服务***************
		&models.Qualification{},
		//&models.QualificationsRange{},

		//***********服务七：订单服务***************
		&models.Orders{},
		&models.OrderItem{},
		&models.OrderLog{},
		&models.OrderEvaluate{},
		&models.OrderItemPartServant{},

		//***********服务八：购物车服务***************
		&models.Cart{},
		//***********服务九：优惠卷服务***************
		&models.Prop{},
		&models.PropLog{},
	)
	fmt.Println("数据库初始化成功")
}
