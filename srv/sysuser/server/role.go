package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IRole interface {
	EditRole(req *sysuser.RoleReq, resp *sysuser.EditResp) error
	DelRole(req *sysuser.DelReq, resp *sysuser.EditResp) error
}

func NewRole() IRole {
	return &Role{}
}

type Role struct{}

func (*Role) DelRole(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysRole{}, req.Id).Error
}

func (*Role) EditRole(req *sysuser.RoleReq, resp *sysuser.EditResp) error {
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
		resp.Id = Role.Id
		db.Model(&Role).Association("SysAPIs").Clear() //先清空关联再插入
		db.Model(&Role).Association("SysSrvs").Clear()
		db.Model(&Role).Association("Menus").Clear()

		if len(req.ApiId) != 0 {
			db.Where(&req.ApiId).Find(&Role.SysAPIs)
		}
		if len(req.SrvId) != 0 {
			db.Where(&req.SrvId).Find(&Role.SysSrvs)
		}
		if len(req.MenId) != 0 {
			db.Where(&req.MenId).Find(&Role.Menus)
		}
		return db.Updates(Role).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Role)
		Role.Id = mzjuuid.WorkerDefault()
		resp.Id = Role.Id
		if len(req.ApiId) != 0 {
			db.Where(&req.ApiId).Find(&Role.SysAPIs)
		}
		if len(req.SrvId) != 0 {
			db.Where(&req.SrvId).Find(&Role.SysSrvs)
		}
		if len(req.MenId) != 0 {
			db.Where(&req.MenId).Find(&Role.Menus)
		}
		return db.Create(Role).Error
	}
}
