package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ISrv interface {
	EditSrv(req *sysuser.SrvReq, resp *sysuser.EditResp) error
	DelSrv(req *sysuser.DelReq, resp *sysuser.EditResp) error
}

func NewSrv() ISrv {
	return &Srv{}
}

type Srv struct{}

func (*Srv) DelSrv(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysSrv{}, req.Id).Error
}

func (*Srv) EditSrv(req *sysuser.SrvReq, resp *sysuser.EditResp) error {
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
