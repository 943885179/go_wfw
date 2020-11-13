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

type IQualifications interface {
	EditQualifications(req *dbmodel.Qualification, resp *dbmodel.Id) error
	DelQualifications(req *dbmodel.Id, resp *dbmodel.Id) error
	QualificationsList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	QualificationsById(id *dbmodel.Id, Qualifications *dbmodel.Qualification) error
}

func NewQualifications() IQualifications {
	return &Qualifications{}
}

type Qualifications struct{}

func (a *Qualifications) QualificationsById(id *dbmodel.Id, Qualifications *dbmodel.Qualification) error {
	return Conf.DbConfig.New().Model(&models.Qualification{}).First(Qualifications, id.Id).Error
}

func (a *Qualifications) QualificationsList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
	var t []models.Qualification
	db := Conf.DbConfig.New().Model(&models.Qualification{}).Order("id")
	db.Count(&resp.Total)
	req.Page -= 1 //分页查询页码减1
	if resp.Total == 0 {
		return nil
	}
	db.Limit(int(req.Row)).Offset(int(req.Page * req.Row)).Find(&t)
	for _, a := range t {
		var r dbmodel.Qualification
		mzjstruct.CopyStruct(&a, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}

func (*Qualifications) DelQualifications(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.Qualification{}, req.Id).Error
}
func (*Qualifications) EditQualifications(req *dbmodel.Qualification, resp *dbmodel.Id) error {
	var StartTime, EndTime time.Time
	if strings.Contains(req.StartTime, "-") {
		StartTime, _ = mzjtime.ParseInlocation(req.StartTime, mzjtime.YYYYMMDD_HORIZONTAL)
	} else if strings.Contains(req.StartTime, "/") {
		StartTime, _ = mzjtime.ParseInlocation(req.StartTime, mzjtime.YYYYMMDD_SLASH)
	} else if strings.Contains(req.StartTime, ".") {
		StartTime, _ = mzjtime.ParseInlocation(req.StartTime, mzjtime.YYYYMMDD_SPOT)
	}
	if strings.Contains(req.EndTime, "-") {
		EndTime, _ = mzjtime.ParseInlocation(req.EndTime, mzjtime.YYYYMMDD_HORIZONTAL)
	} else if strings.Contains(req.StartTime, "/") {
		EndTime, _ = mzjtime.ParseInlocation(req.EndTime, mzjtime.YYYYMMDD_SLASH)
	} else if strings.Contains(req.StartTime, ".") {
		EndTime, _ = mzjtime.ParseInlocation(req.EndTime, mzjtime.YYYYMMDD_SPOT)
	}
	db := Conf.DbConfig.New()
	//defer db.Close()
	Qualifications := &models.Qualification{}
	if len(req.Id) > 0 { //修改0
		if err := db.First(Qualifications, req.Id).Error; err != nil {
			if Conf.DbConfig.IsErrRecordNotFound(err) {
				return errors.New("修改失败,数据不存在")
			}
			return err
		}
		mzjstruct.CopyStruct(req, Qualifications)
		db.Model(&Qualifications).Association("QuaFiles").Clear()
		if req.QuaFiles != nil && len(req.QuaFiles) != 0 {
			var ids []string
			for _, a := range req.QuaFiles {
				ids = append(ids, a.Id)
			}
			if len(ids) > 0 {
				db.Where(ids).Find(&Qualifications.QuaFiles)
			}
		}
		Qualifications.StartTime = StartTime
		Qualifications.EndTime = EndTime
		if len(Qualifications.QuaNumber) == 0 {
			Qualifications.QuaNumber = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		}
		resp.Id = Qualifications.Id
		return db.Updates(Qualifications).Error
	} else { //添加
		mzjstruct.CopyStruct(req, Qualifications)
		Qualifications.StartTime = StartTime
		Qualifications.EndTime = EndTime
		Qualifications.Id = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		if len(Qualifications.QuaNumber) == 0 {
			Qualifications.QuaNumber = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		}
		resp.Id = Qualifications.Id
		return db.Create(Qualifications).Error
	}
}
