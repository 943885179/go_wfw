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
	EditApi(req *sysuser.ApiReq, resp *sysuser.EditResp) error
	DelApi(req *sysuser.DelReq, resp *sysuser.EditResp) error
}

func NewAPI() IApi {
	return &Api{}
}

type Api struct{}

func (*Api) DelApi(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysAPI{}, req.Id).Error
}
func (*Api) EditApi(req *sysuser.ApiReq, resp *sysuser.EditResp) error {
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
		resp.Id = api.Id
		return db.Updates(api).Error
	} else { //添加
		mzjstruct.CopyStruct(req, api)
		api.Id = mzjuuid.WorkerDefault()
		resp.Id = api.Id
		return db.Create(api).Error
	}
}
