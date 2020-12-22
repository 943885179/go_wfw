package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"

	"github.com/golang/protobuf/ptypes"
)

type IApi interface {
	EditApi(req *dbmodel.SysApi, resp *dbmodel.Id) error
	DelApi(req *dbmodel.Id, resp *dbmodel.Id) error
	ApiList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	ApiById(id *dbmodel.Id, api *dbmodel.SysApi) error
	ApiListByUser(user *dbmodel.SysUser, api *dbmodel.OnlyApi) error
}

func NewAPI() IApi {
	return &Api{}
}

type Api struct{}

func (a *Api) ApiListByUser(req *dbmodel.SysUser, resp *dbmodel.OnlyApi) error {
	var all []models.SysApi
	Conf.DbConfig.New().Find(&all)
	/*var ut models.SysTree
	Conf.DbConfig.New().First(&ut, req.UserTypeId)*/
	var hasIds []string
	for _, group := range req.Groups {
		for _, role := range group.Roles {
			for _, a := range role.Apis {
				hasIds = append(hasIds, a.Id)
			}
		}
	}
	for _, role := range req.Roles {
		for _, a := range role.Apis {
			hasIds = append(hasIds, a.Id)
		}
	}
	for _, a := range all {
		for _, id := range hasIds {
			if id == a.Id {
				var d dbmodel.SysApi
				mzjstruct.CopyStruct(&a, &d)
				resp.Apis = append(resp.Apis, &d)
				break
			}
		}
	}
	return nil
}

func (a *Api) ApiById(id *dbmodel.Id, api *dbmodel.SysApi) error {

	return Conf.DbConfig.New().Model(&models.SysApi{}).First(api, id.Id).Error
}

func (a *Api) ApiList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.SysApi
	db := Conf.DbConfig.New().Model(&models.SysApi{}).Order("id")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, a := range t {
		var r dbmodel.SysApi
		mzjstruct.CopyStruct(&a, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Api) DelApi(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysApi{}, req.Id).Error
}
func (*Api) EditApi(req *dbmodel.SysApi, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	api := &models.SysApi{}
	if len(req.Id) > 0 { //修改0
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
	} else {
		mzjstruct.CopyStruct(req, api)
		api.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		resp.Id = api.Id
		return db.Create(api).Error
	}
}
