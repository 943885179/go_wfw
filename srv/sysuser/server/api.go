package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IApi interface {
	EditApi(req *sysuser.SysApi, resp *sysuser.EditResp) error
	DelApi(req *sysuser.DelReq, resp *sysuser.EditResp) error
	ApiList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewAPI() IApi {
	return &Api{}
}

type Api struct{}

func (a *Api) ApiList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var t []models.SysApi
	db := Conf.DbConfig.New().Model(&models.SysApi{})
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r sysuser.SysApi
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Api) DelApi(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysApi{}, req.Id).Error
}
func (*Api) EditApi(req *sysuser.SysApi, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	api := &models.SysApi{}
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
