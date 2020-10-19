package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/utils/mzjinit"
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
