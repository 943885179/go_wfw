package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjuuid"

	"github.com/golang/protobuf/ptypes"
)

type IArea interface {
	EditArea(req *dbmodel.SysArea, resp *dbmodel.Id) error
	DelArea(req *dbmodel.Id, resp *dbmodel.Id) error
	AreaList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	AreaById(id *dbmodel.Id, Area *dbmodel.SysArea) error
	AreaTree(resp *dbmodel.TreeResp) error
}

func NewArea() IArea {
	return &Area{}
}

type Area struct{}

func (r *Area) AreaTree(resp *dbmodel.TreeResp) error {
	var data []models.SysArea
	db := Conf.DbConfig.New().Model(&models.SysArea{}).Where("p_id=''")
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

func (g *Area) AreaById(id *dbmodel.Id, Area *dbmodel.SysArea) error {
	return Conf.DbConfig.New().Model(&models.SysArea{}).First(Area, id.Id).Error
}
func (t *Area) AreaList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var ts []models.SysArea
	db := Conf.DbConfig.New().Model(&models.SysArea{}) //.Where("p_id=0")
	if len(req.Text) > 0 {
		db = db.Where("text like ?", "%"+req.Text+"%")
	}
	db.Count(&resp.Total)
	req.Page = req.Page - 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db = db.Preload("Children")
	/*db = db.Preload("Children.Children")
	db = db.Preload("Children.Children.Children")
	db = db.Preload("Children.Children.Children.Children")*/
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&ts)
	for _, role := range ts {
		var r dbmodel.SysArea
		mzjstruct.CopyStruct(&role, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Area) DelArea(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.SysArea{}, req.Id).Error
}

func (*Area) EditArea(req *dbmodel.SysArea, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	Area := &models.SysArea{}
	if len(req.Id) > 0 { //修改0
		if err := db.First(Area, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Area)
		resp.Id = Area.Id
		Area.Title = Area.Text
		return db.Updates(Area).Error
	} else {
		mzjstruct.CopyStruct(req, Area)
		Area.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		Area.Key = Area.Id
		resp.Id = Area.Id
		Area.Title = Area.Text
		return db.Create(Area).Error
	}
}
