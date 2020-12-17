package main

import (
	"fmt"
	"qshapi/models"
	"qshapi/utils/mzjjson"
)

func initArea() {
	db := conf.DbConfig.New()
	area := []models.SysArea{}
	mzjjson.JSONReadEntity("area.json", &area)
	fmt.Println(len(area))
	x := [][]models.SysArea{}
	ars := []models.SysArea{}
	for i := 0; i < len(area); i++ {
		ars = append(ars, area[i])
		if i%50 == 0 || i == len(area)-1 {
			x = append(x, ars)
			ars = []models.SysArea{}
		}
	}
	for _, r := range x {
		db.Create(&r)
	}
}
