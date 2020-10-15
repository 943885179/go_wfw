package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IMenu interface {
	EditMenu(req *sysuser.MenuReq) error
	DelMenu(req *sysuser.DelReq) error
}

func NewMenu() IMenu {
	return &Menu{}
}

type Menu struct{}

func (*Menu) DelMenu(req *sysuser.DelReq) error {
	db := Conf.DbConfig.New()
	return db.Delete(models.SysMenu{}, req.Id).Error
}

func (*Menu) EditMenu(req *sysuser.MenuReq) error {
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
		return db.Updates(menu).Error
	} else { //添加
		menu.ID = mzjuuid.WorkerDefault()
		mzjstruct.CopyStruct(req, menu)
		return db.Create(menu).Error
	}
}
