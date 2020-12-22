package db

import "time"

//Qualification 资质表
type Qualification struct {
	Id                 string              `gorm:"primary_key;type:varchar(50);"`
	QuaTypeId          string              `gorm:"index;column:qua_type_id;not null;comment:'资质类型'"json:"qua_type_id"`
	QuaType            SysTree             `gorm:"foreignKey:qua_type_id" json:"qua_type"`
	QualificationFiles []QualificationFile `gorm:"foreignKey:qua_id"  json:"qualification_files"` //资质文件
	QuaExplain         string              `gorm:"column:qua_explain;type:json;comment:'资质描述'" json:"qua_explain"`
	StartTime          time.Time           `gorm:"column:start_time;comment:'注册日期'" json:"start_time"`
	EndTime            time.Time           `gorm:"column:end_time;comment:'过期日期'" json:"end_time"`
	QuaNumber          string              `gorm:"column:qua_number;comment:'资质编号'" json:"qua_number"`

	//外键关联
	ForeignId string `gorm:"index;column:foreign_id;comment:'用户Sys_user/店铺sys_shop/商品ID product'" json:"foreign_id"`

	Model
}

type QualificationFile struct {
	Id     string `gorm:"primary_key;type:varchar(50);"`
	QuaId  string `gorm:"index;column:qua_id;comment:'资质表id(Qualification)'" json:"qua_id"`
	FileId string `gorm:"index;column:file_id;comment:'文件id(sys_file)'" json:"file_id"`
}
