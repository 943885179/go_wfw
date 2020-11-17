package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
	"strings"
)

type ISrv interface {
	EditSrv(req *dbmodel.SysSrv, resp *dbmodel.Id) error
	DelSrv(req *dbmodel.Id, resp *dbmodel.Id) error
	SrvList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	SrvById(id *dbmodel.Id, srv *dbmodel.SysSrv) error
	SrvListByUser(user *dbmodel.SysUser, srv *dbmodel.OnlySrv) error
}

func NewSrv() ISrv {
	return &Srv{}
}

type Srv struct{}

func (a *Srv) SrvListByUser(req *dbmodel.SysUser, resp *dbmodel.OnlySrv) error {
	var all []models.SysSrv
	Conf.DbConfig.New().Find(&all)
	if strings.ToLower(req.UserType.Code) == strings.ToLower("admin") { // 超级管理员有所有的菜单权限，不受约束
		for _, a := range all {
			var d dbmodel.SysSrv
			mzjstruct.CopyStruct(&a, &d)
			resp.Srvs = append(resp.Srvs, &d)
		}
		return nil
	} else {
		var hasIds []string
		for _, group := range req.Groups {
			for _, role := range group.Roles {
				for _, a := range role.Srvs {
					hasIds = append(hasIds, a.Id)
				}
			}
		}
		for _, role := range req.Roles {
			for _, a := range role.Srvs {
				hasIds = append(hasIds, a.Id)
			}
		}
		for _, a := range all {
			for _, id := range hasIds {
				if id == a.Id {
					var d dbmodel.SysSrv
					mzjstruct.CopyStruct(&a, &d)
					resp.Srvs = append(resp.Srvs, &d)
					break
				}
			}
		}
		return nil
	}
}
func (a *Srv) SrvById(id *dbmodel.Id, srv *dbmodel.SysSrv) error {
	return Conf.DbConfig.New().Model(&models.SysSrv{}).First(srv, id.Id).Error
}

func (s *Srv) SrvList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.SysSrv
	db := Conf.DbConfig.New().Model(&models.SysSrv{})
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r dbmodel.SysSrv
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Srv) DelSrv(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysSrv{}, req.Id).Error
}

func (*Srv) EditSrv(req *dbmodel.SysSrv, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	Srv := &models.SysSrv{}
	if len(req.Id) > 0 { //修改0
		if err := db.First(Srv, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Srv)
		resp.Id = Srv.Id
		return db.Updates(Srv).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Srv)
		Srv.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		resp.Id = Srv.Id
		return db.Create(Srv).Error
	}
}
