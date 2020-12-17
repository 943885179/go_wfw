package main

import (
	"log"
	"qshapi/models"
	"qshapi/utils/mzjinit"
)

var (
	conf models.APIConfig
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// dbInit()
	// initArea()
	initValue()
	/*initTree()
	initMenu()
	initApi()
	initSrv()
	initAdmin()
	//test()
	db := configs.DbConfig.New().Model(&models.Product{})
	for i := 1; i < 1000; i++ {
		prd := models.Product{
			Id:          mzjuuid.WorkerDefaultStr(configs.WorkerId),
			GoodsCode:   fmt.Sprintf("%08d", i),
			GoodsName:   "商品" + strconv.Itoa(i),
			GoodsByname: "商品" + strconv.Itoa(i),
			Sort:        i,
		}
		db.Create(&prd)
	}*/
}
