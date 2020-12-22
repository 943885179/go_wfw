package db

//ServiceLog 请求服务日志表
type ServiceLog struct {
	Id string `gorm:"primary_key;type:varchar(50);" `
	Model
}
