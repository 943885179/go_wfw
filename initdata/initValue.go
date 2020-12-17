package main

import (
	"qshapi/models"
	"qshapi/proto/dbmodel"
)

func initValue() {
	initValue_sex()
	initValue_quatype()
	initValue_filetype()
	initValue_dataReadRole()
	initValue_dataInsertRole()
	initValue_dataUpdateRole()
	initValue_dataDelRole()
}

//性别初始化
func initValue_sex() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100001", Code: "SEX_MAN", Value: "男", ValueExplain: "性别", ValueType: dbmodel.ValueType_SEX},
		{Id: "100002", Code: "SEX_WOMEN", Value: "女", ValueExplain: "性别", ValueType: dbmodel.ValueType_SEX},
	}
	db.Create(&v)
}

//资质初始化
func initValue_quatype() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100003", Code: "QUA_IDCARD", Value: "身份证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100004", Code: "QUA_BIZLICENSE", Value: "营业执照/营业执照证书(三证合一)", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100005", Code: "QUA_BANCKCARD", Value: "银行卡", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100006", Code: "QUA_DMC", Value: "药品生产许可证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100007", Code: "QUA_DSC", Value: "药品经营许可证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100008", Code: "QUA_GSP", Value: "药品经营质量管理规范", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100009", Code: "QUA_GMP", Value: "药品生产质量管理规范", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100011", Code: "QUA_PPC", Value: "医疗机构制剂许可证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100012", Code: "QUA_IDL", Value: "进口药品注册证书", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100013", Code: "QUA_CT", Value: "临床试验", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100014", Code: "QUA_NDC", Value: "新药证书", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100015", Code: "QUA_ATTORNEY", Value: "授权委托书", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100016", Code: "QUA_GTQ", Value: "一般纳税人资质证明", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100017", Code: "QUA_SQC", Value: "SC/QC食品生产许可证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100018", Code: "QUA_QS", Value: "食品质量安全生产许可证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100019", Code: "QUA_HZP", Value: "化妆品生产许可", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100020", Code: "QUA_GMPC", Value: "化妆品GMPC", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100021", Code: "QUA_3C", Value: "3C认证", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100022", Code: "QUA_SHOPURL", Value: "网店完整的域名截图照片（含店铺信息）", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100023", Code: "QUA_BIZLICENSEADREES", Value: "商户营业场所照片", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100024", Code: "QUA_VIOD", Value: "开户视频", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
		{Id: "100025", Code: "QUA_OTHER", Value: "其他资质", ValueExplain: "资质", ValueType: dbmodel.ValueType_QUATYPE},
	}
	db.Create(&v)
}

//文件类型初始化
func initValue_filetype() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100026", Code: "FILE_USERLOG", Value: "用户头像", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100027", Code: "FILE_SHOPLOG", Value: "店铺头像", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100028", Code: "FILE_PLATFORMIMG", Value: "平台图片", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100029", Code: "FILE_INDEXLBTIMG", Value: "首页轮播图", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100030", Code: "FILE_SHOPINDEXLBTIMG", Value: "店铺轮播图", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100031", Code: "FILE_PRODUCTIMG", Value: "商品图片", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100032", Code: "FILE_QUAFILE", Value: "资质文件", ValueExplain: "文件类型", ValueType: dbmodel.ValueType_FILETYPE},
		{Id: "100033", Code: "FILE_OTHERIMG", Value: "资质文件", ValueExplain: "其他图片", ValueType: dbmodel.ValueType_FILETYPE},
	}
	db.Create(&v)
}

//数据读取权限初始化
func initValue_dataReadRole() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100034", Code: "DATAREADROLE_USER", Value: "本人", ValueExplain: "数据权限——只能读取个人数据", ValueType: dbmodel.ValueType_DATAREADROLE},
		{Id: "100035", Code: "DATAREADROLE_Group", Value: "本人所在组织和下级组织", ValueExplain: "数据权限——只能读取所在组织和下级组织数据", ValueType: dbmodel.ValueType_DATAREADROLE},
		{Id: "100036", Code: "DATAREADROLE_AREA", Value: "本人所在地区和下级地区", ValueExplain: "数据权限——只能读取本人所在地区和下级地区数据", ValueType: dbmodel.ValueType_DATAREADROLE},
		{Id: "100037", Code: "DATAREADROLE_ALL", Value: "所有数据", ValueExplain: "数据权限——能读取所有数据数据", ValueType: dbmodel.ValueType_DATAREADROLE},
	}
	db.Create(&v)
}

//数据插入权限初始化
func initValue_dataInsertRole() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100038", Code: "DATAUPDATEROLE_USER", Value: "可插入", ValueExplain: "数据添加权限——有权", ValueType: dbmodel.ValueType_DATAUPDATEROLE},
		{Id: "100039", Code: "DDATAUPDATEROLE_Group", Value: "不可插入", ValueExplain: "数据添加权限——无权", ValueType: dbmodel.ValueType_DATAUPDATEROLE},
	}
	db.Create(&v)
}

//数据修改权限初始化
func initValue_dataUpdateRole() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100040", Code: "DATAREADROLE_USER", Value: "可修改", ValueExplain: "数据修改权限——有权", ValueType: dbmodel.ValueType_DATAINSERTROLE},
		{Id: "100041", Code: "DATAREADROLE_Group", Value: "不可修改", ValueExplain: "数据修改权限——无权", ValueType: dbmodel.ValueType_DATAINSERTROLE},
	}
	db.Create(&v)
}

//数据读取权限初始化
func initValue_dataDelRole() {
	db := conf.DbConfig.New()
	v := []models.SysValue{
		{Id: "100042", Code: "DATADELROLE_USER", Value: "可删除", ValueExplain: "数据删除权限——有权", ValueType: dbmodel.ValueType_DATADELROLE},
		{Id: "100043", Code: "DATADELROLE_Group", Value: "不可删除", ValueExplain: "数据删除权限——无权", ValueType: dbmodel.ValueType_DATADELROLE},
	}
	db.Create(&v)
}
