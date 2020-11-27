package server

import (
	"errors"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjtime"
	"qshapi/utils/mzjuuid"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type IQualification interface {
	EditQualification(req *dbmodel.Qualification, resp *dbmodel.Id) error
	DelQualification(req *dbmodel.Id, resp *dbmodel.Id) error
	QualificationsList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error
	QualificationsById(id *dbmodel.Id, Qualifications *dbmodel.Qualification) error
	QualificationByForeignId(id *dbmodel.Id, qualifications *basic.Qualifications) error
	EditQualifications(qualifications *basic.Qualifications) error
}

func NewQualification() IQualification {
	return &Qualification{}
}

type Qualification struct{}

func (a *Qualification) EditQualifications(qualifications *basic.Qualifications) error {
	/*var quas = []models.Qualification{}
	mzjstruct.CopyStruct(&qualifications.Data, &quas)

	return Conf.DbConfig.New().Model(&models.Qualification{}).Create(quas).Error*/
	for _, q := range qualifications.Data {
		a.EditQualification(q, &dbmodel.Id{})
	}
	return nil
}
func (a *Qualification) QualificationByForeignId(id *dbmodel.Id, qualifications *basic.Qualifications) error {
	var qua = []models.Qualification{}
	db := Conf.DbConfig.New().Model(&models.Qualification{})
	db = db.Preload("QuaFiles").Where("foreign_id=?", id.Id)
	err := db.Find(&qua).Error
	if err != nil {
		return err
	}
	mzjstruct.CopyStruct(&qua, &qualifications.Data)
	return nil
}
func (a *Qualification) QualificationsById(id *dbmodel.Id, Qualifications *dbmodel.Qualification) error {
	return Conf.DbConfig.New().Model(&models.Qualification{}).First(Qualifications, id.Id).Error
}

func (a *Qualification) QualificationsList(req *dbmodel.PageReq, resp *dbmodel.PageResp) error {
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

func (*Qualification) DelQualification(req *dbmodel.Id, resp *dbmodel.Id) error {
	db := Conf.DbConfig.New()
	resp.Id = req.Id
	return db.Delete(models.Qualification{}, req.Id).Error
}
func (*Qualification) EditQualification(req *dbmodel.Qualification, resp *dbmodel.Id) error {
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
	} else if strings.Contains(req.EndTime, "/") {
		EndTime, _ = mzjtime.ParseInlocation(req.EndTime, mzjtime.YYYYMMDD_SLASH)
	} else if strings.Contains(req.EndTime, ".") {
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
