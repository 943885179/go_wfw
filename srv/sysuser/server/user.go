package server

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"qshapi/models"
	"qshapi/proto/sysuser"
	"qshapi/utils/mzjmd5"
	"qshapi/utils/mzjstruct"
)

type IUser interface {
	ChangePassword(req *sysuser.ChangePasswordReq, resp *sysuser.EditResp) error
	UserInfoList(req *sysuser.UserInfoListReq, resp *sysuser.PageResp) error
	EditUser(user *sysuser.SysUser, resp *sysuser.EditResp) error
}

func NewUser() IUser {
	return &User{}
}

type User struct{}

func (u User) EditUser(req *sysuser.SysUser, resp *sysuser.EditResp) error {
	if req.Id == 0 {
		return errors.New("不存在该客户")
	}
	db := Conf.DbConfig.New()
	api := &models.SysUser{}
	if err := db.First(api, req.Id).Error; err != nil {
		if Conf.DbConfig.IsErrRecordNotFound(err) {
			return errors.New("修改失败,数据不存在")
		}
		return err
	}
	mzjstruct.CopyStruct(req, api)
	resp.Id = api.Id
	return db.Updates(api).Error
}

func (u User) ChangePassword(req *sysuser.ChangePasswordReq, resp *sysuser.EditResp) error {
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

func (u User) UserInfoList(req *sysuser.UserInfoListReq, resp *sysuser.PageResp) error {

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
		var r sysuser.SysUser
		mzjstruct.CopyStruct(&user, &r)
		any, _ := ptypes.MarshalAny(&r)
		resp.Data = append(resp.Data, any)
	}
	//bt, _ := json.Marshal(users)
	/*xd, _ := json.MarshalIndent(users, "", "    ")
	fmt.Println(string(xd))*/
	//resp.Data = &any.Any{Value: bt}
	//google.protobuf.ListValue=
	/*resp.Data, _ = ptypes.MarshalAny(&any.Any{
		//TypeUrl: users,
		Value: users,
	})*/
	return nil
}
