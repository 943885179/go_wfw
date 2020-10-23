package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IMenu interface {
	EditMenu(req *dbmodel.SysMenu, resp *dbmodel.Id) error
	DelMenu(req *dbmodel.Id, resp *dbmodel.Id) error
	MenuList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	MenuListByUser(userMenu ...[]models.SysMenu) []models.SysMenu
}

func NewMenu() IMenu {
	return &Menu{}
}

type Menu struct{}

func (m *Menu) MenuListByUser(userMenu ...[]models.SysMenu) []models.SysMenu {
	onlyMenu := []models.SysMenu{}
	for _, menus := range userMenu {
		for _, menu := range menus {
			isAdd := true
			for _, o := range onlyMenu {
				if menu.Id == o.Id {
					isAdd = false
					break
				}
			}
			if isAdd {
				onlyMenu = append(onlyMenu, menu)
			}
		}
	}
	var ms []models.SysMenu
	db := Conf.DbConfig.New().Model(&models.SysMenu{}).Where("p_id=0")
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	db.Find(&ms)
	return menuTree(onlyMenu, ms)
}
func menuTree(userMenu, menus []models.SysMenu) []models.SysMenu {
	var result []models.SysMenu
	for _, menu := range menus {
		if len(menu.Children) > 0 {
			menu.Children = menuTree(userMenu, menu.Children)
			if len(menu.Children) > 0 {
				result = append(result, menu)
			}
		} else {
			for _, sysMenu := range userMenu {
				if sysMenu.Id == menu.Id {
					result = append(result, menu)
					break
				}
			}
		}
	}
	return result
}

func (m *Menu) MenuList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
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
		var r dbmodel.SysMenu
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Menu) DelMenu(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysMenu{}, req.Id).Error
}

func (*Menu) EditMenu(req *dbmodel.SysMenu, resp *dbmodel.Id) error {
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
