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
}

func NewShop() IShop {
	return &Shop{}
}

type Shop struct{}

func (a *Shop) ShopList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.SysShop
	db := Conf.DbConfig.New().Model(&models.SysShop{})
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
	if req.Id > 0 { //修改0
		//if db.FirstOrInit(Shop, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(Shop, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Shop)
		resp.Id = Shop.Id
		return db.Updates(Shop).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Shop)
		Shop.Id = mzjuuid.WorkerDefault()
		resp.Id = Shop.Id
		return db.Create(Shop).Error
	}
}
