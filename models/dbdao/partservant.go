package db

//PartServant 分佣表
type PartServant struct {
	Id        string  `gorm:"primary_key;type:varchar(50);"`
	PartType  string  `gorm:"index;column:part_type;not null;default:0;comment:'分佣方式（1固定金额，0百分比）'" json:"part_type"`
	PartValue float64 `gorm:"column:part_value;not null;comment:'分佣值'" json:"part_value"`
	//ProductId                   string    `gorm:"index;column:product_id;not null;comment:'商品编码'" json:"product_id"`
	//Goods           Product `gorm:"foreignKey:product_id"`
	PartPriceTypeId string  `gorm:"index;column:part_price_type_id;not null;comment:'费用类型（平台推广费。。。）'" json:"part_price_type_id"`
	PartPriceType   SysTree `gorm:"foreignKey:part_price_type_id"`
	AreaId          string  `gorm:"index;column:area_id;not null;comment:'地址'" json:"area_id"`
	Area            SysTree `gorm:"foreignKey:area_id"`
	Model
}
