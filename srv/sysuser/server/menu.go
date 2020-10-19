package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IMenu interface {
	EditMenu(req *sysuser.SysMenu, resp *sysuser.EditResp) error
	DelMenu(req *sysuser.DelReq, resp *sysuser.EditResp) error
	MenuList(req *sysuser.PageReq, resp *sysuser.PageResp) error
}

func NewMenu() IMenu {
	return &Menu{}
}

type Menu struct{}

func (m *Menu) MenuList(req *sysuser.PageReq, resp *sysuser.PageResp) error {
	var t []models.SysMenu
	db := Conf.DbConfig.New().Model(&models.SysMenu{}).Where("p_id=0")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r sysuser.SysMenu
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Menu) DelMenu(req *sysuser.DelReq, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysMenu{}, req.Id).Error
}

func (*Menu) EditMenu(req *sysuser.SysMenu, resp *sysuser.EditResp) error {
	db := Conf.DbConfig.New()
	menu := &models.SysMenu{}
	if req.Id > 0 { //修改0
		if err := db.First(menu, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, menu)
		resp.Id = menu.Id
		return db.Updates(menu).Error
	} else { //添加
		mzjstruct.CopyStruct(req, menu)
		menu.Id = mzjuuid.WorkerDefault()
		resp.Id = menu.Id
		return db.Create(menu).Error
	}
}
