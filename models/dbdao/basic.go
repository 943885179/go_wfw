package db

import "qshapi/proto/dbmodel"

//SysUser 用户表
type SysUser struct {
	Id           string  `gorm:"primary_key;type:varchar(50);"`
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
	Address      string  `gorm:"column:address;" json:"address"`
	Vip          int     `gorm:"index;column:vip;comment:'vip等级'" json:"vip"`
	BankName     string  `gorm:"column:bank_name;comment:'银行名称'" json:"bank_name"`
	BranchName   string  `gorm:"column:branch_name;comment:'银行分行名称'" json:"branch_name"`
	Bankcard     string  `gorm:"column:bankcard;comment:'银行卡号'" json:"bankcard"`
	BankCardName string  `gorm:"column:bank_card_name;comment:'持卡人/单位'" json:"bank_card_name"`

	AreaId  string  `gorm:"index;column:area_id;comment:'所属地区id'" json:"area_id"`
	SysArea SysArea `gorm:"foreignKey:area_id"  json:"sys_area"`

	SysDeps       []SysDep       `gorm:"many2many:sys_user_sys_deps"  json:"sys_deps"`              //组织 (多对多会自动创建一个表，或者用foreignkey，手动创建中间表)
	SysRoles      []SysRole      `gorm:"many2many:sys_user_sys_roles" json:"sys_roles"`             //角色
	SysUserGroups []SysUserGroup `gorm:"many2many:sys_user_sys_user_groups" json:"sys_user_groups"` //用户组
	SysPositions  []SysPosition  `gorm:"many2many:sys_user_sys_positions"  json:"sys_positions"`    //职位管理

	UserIcon string `gorm:"index;column:user_icon;comment:'用户头像'" json:"user_icon"` //关联了另外一个服务（文件服务）
	Model
}

//SysPosition 职位表
type SysPosition struct {
	Id           string     `gorm:"primary_key;type:varchar(50);"`
	Code         string     `gorm:"column:code;not null;comment:'职位编码'" json:"code"`
	Value        string     `gorm:"column:value;not null;comment:'职位名称'" json:"value"`
	ValueExplain string     `gorm:"column:value_explain;not null;comment:'职位说明'"  json:"value_explain"`
	DataRole     string     `gorm:"column:value_explain;not null;comment:'职位数据权限'"json:"data_role"`
	SysValues    []SysValue `gorm:"many2many:sys_position_sys_value" json:"sys_values"` //数据范围权限管理（个人/全部/部门/区域）

	SysUser []SysUser `gorm:"many2many:sys_user_sys_positions"   json:"sys_user"` //职位管理
	Model
}

//SysGroup  用户组
type SysUserGroup struct {
	Id           string    `gorm:"primary_key;type:varchar(50);"`
	GroupName    string    `gorm:"column:group_name;not null;comment:'用户组名称'" json:"group_name"`
	GroupExplain string    `gorm:"column:group_explain;comment:'用户组描述'" json:"group_explain"`
	SysRoles     []SysRole `gorm:"many2many:sys_user_group_sys_roles" json:"sys_roles"`

	SysUsers []SysUser `gorm:"many2many:sys_user_sys_user_groups" json:"sys_users"`
	Model
}

