package db

//SysShop 商家店铺基础信息表
type SysShop struct {
	Id          string  `gorm:"primary_key;type:varchar(50);"`
	ShopName    string  `gorm:"column:shop_name;not null;comment:'店铺名称';unique" json:"shop_name"`
	ShopExplain string  `gorm:"column:shop_explain;size:800;comment:'公告描述'" json:"shop_explain"`
	AppId       string  `gorm:"column:app_id;" json:"app_id"`
	Appsecret   string  `gorm:"column:appsecret;" json:"appsecret"`
	Grade       float64 `gorm:"index;column:grade;default:5;comment:'评分总数'" json:"grade"`
	GradeWl     float64 `gorm:"index;column:grade_wl;default:5;comment:'物流评分总数'" json:"grade_wl"`
	GradeFw     float64 `gorm:"index;column:grade_fw;default:5;comment:'服务评分总数'" json:"grade_fw"`
	GradeMs     float64 `gorm:"index;column:grade_ms;default:5;comment:'描述评分总数'" json:"grade_ms"`
	Cash        int     `gorm:"index;column:cash;comment:'保证金'" json:"cash"`
	Sort        int     `gorm:"index;column:sort;comment:'排序';default:100" json:"sort"`
	Comments    int     `gorm:"index;column:comments;comment:'评价次数'" json:"comments"`
	Point       float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	Vip         int     `gorm:"index;column:vip;comment:'vip等级'" json:"vip"`

	AreaId string `gorm:"index;column:area_id;comment:'区'" json:"area_id"`
	LogoId string `gorm:"index;column:logo_id;comment:'店铺logo'" json:"logo_id"`
	UserId string `gorm:"index;column:user_id;comment:'用户id'" json:"user_id"`

	SysShopBusinessRanges []SysShopBusinessRange `gorm:"foreignKey:sys_shop_id" json:"sys_shop_business_ranges"`
	SysShopImgs           []SysShopImg           `gorm:"foreignKey:sys_shop_id"   json:"sys_shop_imgs"`
	SysShopQualifications []SysShopQualification `gorm:"foreignKey:sys_shop_id"   json:"sys_shop_qualifications"` //资质
	Model
}

//SysShopCustomer 商家客户对照表
type SysShopCustomer struct {
	Id           string  `gorm:"primary_key;type:varchar(50);"`
	Point        float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	Price        float64 `gorm:"column:price;default:0;comment:'店铺余额'" json:"price"`
	HasPrice     float64 `gorm:"column:has_price;default:0;comment:'剩余积分'" json:"has_price"`
	HasIntergral float64 `gorm:"column:has_intergral;default:0;comment:'剩余积分'" json:"has_intergral"`

	SysShopId string  `gorm:"index;column:sys_shop_id;not null;comment:'店铺id'" json:"sys_shop_id"`
	SysShop   SysShop `gorm:"foreignKey:sys_shop_id" json:"sys_shop"`

	CustomerId string `gorm:"index;column:customer_id;not null;comment:'客户id'" json:"customer_id"`
	Model
}

type SysShopQualification struct {
	Id        string `gorm:"primary_key;type:varchar(50);"`
	SysShopId string `gorm:"index;column:sys_shop_id;comment:'店铺id'" json:"sys_shop_id"`
	QuaId     string `gorm:"index;column:qua_id;comment:'资质id'" json:"qua_id"`
}

type SysShopImg struct {
	Id        string `gorm:"primary_key;type:varchar(50);"`
	SysShopId string `gorm:"index;column:sys_shop_id;comment:'店铺id'" json:"sys_shop_id"`
	ImgId     string `gorm:"index;column:img_id;comment:'资质id'" json:"img_id"`
}
type SysShopBusinessRange struct {
	Id              string `gorm:"primary_key;type:varchar(50);"`
	SysShopId       string `gorm:"index;column:sys_shop_id;comment:'店铺id'" json:"sys_shop_id"`
	BusinessRangeId string `gorm:"index;column:img_id;comment:'经营范围id'"  json:"business_range_id"`
}

//Express 快递公司设置表
type Express struct {
	Id             string    `gorm:"primary_key;type:varchar(50);"`
	ExpressName    int       `gorm:"column:express_name;not null;comment:'快递公司名称'" json:"express_name"`
	ExpressURL     int       `gorm:"column:express_url;not null;comment:'快递公司网址'" json:"express_url"`
	ExpressExplain string    `gorm:"column:express_explain;comment:'说明'" json:"express_explain"`
	Freight        []Freight `gorm:"foreignKey:express_id"`
	ShopId         int       `gorm:"index;column:shop_id;not null;comment:'店铺iD'" json:"shop_id"`
	Shop           SysShop   `gorm:"foreignKey:shop_id"`
	Model
}

//Freight 运费设置
type Freight struct {
	Id           string  `gorm:"primary_key;type:varchar(50);"`
	FirstPrice   float64 `gorm:"column:first_price;not null;comment:'首重价格'" json:"first_price"`
	FirstWeight  float64 `gorm:"column:first_weight;not null;comment:'首重重量(克)'" json:"first_weight"`
	SecondWeight float64 `gorm:"column:second_weight;not null;comment:'续重重量(克)'" json:"second_weight"`
	SecondPrice  float64 `gorm:"column:second_price;not null;comment:'续重价格'" json:"second_price"`
	ExpressId    string  `gorm:"index;column:express_id;not null;comment:'物流公司ID'" json:"express_id"`
	Express      Express `gorm:"foreignKey:express_id"`
	AreaId       string  `gorm:"index;column:area_id;not null;comment:'地址" json:"area_id"`
	Area         SysTree `gorm:"foreignKey:area_id"`
	IsDefault    bool    `gorm:"column:is_default;not null;comment:'是否为默认（只能有一个是默认的，默认地址支持全部可售范围）'" json:"is_default"`
	Model
}
