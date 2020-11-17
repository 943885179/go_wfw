package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IShop interface {
	EditShop(req *dbmodel.SysShop, resp *dbmodel.Id) error
	DelShop(req *dbmodel.Id, resp *dbmodel.Id) error
	ShopList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	ShopById(id *dbmodel.Id, shop *dbmodel.SysShop) error
}

func NewShop() IShop {
	return &Shop{}
}

type Shop struct{}

func (a *Shop) ShopById(id *dbmodel.Id, shop *dbmodel.SysShop) error {
	//return Conf.DbConfig.New().Model(&models.SysShop{}).First(shop, id.Id).Error

	db := Conf.DbConfig.New().Model(&models.SysShop{}).Preload("Logo").Preload("Imgs")
	db = db.Preload("Qualifications").Preload("Qualifications.QuaFiles").Preload("Qualifications.QuaType")
	var dbs models.SysShop
	if err := db.First(&dbs, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbs, &shop)
	return nil
}
func (a *Shop) ShopList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.SysShop
	db := Conf.DbConfig.New().Model(&models.SysShop{}).Preload("Logo").Preload("Imgs")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r dbmodel.SysShop
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Shop) DelShop(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysShop{}, req.Id).Error
}
func (*Shop) EditShop(req *dbmodel.SysShop, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	Shop := &models.SysShop{}
	if len(req.Id) > 0 { //修改0
		//if db.FirstOrInit(Shop, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(Shop, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}

		oldShop := &models.SysShop{}
		mzjstruct.CopyStruct(Shop, oldShop)
		resp.Id = Shop.Id
		mzjstruct.CopyStruct(req, Shop)
		if req.Logo != nil {
			Shop.LogoId = req.Logo.Id
		}
		db.Model(&Shop).Association("Imgs").Clear()
		if req.Imgs != nil && len(req.Imgs) != 0 {
			var ids []string
			for _, a := range req.Imgs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Shop.Imgs)
			}
		}
		var q = NewQualifications()
		for _, qualification := range req.Qualifications { //添加资质
			qualification.ShopId = Shop.Id
			q.EditQualifications(qualification, &dbmodel.Id{})
		}
		return db.Updates(Shop).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Shop)
		Shop.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		resp.Id = Shop.Id
		if req.Logo != nil {
			Shop.LogoId = req.Logo.Id
		}
		var q = NewQualifications()
		for _, qualification := range req.Qualifications { //添加资质
			qualification.ShopId = Shop.Id
			q.EditQualifications(qualification, &dbmodel.Id{})
		}
		return db.Create(Shop).Error
	}
}
