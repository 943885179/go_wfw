package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IRole interface {
	EditRole(req *dbmodel.SysRole, resp *dbmodel.Id) error
	DelRole(req *dbmodel.Id, resp *dbmodel.Id) error
	RoleList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	RoleById(id *dbmodel.Id, role *dbmodel.SysRole) error
	RoleTree(empty *empty.Empty, resp *dbmodel.TreeResp) error
}

func NewRole() IRole {
	return &Role{}
}

type Role struct{}

func (r *Role) RoleTree(empty *empty.Empty, resp *dbmodel.TreeResp) error {
	var data []models.SysRole
	db := Conf.DbConfig.New().Model(&models.SysMenu{}).Where("p_id=0")
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	if err := db.Find(&data).Error; err != nil {
		return err
	}

	for _, m := range data {
		var r dbmodel.Tree
		mzjstruct.CopyStruct(&m, &r)
		resp.Data = append(resp.Data, &r)
	}
	return nil
}

func (r *Role) RoleById(id *dbmodel.Id, role *dbmodel.SysRole) error {
	db := Conf.DbConfig.New().Model(&models.SysRole{})
	db = db.Preload("Srvs").Preload("Apis").Preload("Menus")
	db = db.Preload("Children").Preload("Children.Srvs").Preload("Children.Apis").Preload("Children.Menus")
	db = db.Preload("Children.Children").Preload("Children.Children.Srvs").Preload("Children.Children.Apis").Preload("Children.Children.Menus")
	db = db.Preload("Children.Children.Children").Preload("Children.Children.Children.Srvs").Preload("Children.Children.Children.Apis").Preload("Children.Children.Children.Menus")
	db = db.Preload("Children.Children.Children.Children").Preload("Children.Children.Children.Children.Srvs").Preload("Children.Children.Children.Children.Apis").Preload("Children.Children.Children.Children.Menus")
	var dbrole models.SysRole
	if err := db.First(&dbrole, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbrole, &role)
	return nil
}

func (r *Role) RoleList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
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
		var r dbmodel.SysRole
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Role) DelRole(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysRole{}, req.Id).Error
}

func (*Role) EditRole(req *dbmodel.SysRole, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	Role := &models.SysRole{}
	if len(req.Id) > 0 { //修改0
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
			var ids []string
			for _, a := range req.Apis {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Apis)
			}
		}
		if req.Srvs != nil && len(req.Srvs) != 0 {
			var ids []string
			for _, a := range req.Srvs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Srvs)
			}
		}
		if req.Menus != nil && len(req.Menus) != 0 {
			var ids []string
			for _, a := range req.Menus {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Menus)
			}
		}
		Role.Title = Role.RoleName
		return db.Updates(Role).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Role)
		Role.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		Role.Key = Role.Id
		Role.Title = Role.RoleName
		resp.Id = Role.Id
		if req.Apis != nil && len(req.Apis) != 0 {
			var ids []string
			for _, a := range req.Apis {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Apis)
			}
		}
		if req.Srvs != nil && len(req.Srvs) != 0 {
			var ids []string
			for _, a := range req.Srvs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Role.Srvs)
			}
		}
		if req.Menus != nil && len(req.Menus) != 0 {
			var ids []string
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
