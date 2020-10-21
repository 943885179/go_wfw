package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"
)

type IOrder interface {
	EditOrder(req *dbmodel.Orders, resp *dbmodel.Id) error
	DelOrder(req *dbmodel.Id, resp *dbmodel.Id) error
	OrderList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
}

func NewOrder() IOrder {
	return &Order{}
}

type Order struct{}

func (a *Order) OrderList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.Orders
	db := Conf.DbConfig.New().Model(&models.Orders{})
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, role := range t {
		var r dbmodel.Orders
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Order) DelOrder(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.Orders{}, req.Id).Error
}
func (*Order) EditOrder(req *dbmodel.Orders, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	//defer db.Close()
	Order := &models.Orders{}
	if req.Id > 0 { //修改0
		//if db.FirstOrInit(Order, req.Id).RecordNotFound() { v2版本移除了
		if err := db.First(Order, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Order)
		resp.Id = Order.Id
		return db.Updates(Order).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Order)
		Order.Id = mzjuuid.WorkerDefault()
		resp.Id = Order.Id
		return db.Create(Order).Error
	}
}
