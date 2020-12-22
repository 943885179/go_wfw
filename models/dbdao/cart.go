package db

//Cart 购物车
type Cart struct {
	Id       string  `gorm:"primary_key;type:varchar(50);"`
	UserId   string  `gorm:"index;column:userid;not null" json:"userid"`
	PrdSkuId string  `gorm:"index;column:prd_sku_id;not null" json:"prd_sku_id"`
	Quantity float64 `gorm:"column:quantity;not null;default:1;comment:'数量'" json:"quantity"`
	Model
}
