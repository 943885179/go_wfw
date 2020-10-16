package server

import (
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IUserGroup interface {
	EditUserGroup(req *sysuser.UserGroupReq, resp *sysuser.EditResp) error
	DelUserGroup(req *sysuser.DelReq, resp *sysuser.EditResp) error
}

func NewUserGroup() IUserGroup {
	return &UserGroup{}
}

type UserGroup struct{}

func (*UserGroup) DelUserGroup(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	//defer db.Close()
	return db.Delete(models.SysGroup{}, req.Id).Error
}

func (*UserGroup) EditUserGroup(req *sysuser.UserGroupReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	UserGroup := &models.SysGroup{}
	if req.Id > 0 { //修改0
		if err := db.First(UserGroup, req.Id).Error; err != nil {
			return err
		}
		mzjstruct.CopyStruct(req, UserGroup)
		resp.Id = UserGroup.Id
		db.Model(&UserGroup).Association("Roles").Clear() //先清空关联再插入
		if len(req.RoleId) > 0 {
			db.Where(&req.RoleId).Find(&UserGroup.Roles)
		}
		return db.Updates(UserGroup).Error
	} else { //添加
		mzjstruct.CopyStruct(req, UserGroup)
		UserGroup.Id = mzjuuid.WorkerDefault()
		resp.Id = UserGroup.Id
		if len(req.RoleId) > 0 {
			db.Where(&req.RoleId).Find(&UserGroup.Roles)
		}
		return db.Create(UserGroup).Error
	}

}
