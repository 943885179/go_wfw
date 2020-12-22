package db

import "time"

//Product 商品信息表
type Product struct {
	Id           string  `gorm:"primary_key;type:varchar(50);"`
	GoodsCode    string  `gorm:"column:goods_code;not null;comment:'商品编码'" json:"goods_code"`
	GoodsName    string  `gorm:"column:goods_name;comment:'商品名称'" json:"goods_name"`
	GoodsByname  string  `gorm:"column:goods_byname;comment:'商品名称'" json:"goods_byname"`
	Factory      string  `gorm:"column:factory;comment:'生产厂家'" json:"factory"`
	PrdAddress   string  `gorm:"column:prd_address;comment:'生产地址'" json:"prd_address"`
	ApprovalNum  string  `gorm:"column:approval_num;comment:'批准文号'" json:"approval_num"`
	Spec         string  `gorm:"column:spec;comment:'药品规格（12粒*2版）'" json:"spec"`
	DosageForm   string  `gorm:"column:dosage_form;comment:'剂型（胶囊剂。。。）'" json:"dosage_form"`
	Unit         string  `gorm:"column:unit;comment:'单位（盒，瓶。。。）'" json:"unit"`
	Opcode       string  `gorm:"column:opcode;comment:'拼音'" json:"opcode"`
	MpackTotal   float64 `gorm:"column:mpack_total;comment:'中包装数量'" json:"mpack_total"`
	PackTotal    float64 `gorm:"column:pack_total;comment:'件包装数量'" json:"pack_total"`
	IsUnbundled  bool    `gorm:"column:is_unbundled;default:0;comment:'是否可拆零'" json:"is_unbundled"`
	IsStop       bool    `gorm:"column:is_stop;default:0;comment:'是否停售'" json:"is_stop"`
	GoodsExplain string  `gorm:"column:goods_explain;size:800;comment:'商品描述'" json:"goods_explain"`

	Visit        int     `gorm:"column:visit;comment:'浏览次数'" json:"visit"`
	Favorite     int     `gorm:"column:favorite;comment:'收藏次数'" json:"favorite"`
	Comments     int     `gorm:"column:comments;comment:'评论次数'" json:"comments"`
	Sale         int     `gorm:"column:sale;comment:'销量'" json:"sale"`
	Sort         int     `gorm:"column:sort;comment:'排序';default:100;" json:"sort"`
	Keywords     string  `gorm:"column:keywords;comment:'SEO关键词'" json:"keywords"`
	Config       string  `gorm:"column:config;comment:'商品配置'" json:"config"`
	Stock        float64 `gorm:"column:stock;not null;comment:'库存'" json:"stock"`
	StockEarly   float64 `gorm:"column:stock_early;not null;comment:'库存预警值'" json:"stock_early"`
	SellPriceMin float64 `gorm:"column:sell_price_min;default:9999;comment:'销售价格'" json:"sell_price_min"`
	SellPriceMax float64 `gorm:"column:sell_price_max;default:9999;comment:'销售价格'" json:"sell_price_max"`
	SalePriceMin float64 `gorm:"column:sale_price_min;default:9999;comment:'批发价格'" json:"sale_price_min"`
	SalePriceMax float64 `gorm:"column:sale_price_max;default:9999;comment:'批发价格'" json:"sale_price_max"`

	GoodsImg               string       `gorm:"index;column:goods_img;comment:'商品图片id'" json:"goods_img"`
	ImgFile                SysFile      `gorm:"foreignKey:goods_img"`
	ProductClassifyId      string       `gorm:"index;column:product_classify_id;not null;comment:'商品分类（平台统一）'" json:"product_classify_id"`
	ProductClassify        SysTree      `gorm:"foreignKey:product_classify_id" json:"product_classify"`
	ShopClassifyId         string       `gorm:"index;column:shop_classify_id;not null;comment:'商家分类（店铺可编辑）'" json:"shop_classify_id"`
	ShopClassify           SysTree      `gorm:"foreignKey:shop_classify_id"`
	DistributionProportion string       `gorm:"column:distribution_proportion;not null;default:0;comment:'分销类型，0千分比，1.固定金额，2百分比'" json:"distribution_proportion"`
	DistributionNumber     float64      `gorm:"column:distribution_number;default:0;comment:'0'" json:"distribution_number"`
	ProductSkus            []ProductSku `gorm:"foreignKey:product_id" json:"product_skus"`
	PrdType                string       `gorm:"index,column:'prd_type',not null;comment:'商品审核状态';default:0" json:"prd_type"`

	ProductBusinessRanges []ProductBusinessRange `gorm:"foreignKey:product_id" json:"product_business_ranges"`
	ProductImgs           []ProductImg           `gorm:"foreignKey:product_id"  json:"product_imgs"`
	ProductQualifications []ProductQualification `gorm:"foreignKey:product_id"   json:"product_qualifications""` //资质

	ShopId string `gorm:"index;column:shop_id;not null;comment:'商家编号'" json:"shop_id"`
	AreaId string `gorm:"index;column:area_id;comment:'区id'" json:"area_id"`
	Model
}