//SysRole 角色表
type SysRole struct {
	Id          string `gorm:"primary_key;type:varchar(50);"`
	RoleName    string `gorm:"column:role_name;not null;comment:'角色名称'" json:"role_name"`
	RoleExplain string `gorm:"column:role_explain;comment:'角色描述'" json:"role_explain"`

	SysMenus   []SysMenu   `gorm:"many2many:sys_role_menu" json:"sys_menus"`
	SysSrvs    []SysSrv    `gorm:"many2many:sys_role_sys_srvs"  json:"sys_srvs"`
	SysWebApis []SysWebApi `gorm:"many2many:sys_role_sys_web_apis" json:"sys_web_apis"`
	//SysFiles     []SysFile     `gorm:"many2many:sys_role_sys_files"  json:"sys_files"`
	SysViewRoles []SysViewRole `gorm:"many2many:sys_role_sys_view_roles" json:"sys_view_roles"`
	SysResources []SysResource `gorm:"many2many:sys_role_sys_resources"  json:"sys_resources"`

	SysUsers []SysUser `gorm:"many2many:sys_user_sys_roles" json:"sys_users"`

	PId      string    `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Children []SysRole `gorm:"foreignKey:p_id"  json:"children"`
	Model
}

//SysDep 组织机构管理
type SysDep struct {
	Id       string           `gorm:"primary_key;type:varchar(50);"`
	Code     string           `gorm:"column:code;comment:'编码';unique" json:"code"`
	Text     string           `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	Sort     int32            `gorm:"column:sort;comment:'排序';default:100" json:"sort"`
	PId      string           `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Type     dbmodel.TreeType `gorm:"column:type;comment:'树类型'" json:"type"`
	Children []SysDep         `gorm:"foreignKey:p_id"  json:"children"`

	SysUsers []SysUser `gorm:"many2many:sys_user_sys_deps" json:"sys_users"` //组织 (多对多会自动创建一个表，或者用foreignkey，手动创建中间表)
	Model
}

//SysViewRole 视图权限
type SysViewRole struct {
	Id    string `gorm:"primary_key;type:varchar(50);"`
	Code  string `gorm:"column:code;not null;comment:'资源编码'" json:"code"`
	Value string `gorm:"column:value;not null;comment:'资源名称'" json:"value"`
	Model
}

//SysResources 资源表
type SysResource struct {
	Id           string `gorm:"primary_key;type:varchar(50);"`
	Code         string `gorm:"column:code;not null;comment:'资源编码'" json:"code"`
	Value        string `gorm:"column:value;not null;comment:'资源名称'" json:"value"`
	ValueExplain string `gorm:"column:value_explain;not null;comment:'说明'"  json:"value_explain"`

	Model
}

//SysValue 常量
type SysValue struct {
	Id           string            `gorm:"primary_key;type:varchar(50);"`
	Code         string            `gorm:"column:code;not null;unique;comment:'标题'" json:"code"`
	Value        string            `gorm:"column:value;not null;comment:'名称'" json:"value"`
	ValueExplain string            `gorm:"column:value_explain;not null;comment:'说明'"  json:"value_explain"`
	ValueType    dbmodel.ValueType `gorm:"column:value_type;not null;comment:'分类'" json:"value_type"`

	//SysPositions    []SysPosition `gorm:"many2many:sys_position_sys_value"  json:"sys_positions"`
	Model
}

//SysTree 树管理
type SysTree struct {
	Id       string           `gorm:"primary_key;type:varchar(50);"`
	Code     string           `gorm:"column:code;comment:'编码';unique" json:"code"`
	Text     string           `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	Sort     int32            `gorm:"column:sort;comment:'排序';default:100" json:"sort"`
	PId      string           `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	TreeType dbmodel.TreeType `gorm:"column:tree_type;comment:'树类型'" json:"tree_type"`
	Children []SysTree        `gorm:"foreignKey:p_id"  json:"children"`
	Model
}

//SysArea 地址树管理
type SysArea struct {
	Id         string    `gorm:"primary_key;type:varchar(50);"`
	Code       string    `gorm:"column:code;comment:'编码'" json:"code"`
	Text       string    `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	Sort       int32     `gorm:"column:sort;comment:'排序';default:100" json:"sort"`
	ShortName  string    `gorm:"column:short_name;comment:'简称'" json:"short_name"`
	CityCode   string    `gorm:"column:city_code;comment:'区号'"  json:"city_code"`
	ZipCode    string    `gorm:"column:zip_code;comment:'邮编'" json:"zip_code"`
	MergerName string    `gorm:"column:merger_name;comment:'完整地址'" json:"merger_name"`
	PId        string    `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Lng        float64   `gorm:"column:lng;comment:'横坐标'"  json:"lng"`
	Lat        float64   `gorm:"column:lat;comment:'纵坐标'" json:"lat"`
	Pinyin     string    `gorm:"column:pinyin;comment:'拼音'" json:"pinyin"`
	Children   []SysArea `gorm:"foreignKey:p_id"  json:"children"`
	Model
}

//API接口管理
type SysWebApi struct {
	Id         string `gorm:"primary_key;type:varchar(50);"`
	Service    string `gorm:"column:service;comment:'api接口服务'" json:"service"`
	Method     string `gorm:"column:method;comment:'api接口名称'" json:"method"`
	ApiExplain string `gorm:"column:api_explain;comment:'说明'" json:"api_explain"`
	Model
}

//Srv接口管理
type SysSrv struct {
	Id         string `gorm:"primary_key;type:varchar(50);"`
	Service    string `gorm:"column:service;comment:'api接口服务'" json:"service"`
	Method     string `gorm:"column:method;comment:'api接口名称'" json:"method"`
	SrvExplain string `gorm:"column:srv_explain;comment:'说明'" json:"srv_explain"`
	Model
}

// SysMenu 菜单管理
type SysMenu struct {
	Id               string    `gorm:"primary_key;type:varchar(50);"`
	Text             string    `gorm:"column:text;not null;comment:'树名称'" json:"text"`
	I18N             string    `gorm:"column:i18n;not null;comment:'i18n主键（支持HTML）'" json:"i_18_n"`
	Group            bool      `gorm:"column:group;tinyint(1);default:0;comment:'	是否显示分组名，指示例中的【主导航】字样'" json:"group,omitempty"`
	Link             string    `gorm:"column:link;comment:'路由，link、externalLink 二选其一'" json:"link"`
	ExternalLink     string    `gorm:"column:external_link;comment:'外部链接'" json:"external_link"`
	Target           string    `gorm:"column:target;comment:'链接 target_blank,_self,_parent,_top'" json:"target"`
	Sort             int       `gorm:"column:sort;comment:'排序';default:100" json:"sort"`
	Badge            int       `gorm:"column:badge;comment:'标签数量'" json:"badge"`
	BadgeDoc         string    `gorm:"column:badge_doc;comment:'标签文字'" json:"badgeDot"`
	BadgeStatus      string    `gorm:"column:badge_status;comment:'徽标 Badge 颜色'" json:"badgeStatus"`
	Hide             bool      `gorm:"column:hide;tinyint(1);default:0;comment:'是否掩藏'" json:"hide"`
	Disabled         bool      `gorm:"column:disabled;tinyint(1);default:0;comment:'是否禁用'" json:"disabled"`
	HideInBreadcrumb bool      `gorm:"column:hideInBreadcrumb;tinyint(1);default:0;comment:'隐藏面包屑，指 page-header 组件的自动生成面包屑时有效'" json:"hideInBreadcrumb"`
	ACL              string    `gorm:"column:acl;comment:'ACL配置若导入 @delon/acl 时自动有效'" json:"acl"`
	Shortcut         bool      `gorm:"column:shortcut;tinyint(1);default:0;comment:'是否快捷菜单项'" json:"shortcut"`
	ShortcutRoot     bool      `gorm:"column:shortcut_root;tinyint(1);default:0;comment:'是否禁用'" json:"shortcutRoot"`
	Reuse            bool      `gorm:"column:reuse;tinyint(1);default:0;comment:'是否允许复用，需配合 reuse-tab 组件'" json:"reuse"`
	Icon             string    `gorm:"column:icon;default:'anticon-dashboard';comment:'图标图标'" json:"icon"`
	PId              string    `gorm:"column:p_id;comment:'上级id，为0表示没有上级'" json:"p_id"`
	Children         []SysMenu `gorm:"foreignKey:p_id"  json:"children,omitempty"`
	Model
}
