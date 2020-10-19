package server

import (
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IUserGroup interface {
	EditUserGroup(req *sysuser.SysGroup, resp *sysuser.EditResp) error
	DelUserGroup(req *sysuser.DelReq, resp *sysuser.EditResp) error
	UserGroupList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewUserGroup() IUserGroup {
	return &UserGroup{}
}

type UserGroup struct{}

func (g *UserGroup) UserGroupList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var ts []models.SysGroup
	db := Conf.DbConfig.New().Model(&models.SysGroup{})
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&ts)
	for _, role := range ts {
		var r sysuser.SysGroup
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*UserGroup) DelUserGroup(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	//defer db.Close()
	return db.Delete(models.SysGroup{}, req.Id).Error
}

func (*UserGroup) EditUserGroup(req *sysuser.SysGroup, resp *sysuser.EditResp) error {
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
		if len(req.Roles) > 0 {
			db.Find(&UserGroup.Roles)
		}
		return db.Updates(UserGroup).Error
	} else { //添加
		mzjstruct.CopyStruct(req, UserGroup)
		UserGroup.Id = mzjuuid.WorkerDefault()
		resp.Id = UserGroup.Id
		if len(req.Roles) > 0 {
			db.Find(&UserGroup.Roles)
		}
		return db.Create(UserGroup).Error
	}

}
