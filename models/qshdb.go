package models

import (
	"qshapi/proto/dbmodel"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	//Id                  uint `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt *time.Time `sql:"index"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
	//Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间
	Updated int64 `gorm:"autoUpdateTime:milli" json:"updated"` // 使用时间戳毫秒数填充更新时间
	Created int64 `gorm:"autoCreateTime" json:"created"`       // 使用时间戳秒数填充创建时间
}

/**
 * @Author mzj
 * @Description  资质类型表
 * @Date 下午 2:26 2020/10/19 0019
 * @Param
 * @return
 **/
type Qualification struct {
	Id                   int64                 `gorm:"primary_key"`
	QuaTypeId            int64                 `gorm:"index;column:qua_type_id;not null;comment:'资质类型（身份证正面，身份证背面，营业执照，。。。）'" json:"qua_type_id"`
	QuaType              SysTree               `gorm:"foreignKey:qua_type_id" json:"qua_type"`
	UserId               int64                 `gorm:"index;column:user_id;not null;comment:'用户ID'" json:"user_id"`
	User                 SysUser               `gorm:"foreignKey:user_id"`
	QuaFileId            int64                 `gorm:"index;column:qua_file_id;not null;comment:'资质文件对应id'" json:"qua_file_id"`
	QuaFile              SysFile               `gorm:"foreignKey:qua_file_id" json:"qua_file"`
	QuaExplain           string                `gorm:"column:qua_explain;comment:'资质描述'" json:"qua_explain"`
	StartTime            time.Time             `gorm:"column:start_time;comment:'注册日期'" json:"start_time"`
	EndTime              time.Time             `gorm:"column:end_time;comment:'过期日期'" json:"end_time"`
	QuaNumber            string                `gorm:"column:qua_number;comment:'资质编号'" json:"qua_number"`
	QualificationsRanges []QualificationsRange `gorm:"foreignKey:qualifications_id"`
	Model
}

//SysUser 用户表
type SysUser struct {
	Id           int64   `gorm:"primary_key"`
	UserName     string  `gorm:"column:user_name;not null;comment:'登录名';unique" json:"user_name"` //unique唯一
	TrueName     string  `gorm:"column:true_name;comment:'真实姓名'" json:"true_name"`
	UserPassword string  `gorm:"column:user_password;not null;comment:'登录密码'" json:"user_password"`
	UserPhone    string  `gorm:"column:user_phone;not null;comment:'电话号码';unique" json:"user_phone"`
	UserWx       string  `gorm:"column:user_wx;comment:'微信'" json:"user_wx"`
	WxOpenId     string  `gorm:"column:wx_open_id;comment:'微信公众号openid'" json:"wx_open_id"`
	UserEmail    string  `gorm:"column:user_email;comment:'邮箱'" json:"user_email"`
	UserTel      string  `gorm:"column:user_tel;comment:'电话'" json:"user_tel"`
	UserQq       string  `gorm:"column:user_qq;comment:'QQ'" json:"user_qq"`
	Balance      float64 `gorm:"column:balance;comment:'账户余额'" json:"balance"`
	Point        float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	IdCard       string  `gorm:"column:id_card;comment:'身份证号码'" json:"id_card"`

	ProvinceId int64   `gorm:"index;column:province_id;comment:'省'" json:"province_id"`
	Province   SysTree `gorm:"foreignKey:province_id" json:"province"`
	CityId     int64   `gorm:"index;column:city_id;comment:'市'" json:"city_id"`
	City       SysTree `gorm:"foreignKey:city_id"`
	AreaId     int64   `gorm:"index;column:area_id;comment:'区'" json:"area_id"`
	Area       SysTree `gorm:"foreignKey:area_id"`
	Address    string  `gorm:"column:address;" json:"address"`
	Vip        int     `gorm:"index;column:vip;comment:'vip等级'" json:"vip"`

	BankName     string `gorm:"column:bank_name;comment:'银行名称'" json:"bank_name"`
	BranchName   string `gorm:"column:branch_name;comment:'银行分行名称'" json:"branch_name"`
	Bankcard     string `gorm:"column:bankcard;comment:'银行卡号'" json:"bankcard"`
	BankCardName string `gorm:"column:bank_card_name;comment:'持卡人/单位'" json:"bank_card_name"`

	UserIcon          int64              `gorm:"index;column:user_icon;comment:'用户头像'" json:"user_icon"`
	Icon              SysFile            `gorm:"foreignKey:UserIcon"` //;AssociationforeignKey:Code
	Roles             []SysRole          `gorm:"many2many:sys_user_role" json:"roles"`
	Groups            []SysGroup         `gorm:"many2many:sys_user_group" json:"groups"`
	LogisticsAddresss []LogisticsAddress `gorm:"foreignKey:user_id" json:"logistics_addresss"` //地址管理
	Qualifications    []Qualification    `gorm:"foreignKey:user_id" json:"qualifications"`
	UserType          dbmodel.UserType   `gorm:"column:user_type;comment:'用户类型'" json:"user_type"`
	Model
}

