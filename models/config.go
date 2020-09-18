package models

import "time"

//APIConfig erp基础配置
type APIConfig struct {
	//单库也使用group管理
	//IsOneDb	bool `json:"isOneDb"`//是否是单库
	DbConfig   DbConfig    `json:"dbConfig"`    //数据库配置
	FilePath    string      `json:"filePath"`    //文件基础路径
	RedisConfig RedisConfig `json:"redisConfig"` //redis配置
	TxOcrAPI    TxOcrAPI    `json:"txOcrApi"`    //腾讯Orcapi
	WxPayConfig WxPayConfig `json:"wxPayConfig"` //微信支付config
	Jwt Jwt `json:"jwt"` //jwt
	Services map[string]Service `json:"services"`
}
//Service 服务
type Service struct {
	Ip string `json:"ip"` //Ip地址
	Port int        `json:"port"`//端口
	Version string  `json:"version"`//版本
	Name string     `json:"name"`//服务名称
	Describe string `json:"describe"`//叙述
	Etcd string `json:"etcd"`//注入的etcd地址
}
type Jwt struct {
	Secret string `json:"secret"`//jwt加密字段
	TimeOut    time.Duration `json:"timeOut"`    //过期时长
	Token   string `json:"token"`    //token
}
//DbConfig 数据库配置
type DbConfig struct {
	DriverType int `json:"driverType"` //驱动类型（这个是我自定义的）
	Server     string     `json:"server"`     //服务器
	Port       int     `json:"port"`       //端口
	User       string     `json:"user"`       //用户名
	Password   string     `json:"password"`   //密码
	Database   string     `json:"database"`   //数据库
	Source     string     `json:"source"`     //完整连接（优先读取）
	IsDebug    bool     `json:"isDebug"`    //是否为调试模式
}

//RedisConfig RedisConfig
type RedisConfig struct {
	Network  string `json:"network"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	IsDebug    bool     `json:"isDebug"`    //是否为调试模式
}

//TxOcrAPI 腾讯文字OrC
type TxOcrAPI struct {
	Region    string `json:"region"`
	SecretID  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"endpoint"`
	IsDebug    bool     `json:"isDebug"`    //是否为调试模式
}

//WxPayConfig 微信支付基础配置
type WxPayConfig struct {
	AppID  string `json:"appId"`  //应用id
	MchID  string `json:"mchId"`  //商户id
	APIKey string `json:"apikey"` //API密钥
	IsProd bool   `json:"isProd"` //是否是正式环境
	IsDebug    bool     `json:"isDebug"`    //是否为调试模式
}
