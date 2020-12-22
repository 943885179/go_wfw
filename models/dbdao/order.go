package db

import "time"

//Orders 订单主表
type Orders struct {
	Id           string `gorm:"primary_key;type:varchar(50);;"`
	OrderNumber  string `gorm:"column:order_number;not null;comment:'订单号'" json:"order_number"`
	SerialNumber string `gorm:"column:serial_number;comment:'三方支付流水号'" json:"serial_number"`

	ExpressId     string  `gorm:"index;column:express_id;not null;comment:'物流公司ID'" json:"express_id"`
	Express       Express `gorm:"foreignKey:express_id"`
	ExpressNumber string  `gorm:"column:logistics_number;comment:'快递单号'" json:"logistics_number"`

	PayableAmount  float64 `gorm:"column:payable_amount;not null;comment:'应付商品总金额'" json:"payable_amount"`
	RealAmount     float64 `gorm:"column:real_amount;comment:'实付商品总金额(会员折扣,促销规则折扣)'" json:"real_amount"`
	PayableFreight float64 `gorm:"column:payable_freight;not null;comment:'总运费金额'" json:"payable_freight"`
	RealFreight    float64 `gorm:"column:real_freight;comment:'实付运费'" json:"real_freight"`
	Taxes          float64 `gorm:"column:taxes;comment:'税金'" json:"taxes"`
	Promotions     float64 `gorm:"column:promotions;comment:'促销优惠金额和会员折扣'" json:"promotions"`
	Discount       float64 `gorm:"column:discount;comment:'订单折扣或涨价'" json:"discount"`
	OrderAmount    float64 `gorm:"column:order_amount;comment:'订单总金额'" json:"order_amount"`
	Point          float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`

	CreateTime     time.Time `gorm:"column:create_time;not null;comment:'下单时间'" json:"create_time"`
	PayTime        time.Time `gorm:"column:pay_time;comment:'付款时间'" json:"pay_time"`
	SendTime       time.Time `gorm:"column:send_time;comment:'发货时间'" json:"send_time"`
	CompletionTime time.Time `gorm:"column:completion_time;comment:'订单完成时间'" json:"completion_time"`

	Invoice         bool        `gorm:"column:invoice;not null;default:0;comment:'发票：0不索要1索要'" json:"invoice"`
	InvoiceHeard    string      `gorm:"column:invoice_heard;not null;comment:'发票抬头'" json:"invoice_heard"`
	ReceiverName    string      `gorm:"column:receiver_name;not null;comment:'收货人姓名'" json:"receiver_name"`
	ReceiverMobile  string      `gorm:"column:receiver_mobile;not null;comment:'收货人电话'" json:"receiver_mobile"`
	ReceiverAddress string      `gorm:"column:receiver_address;not null;comment:'收货人地址'" json:"receiver_address"`
	ReceiverZipcode string      `gorm:"column:receiver_zipcode;comment:'收货人邮编'" json:"receiver_zipcode"`
	Remark          string      `gorm:"column:remark;comment:'订单备注'" json:"remark"`
	Note            string      `gorm:"column:note;comment:'管理员备注和促销规则描述'" json:"note"`
	Address         string      `gorm:"column:address;" json:"address"`
	Version         int         `gorm:"column:version;" json:"version"`
	OrderItems      []OrderItem `gorm:"foreignKey:order_id" json:"order_items"`

	AreaId string `gorm:"index;column:area_id;comment:'区id'" json:"area_id"`
	ShopId string `gorm:"column:shop_id;comment:'商家'" json:"shop_id"`
	UserId string `gorm:"column:user_id;comment:'用户id'" json:"user_id"`
	PropId string `gorm:"column:prop_id;comment:'使用的道具id'"json:"prop_id"`

	OrderStatusCode        string `gorm:"column:order_status_code;comment:'订单状态(sys_value  生成订单,支付订单,取消订单(客户触发),作废订单(管理员触发),完成订单,退款(订单完成后),部分退款(订单完成后))'  json:"order_status_code"`
	PayTypeCode            string `gorm:"column:pay_type_code;comment:'支付方式(sys_value 微信，支付宝,快捷通，。。。)' json:"pay_type_code"`
	PayStatusCode          string `gorm:"column:pay_status_code;comment:'支付状态(sys_value 0未支付,1已支付)'" json:"pay_status_code"`
	DistributionStatusCode string `gorm:"column:distribution_status_code;comment:'配送状态(sys_value 0：未发送,1：已发送,2：部分发送)'"json:"distribution_status_code"`
	InvoiceTypCode         string `gorm:"column:invoice_typ_code;comment:'发票类型,增值税发票，普通发票'" json:"invoice_typ_code"`
	Model
}

//OrderItem 订单明细表
type OrderItem struct {
	Id           string  `gorm:"primary_key;type:varchar(50);"`
	UnitPrice    float64 `gorm:"column:unit_price;not null;comment:'商品价格（单价（不打折前价格））'" json:"unit_price"`
	PayUnitPrice float64 `gorm:"column:pay_unit_price;not null;comment:'商品价格（支付单价）'" json:"pay_unit_price"`
	Quantity     float64 `gorm:"column:quantity;not null;comment:'购买数量'" json:"quantity"`
	SendQuantity float64 `gorm:"column:send_quantity;not null;comment:'实发数量'" json:"send_quantity"`
	TotalPrice   float64 `gorm:"column:total_price;not null;comment:'总价'" json:"total_price"`

	OrderId string `gorm:"index;column:order_id;" json:"order_id"`
	//Orders       Orders  `gorm:"foreignKey:order_id"`
	PrdSkuId string `gorm:"index;column:prd_sku_id;not null;comment:'商品编码'" json:"prd_sku_id"`

	PrdTypCode string `gorm:"column:prd_typ_code;comment:'商品类型，药品，包装费，邮费，其他等' json:"prd_typ_code"`
	Model
}
