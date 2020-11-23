package server

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/utils/mzjinit"
)

var (
	cliName     = "basicCli"
	svName      = "basicSrv"
	basicClient basic.BasicSrvService
	Conf        models.APIConfig
)

func init() {
	if err := mzjinit.Default(&Conf); err != nil {
		log.Fatal(err)
	}
	service := Conf.Services[cliName]
	basicClient = basic.NewBasicSrvService(Conf.Services[svName].Name, service.NewRoundSrv().Options().Client)
}

/*
func CodeVerify(emailOrPhone, code string) (bool, error) {
	req := &send.CodeVerifyReq{
		Code:         code,
		EmailOrPhone: emailOrPhone,
	}
	resp, err := sendClient.CodeVerify(context.Background(), req)
	if err != nil {
		fmt.Println("读取服务失败", err)
		return false, err
	}
	return resp.Verify, nil
}*/
