package db

import "time"

//Prop  道具表（优惠券等）
type Prop struct {
	Id            string    `gorm:"primary_key;type:varchar(50);"`
	PropName      string    `gorm:"column:prop_name;not null;comment:'道具名称'" json:"prop_name"`
	CardName      string    `gorm:"column:card_name;comment:'道具的卡号'" json:"card_name"`
	CardPwd       string    `gorm:"column:card_pwd;comment:'道具的密码'" json:"card_pwd"`
	StartTime     time.Time `gorm:"column:start_time;not null;comment:'开始时间'" json:"start_time"`
	EndTime       time.Time `gorm:"column:end_time;not null;comment:'结束时间'" json:"end_time"`
	CouponExplain string    `gorm:"column:coupon_explain;not null;comment:'描述'" json:"coupon_explain"`

	Value      float64 `gorm:"column:value;not null;comment:'面值'" json:"value"`
	LimitSum   float64 `gorm:"column:limit_sum;comment:'满多少元可用'" json:"limit_sum"`
	Point      float64 `gorm:"column:point;default:0;comment:'兑换所需积分（0表示不需要积分兑换）'" json:"point"`
	Condition  string  `gorm:"column:condition;size:400;comment:'条件数据'" json:"condition"`
	IsUserd    bool    `gorm:"column:is_userd;tinyint(1);default:0;comment:'是否启用'" json:"is_userd"`
	PropNumber int     `gorm:"column:prop_number;not null;default:99999;comment:'领取剩余数量'" json:"prop_number"`

	PropPrds []PropPrd `gorm:"foreignKey:prop_id" json:"prop_prds"` //针对特定产品使用

	ImgId  string `gorm:"column:img_id;comment:'道具图片'" json:"img_id"`
	ShopId string `gorm:"column:shop_id;comment:'商家'" json:"shop_id"`

	TypeCode string `gorm:"column:type_code;comment:'道具类型:优惠券'" json:"type_code"`
	Model
}
type PropPrd struct {
	Id     string `gorm:"primary_key;type:varchar(50);"`
	PropId string `gorm:"index;column:prop_id;comment:'道具id'" json:"prop_id"`
	PrdId  string `gorm:"index;column:prd_id;comment:'商品id'" json:"prd_id"`
}
