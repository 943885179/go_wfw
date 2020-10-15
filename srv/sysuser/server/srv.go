package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type ISrv interface {
	EditSrv(req *sysuser.SrvReq) error
	DelSrv(req *sysuser.DelReq) error
}

func NewSrv() ISrv {
	return &Srv{}
}

type Srv struct{}

func (*Srv) DelSrv(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	return db.Delete(models.SysSrv{}, req.Id).Error
}

func (*Srv) EditSrv(req *sysuser.SrvReq) error {
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
		return db.Updates(Srv).Error
	} else { //添加
		Srv.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, Srv)
		return db.Create(Srv).Error
	}
}
