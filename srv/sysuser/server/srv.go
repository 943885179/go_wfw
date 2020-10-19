package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ISrv interface {
	EditSrv(req *sysuser.SysSrv, resp *sysuser.EditResp) error
	DelSrv(req *sysuser.DelReq, resp *sysuser.EditResp) error
	SrvList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewSrv() ISrv {
	return &Srv{}
}

type Srv struct{}

func (s *Srv) SrvList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var t []models.SysSrv
	db := Conf.DbConfig.New().Model(&models.SysSrv{})
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r sysuser.SysSrv
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Srv) DelSrv(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysSrv{}, req.Id).Error
}

func (*Srv) EditSrv(req *sysuser.SysSrv, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	Srv := &models.SysSrv{}
	if req.Id > 0 { //修改0
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
		Srv.Id = mzjuuid.WorkerDefault()
		resp.Id = Srv.Id
		return db.Create(Srv).Error
	}
}
