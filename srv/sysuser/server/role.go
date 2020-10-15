package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IRole interface {
	EditRole(req *sysuser.RoleReq) error
	DelRole(req *sysuser.DelReq) error
}

func NewRole() IRole {
	return &Role{}
}

type Role struct{}

func (*Role) DelRole(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	return db.Delete(models.SysRole{}, req.Id).Error
}

func (*Role) EditRole(req *sysuser.RoleReq) error {
	db := Conf.DbConfig.New()
	Role := &models.SysRole{}
	if req.Id > 0 { //修改0
		if err := db.First(Role, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Role)
		return db.Updates(Role).Error
	} else { //添加
		Role.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, Role)
		return db.Create(Role).Error
	}
}
