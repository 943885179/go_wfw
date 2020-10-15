package server

import (
	"errors"
	"fmt"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IApi interface {
	EditApi(req *sysuser.ApiReq) error
	DelApi(req *sysuser.DelReq) error
}

func NewAPI() IApi {
	return &Api{}
}

type Api struct{}

func (*Api) DelApi(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	return db.Delete(models.SysAPI{}, req.Id).Error
}
func (*Api) EditApi(req *sysuser.ApiReq) error {
	fmt.Println(req)
	db := Conf.DbConfig.New()
	//defer db.Close()
	api := &models.SysAPI{}
	if req.Id > 0 { //修改0
		//if db.FirstOrInit(api, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(api, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, api)
		return db.Updates(api).Error
	} else { //添加
		api.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, api)
		fmt.Println(api)
		return db.Create(api).Error
	}
}