//SysRole 角色表
type SysRole struct {
	Id          int64     `gorm:"primary_key"`
	RoleName    string    `gorm:"column:role_name;not null;comment:'角色名称'" json:"role_name"`
	RoleExplain string    `gorm:"column:role_explain;comment:'角色描述'" json:"role_explain"`
	Menus       []SysMenu `gorm:"many2many:sys_role_menu" json:"menus"`
	Srvs        []SysSrv  `gorm:"many2many:sys_role_syssrv" json:"srvs"`
	Apis        []SysApi  `gorm:"many2many:sys_role_sysapi" json:"apis"`
	//Groups      []SysGroup `gorm:"many2many:sys_group_role"`
	Users    []SysUser `gorm:"many2many:sys_user_role" json:"users"`
	PId      int64     `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Children []SysRole `gorm:"foreignKey:p_id"  json:"children"`
	Model
}

//SysGroup  用户组
type SysGroup struct {
	Id           int64     `gorm:"primary_key"`
	GroupName    string    `gorm:"column:group_name;not null;comment:'用户组名称'" json:"group_name"`
	GroupExplain string    `gorm:"column:group_explain;comment:'用户组描述'" json:"group_explain"`
	Roles        []SysRole `gorm:"many2many:sys_group_role" json:"roles"`
	Users        []SysUser `gorm:"many2many:sys_user_group" json:"users"`
	Model
}

//SysTree 树管理
type SysTree struct {
	Id       int64     `gorm:"primary_key"`
	Code     string    `gorm:"column:code;comment:'编码'" json:"code"`
	Text     string    `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	Sort     int32     `gorm:"column:sort;comment:'排序'" json:"sort"`
	PId      int64     `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Children []SysTree `gorm:"foreignKey:p_id"  json:"children"`
	Model
}

//API接口管理
type SysApi struct {
	Id         int64  `gorm:"primary_key"`
	Service    string `gorm:"column:service;comment:'api接口服务'" json:"service"`
	Method     string `gorm:"column:method;comment:'api接口名称'" json:"method"`
	ApiExplain string `gorm:"column:api_explain;comment:'说明'" json:"api_explain"`
	Model
}

//Srv接口管理
type SysSrv struct {
	Id         int64  `gorm:"primary_key"`
	Service    string `gorm:"column:service;comment:'api接口服务'" json:"service"`
	Method     string `gorm:"column:method;comment:'api接口名称'" json:"method"`
	SrvExplain string `gorm:"column:srv_explain;comment:'说明'" json:"srv_explain"`
	Model
}

//SysMenu 菜单管理
type SysMenu struct {
	Id           int64  `gorm:"primary_key"`
	Key          string `gorm:"column:key;not null;comment:'菜单项唯一标识符，可用于 getItem、setItem 来更新某个菜单'" json:"key"`
	Text         string `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	I18n         string `gorm:"column:i18n;not null;comment:'i18n主键（支持HTML）'" json:"i18n"`
	Group        bool   `gorm:"column:group;tinyint(1);default:0;comment:'	是否显示分组名，指示例中的【主导航】字样'" json:"group,omitempty"`
	Link         string `gorm:"column:link;comment:'路由，link、externalLink 二选其一'" json:"link"`
	ExternalLink string `gorm:"column:external_link;comment:'外部链接'" json:"externalLink"`
	Target       string `gorm:"column:target;comment:'链接 target_blank,_self,_parent,_top'" json:"target"`

	Sort int `gorm:"column:sort;comment:'排序'" json:"sort"`

	Badge       int    `gorm:"column:badge;comment:'标签数量'" json:"badge"`
	BadgeDoc    string `gorm:"column:badge_doc;comment:'标签文字'" json:"badgeDot"`
	BadgeStatus string `gorm:"column:badge_status;comment:'徽标 Badge 颜色'" json:"badgeStatus"`

	Hide             bool   `gorm:"column:hide;tinyint(1);default:0;comment:'是否掩藏'" json:"hide"`
	Disabled         bool   `gorm:"column:disabled;tinyint(1);default:0;comment:'是否禁用'" json:"disabled"`
	HideInBreadcrumb bool   `gorm:"column:hideInBreadcrumb;tinyint(1);default:0;comment:'隐藏面包屑，指 page-header 组件的自动生成面包屑时有效'" json:"hideInBreadcrumb"`
	ACL              string `gorm:"column:acl;comment:'ACL配置若导入 @delon/acl 时自动有效'" json:"acl"`
	Shortcut         bool   `gorm:"column:shortcut;tinyint(1);default:0;comment:'是否快捷菜单项'" json:"shortcut"`
	ShortcutRoot     bool   `gorm:"column:shortcut_root;tinyint(1);default:0;comment:'是否禁用'" json:"shortcutRoot"`
	Reuse            bool   `gorm:"column:reuse;tinyint(1);default:0;comment:'是否允许复用，需配合 reuse-tab 组件'" json:"reuse"`
	Icon             string `gorm:"column:icon;default:'anticon-dashboard';comment:'图标图标'" json:"icon"`

	PId      int64     `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Children []SysMenu `gorm:"foreignKey:p_id"  json:"children,omitempty"`
	Model
}

//SysFile 资源表
type SysFile struct {
	Id          int64            `gorm:"primary_key"`
	Path        string           `gorm:"column:path;not null;comment:'路径'" json:"path"`
	Name        string           `gorm:"column:name;not null;comment:'文件名称（一般是id+后缀）'" json:"name"`
	Size        int64            `gorm:"column:size;comment:'大小'" json:"size"`
	FileExplain string           `gorm:"column:file_explain;comment:'描述'" json:"file_explain"`
	FileType    dbmodel.FileType `gorm:"index;column:file_type;not null;comment:'商业用途（头像，店铺logo，商品图片等）'"json:"file_type"`
	FileSuffix  string           `gorm:"index;column:file_suffix;not null;comment:'文件后缀（.img,.png等）'" json:"file_suffix"`
	Sort        int32            `gorm:"column:sort;coment:'排序'" json:"sort"`
	Model
}

//SysShop 商家店铺基础信息表
type SysShop struct {
	Id          int64  `gorm:"primary_key"`
	ShopName    string `gorm:"column:shop_name;not null;comment:'店铺名称'" json:"shop_name"`
	ShopExplain string `gorm:"column:shop_explain;size:800;comment:'公告描述'" json:"shop_explain"`
	IsSht       bool   `gorm:"column:is_sht;not null;default:0;comment:'四海通认证状态'" json:"is_sht"`
	ShtExplain  string `gorm:"column:sht_explain;comment:'四海通认证返回'" json:"sht_explain"`

	AppId     string `gorm:"column:app_id;" json:"app_id"`
	Appsecret string `gorm:"column:appsecret;" json:"appsecret"`

	Grade    float64 `gorm:"index;column:grade;default:5;comment:'评分总数'" json:"grade"`
	GradeWl  float64 `gorm:"index;column:grade_wl;default:5;comment:'物流评分总数'" json:"grade_wl"`
	GradeFw  float64 `gorm:"index;column:grade_fw;default:5;comment:'服务评分总数'" json:"grade_fw"`
	GradeMs  float64 `gorm:"index;column:grade_ms;default:5;comment:'描述评分总数'" json:"grade_ms"`
	Cash     int     `gorm:"index;column:cash;comment:'保证金'" json:"cash"`
	Sort     int     `gorm:"index;column:sort;comment:'排序'" json:"sort"`
	Comments int     `gorm:"index;column:comments;comment:'评价次数'" json:"comments"`
	Point    float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	Vip      int     `gorm:"index;column:vip;comment:'vip等级'" json:"vip"`

	LogoId int64   `gorm:"index;column:logo_id;comment:'店铺logo'" json:"logo_id"`
	Logo   SysFile `gorm:"foreignKey:logo_id" json:"logo"`

	Classify []SysTree `gorm:"many2many:sys_shop_classify" json:"classify"` //商家分类
	User     []SysUser `gorm:"many2many:sys_shop_user" json:"user"`
	Imgs     []SysFile `gorm:"many2many:sys_shop_imgs" json:"imgs"`
	Model
}

//SysShopCustomer 商家客户对照表
type SysShopCustomer struct {
	Id           int64   `gorm:"primary_key"`
	Point        float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	Price        float64 `gorm:"column:price;default:0;comment:'店铺余额'" json:"price"`
	HasPrice     float64 `gorm:"column:has_price;default:0;comment:'剩余积分'" json:"has_price"`
	HasIntergral float64 `gorm:"column:has_intergral;default:0;comment:'剩余积分'" json:"has_intergral"`

	ShopId     int64   `gorm:"index;column:shop_id;not null;comment:'店铺id'" json:"shop_id"`
	Shop       SysShop `gorm:"foreignKey:shop_id"`
	CustomerId int64   `gorm:"index;column:customer_id;not null;comment:'客户id'" json:"customer_id"`
	Customer   SysUser `gorm:"foreignKey:customer_id"`
	Model
}

//LogisticsAddress 发收货地址管理
type LogisticsAddress struct {
	Id         int64   `gorm:"primary_key"`
	IsDefault  bool    `gorm:"column:is_default;not null;default:0;comment:'是否默认收发货地址'" json:"is_default"`
	UserId     int64   `gorm:"index;column:user_id;not null;comment:'用户id'" json:"user_id"`
	User       SysUser `gorm:"foreignKey:user_id"`
	ProvinceId int64   `gorm:"index;column:province_id;comment:'省id'" json:"province_id"`
	Provuince  SysTree `gorm:"foreignKey:province_id" json:"provuince"`
	CityId     int64   `gorm:"index;column:city_id;comment:'市id'" json:"city_id"`
	City       SysTree `gorm:"foreignKey:city_id"`
	AreaId     int64   `gorm:"index;column:area_id;comment:'区id'"json:"area_id"`
	Area       SysTree `gorm:"foreignKey:area_id" json:"area"`
	Address    string  `gorm:"column:address;" json:"address"`
	Model
}

//Product 商品信息表
type Product struct {
	Id          int64  `gorm:"primary_key"`
	GoodsCode   string `gorm:"column:goods_code;not null;comment:'商品编码'" json:"goods_code"`
	GoodsName   string `gorm:"column:goods_name;comment:'商品名称'" json:"goods_name"`
	Factory     string `gorm:"column:factory;comment:'生产厂家'" json:"factory"`
	ProAddress  string `gorm:"column:pro_address;comment:'生产地址'" json:"pro_address"`
	ApprovalNum string `gorm:"column:approval_num;comment:'批准文号'" json:"approval_num"`
	Spec        string `gorm:"column:spec;comment:'药品规格（12粒*2版）'" json:"spec"`
	DosageForm  string `gorm:"column:dosage_form;comment:'剂型（胶囊剂。。。）'" json:"dosage_form"`
	Unit        string `gorm:"column:unit;comment:'单位（盒，瓶。。。）'" json:"unit"`
	Opcode      string `gorm:"column:opcode;comment:'拼音'" json:"opcode"`
	MpackTotal  int    `gorm:"column:mpack_total;comment:'中包装数量'" json:"mpack_total"`
	PackTotal   int    `gorm:"column:pack_total;comment:'件包装数量'" json:"pack_total"`
	//BatchNumber   string    `gorm:"column:batch_number;comment:'批号'" json:"batch_number"`
	//ProdutionDate time.Time `gorm:"column:prodution_date;date;comment:'生产日期'" json:"prodution_date"`
	//EffectiveDate time.Time `gorm:"column:effective_date;date;comment:'有效期至'" json:"effective_date"`
	IsUnbundled  bool   `gorm:"column:is_unbundled;default:0;comment:'是否可拆零'" json:"is_unbundled"`
	IsStop       bool   `gorm:"column:is_stop;default:0;comment:'是否停售'" json:"is_stop"`
	GoodsExplain string `gorm:"column:goods_explain;size:800;comment:'商品描述'" json:"goods_explain"`

	Visit        int     `gorm:"column:visit;comment:'浏览次数'" json:"visit"`
	Favorite     int     `gorm:"column:favorite;comment:'收藏次数'" json:"favorite"`
	Comments     int     `gorm:"column:comments;comment:'评论次数'" json:"comments"`
	Sale         int     `gorm:"column:sale;comment:'销量'" json:"sale"`
	Sort         int     `gorm:"column:sort;comment:'排序'" json:"sort"`
	Keywords     int     `gorm:"column:keywords;comment:'SEO关键词'" json:"keywords"`
	Config       int     `gorm:"column:config;comment:'商品配置'" json:"config"`
	Stock        float64 `gorm:"column:stock;not null;comment:'库存'" json:"stock"`
	StockEarly   float64 `gorm:"column:stock_early;not null;comment:'库存预警值'" json:"stock_early"`
	SellPriceMin float64 `gorm:"column:sell_price_min;default:9999;comment:'销售价格'" json:"sell_price_min"`
	SellPriceMax float64 `gorm:"column:sell_price_max;default:9999;comment:'销售价格'" json:"sell_price_max"`
	SalePriceMin float64 `gorm:"column:sale_price_min;default:9999;comment:'批发价格'" json:"sale_price_min"`
	SalePriceMax float64 `gorm:"column:sale_price_max;default:9999;comment:'批发价格'" json:"sale_price_max"`

	ShopId                 int64        `gorm:"index;column:shop_id;not null;comment:'商家编号'" json:"shop_id"`
	Shop                   SysShop      `gorm:"foreignKey:shop_id"`
	GoodsImg               int64        `gorm:"index;column:goods_img;comment:'商品图片id'" json:"goods_img"`
	ImgFile                SysFile      `gorm:"foreignKey:goods_img"`
	ProductClassifyId      int64        `gorm:"index;column:product_classify_id;not null;comment:'商品分类（平台统一）'" json:"product_classify_id"`
	ProductClassify        SysTree      `gorm:"foreignKey:product_classify_id"`
	ShopClassifyId         int64        `gorm:"index;column:shop_classify_id;not null;comment:'商家分类（店铺可编辑）'" json:"shop_classify_id"`
	ShopClassify           SysTree      `gorm:"foreignKey:shop_classify_id"`
	DistributionProportion int64        `gorm:"column:distribution_proportion;not null;default:0;comment:'分销类型，0千分比，1.固定金额，2百分比'" json:"distribution_proportion"`
	DistributionNumber     float64      `gorm:"column:distribution_number;default:0;comment:'分销值'" json:"distribution_number"`
	Imgs                   []SysFile    `gorm:"many2many:product_img"`
	BusinessRange          []SysTree    `gorm:"many2many:product_range;"`
	ProductSkus            []ProductSku `gorm:"foreignKey:goods_id"`
	//ProductPartServants    []ProductPartServant `gorm:"foreignKey:goods_id"` //商品分佣表

	ProductLogs []ProductLog `gorm:"foreignKey:goods_id"`
	Model
}

//ProductSku 商品规格表
type ProductSku struct {
	Id           int64   `gorm:"primary_key"`
	SkuName      string  `gorm:"column:sku_name;not null;comment:'Sku值（医药批发多批号，所以默认为批号）'" json:"sku_name"`
	AttriList    string  `gorm:"column:attri_list;size:800;not null;comment:'Sku描述（这里还没想好，初步打算存放json）'" json:"attri_list"`
	Point        float64 `gorm:"column:point;default:0;comment:'积分'" json:"point"`
	SellPrice    float64 `gorm:"column:sell_price;default:9999;comment:'销售价格'" json:"sell_price"`
	MarketPrice  float64 `gorm:"column:market_price;default:9999;comment:'市场价格'" json:"market_price"`
	CostPrice    float64 `gorm:"column:cost_price;default:9999;comment:'成本价格'" json:"cost_price"`
	SalePrice    float64 `gorm:"column:sale_price;default:9999;comment:'批发价格'" json:"sale_price"`
	SalePriceErp float64 `gorm:"column:sale_price_erp;default:9999;comment:'erp批发价格(当批发价为0或者9999读取erp价格)'" json:"sale_price_erp"`
	Stock        float64 `gorm:"column:stock;not null;comment:'库存'" json:"stock"`

	IsChecked bool    `gorm:"column:is_checked;not null;comment:'是否选择'" json:"is_checked"`
	GoodsId   int     `gorm:"index;column:goods_id;not null" json:"goods_id"`
	Goods     Product `gorm:"foreignKey:goods_id"`
	Model
}

//ProductLog 商品日志表
type ProductLog struct {
	Id      int64   `gorm:"primary_key"`
	GoodsId int     `gorm:"index;column:goods_id;not null" json:"goods_id"`
	Goods   Product `gorm:"foreignKey:goods_id"`
	Action  string  `gorm:"column:action;comment:'操作内容（比如发货之类的）'" json:"action"`
	UserId  int64   `gorm:"index;column:user_id;not null" json:"user_id"`
	User    SysUser `gorm:"foreignKey:user_id"`
	Model
}

//PartServant 分佣表
type PartServant struct {
	Id        int64   `gorm:"primary_key"`
	PartType  int64   `gorm:"index;column:part_type;not null;default:0;comment:'分佣方式（1固定金额，0百分比）'" json:"part_type"`
	PartValue float64 `gorm:"column:part_value;not null;comment:'分佣值'" json:"part_value"`
	//GoodsId                   int64    `gorm:"index;column:goods_id;not null;comment:'商品编码'" json:"goods_id"`
	//Goods           Product `gorm:"foreignKey:goods_id"`
	PartPriceTypeId int64     `gorm:"index;column:part_price_type_id;not null;comment:'费用类型（平台推广费。。。）'" json:"part_price_type_id"`
	PartPriceType   SysTree   `gorm:"foreignKey:part_price_type_id"`
	Area            []SysTree `gorm:"many2many:part_servant_"`
	Model
}

//QualificationsRange 资质对应范围
type QualificationsRange struct {
	Id               int64   `gorm:"primary_key"`
	QualificationsId int64   `gorm:"index;column:qualifications_id;not null;comment:'资质对应范围ID'" json:"qualifications_id"`
	GpmId            int64   `gorm:"index;column:gpm_id;not null;comment:'资质类型（生产范围剂型，经营范围，诊疗机构等）'" json:"gpm_id"`
	Gpm              SysTree `gorm:"foreignKey:gpm_id"`
	QualId           int64   `gorm:"index;column:qual_id;not null;comment:'资质范围类型（根据资质类型生产的明细 中成药，等）'"json:"qual_id"`
	Qual             SysTree `gorm:"foreignKey:qual_id"`
	Model
}

//Express 快递公司设置表
type Express struct {
	Id             int64     `gorm:"primary_key"`
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
	Id           int64   `gorm:"primary_key"`
	FirstPrice   float64 `gorm:"column:first_price;not null;comment:'首重价格'" json:"first_price"`
	FirstWeight  float64 `gorm:"column:first_weight;not null;comment:'首重重量(克)'" json:"first_weight"`
	SecondWeight float64 `gorm:"column:second_weight;not null;comment:'续重重量(克)'" json:"second_weight"`
	SecondPrice  float64 `gorm:"column:second_price;not null;comment:'续重价格'" json:"second_price"`
	ExpressId    int64   `gorm:"index;column:express_id;not null;comment:'物流公司ID'" json:"express_id"`
	Express      Express `gorm:"foreignKey:express_id"`
	AreaId       int64   `gorm:"index;column:area_id;not null;comment:'物流公司ID'" json:"area_id"`
	Area         SysTree `gorm:"foreignKey:area_id"`
	IsDefault    bool    `gorm:"column:is_default;not null;comment:'是否为默认（只能有一个是默认的，默认地址支持全部可售范围）'" json:"is_default"`
	Model
}

//Orders 订单主表
type Orders struct {
	Id           int64  `gorm:"primary_key"`
	OrderNumber  string `gorm:"column:order_number;not null;comment:'订单号'" json:"order_number"`
	SerialNumber string `gorm:"column:serial_number;comment:'三方支付流水号'" json:"serial_number"`

	ExpressId     int64   `gorm:"index;column:express_id;not null;comment:'物流公司ID'" json:"express_id"`
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

	Invoice       bool    `gorm:"column:invoice;not null;default:0;comment:'发票：0不索要1索要'" json:"invoice"`
	InvoiceTypeId int     `gorm:"column:invoice_type_id;default:0;comment:'发票类型'" json:"invoice_type_id"`
	InvoiceType   SysTree `gorm:"foreignKey:invoice_type_id"`
	InvoiceHeard  string  `gorm:"column:invoice_heard;not null;comment:'发票抬头'" json:"invoice_heard"`

	ReceiverName    string `gorm:"column:receiver_name;not null;comment:'收货人姓名'" json:"receiver_name"`
	ReceiverMobile  string `gorm:"column:receiver_mobile;not null;comment:'收货人电话'" json:"receiver_mobile"`
	ReceiverAddress string `gorm:"column:receiver_address;not null;comment:'收货人地址'" json:"receiver_address"`
	ReceiverZipcode string `gorm:"column:receiver_zipcode;comment:'收货人邮编'" json:"receiver_zipcode"`

	Remark string `gorm:"column:remark;comment:'订单备注'" json:"remark"`
	Note   string `gorm:"column:note;comment:'管理员备注和促销规则描述'" json:"note"`

	ProvinceId int64   `gorm:"index;column:province_id;comment:'省id'"json:"province_id"`
	Provuince  SysTree `gorm:"foreignKey:province_id"`
	CityId     int64   `gorm:"index;column:city_id;comment:'市id'" json:"city_id"`
	City       SysTree `gorm:"foreignKey:city_id"`
	AreaId     int64   `gorm:"index;column:area_id;comment:'区id'" json:"area_id"`
	Area       SysTree `gorm:"foreignKey:area_id"`
	Address    string  `gorm:"column:address;" json:"address"`

	Version int `gorm:"column:version;" json:"version"`

	UserId    int64   `gorm:"column:user_id;comment:'用户id'" json:"user_id"`
	User      SysUser `gorm:"foreignKey:user_id"`
	PayTypeId int64   `gorm:"column:pay_type_id;comment:'支付方式'"json:"pay_type_id"`
	PayType   SysTree `gorm:"foreignKey:pay_type_id"`
	ShopId    int64   `gorm:"column:shop_id;comment:'商家'" json:"shop_id"`
	Shop      SysShop `gorm:"foreignKey:shop_id"`
	PropId    int64   `gorm:"column:prop_id;comment:'使用的道具id'"json:"prop_id"`
	Prop      PropLog `gorm:"foreignKey:prop_id"`

	OrderStatus int64 `gorm:"column:order_status;not null;comment:'订单状态 1生成订单,2支付订单,3取消订单(客户触发),4作废订单(管理员触发),5完成订单,6退款(订单完成后),7部分退款(订单完成后)'" json:"order_status"`
	//OrderStatus   SysTree `gorm:"foreignKey:order_status"`
	PayStatus int64 `gorm:"column:pay_status;not null;comment:'支付状态 0未支付,1已支付'" json:"pay_status"`
	//PayStatus     SysTree `gorm:"foreignKey:pay_status"`
	DistributionStatus int64 `gorm:"column:distribution_status;not null;comment:'配送状态 0：未发送,1：已发送,2：部分发送'" json:"distribution_status"`
	Model
}

//OrderItem 订单明细表
type OrderItem struct {
	Id                    int64                  `gorm:"primary_key"`
	GoodsSkuId            int64                  `gorm:"index;column:goods_sku_id;not null;comment:'商品编码'" json:"goods_sku_id"`
	OrderId               int64                  `gorm:"index;column:order_id;" json:"order_id"`
	Orders                Orders                 `gorm:"foreignKey:order_id"`
	UnitPrice             float64                `gorm:"column:unit_price;not null;comment:'商品价格（单价（不打折前价格））'" json:"unit_price"`
	PayUnitPrice          float64                `gorm:"column:pay_unit_price;not null;comment:'商品价格（支付单价）'" json:"pay_unit_price"`
	Quantity              float64                `gorm:"column:quantity;not null;comment:'购买数量'" json:"quantity"`
	SendQuantity          float64                `gorm:"column:send_quantity;not null;comment:'实发数量'" json:"send_quantity"`
	TotalPrice            float64                `gorm:"column:total_price;not null;comment:'总价'" json:"total_price"`
	OrderItemPartServants []OrderItemPartServant `gorm:"foreignKey:order_item_id"`
	Model
}

//OrderLog 订单日志表
type OrderLog struct {
	Id      int64   `gorm:"primary_key"`
	OrderId int64   `gorm:"index;column:order_id;" json:"order_id"`
	Orders  Orders  `gorm:"foreignKey:order_id"`
	Action  string  `gorm:"column:action;comment:'操作内容（比如发货之类的）'" json:"action"`
	UserId  int64   `gorm:"index;column:user_id;not null" json:"user_id"`
	User    SysUser `gorm:"foreignKey:user_id"`
	Model
}

//OrderEvaluate 订单评价表
type OrderEvaluate struct {
	Id      int64   `gorm:"primary_key"`
	OrderId int64   `gorm:"index;column:order_id;" json:"order_id"`
	Orders  Orders  `gorm:"foreignKey:order_id"`
	GradeWl float64 `gorm:"index;column:grade_wl;comment:'物流评分总数'" json:"grade_wl"`
	GradeFw float64 `gorm:"index;column:grade_fw;comment:'服务评分总数'" json:"grade_fw"`
	GradeMs float64 `gorm:"index;column:grade_ms;comment:'描述评分总数'" json:"grade_ms"`

	Content string    `gorm:"column:content;size:800;comment:'内容'" json:"content"`
	Imgs    []SysFile `gorm:"many2many:order_evaluate_imgs"`
	Model
}

//OrderItemPartServant 商品分佣明细表(每个)
type OrderItemPartServant struct {
	Id            int64       `gorm:"primary_key"`
	OrderItemId   int64       `gorm:"index;column:order_item_id;not null;comment:'订单明细id'" json:"order_item_id"`
	OrderItem     OrderItem   `gorm:"foreignKey:order_item_id"`
	PartServantId int64       `gorm:"index;column:part_servant_id;not null;comment:'分佣明细id'" json:"part_servant_id"`
	PartServant   PartServant `gorm:"foreignKey:part_servant_id"`
	Model
}

//Cart 购物车
type Cart struct {
	Id         int64   `gorm:"primary_key"`
	UserId     int64   `gorm:"index;column:userid;not null" json:"userid"`
	GoodsSkuId int64   `gorm:"index;column:goods_sku_id;not null" json:"goods_sku_id"`
	Quantity   float64 `gorm:"column:quantity;not null" json:"quantity"`
	IsChecked  bool    `gorm:"column:is_checked;not null" json:"is_checked"`
	Model
}

//Prop  道具表（优惠券等）
type Prop struct {
	Id            int64     `gorm:"primary_key"`
	PropName      string    `gorm:"column:prop_name;not null;comment:'道具名称'" json:"prop_name"`
	CardName      string    `gorm:"column:card_name;comment:'道具的卡号'" json:"card_name"`
	CardPwd       string    `gorm:"column:card_pwd;comment:'道具的密码'" json:"card_pwd"`
	StartTime     time.Time `gorm:"column:start_time;not null;comment:'开始时间'" json:"start_time"`
	EndTime       time.Time `gorm:"column:end_time;not null;comment:'结束时间'" json:"end_time"`
	CouponExplain string    `gorm:"column:coupon_explain;not null;comment:'描述'" json:"coupon_explain"`

	Value      float64   `gorm:"column:value;not null;comment:'面值'" json:"value"`
	LimitSum   float64   `gorm:"column:limit_sum;comment:'满多少元可用'" json:"limit_sum"`
	Point      float64   `gorm:"column:point;default:0;comment:'兑换所需积分（0表示不需要积分兑换）'" json:"point"`
	Condition  string    `gorm:"column:condition;size:400;comment:'条件数据'" json:"condition"`
	Type       int       `gorm:"index;column:type;not null;default:0;comment:'道具类型 0:优惠券'" json:"type"`
	IsUserd    bool      `gorm:"column:is_userd;tinyint(1);default:0;comment:'是否启用'" json:"is_userd"`
	PropNumber int       `gorm:"column:prop_number;not null;default:99999;comment:'领取剩余数量'" json:"prop_number"`
	ImgId      int64     `gorm:"column:img_id;comment:'道具图片'" json:"img_id"`
	Img        SysFile   `gorm:"foreignKey:img_id"`
	ShopId     int64     `gorm:"column:shop_id;comment:'商家'" json:"shop_id"`
	Shop       SysShop   `gorm:"foreignKey:shop_id"`
	Goods      []Product `gorm:"many2many:prop_goods"` //针对特定产品使用
	Model
}

//PropLog 优惠券领取记录
type PropLog struct {
	Id       int64   `gorm:"primary_key"`
	PropId   int64   `gorm:"index;column:prop_id;not null" json:"prop_id"`
	Prop     Prop    `gorm:"foreignKey:prop_id"`
	UserId   int64   `gorm:"index;column:user_id;not null" json:"user_id"`
	User     SysUser `gorm:"foreignKey:user_id"`
	IsExpire bool    `gorm:"column:is_expire;not null;comment:'是否过期'" json:"is_expire"`
	IsUserd  bool    `gorm:"column:is_userd;tinyint(1);default:0;comment:'是否使用'" json:"is_userd"`
	Model
}
