package server

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
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

/*func TokenResp(token string) (basic.LoginResp, error) {
	return mzjgin.TokenResp(token)
}*/

func EditQualifications(req *dbmodel.Qualification, resp *dbmodel.Id) error {
	resp, err := basicClient.EditQualification(context.TODO(), req)
	return err
}
