package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjtime"
	"qshapi/utils/mzjuuid"
	"strings"
	"time"
)

type IProductSku interface {
	EditProductSku(req *dbmodel.ProductSku, resp *dbmodel.Id) error
	DelProductSku(req *dbmodel.Id, resp *dbmodel.Id) error
	ProductSkuList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	ProductSkuById(id *dbmodel.Id, ProductSku *dbmodel.ProductSku) error
}

func NewProductSku() IProductSku {
	return &ProductSku{}
}

type ProductSku struct{}

func (*ProductSku) ProductSkuById(id *dbmodel.Id, ProductSku *dbmodel.ProductSku) error {
	db := Conf.DbConfig.New().Model(&models.ProductSku{}).Preload("Imgs")
	var dbs models.ProductSku
	if err := db.First(&dbs, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbs, &ProductSku)
	return nil
}
func (*ProductSku) ProductSkuList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.ProductSku
	db := Conf.DbConfig.New().Model(&models.ProductSku{}).Preload("Imgs")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r dbmodel.ProductSku
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*ProductSku) DelProductSku(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.ProductSku{}, req.Id).Error
}
func (*ProductSku) EditProductSku(req *dbmodel.ProductSku, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	ProductSku := &models.ProductSku{}
	if len(req.SkuName) == 0 {
		req.SkuName = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
	}
	if len(req.BatchNumber) == 0 {
		req.BatchNumber = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
	}
	var EffectiveDate, ProdutionDate time.Time
	if strings.Contains(req.EffectiveDate, "-") {
		EffectiveDate, _ = mzjtime.ParseInlocation(req.EffectiveDate, mzjtime.YYYYMMDD_HORIZONTAL)
	} else if strings.Contains(req.EffectiveDate, "/") {
		EffectiveDate, _ = mzjtime.ParseInlocation(req.EffectiveDate, mzjtime.YYYYMMDD_SLASH)
	} else if strings.Contains(req.EffectiveDate, ".") {
		EffectiveDate, _ = mzjtime.ParseInlocation(req.EffectiveDate, mzjtime.YYYYMMDD_SPOT)
	}
	if strings.Contains(req.ProdutionDate, "-") {
		ProdutionDate, _ = mzjtime.ParseInlocation(req.ProdutionDate, mzjtime.YYYYMMDD_HORIZONTAL)
	} else if strings.Contains(req.ProdutionDate, "/") {
		ProdutionDate, _ = mzjtime.ParseInlocation(req.ProdutionDate, mzjtime.YYYYMMDD_SLASH)
	} else if strings.Contains(req.ProdutionDate, ".") {
		ProdutionDate, _ = mzjtime.ParseInlocation(req.ProdutionDate, mzjtime.YYYYMMDD_SPOT)
	}
	if len(req.Id) > 0 { //修改0
		//if db.FirstOrInit(ProductSku, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(ProductSku, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, ProductSku)
		db.Model(&ProductSku).Association("Imgs").Clear()
		if req.Imgs != nil && len(req.Imgs) != 0 {
			var ids []string
			for _, a := range req.Imgs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&ProductSku.Imgs)
			}
		}
		ProductSku.EffectiveDate = EffectiveDate
		ProductSku.ProdutionDate = ProdutionDate
		resp.Id = ProductSku.Id
		return db.Updates(ProductSku).Error
	} else { //添加
		mzjstruct.CopyStruct(req, ProductSku)
		ProductSku.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		ProductSku.ProductId = req.ProductId
		ProductSku.EffectiveDate = EffectiveDate
		ProductSku.ProdutionDate = ProdutionDate
		resp.Id = ProductSku.Id
		db.Model(&ProductSku).Association("Imgs").Clear()
		if req.Imgs != nil && len(req.Imgs) != 0 {
			var ids []string
			for _, a := range req.Imgs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&ProductSku.Imgs)
			}
		}
		return db.Create(ProductSku).Error
	}
}
