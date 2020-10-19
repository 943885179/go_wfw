package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IRole interface {
	EditRole(req *sysuser.SysRole, resp *sysuser.EditResp) error
	DelRole(req *sysuser.DelReq, resp *sysuser.EditResp) error
	RoleList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewRole() IRole {
	return &Role{}
}

type Role struct{}

func (r *Role) RoleList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var roles []models.SysRole
	db := Conf.DbConfig.New().Model(&models.SysRole{}).Where("p_id=0")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db = db.Preload("Srvs").Preload("Apis").Preload("Menus")
	db = db.Preload("Children").Preload("Children.Srvs").Preload("Children.Apis").Preload("Children.Menus")
	db = db.Preload("Children.Children").Preload("Children.Children.Srvs").Preload("Children.Children.Apis").Preload("Children.Children.Menus")
	db = db.Preload("Children.Children.Children").Preload("Children.Children.Children.Srvs").Preload("Children.Children.Children.Apis").Preload("Children.Children.Children.Menus")
	db = db.Preload("Children.Children.Children.Children").Preload("Children.Children.Children.Children.Srvs").Preload("Children.Children.Children.Children.Apis").Preload("Children.Children.Children.Children.Menus")
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&roles)
	/*var item = &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name": &structpb.Value{
				Kind: &structpb.Value_StringValue{
					StringValue: "Anuj",
				},
			},
			"age": &structpb.Value{
				Kind: &structpb.Value_StringValue{
					StringValue: "Anuj",
				},
			},
		},
	}
	resp.Data = append(resp.Data, item)
	*/
	for _, role := range roles {
		var r sysuser.SysRole
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Role) DelRole(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysRole{}, req.Id).Error
}

func (*Role) EditRole(req *sysuser.SysRole, resp *sysuser.EditResp) error {
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
		db.Model(&Role).Association("Apis").Clear() //先清空关联再插入
		db.Model(&Role).Association("Srvs").Clear()
		db.Model(&Role).Association("Menus").Clear()
		if req.Apis != nil && len(req.Apis) != 0 {
			var ids []int64
			for _, a := range req.Apis {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Apis)
			}
		}
		if req.Srvs != nil && len(req.Srvs) != 0 {
			var ids []int64
			for _, a := range req.Srvs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Srvs)
			}
		}
		if req.Menus != nil && len(req.Menus) != 0 {
			var ids []int64
			for _, a := range req.Menus {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Menus)
			}
		}
		return db.Updates(Role).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Role)
		Role.Id = mzjuuid.WorkerDefault()
		resp.Id = Role.Id
		if req.Apis != nil && len(req.Apis) != 0 {
			var ids []int64
			for _, a := range req.Apis {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Apis)
			}
		}
		if req.Srvs != nil && len(req.Srvs) != 0 {
			var ids []int64
			for _, a := range req.Srvs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Srvs)
			}
		}
		if req.Menus != nil && len(req.Menus) != 0 {
			var ids []int64
			for _, a := range req.Menus {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Menus)
			}
		}
		return db.Create(Role).Error
	}
}
