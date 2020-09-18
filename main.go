

package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/utils/mzjgorm"
	"qshapi/utils/mzjinit"
)
var (
	conf models.APIConfig
	dbConfig *mzjgorm.DbConfig
)
type Server struct {

}
func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
	dbConfig=mzjgorm.NewDbConfig(conf.DbConfig)
	if dbConfig == nil {
		log.Fatal("数据库连接初始化失败",conf.DbConfig)
	}
}
func main() {
	dbInit()
	/*w:=web.NewService(web.Address(":8000"),web.Icon("favicon.ico"))
	w.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("你好"))
	})
	w.Run()*/
	/*s:=v2.Service{}
	web:=s.NewRoundWeb()
	web.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("你好"))
	})
	web.Run()*/
	/*s:=v2.Service{}
	sv:=s.NewRoundSrv()
	sv.Run()*/
	/*s:=v2.Service{
		Name: "com.weixiao.test",
		Port: 8875,
		Ip: "127.0.0.1",
		Describe: "这是个描述",
		Version: "1.0.0",
	}
	web:=s.NewWeb()
	web.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("你好"))
	})
	web.Run()*/
	/*s:=v2.Service{
		Name: "com.weixiao.test",
		Port: 8875,
		//Ip: "127.0.0.1",
		Describe: "这是个描述",
		Version: "1.0.0",
	}
	sv:=s.NewSrv()
	sv.Run()*/
}

func dbInit(){
	db:=dbConfig.New()
	defer db.Close()
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
	db.Debug().AutoMigrate(
		&models.SysMenu{},
		&models.SysUser{},
		&models.SysRole{},
		&models.SysGroup{},
		&models.SysTree{},
		&models.SysFile{},
		&models.SysShop{},
		&models.SysShopCustomer{},
		&models.LogisticsAddress{},
		&models.Product{},
		&models.ProductSku{},
		&models.ProductLog{},
		&models.PartServant{},
		&models.Qualifications{},
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