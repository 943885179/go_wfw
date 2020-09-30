package models

import (
	"time"
)

type Model struct {
	//ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//SysUser 用户表
type SysUser struct {
	ID        int64 `gorm:"primary_key"`
	UserName     string  `gorm:"column:user_name;type:varchar(50);not null;comment:'登录名';unique" json:"user_name"` //unique唯一
	TrueName     string  `gorm:"column:true_name;type:varchar(50);comment:'真实姓名'" json:"true_name"`
	UserPassword string  `gorm:"column:user_password;type:varchar(80);not null;comment:'登录密码'" json:"user_password"`
	UserPhone    string  `gorm:"column:user_phone;type:varchar(20);not null;comment:'电话号码';unique" json:"user_phone"`
	UserWx       string  `gorm:"column:user_wx;type:varchar(20);comment:'微信'" json:"user_wx"`
	UserIcon     int     `gorm:"index;column:user_icon;type:int;comment:'用户头像'" json:"user_icon"`
	Icon         SysFile `gorm:"ForeignKey:UserIcon"` //;AssociationForeignKey:Code

	WxOpenid  string  `gorm:"column:wx_openid;type:varchar(50);comment:'微信公众号openid'" json:"wx_openid"`
	UserEmail string  `gorm:"column:user_email;type:varchar(50);comment:'邮箱'" json:"user_email"`
	UserTel   string  `gorm:"column:user_tel;type:varchar(20);comment:'电话'" json:"user_tel"`
	Qq        string  `gorm:"column:qq;type:varchar(20);comment:'QQ'" json:"qq"`
	Balance   float64 `gorm:"column:balance;type:decimal(18,4);comment:'账户余额'" json:"balance"`
	Point     float64 `gorm:"column:point;type:decimal(18,4);default:0;comment:'积分'" json:"point"`
	IDcard    string  `gorm:"column:idcard;type:varchar(50);comment:'身份证号码'" json:"idcard"`

	ProvinceID uint    `gorm:"index;column:province_id;type:int;comment:'省'" json:"province_id"`
	Province   SysTree `gorm:"Foreignkey:province_id"`
	CityID     uint    `gorm:"index;column:city_id;type:int;comment:'市'" json:"city_id"`
	City       SysTree `gorm:"Foreignkey:city_id"`
	AreaID     uint    `gorm:"index;column:area_id;type:int;comment:'区'" json:"area_id"`
	Area       SysTree `gorm:"Foreignkey:area_id"`
	Address    string  `gorm:"column:address;type:varchar(200)" json:"address"`
	Vip        int     `gorm:"index;column:vip;type:int;comment:'vip等级'" json:"vip"`

	BankName     string `gorm:"column:bank_name;type:varchar(100);comment:'银行名称'" json:"bank_name"`
	BranchName   string `gorm:"column:branch_name;type:varchar(100);comment:'银行分行名称'" json:"branch_name"`
	Bankcard     string `gorm:"column:bankcard;type:varchar(30);comment:'银行卡号'" json:"bankcard"`
	BankCardName string `gorm:"column:bank_card_name;type:varchar(100);comment:'持卡人/单位'" json:"bank_card_name"`

	Roles             []SysRole          `gorm:"many2many:sys_user_role"`
	Groups            []SysGroup         `gorm:"many2many:sys_user_group"`
	LogisticsAddresss []LogisticsAddress `gorm:"Foreignkey:user_id"` //地址管理
	Qualificationss   []Qualifications   `gorm:"Foreignkey:user_id"` //资质管理
	Model
}

//SysRole 角色表
type SysRole struct {
	ID        int64 `gorm:"primary_key"`
	RoleName    string     `gorm:"column:role_name;type:varchar(50);not null;comment:'角色名称'" json:"role_name"`
	RoleExplain string     `gorm:"column:role_explain;type:varchar(200);comment:'角色描述'" json:"role_explain"`
	Menus       []SysMenu  `gorm:"many2many:sys_role_menu"`
	Groups      []SysGroup `gorm:"many2many:sys_group_role"`
	Users       []SysUser  `gorm:"many2many:sys_user_role"`
	Model
}

//SysGroup  用户组
type SysGroup struct {
	ID        int64 `gorm:"primary_key"`
	GroupName    string    `gorm:"column:group_name;type:varchar(40);not null;comment:'用户组名称'" json:"group_name"`
	GroupExplain string    `gorm:"column:group_explain;type:varchar(200);comment:'用户组描述'" json:"group_explain"`
	Roles        []SysRole `gorm:"many2many:sys_group_role"`
	Users        []SysUser `gorm:"many2many:sys_user_group"`
	Model
}

//SysTree 树管理
type SysTree struct {
	ID        int64 `gorm:"primary_key"`
	Code     string    `gorm:"column:code;type:char(20);comment:'编码'" json:"code"`
	Text     string    `gorm:"column:text;type:varchar(20);not null;comment:'树名称'" json:"text"`
	Sort     int       `gorm:"column:sort;type:int;comment:'排序'" json:"sort"`
	Pid      uint      `gorm:"column:pid;type:int;comment:'上级id，为0表示没有上级'" json:"pid"`
	Children []SysTree `gorm:"ForeignKey:pid"  json:"children"`
	Model
}

//SysMenu 菜单管理
type SysMenu struct {
	ID        int64 `gorm:"primary_key"`
	Key          string `gorm:"column:key;type:varchar(20);not null;comment:'菜单项唯一标识符，可用于 getItem、setItem 来更新某个菜单'" json:"key"`
	Text         string `gorm:"column:text;type:varchar(20);not null;comment:'树名称'" json:"text"`
	I18n         string `gorm:"column:i18n;type:varchar(20);not null;comment:'i18n主键（支持HTML）'" json:"i18n"`
	Group        bool   `gorm:"column:group;tinyint(1);default:0;comment:'	是否显示分组名，指示例中的【主导航】字样'" json:"group,omitempty"`
	Link         string `gorm:"column:link;type:varchar(100);comment:'路由，link、externalLink 二选其一'" json:"link"`
	ExternalLink string `gorm:"column:external_link;type:varchar(100);comment:'外部链接'" json:"externalLink"`
	Target       string `gorm:"column:target;type:varchar(200);comment:'链接 target_blank,_self,_parent,_top'" json:"target"`

	Sort int `gorm:"column:sort;type:int;comment:'排序'" json:"sort"`

	Badge       int    `gorm:"column:badge;type:int;comment:'标签数量'" json:"badge"`
	BadgeDoc    string `gorm:"column:badge_doc;type:varchar(200);comment:'标签文字'" json:"badgeDot"`
	BadgeStatus string `gorm:"column:badge_status;type:varchar(200);comment:'徽标 Badge 颜色'" json:"badgeStatus"`

	Hide             bool   `gorm:"column:hide;tinyint(1);default:0;comment:'是否掩藏'" json:"hide"`
	Disabled         bool   `gorm:"column:disabled;tinyint(1);default:0;comment:'是否禁用'" json:"disabled"`
	HideInBreadcrumb bool   `gorm:"column:hideInBreadcrumb;tinyint(1);default:0;comment:'隐藏面包屑，指 page-header 组件的自动生成面包屑时有效'" json:"hideInBreadcrumb"`
	ACL              string `gorm:"column:acl;type:varchar(200);comment:'ACL配置若导入 @delon/acl 时自动有效'" json:"acl"`
	Shortcut         bool   `gorm:"column:shortcut;tinyint(1);default:0;comment:'是否快捷菜单项'" json:"shortcut"`
	ShortcutRoot     bool   `gorm:"column:shortcut_root;tinyint(1);default:0;comment:'是否禁用'" json:"shortcutRoot"`
	Reuse            bool   `gorm:"column:reuse;tinyint(1);default:0;comment:'是否允许复用，需配合 reuse-tab 组件'" json:"reuse"`
	Icon             string `gorm:"column:icon;type:varchar(50);default:'anticon-dashboard';comment:'图标图标'" json:"icon"`

	Pid      uint      `gorm:"column:pid;type:int;comment:'上级id，为0表示没有上级'" json:"pid"`
	Children []SysMenu `gorm:"ForeignKey:pid"  json:"children,omitempty"`
	Model
}

//SysFile 资源表
type SysFile struct {
	ID        int64 `gorm:"primary_key"`
	Path         string  `gorm:"column:path;type:varchar(200);not null;comment:'路径'" json:"path"`
	Size         int64 `gorm:"column:size;type:bigint;comment:'大小'" json:"size"`
	FileExplain  string  `gorm:"column:file_explain;type:varchar(100);comment:'描述'" json:"file_explain"`
	FileType int32    `gorm:"index;column:file_type;type:int;not null;comment:'商业用途（头像，店铺logo，商品图片等）'"json:"file_type"`
	FileSuffix string `gorm:"index;column:file_suffix;type:varchar(10);not null;comment:'文件后缀（.img,.png等）'" json:"file_suffix"`
	Sort int32 `gorm:"column:sort;type:int;coment:'排序'" json:"sort"`
	Model
}

//SysShop 商家店铺基础信息表
type SysShop struct {
	ID        int64 `gorm:"primary_key"`
	ShopName    string `gorm:"column:shop_name;type:varchar(100);not null;comment:'店铺名称'" json:"shop_name"`
	ShopExplain string `gorm:"column:shop_explain;type:text;comment:'公告描述'" json:"shop_explain"`
	IsSht       bool   `gorm:"column:is_sht;type:tinyint(1);not null;default:0;comment:'四海通认证状态'" json:"is_sht"`
	ShtExplain  string `gorm:"column:sht_explain;type:varchar(200);comment:'四海通认证返回'" json:"sht_explain"`

	Appid     string `gorm:"column:appid;type:varchar(100)" json:"appid"`
	Appsecret string `gorm:"column:appsecret;type:varchar(100)" json:"appsecret"`

	Grade    float64 `gorm:"index;column:grade;type:decimal(18,4);comment:'评分总数'" json:"grade"`
	GradeWl  float64 `gorm:"index;column:grade_wl;type:decimal(18,4);comment:'物流评分总数'" json:"grade_wl"`
	GradeFw  float64 `gorm:"index;column:grade_fw;type:decimal(18,4);comment:'服务评分总数'" json:"grade_fw"`
	GradeMs  float64 `gorm:"index;column:grade_ms;type:decimal(18,4);comment:'描述评分总数'" json:"grade_ms"`
	Cash     int     `gorm:"index;column:cash;type:int;comment:'保证金'" json:"cash"`
	Sort     int     `gorm:"index;column:sort;type:int;comment:'排序'" json:"sort"`
	Comments int     `gorm:"index;column:comments;type:int;comment:'评价次数'" json:"comments"`
	Point    float64 `gorm:"column:point;type:decimal(18,4);default:0;comment:'积分'" json:"point"`
	Vip      int     `gorm:"index;column:vip;type:int;comment:'vip等级'" json:"vip"`

	LogoID uint    `gorm:"index;column:logo_id;type:int;comment:'店铺logo'" json:"logo_id"`
	Logo   SysFile `gorm:"Foreignkey:logo_id"`

	Classify []SysTree `gorm:"many2many:sys_shop_classify"` //商家分类
	User     []SysUser `gorm:"many2many:sys_shop_user"`
	Imgs     []SysFile `gorm:"many2many:sys_shop_imgs"`
	Model
}

//SysShopCustomer 商家客户对照表
type SysShopCustomer struct {
	ID        int64 `gorm:"primary_key"`
	Point        float64 `gorm:"column:point;type:decimal(18,4);default:0;comment:'积分'" json:"point"`
	Price        float64 `gorm:"column:price;type:decimal(18,4);default:0;comment:'店铺余额'" json:"price"`
	HasPrice     float64 `gorm:"column:has_price;type:decimal(18,4);default:0;comment:'剩余积分'" json:"has_price"`
	HasIntergral float64 `gorm:"column:has_intergral;type:decimal(18,4);default:0;comment:'剩余积分'" json:"has_intergral"`

	Shopid     uint    `gorm:"index;column:shopid;type:int;not null;comment:'店铺id'" json:"shopid"`
	Shop       SysShop `gorm:"Foreignkey:shopid"`
	CustomerID uint    `gorm:"index;column:customer_id;type:int;not null;comment:'客户id'" json:"customer_id"`
	Customer   SysUser `gorm:"Foreignkey:customer_id"`
	Model
}

//LogisticsAddress 发收货地址管理
type LogisticsAddress struct {
	ID        int64 `gorm:"primary_key"`
	IsDefault  bool    `gorm:"column:is_default;type:tinyint(1);not null;default:0;comment:'是否默认收发货地址'" json:"is_default"`
	UserID     uint    `gorm:"index;column:user_id;type:int;not null;comment:'用户id'" json:"user_id"`
	User       SysUser `gorm:"Foreignkey:user_id"`
	ProvinceID uint    `gorm:"index;column:provinceid;type:int;comment:'省id'" json:"province"`
	Provuince  SysTree `gorm:"Foreignkey:provinceid"`
	CityID     uint    `gorm:"index;column:cityid;type:int;comment:'市id'" json:"city"`
	City       SysTree `gorm:"Foreignkey:cityid"`
	AreaID     uint    `gorm:"index;column:areaid;type:int;comment:'区id'" json:"area"`
	Area       SysTree `gorm:"Foreignkey:areaid"`
	Address    string  `gorm:"column:address;type:varchar(200)" json:"address"`
	Model
}

//Product 商品信息表
type Product struct {
	ID        int64 `gorm:"primary_key"`
	GoodsCode   string `gorm:"column:goods_code;type:varchar(50);not null;comment:'商品编码'" json:"goods_code"`
	GoodsName   string `gorm:"column:goods_name;type:varchar(100);comment:'商品名称'" json:"goods_name"`
	Factory     string `gorm:"column:factory;type:varchar(100);comment:'生产厂家'" json:"factory"`
	ProAddress  string `gorm:"column:pro_address;type:varchar(100);comment:'生产地址'" json:"pro_address"`
	ApprovalNum string `gorm:"column:approval_num;type:varchar(100);comment:'批准文号'" json:"approval_num"`
	Spec        string `gorm:"column:spec;type:varchar(50);comment:'药品规格（12粒*2版）'" json:"spec"`
	DosageForm  string `gorm:"column:dosage_form;type:varchar(10);comment:'剂型（胶囊剂。。。）'" json:"dosage_form"`
	Unit        string `gorm:"column:unit;type:varchar(10);comment:'单位（盒，瓶。。。）'" json:"unit"`
	Opcode      string `gorm:"column:opcode;type:varchar(100);comment:'拼音'" json:"opcode"`
	MpackTotal  int    `gorm:"column:mpack_total;type:int;comment:'中包装数量'" json:"mpack_total"`
	PackTotal   int    `gorm:"column:pack_total;type:int;comment:'件包装数量'" json:"pack_total"`
	//BatchNumber   string    `gorm:"column:batch_number;type:varchar(50);comment:'批号'" json:"batch_number"`
	//ProdutionDate time.Time `gorm:"column:prodution_date;type:date;comment:'生产日期'" json:"prodution_date"`
	//EffectiveDate time.Time `gorm:"column:effective_date;type:date;comment:'有效期至'" json:"effective_date"`
	IsUnbundled  bool   `gorm:"column:is_unbundled;type:tinyint(1);default:0;comment:'是否可拆零'" json:"is_unbundled"`
	IsStop       bool   `gorm:"column:is_stop;type:tinyint(1);default:0;comment:'是否停售'" json:"is_stop"`
	GoodsExplain string `gorm:"column:goods_explain;type:text;comment:'商品描述'" json:"goods_explain"`

	Visit        int     `gorm:"column:visit;type:int;comment:'浏览次数'" json:"visit"`
	Favorite     int     `gorm:"column:favorite;type:int;comment:'收藏次数'" json:"favorite"`
	Comments     int     `gorm:"column:comments;type:int;comment:'评论次数'" json:"comments"`
	Sale         int     `gorm:"column:sale;type:int;comment:'销量'" json:"sale"`
	Sort         int     `gorm:"column:sort;type:int;comment:'排序'" json:"sort"`
	Keywords     int     `gorm:"column:keywords;type:varchar(200);comment:'SEO关键词'" json:"keywords"`
	Config       int     `gorm:"column:config;type:varchar(200);comment:'商品配置'" json:"config"`
	Stock        float64 `gorm:"column:stock;type:decimal(18,4);not null;comment:'库存'" json:"stock"`
	StockEarly   float64 `gorm:"column:stock_early;type:decimal(18,4);not null;comment:'库存预警值'" json:"stock_early"`
	SellPriceMin float64 `gorm:"column:sell_price_min;type:decimal(18,4);default:9999;comment:'销售价格'" json:"sell_price_min"`
	SellPriceMax float64 `gorm:"column:sell_price_max;type:decimal(18,4);default:9999;comment:'销售价格'" json:"sell_price_max"`
	SalePriceMin float64 `gorm:"column:sale_price_min;type:decimal(18,4);default:9999;comment:'批发价格'" json:"sale_price_min"`
	SalePriceMax float64 `gorm:"column:sale_price_max;type:decimal(18,4);default:9999;comment:'批发价格'" json:"sale_price_max"`

	ShopID                 uint         `gorm:"index;column:shop_id;type:int;not null;comment:'商家编号'" json:"shop_id"`
	Shop                   SysShop      `gorm:"Foreignkey:shop_id"`
	GoodsImg               uint         `gorm:"index;column:goods_img;type:int;comment:'商品图片id'" json:"goods_img"`
	ImgFile                SysFile      `gorm:"Foreignkey:goods_img"`
	ProductClassifyID      uint         `gorm:"index;column:product_classify_id;type:int;not null;comment:'商品分类（平台统一）'" json:"product_classify_id"`
	ProductClassify        SysTree      `gorm:"Foreignkey:product_classify_id"`
	ShopClassifyID         uint         `gorm:"index;column:shop_classify_id;type:int;not null;comment:'商家分类（店铺可编辑）'" json:"shop_classify_id"`
	ShopClassify           SysTree      `gorm:"Foreignkey:shop_classify_id"`
	DistributionProportion uint         `gorm:"column:distribution_proportion;type:int;not null;default:0;comment:'分销类型，0千分比，1.固定金额，2百分比'" json:"distribution_proportion"`
	DistributionNumber     float64      `gorm:"column:distribution_number;type:decimal(18,4);default:0;comment:'分销值'" json:"distribution_number"`
	Imgs                   []SysFile    `gorm:"many2many:product_img"`
	BusinessRange          []SysTree    `gorm:"many2many:product_range;"`
	ProductSkus            []ProductSku `gorm:"Foreignkey:goods_id"`
	//ProductPartServants    []ProductPartServant `gorm:"Foreignkey:goods_id"` //商品分佣表

	ProductLogs []ProductLog `gorm:"Foreignkey:goods_id"`
	Model
}

//ProductSku 商品规格表
type ProductSku struct {
	ID        int64 `gorm:"primary_key"`
	SkuName      string  `gorm:"column:sku_name;type:varchar(50);not null;comment:'Sku值（医药批发多批号，所以默认为批号）'" json:"sku_name"`
	AttriList    string  `gorm:"column:attri_list;type:text;not null;comment:'Sku描述（这里还没想好，初步打算存放json）'" json:"attri_list"`
	Point        float64 `gorm:"column:point;type:decimal(18,4);default:0;comment:'积分'" json:"point"`
	SellPrice    float64 `gorm:"column:sell_price;type:decimal(18,4);default:9999;comment:'销售价格'" json:"sell_price"`
	MarketPrice  float64 `gorm:"column:market_price;type:decimal(18,4);default:9999;comment:'市场价格'" json:"market_price"`
	CostPrice    float64 `gorm:"column:cost_price;type:decimal(18,4);default:9999;comment:'成本价格'" json:"cost_price"`
	SalePrice    float64 `gorm:"column:sale_price;type:decimal(18,4);default:9999;comment:'批发价格'" json:"sale_price"`
	SalePriceErp float64 `gorm:"column:sale_price_erp;type:decimal(18,4);default:9999;comment:'erp批发价格(当批发价为0或者9999读取erp价格)'" json:"sale_price_erp"`
	Stock        float64 `gorm:"column:stock;type:decimal(18,4);not null;comment:'库存'" json:"stock"`

	IsChecked bool    `gorm:"column:is_checked;type:tinyint(1);not null;comment:'是否选择'" json:"is_checked"`
	GoodsID   int     `gorm:"index;column:goods_id;type:int;not null" json:"goods_id"`
	Goods     Product `gorm:"Foreignkey:goods_id"`
	Model
}

//ProductLog 商品日志表
type ProductLog struct {
	ID        int64 `gorm:"primary_key"`
	GoodsID int     `gorm:"index;column:goods_id;type:int;not null" json:"goods_id"`
	Goods   Product `gorm:"Foreignkey:goods_id"`
	Action  string  `gorm:"column:action;type:varchar(40);comment:'操作内容（比如发货之类的）'" json:"action"`
	UserID  uint    `gorm:"index;column:user_id;type:int;not null" json:"user_id"`
	User    SysUser `gorm:"Foreignkey:user_id"`
	Model
}

//PartServant 分佣表
type PartServant struct {
	ID        int64 `gorm:"primary_key"`
	PartType  uint    `gorm:"index;column:part_type;type:int;not null;default:0;comment:'分佣方式（1固定金额，0百分比）'" json:"part_type"`
	PartValue float64 `gorm:"column:part_value;type:decimal(18,2);not null;comment:'分佣值'" json:"part_value"`
	//GoodsID         uint    `gorm:"index;column:goods_id;type:int;not null;comment:'商品编码'" json:"goods_id"`
	//Goods           Product `gorm:"Foreignkey:goods_id"`
	PartPriceTypeID uint      `gorm:"index;column:part_price_type_id;type:int;not null;comment:'费用类型（平台推广费。。。）'" json:"part_price_type_id"`
	PartPriceType   SysTree   `gorm:"Foreignkey:part_price_type_id"`
	Area            []SysTree `gorm:"many2many:part_servant_"`
	Model
}

//Qualifications 资质表
type Qualifications struct {
	ID        int64 `gorm:"primary_key"`
	QuaTypeID            uint                  `gorm:"index;column:qua_type_id;type:int;not null;comment:'资质类型（身份证正面，身份证背面，营业执照，。。。）'" json:"qua_type_id"`
	QuaType              SysTree               `gorm:"Foreignkey:qua_type_id"`
	UserID               uint                  `gorm:"index;column:user_id;type:int;not null;comment:'用户ID'" json:"user_id"`
	User                 SysUser               `gorm:"Foreignkey:userid"`
	QuaFileID            uint                  `gorm:"index;column:qua_file_id;type:int;not null;comment:'资质文件对应id'" json:"qua_file_id"`
	QuaFile              SysFile               `gorm:"Foreignkey:qua_file_id"`
	QuaExplain           string                `gorm:"column:qua_explain;type:varchar(100);comment:'资质描述'" json:"qua_explain"`
	StartTime            time.Time             `gorm:"column:start_time;type:datetime;comment:'注册日期'" json:"start_time"`
	EndTime              time.Time             `gorm:"column:end_time;type:datetime;comment:'过期日期'" json:"end_time"`
	QuaNumber            string                `gorm:"column:qua_number;type:varchar(40);comment:'资质编号'" json:"qua_number"`
	QualificationsRanges []QualificationsRange `gorm:"Foreignkey:qualifications_id"`
	Model
}

//QualificationsRange 资质对应范围
type QualificationsRange struct {
	ID        int64 `gorm:"primary_key"`
	QualificationsID uint    `gorm:"index;column:qualifications_id;type:int;not null;comment:'资质对应范围ID'" json:"qualifications_id"`
	GpmID            uint    `gorm:"index;column:gpmid;type:int;not null;comment:'资质类型（生产范围剂型，经营范围，诊疗机构等）'" json:"gpmid"`
	Gpm              SysTree `gorm:"Foreignkey:gpmid"`
	QualID           uint    `gorm:"index;column:qualid;type:int;not null;comment:'资质范围类型（根据资质类型生产的明细 中成药，等）'" json:"qualid"`
	Qual             SysTree `gorm:"Foreignkey:qualid"`
	Model
}

//Express 快递公司设置表
type Express struct {
	ID        int64 `gorm:"primary_key"`
	ExpressName    int       `gorm:"column:express_name;type:varchar(50);not null;comment:'快递公司名称'" json:"express_name"`
	ExpressURL     int       `gorm:"column:express_url;type:varchar(50);not null;comment:'快递公司网址'" json:"express_url"`
	ShopID         int       `gorm:"index;column:shop_id;type:int;not null;comment:'店铺iD'" json:"shop_id"`
	Shop           SysShop   `gorm:"Foreignkey:shopid"`
	ExpressExplain string    `gorm:"column:express_explain;type:varchar(200);comment:'说明'" json:"express_explain"`
	Freight        []Freight `gorm:"Foreignkey:express_id"`
	Model
}

//Freight 运费设置
type Freight struct {
	ID        int64 `gorm:"primary_key"`
	FirstPrice   float64 `gorm:"column:first_price;type:decimal(18,4);not null;comment:'首重价格'" json:"first_price"`
	FirstWeight  float64 `gorm:"column:first_weight;type:decimal(18,2);not null;comment:'首重重量(克)'" json:"first_weight"`
	SecondWeight float64 `gorm:"column:second_weight;type:decimal(18,2);not null;comment:'续重重量(克)'" json:"second_weight"`
	SecondPrice  float64 `gorm:"column:second_price;type:decimal(18,4);not null;comment:'续重价格'" json:"second_price"`
	ExpressID    uint    `gorm:"index;column:express_id;type:int;not null;comment:'物流公司ID'" json:"express_id"`
	Express      Express `gorm:"Foreignkey:expressid"`
	AreaID       uint    `gorm:"index;column:area_id;type:int;not null;comment:'物流公司ID'" json:"area_id"`
	Area         SysTree `gorm:"Foreignkey:area_id"`
	IsDefault    bool    `gorm:"column:is_default;type:tinyint(1);not null;comment:'是否为默认（只能有一个是默认的，默认地址支持全部可售范围）'" json:"is_default"`
	Model
}

//Orders 订单主表
type Orders struct {
	ID        int64 `gorm:"primary_key"`
	OrderNumber  string `gorm:"column:order_number;type:varchar(40);not null;comment:'订单号'" json:"order_number"`
	SerialNumber string `gorm:"column:serial_number;type:varchar(40);comment:'三方支付流水号'" json:"serial_number"`

	ExpressID     uint    `gorm:"index;column:express_id;type:int;not null;comment:'物流公司ID'" json:"express_id"`
	Express       Express `gorm:"Foreignkey:express_id"`
	ExpressNumber string  `gorm:"column:logistics_number;type:varchar(40);comment:'快递单号'" json:"logistics_number"`

	PayableAmount  float64 `gorm:"column:payable_amount;type:decimal(18,4);not null;comment:'应付商品总金额'" json:"payable_amount"`
	RealAmount     float64 `gorm:"column:real_amount;type:decimal(18,4);comment:'实付商品总金额(会员折扣,促销规则折扣)'" json:"real_amount"`
	PayableFreight float64 `gorm:"column:payable_freight;type:decimal(18,4);not null;comment:'总运费金额'" json:"payable_freight"`
	RealFreight    float64 `gorm:"column:real_freight;type:decimal(18,4);comment:'实付运费'" json:"real_freight"`
	Taxes          float64 `gorm:"column:taxes;type:decimal(18,4);comment:'税金'" json:"taxes"`
	Promotions     float64 `gorm:"column:promotions;type:decimal(18,4);comment:'促销优惠金额和会员折扣'" json:"promotions"`
	Discount       float64 `gorm:"column:discount;type:decimal(18,4);comment:'订单折扣或涨价'" json:"discount"`
	OrderAmount    float64 `gorm:"column:order_amount;type:decimal(18,4);comment:'订单总金额'" json:"order_amount"`
	Point          float64 `gorm:"column:point;type:decimal(18,4);default:0;comment:'积分'" json:"point"`

	CreateTime     time.Time `gorm:"column:create_time;type:datetime;not null;comment:'下单时间'" json:"create_time"`
	PayTime        time.Time `gorm:"column:pay_time;type:datetime;comment:'付款时间'" json:"pay_time"`
	SendTime       time.Time `gorm:"column:send_time;type:datetime;comment:'发货时间'" json:"send_time"`
	CompletionTime time.Time `gorm:"column:completion_time;type:datetime;comment:'订单完成时间'" json:"completion_time"`

	Invoice       bool    `gorm:"column:invoice;type:tinyint(1);not null;default:0;comment:'发票：0不索要1索要'" json:"invoice"`
	InvoiceTypeID uint    `gorm:"column:invoice_type;type:tinyint(1);default:0;comment:'发票类型'" json:"invoice_type"`
	InvoiceType   SysTree `gorm:"Foreignkey:invoice_type"`
	InvoiceHeard  string  `gorm:"column:invoice_heard;type:varchar(40);not null;comment:'发票抬头'" json:"invoice_heard"`

	ReceiverName    string `gorm:"column:receiver_name;type:varchar(40);not null;comment:'收货人姓名'" json:"receiver_name"`
	ReceiverMobile  string `gorm:"column:receiver_mobile;type:varchar(20);not null;comment:'收货人电话'" json:"receiver_mobile"`
	ReceiverAddress string `gorm:"column:receiver_address;type:varchar(250);not null;comment:'收货人地址'" json:"receiver_address"`
	ReceiverZipcode string `gorm:"column:receiver_zipcode;type:varchar(50);comment:'收货人邮编'" json:"receiver_zipcode"`

	Remark string `gorm:"column:remark;type:varchar(250);comment:'订单备注'" json:"remark"`
	Note   string `gorm:"column:note;type:varchar(250);comment:'管理员备注和促销规则描述'" json:"note"`

	ProvinceID uint    `gorm:"index;column:provinceid;type:int;comment:'省id'" json:"province"`
	Provuince  SysTree `gorm:"Foreignkey:provinceid"`
	CityID     uint    `gorm:"index;column:cityid;type:int;comment:'市id'" json:"city"`
	City       SysTree `gorm:"Foreignkey:cityid"`
	AreaID     uint    `gorm:"index;column:areaid;type:int;comment:'区id'" json:"area"`
	Area       SysTree `gorm:"Foreignkey:areaid"`
	Address    string  `gorm:"column:address;type:varchar(200)" json:"address"`

	Version int `gorm:"column:version;type:int" json:"version"`

	UserID    uint    `gorm:"column:user_id;type:int;comment:'用户id'" json:"user_id"`
	User      SysUser `gorm:"Foreignkey:user_id"`
	PayTypeID uint    `gorm:"column:pay_type;type:int;comment:'支付方式'" json:"pay_type"`
	PayType   SysTree `gorm:"Foreignkey:pay_type"`
	ShopID    uint    `gorm:"column:shop_id;type:int;comment:'商家'" json:"shop_id"`
	Shop      SysShop `gorm:"Foreignkey:shop_id"`
	PropID    uint    `gorm:"column:prop;type:int;comment:'使用的道具id'" json:"prop"`
	Prop      PropLog `gorm:"Foreignkey:prop"`

	OrderStatus uint `gorm:"column:order_status;type:int;not null;comment:'订单状态 1生成订单,2支付订单,3取消订单(客户触发),4作废订单(管理员触发),5完成订单,6退款(订单完成后),7部分退款(订单完成后)'" json:"order_status"`
	//OrderStatus   SysTree `gorm:"Foreignkey:order_status"`
	PayStatus uint `gorm:"column:pay_status;type:int;not null;comment:'支付状态 0未支付,1已支付'" json:"pay_status"`
	//PayStatus     SysTree `gorm:"Foreignkey:pay_status"`
	DistributionStatus uint `gorm:"column:distribution_status;type:int;not null;comment:'配送状态 0：未发送,1：已发送,2：部分发送'" json:"distribution_status"`
	Model
}

//OrderItem 订单明细表
type OrderItem struct {
	ID        int64 `gorm:"primary_key"`
	GoodsSkuID            int                    `gorm:"index;column:goods_sku_id;type:int;not null;comment:'商品编码'" json:"goods_sku_id"`
	OrderID               int                    `gorm:"index;column:order_id;type:int" json:"order_id"`
	Orders                Orders                 `gorm:"Foreignkey:order_id"`
	UnitPrice             float64                `gorm:"column:unit_price;type:decimal(18,4);not null;comment:'商品价格（单价（不打折前价格））'" json:"unit_price"`
	PayUnitPrice          float64                `gorm:"column:pay_unit_price;type:decimal(18,4);not null;comment:'商品价格（支付单价）'" json:"pay_unit_price"`
	Quantity              float64                `gorm:"column:quantity;type:decimal(18,4);not null;comment:'购买数量'" json:"quantity"`
	SendQuantity          float64                `gorm:"column:send_quantity;type:decimal(18,4);not null;comment:'实发数量'" json:"send_quantity"`
	TotalPrice            float64                `gorm:"column:total_price;type:decimal(18,4);not null;comment:'总价'" json:"total_price"`
	OrderItemPartServants []OrderItemPartServant `gorm:"Foreignkey:order_item_id"`
	Model
}

//OrderLog 订单日志表
type OrderLog struct {
	ID        int64 `gorm:"primary_key"`
	OrderID uint    `gorm:"index;column:order_id;type:int" json:"order_id"`
	Orders  Orders  `gorm:"Foreignkey:order_id"`
	Action  string  `gorm:"column:action;type:varchar(40);comment:'操作内容（比如发货之类的）'" json:"action"`
	UserID  uint    `gorm:"index;column:user_id;type:int;not null" json:"user_id"`
	User    SysUser `gorm:"Foreignkey:user_id"`
	Model
}

//OrderEvaluate 订单评价表
type OrderEvaluate struct {
	ID        int64 `gorm:"primary_key"`
	Orders  Orders  `gorm:"Foreignkey:order_id"`
	GradeWl float64 `gorm:"index;column:grade_wl;type:decimal(18,4);comment:'物流评分总数'" json:"grade_wl"`
	GradeFw float64 `gorm:"index;column:grade_fw;type:decimal(18,4);comment:'服务评分总数'" json:"grade_fw"`
	GradeMs float64 `gorm:"index;column:grade_ms;type:decimal(18,4);comment:'描述评分总数'" json:"grade_ms"`

	Content string    `gorm:"column:content;type:text;comment:'内容'" json:"content"`
	Imgs    []SysFile `gorm:"many2many:order_evaluate_imgs"`
	Model
}

//OrderItemPartServant 商品分佣明细表(每个)
type OrderItemPartServant struct {
	ID        int64 `gorm:"primary_key"`
	OrderItemID   uint        `gorm:"index;column:order_item_id;type:int;not null;comment:'订单明细id'" json:"order_item_id"`
	OrderItem     OrderItem   `gorm:"Foreignkey:order_item_id"`
	PartServantID uint        `gorm:"index;column:part_servant_id;type:int;not null;comment:'分佣明细id'" json:"part_servant_id"`
	PartServant   PartServant `gorm:"Foreignkey:part_servant_id"`
	Model
}

//Cart 购物车
type Cart struct {
	ID        int64 `gorm:"primary_key"`
	Userid     int     `gorm:"index;column:userid;type:int;not null" json:"userid"`
	GoodsSkuID int     `gorm:"index;column:goods_sku_id;type:int;not null" json:"goods_sku_id"`
	Quantity   float64 `gorm:"column:quantity;type:decimal(18,4);not null" json:"quantity"`
	IsChecked  bool    `gorm:"column:is_checked;type:tinyint(1);not null" json:"is_checked"`
	Model
}

//Prop  道具表（优惠券等）
type Prop struct {
	ID        int64 `gorm:"primary_key"`
	PropName      string    `gorm:"column:prop_name;type:varchar(40);not null;comment:'道具名称'" json:"prop_name"`
	CardName      string    `gorm:"column:card_name;type:varchar(40);comment:'道具的卡号'" json:"card_name"`
	CardPwd       string    `gorm:"column:card_pwd;type:varchar(40);comment:'道具的密码'" json:"card_pwd"`
	StartTime     time.Time `gorm:"column:start_time;type:datetime;not null;comment:'开始时间'" json:"start_time"`
	EndTime       time.Time `gorm:"column:end_time;type:datetime;not null;comment:'结束时间'" json:"end_time"`
	CouponExplain string    `gorm:"column:coupon_explain;type:varchar(200);not null;comment:'描述'" json:"coupon_explain"`

	Value      float64   `gorm:"column:value;type:decimal(18,4);not null;comment:'面值'" json:"value"`
	LimitSum   float64   `gorm:"column:limit_sum;type:decimal(18,4);comment:'满多少元可用'" json:"limit_sum"`
	Point      float64   `gorm:"column:point;type:decimal(18,4);default:0;comment:'兑换所需积分（0表示不需要积分兑换）'" json:"point"`
	Condition  string    `gorm:"column:condition;type:varchar(400);comment:'条件数据'" json:"condition"`
	Type       int       `gorm:"index;column:type;type:int;not null;default:0;comment:'道具类型 0:优惠券'" json:"type"`
	IsUserd    bool      `gorm:"column:is_userd;tinyint(1);default:0;comment:'是否启用'" json:"is_userd"`
	PropNumber int       `gorm:"column:prop_number;type:int;not null;default:99999;comment:'领取剩余数量'" json:"prop_number"`
	ImgID      uint      `gorm:"column:img;type:int;comment:'道具图片'" json:"img"`
	Img        SysFile   `gorm:"Foreignkey:img"`
	ShopID     uint      `gorm:"column:shop_id;type:int;comment:'商家'" json:"shop_id"`
	Shop       SysShop   `gorm:"Foreignkey:shop_id"`
	Goods      []Product `gorm:"many2many:prop_goods"` //针对特定产品使用
	Model
}

//PropLog 优惠券领取记录
type PropLog struct {
	ID        int64 `gorm:"primary_key"`
	PropID   uint    `gorm:"index;column:prop;type:int;not null" json:"prop"`
	Prop     Prop    `gorm:"Foreignkey:prop"`
	UserID   uint    `gorm:"index;column:user_id;type:int;not null" json:"user_id"`
	User     SysUser `gorm:"Foreignkey:user_id"`
	IsExpire bool    `gorm:"column:is_expire;type:tinyint(1);not null;comment:'是否过期'" json:"is_expire"`
	IsUserd  bool    `gorm:"column:is_userd;tinyint(1);default:0;comment:'是否使用'" json:"is_userd"`
	Model
}