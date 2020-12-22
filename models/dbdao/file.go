package db

//SysFile 资源表
type SysFile struct {
	Id          string `gorm:"primary_key;type:varchar(50);"`
	Path        string `gorm:"column:path;not null;comment:'路径'" json:"path"`
	Name        string `gorm:"column:name;not null;comment:'文件名称（一般是id+后缀）'" json:"name"`
	Size        string `gorm:"column:size;comment:'大小'" json:"size"`
	FileExplain string `gorm:"column:file_explain;comment:'描述'" json:"file_explain"`
	FileSuffix  string `gorm:"index;column:file_suffix;not null;comment:'文件后缀（.img,.png等）'" json:"file_suffix"`
	Sort        int32  `gorm:"column:sort;coment:'排序';default:100" json:"sort"`

	FileType string `gorm:"index;column:file_type;not null;comment:'商业用途（头像，店铺logo，商品图片等，读取自sys_value code）'"json:"file_type"`
	Model
}
