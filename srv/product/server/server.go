package server

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/utils/mzjinit"
)

var (
	Conf models.APIConfig
)

func init() {
	if err := mzjinit.Default(&Conf); err != nil {
		log.Fatal(err)
	}
}
