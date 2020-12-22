package db

//LogisticsAddress 收货/发货地址管理
type LogisticsAddress struct {
	Id        string `gorm:"primary_key;type:varchar(50);"`
	IsDefault bool   `gorm:"column:is_default;not null;default:0;comment:'是否默认收发货地址'" json:"is_default"`
	Address   string `gorm:"column:address;" json:"address"`

	UserId string `gorm:"index;column:user_id;not null;comment:'用户id'" json:"user_id"`
	AreaId string `gorm:"index;column:area_id;comment:'区id'"json:"area_id"`

	Model
}
