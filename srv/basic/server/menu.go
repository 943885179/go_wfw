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

type IMenu interface {
	EditMenu(req *dbmodel.SysMenu, resp *dbmodel.Id) error
	DelMenu(req *dbmodel.Id, resp *dbmodel.Id) error
	MenuList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	//MenuListByUser(userMenu ...[]models.SysMenu) []models.SysMenu
	MenuListByUser(user *dbmodel.SysUser, menu *dbmodel.OnlyMenu) error
	MenuById(id *dbmodel.Id, menu *dbmodel.SysMenu) error
	MenuTree(empty *empty.Empty, resp *dbmodel.TreeResp) error
}

func NewMenu() IMenu {
	return &Menu{}
}

type Menu struct{}

func (m *Menu) MenuTree(empty *empty.Empty, resp *dbmodel.TreeResp) error {
	var data []models.SysMenu
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

func (m *Menu) MenuById(id *dbmodel.Id, menu *dbmodel.SysMenu) error {
	//return Conf.DbConfig.New().Model(&models.SysMenu{}).First(menu, id.Id).Error
	//还是使用下面的方法，写参数的时候不要用数字了，感觉有点问题
	var dbmenu models.SysMenu
	if err := Conf.DbConfig.New().Model(&models.SysMenu{}).First(&dbmenu, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbmenu, menu)
	return nil
}
func (m *Menu) MenuListByUser(req *dbmodel.SysUser, resp *dbmodel.OnlyMenu) error {
	var ms []models.SysMenu
	db := Conf.DbConfig.New().Model(&models.SysMenu{}).Where("p_id=0")
	db = db.Preload("Children")
	db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")
	db.Find(&ms)

	var hasIds []string
	if req.UserType == dbmodel.UserType_ADMIN { // 超级管理员有所有的菜单权限，不受约束
		var allmenus []models.SysMenu
		Conf.DbConfig.New().Find(&allmenus)
		for _, menu := range allmenus {
			hasIds = append(hasIds, menu.Id)
		}
	} else {
		for _, group := range req.Groups {
			for _, role := range group.Roles {
				for _, menu := range role.Menus {
					hasIds = append(hasIds, menu.Id)
				}
			}
		}
		for _, role := range req.Roles {
			for _, menu := range role.Menus {
				hasIds = append(hasIds, menu.Id)
			}
		}
	}
	tree := menuTree(hasIds, ms)
	mzjstruct.CopyStruct(&tree, &resp.Menus)
	return nil
}
func menuTree(hasIds []string, menus []models.SysMenu) []models.SysMenu {
	var result []models.SysMenu
	for _, menu := range menus {
		if len(menu.Children) > 0 {
			menu.Children = menuTree(hasIds, menu.Children)
			if len(menu.Children) > 0 {
				result = append(result, menu)
			}
		} else {
			for _, id := range hasIds {
				if menu.Id == id {
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
	db := Conf.DbConfig.New().Model(&models.SysMenu{}) //.Where("p_id=0")
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
	for _, m := range t {
		var r dbmodel.SysMenu
		mzjstruct.CopyStruct(&m, &r)
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
	if len(req.Id) > 0 { //修改0
		if err := db.First(menu, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, menu)
		menu.Title = menu.Text
		resp.Id = menu.Id
		return db.Updates(menu).Error
	} else { //添加
		mzjstruct.CopyStruct(req, menu)
		menu.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		menu.Key = menu.Id
		menu.Title = menu.Text
		resp.Id = menu.Id
		return db.Create(menu).Error
	}
}
