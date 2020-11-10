package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/file"
	"qshapi/srv/file/handler"
	"qshapi/utils/mzjinit"
)

var (
	svName = "fileSrv"
	conf   models.APIConfig
)

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}
func main() {
	service := conf.Services[svName]
	fmt.Println(service)
	s := service.NewSrv()
	file.RegisterFileSrvHandler(s.Server(), handler.Handler{})
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
func majorityElement(nums []int) int {
	var m map[int]int
	for _, r := range nums {
		m[r] = m[r] + 1
	}
	for x, y := range m {
		if y > len(nums)/2 {
			return x
		}
	}
	return 0
}
