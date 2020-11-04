package server

import (
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IUserGroup interface {
	EditUserGroup(req *dbmodel.SysGroup, resp *dbmodel.Id) error
	DelUserGroup(req *dbmodel.Id, resp *dbmodel.Id) error
	UserGroupList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	UserGroupById(id *dbmodel.Id, group *dbmodel.SysGroup) error
}

func NewUserGroup() IUserGroup {
	return &UserGroup{}
}

type UserGroup struct{}

func (g *UserGroup) UserGroupById(id *dbmodel.Id, group *dbmodel.SysGroup) error {
	//return Conf.DbConfig.New().Model(&models.SysGroup{}).Preload("Roles").First(group, id.Id).Error
	db := Conf.DbConfig.New().Model(&models.SysGroup{}).Preload("Roles")
	var dbgroup models.SysGroup
	if err := db.First(&dbgroup, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbgroup, &group)
	return nil
}

func (g *UserGroup) UserGroupList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var ts []models.SysGroup
	db := Conf.DbConfig.New().Model(&models.SysGroup{}).Preload("Roles")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&ts)
	for _, role := range ts {
		var r dbmodel.SysGroup
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*UserGroup) DelUserGroup(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	//defer db.Close()
	return db.Delete(models.SysGroup{}, req.Id).Error
}

func (*UserGroup) EditUserGroup(req *dbmodel.SysGroup, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	UserGroup := &models.SysGroup{}
	if len(req.Id) > 0 { //修改0
		if err := db.First(UserGroup, req.Id).Error; err != nil {
			return err
		}
		mzjstruct.CopyStruct(req, UserGroup)
		resp.Id = UserGroup.Id
		db.Model(&UserGroup).Association("Roles").Clear() //先清空关联再插入
		if req.Roles != nil && len(req.Roles) != 0 {
			var ids []string
			for _, a := range req.Roles {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&UserGroup.Roles)
			}
		}
		return db.Updates(UserGroup).Error
	} else { //添加
		mzjstruct.CopyStruct(req, UserGroup)
		UserGroup.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		resp.Id = UserGroup.Id
		if len(req.Roles) > 0 {
			db.Find(&UserGroup.Roles)
		}
		return db.Create(UserGroup).Error
	}

}
