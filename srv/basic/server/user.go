package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/basic"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjstruct"
	"qshapi/utils/mzjtime"
	"qshapi/utils/mzjuuid"
	"strings"
	"time"
)

type IUser interface {
	ChangePassword(req *basic.ChangePasswordReq, resp *dbmodel.Id) error
	UserInfoList(req *basic.UserInfoListReq, resp *dbmodel.PageResp) error
	EditUser(user *dbmodel.SysUser, resp *dbmodel.Id) error
	UserById(id *dbmodel.Id, user *dbmodel.SysUser) error
}

func NewUser() IUser {
	return &User{}
}

type User struct{}

func (u User) UserById(id *dbmodel.Id, user *dbmodel.SysUser) error {
	//return Conf.DbConfig.New().Model(&models.SysUser{}).First(user, id.Id).Error
	db := Conf.DbConfig.New().Model(&models.SysUser{})
	db = db.Preload("Roles").Preload("Groups").Preload("Groups.Roles")
	db = db.Preload("Qualifications").Preload("Qualifications.QuaFiles")
	var dbu models.SysUser
	if err := db.First(&dbu, id.Id).Error; err != nil {
		return err
	}
	mzjstruct.CopyStruct(&dbu, &user)
	return nil
}

func (u User) EditUser(req *dbmodel.SysUser, resp *dbmodel.Id) error {
	if len(req.Id) == 0 {
		return errors.New("不存在该客户")
	}
	user := &models.SysUser{}
	db := Conf.DbConfig.New()
	if err := db.Model(&models.SysUser{}).Preload("Qualifications").Preload("Qualifications.QuaFiles").First(user, req.Id).Error; err != nil {
		if Conf.DbConfig.IsErrRecordNotFound(err) {
			return errors.New("修改失败,数据不存在")
		}
		return err
	}
	oldUser := &models.SysUser{}
	mzjstruct.CopyStruct(user, oldUser)

	resp.Id = user.Id

	db.Model(&user).Association("Groups").Clear()
	mzjstruct.CopyStruct(req, user)
	if req.Groups != nil && len(req.Groups) != 0 {
		var ids []string
		for _, a := range req.Groups {
			ids = append(ids, a.Id)
		}
		if len(ids) > 0 {
			db.Model(&models.SysGroup{}).Where(ids).Find(&user.Groups)
		}
	}
	db.Model(user).Association("Qualifications").Clear()
	for _, qualification := range req.Qualifications { //添加资质
		if len(strings.Trim(qualification.QuaNumber, "")) == 0 { //如果没有设置资质编号设置默认值
			qualification.QuaNumber = mzjuuid.WorkerDefaultStr(Conf.WorkerId)
		}
		var StartTime, EndTime time.Time
		if strings.Contains(qualification.StartTime, "-") {
			StartTime, _ = mzjtime.ParseInlocation(qualification.StartTime, mzjtime.YYYYMMDD_HORIZONTAL)
		} else if strings.Contains(qualification.StartTime, "/") {
			StartTime, _ = mzjtime.ParseInlocation(qualification.StartTime, mzjtime.YYYYMMDD_SLASH)
		} else if strings.Contains(qualification.StartTime, ".") {
			StartTime, _ = mzjtime.ParseInlocation(qualification.StartTime, mzjtime.YYYYMMDD_SPOT)
		}
		if strings.Contains(qualification.EndTime, "-") {
			EndTime, _ = mzjtime.ParseInlocation(qualification.EndTime, mzjtime.YYYYMMDD_HORIZONTAL)
		} else if strings.Contains(qualification.StartTime, "/") {
			EndTime, _ = mzjtime.ParseInlocation(qualification.EndTime, mzjtime.YYYYMMDD_SLASH)
		} else if strings.Contains(qualification.StartTime, ".") {
			EndTime, _ = mzjtime.ParseInlocation(qualification.EndTime, mzjtime.YYYYMMDD_SPOT)
		}
		var qf []models.SysFile
		mzjstruct.CopyStruct(&qualification.QuaFiles, &qf)
		isAdd := true
		for _, q := range oldUser.Qualifications {
			if qualification.QuaType == q.QuaType { //类型相同 修改
				//db.Model(&models.Qualification{}).Where(user.Qualifications[i].Id).Association("QuaFiles").Clear() //清楚资质和资源关系
				qua := &models.Qualification{}
				_ = db.First(qua, q.Id).Error
				db.Model(qua).Association("QuaFiles").Clear()
				isAdd = false
				qua.StartTime = StartTime
				qua.EndTime = EndTime
				qua.UserId = user.Id
				qua.QuaExplain = qualification.QuaExplain
				qua.QuaFiles = qf
				qua.QuaNumber = qualification.QuaNumber
				db.Updates(qua)
				break
			}
		}
		if isAdd {
			qua := &models.Qualification{
				Id:         mzjuuid.WorkerDefaultStr(Conf.WorkerId),
				StartTime:  StartTime,
				EndTime:    EndTime,
				UserId:     user.Id,
				QuaType:    qualification.QuaType,
				QuaExplain: qualification.QuaExplain,
				QuaFiles:   qf,
				QuaNumber:  qualification.QuaNumber,
			}
			db.Create(qua)
		}
		// user.Qualifications[i].QuaType = req.Qualifications[i].QuaType
	}
	return db.Updates(user).Error
}

func (u User) ChangePassword(req *basic.ChangePasswordReq, resp *dbmodel.Id) error {
	if req.UserPassword != req.UserPasswordAgain {
		return errors.New("密码不一致")
	}
	db := Conf.DbConfig.New()
	var user models.SysUser
	if err := db.First(&user, req.Id).Error; err != nil {
		return err
	}
	user.UserPassword = mzjmd5.MD5(req.UserPassword)
	resp.Id = user.Id
	return db.Updates(user).Error
}

func (u User) UserInfoList(req *basic.UserInfoListReq, resp *dbmodel.PageResp) error {

	db := Conf.DbConfig.New().Model(&models.SysUser{})
	if len(req.UserName) != 0 {
		db = db.Where("user_name like ?", "%"+req.UserName+"%")
	}
	if len(req.UserPhone) != 0 {
		db = db.Where("user_phone like ?", "%"+req.UserPhone+"%")
	}
	if len(req.UserName) != 0 {
		db = db.Where("user_email like ?", "%"+req.UserEmail+"%")
	}
	db.Count(&resp.Total)
	if resp.Total == 0 {
		return nil
	}
	req.PageReq.Page -= 1                                              //分页查询页码减1
	db = db.Preload("Roles").Preload("Groups").Preload("Groups.Roles") //注意大小写
	db = db.Preload("Roles.Srvs").Preload("Roles.Apis").Preload("Roles.Menus").Preload("Roles.Menus.Children").Preload("Roles.Menus.Children.Children")
	db = db.Preload("Groups.Roles.Srvs").Preload("Groups.Roles.Apis").Preload("Groups.Roles.Menus").Preload("Groups.Roles.Menus.Children").Preload("Groups.Roles.Menus.Children.Children")
	db = db.Limit(int(req.PageReq.Row)).Offset(int(req.PageReq.Page * req.PageReq.Row))
	db = db.Preload("Province").Preload("City").Preload("Area") //地址
	db = db.Preload("Icon")                                     //头像
	var us []models.SysUser
	if err := db.Find(&us).Error; err != nil {
		return err
	}
	for _, user := range us {
		var r dbmodel.SysUser
		mzjstruct.CopyStruct(&user, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	return nil
}
