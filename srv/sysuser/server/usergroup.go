package server

import (
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IUserGroup interface {
	EditUserGroup(req *sysuser.UserGroupReq) error
	DelUserGroup(req *sysuser.DelReq) error
}

func NewUserGroup() IUserGroup {
	return &UserGroup{}
}

type UserGroup struct{}

func (*UserGroup) DelUserGroup(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	return db.Delete(models.SysGroup{}, req.Id).Error
}

func (*UserGroup) EditUserGroup(req *sysuser.UserGroupReq) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	UserGroup := &models.SysGroup{}
	if req.Id > 0 { //修改0
		if err := db.First(UserGroup, req.Id).Error; err != nil {
			return err
		}
		mzjstruct.CopyStruct(req, UserGroup)
		return db.Updates(UserGroup).Error
	} else { //添加
		UserGroup.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, UserGroup)
		return db.Create(UserGroup).Error
	}
}
