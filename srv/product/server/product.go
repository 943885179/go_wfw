package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjpinyin"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IProduct interface {
	EditProduct(req *dbmodel.Product, resp *dbmodel.Id) error
	DelProduct(req *dbmodel.Id, resp *dbmodel.Id) error
	ProductList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	ProductById(id *dbmodel.Id, product *dbmodel.Product) error
}

func NewProduct() IProduct {
	return &Product{}
}

type Product struct{}

func (*Product) ProductById(id *dbmodel.Id, product *dbmodel.Product) error {
	db := Conf.DbConfig.New().Model(&models.Product{})
	db = db.Preload("Imgs").Preload("ProductSkus").Preload("ProductSkus.Imgs")
	db = db.Preload("Qualifications").Preload("Qualifications.QuaFiles")
	var dbs models.Product
	if err := db.First(&dbs, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbs, &product)
	return nil
}
func (*Product) ProductList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.Product
	db := Conf.DbConfig.New().Model(&models.Product{}).Preload("Imgs")
	if len(req.Code) > 0 {
		db = db.Where("goods_code like ?", "%"+req.Code+"%")
	}
	if len(req.Name) > 0 {
		db = db.Where("goods_name like ? or goods_byname like ? or opcode like ?", "%"+req.Name+"%", "%"+req.Name+"%", "%"+req.Name+"%")
	}
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r dbmodel.Product
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Product) DelProduct(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.Product{}, req.Id).Error
}
func (*Product) EditProduct(req *dbmodel.Product, resp *dbmodel.Id) (err error) {
	db := Conf.DbConfig.New()
	//defer db.Close()
	Product := &models.Product{}
	if len(req.GoodsCode) == 0 {
		return errors.New("商品编码不能为空")
	}
	if len(req.GoodsName) == 0 {
		return errors.New("商品名称不能为空")
	}
	if len(req.Opcode) == 0 {
		req.Opcode = mzjpinyin.DefalutStrToPinyin(req.GoodsName)
	}
	if len(req.Id) > 0 { //修改0
		//if db.FirstOrInit(Product, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(Product, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Product)
		db.Model(&Product).Association("Imgs").Clear()
		if req.Imgs != nil && len(req.Imgs) != 0 {
			var ids []string
			for _, a := range req.Imgs {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Product.Imgs)
			}
		}
		db.Model(&Product).Association("ProductSkus").Clear()
		resp.Id = Product.Id
		err = db.Updates(Product).Error
		var q = NewProductSku()
		if err != nil {
			return err
		}
		for _, sku := range req.ProductSkus { //添加sku
			sku.ProductId = Product.Id
			err = q.EditProductSku(sku, &dbmodel.Id{})
			if err != nil {
				return err
			}
		}
		/*db.Model(&Product).Association("Qualifications").Clear()
		for _, qualification := range req.Qualifications {
			qualification.ForeignId = Product.Id
			var quaReq = dbmodel.Id{}
			err := EditQualifications(qualification, &quaReq)
			if err != nil {
				return err
			}
		}*/
		return nil
	} else { //添加
		mzjstruct.CopyStruct(req, Product)
		Product.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		db.Model(&Product).Association("ProductSkus").Clear()
		resp.Id = Product.Id
		err = db.Create(Product).Error
		if err != nil {
			return err
		}
		var q = NewProductSku()
		if err != nil {
			return err
		}
		for _, sku := range req.ProductSkus { //添加sku
			sku.ProductId = Product.Id
			err = q.EditProductSku(sku, &dbmodel.Id{})
			if err != nil {
				return err
			}
		}
		/*for _, qualification := range req.Qualifications {
			qualification.ForeignId = Product.Id
			var quaReq = dbmodel.Id{}
			err := EditQualifications(qualification, &quaReq)
			if err != nil {
				return err
			}
		}*/
		return nil
	}
}