//ProductSku 商品规格表
type ProductSku struct {
	Id             string    `gorm:"primary_key;type:varchar(50);"`
	SkuName        string    `gorm:"column:sku_name;not null;comment:'Sku值（医药批发多批号，所以默认为批号）'" json:"sku_name"`
	AttriList      string    `gorm:"column:attri_list;type:text;size:800;not null;comment:'Sku描述（这里还没想好，初步打算存放json）'" json:"attri_list"` //datatypes.JSON
	Point          float64   `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	SellPrice      float64   `gorm:"column:sell_price;default:9999;comment:'销售价格'" json:"sell_price"`
	MarketPrice    float64   `gorm:"column:market_price;default:9999;comment:'市场价格'" json:"market_price"`
	CostPrice      float64   `gorm:"column:cost_price;default:9999;comment:'成本价格'" json:"cost_price"`
	SalePrice      float64   `gorm:"column:sale_price;default:9999;comment:'批发价格'" json:"sale_price"`
	Stock          float64   `gorm:"column:stock;not null;comment:'库存'" json:"stock"`
	BatchNumber    string    `gorm:"column:batch_number;comment:'批号'" json:"batch_number"`
	ProdutionDate  time.Time `gorm:"column:prodution_date;date;comment:'生产日期'" json:"prodution_date"`
	EffectiveDate  time.Time `gorm:"column:effective_date;date;comment:'有效期至'" json:"effective_date"`
	IsChecked      bool      `gorm:"column:is_checked;not null;comment:'是否选择'" json:"is_checked"`
	IsErpSalePrice bool      `gorm:"column:is_erp_sale_price;not null;comment:'是否同步erp价格'"json:"is_erp_sale_price"`
	IsErpStock     bool      `gorm:"column:is_erp_stock;not null;comment:'是否同步erp库存'"  json:"is_erp_stock"`
	ProductId      string    `gorm:"index;column:product_id;not null"  json:"product_id"`
	Product        Product   `gorm:"foreignKey:product_id"`

	ProductSkuImgs []ProductSkuImg `gorm:"foreignKey:product_sku_id"  json:"product_sku_imgs"`
	Model
}

type ProductQualification struct {
	Id        string `gorm:"primary_key;type:varchar(50);"`
	ProductId string `gorm:"index;column:product_id;comment:'店铺id'"json:"product_id"`
	QuaId     string `gorm:"index;column:qua_id;comment:'资质id'" json:"qua_id"`
}

type ProductImg struct {
	Id        string `gorm:"primary_key;type:varchar(50);"`
	ProductId string `gorm:"index;column:product_id;comment:'店铺id'"json:"product_id"`
	ImgId     string `gorm:"index;column:img_id;comment:'文件id'" json:"img_id"`
}
type ProductBusinessRange struct {
	Id              string `gorm:"primary_key;type:varchar(50);"`
	ProductId       string `gorm:"index;column:product_id;comment:'店铺id'"json:"product_id"`
	BusinessRangeId string `gorm:"index;column:img_id;comment:'经营范围id'"  json:"business_range_id"`
}

type ProductSkuImg struct {
	Id           string `gorm:"primary_key;type:varchar(50);"`
	ProductSkuId string `gorm:"index;column:product_sku_id;comment:'店铺id'" json:"product_sku_id"`
	ImgId        string `gorm:"index;column:img_id;comment:'文件id'" json:"img_id"`
}
