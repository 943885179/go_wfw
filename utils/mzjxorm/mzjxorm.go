package mzjxorm

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //mysql 数据库连接
	"github.com/go-xorm/xorm"
)

//DriverType 数据库驱动
type DriverType int

const (
	mysql            DriverType = iota //0
	sqlserver                          //1
	sqlserverwindows                   //2
	mssql                              //3
	oracle                             //3
	sqllite                            //5
	postgresql                         //6
	postgresql1                        //7
	postgresql2                        //8
)

func (d DriverType) String() string {
	switch d {
	case mysql:
		return "mysql" //go get github.com/go-sql-driver/mysql
	case sqlserver: //sql server使用adodb驱动
		return "adodb" //go get github.com/mattn/go-adodb（gorm不能识别该驱动,弃用）
	case sqlserverwindows:
		return "adodb" //go get github.com/mattn/go-adodb（gorm不能识别该驱动,弃用）
	case mssql:
		return "mssql" //go get github.com/denisenkom/go-mssqldb 这个驱动可以连接sqlserver2019,但是好像不能够连接sql server 2008一下
	case oracle:
		return "oci8" //go get github.com/mattn/go-oci8
	case sqllite:
		return "sqlite3" //go get github.com/mattn/go-sqlite3（支持database/sql所以我就使用了它） github.com/feyeleanor/gosqlite3（不支持database/sql）github.com/phf/go-sqlite3（不支持database/sql）
	case postgresql:
		return "postgres" //go get github.com/bmizerany/pq
	case postgresql1:
		return "postgres" //go get github.com/jbarham/gopgsqldriver
	case postgresql2:
		return "postgres" //go get github.com/lxn/go-pgsql
	default:
		return "暂时没有设置该驱动"
	}
}

//DbConfig 数据库配置
type DbConfig struct {
	DriverType DriverType `json:"driverType"` //驱动类型（这个是我自定义的）
	Server     string     `json:"server"`     //服务器
	Port       int        `json:"port"`       //端口
	User       string     `json:"user"`       //用户名
	Password   string     `json:"password"`   //密码
	Database   string     `json:"database"`   //数据库
	Source     string     `json:"source"`     //完整连接（优先读取）
	Sources    []string   `json:"sources"`    //完整连接（优先读取）多个，读写分离
	IsDebug    bool       `json:"isDebug"`    //是否为调试模式
	//DB         *sql.DB    `json:"-"`          //db
}

func (c *DbConfig) init() {
	if c.Source != "" { //如果有数据库链接了那么直接使用
		return
	}
	switch c.DriverType {
	case mysql:
		c.Source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	case sqlserver:
		if c.Port == 1433 {
			c.Source = fmt.Sprintf("Provider=SQLOLEDB;Data Source=%s;Initial Catalog=%s;user id=%s;password=%s;Connection Timeout=3600;Connect Timeout=3600;", c.Server, c.Database, c.User, c.Password)
		} else {
			c.Source = fmt.Sprintf("Provider=SQLOLEDB;Data Source=%s,%d;Initial Catalog=%s;user id=%s;password=%s;Connection Timeout=3600;Connect Timeout=3600;", c.Server, c.Port, c.Database, c.User, c.Password)
		}
	case sqlserverwindows:
		c.Source = fmt.Sprintf("Provider=SQLOLEDB;Data Source=%s;integrated security=SSPI;Initial Catalog=%s;", c.Server, c.Database)
	case mssql:
		c.Source = fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", c.Server, c.Port, c.Database, c.User, c.Password)
	case oracle:
		c.Source = fmt.Sprintf("%s/%s@%s:%d/%s", c.User, c.Password, c.Server, c.Port, c.Database)
	case sqllite:
		c.Source = "./foo.db" //设置默认的db
	case postgresql:
		//c.Source = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", c.User, c.Password, c.Server, c.Database)
		c.Source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=verify-full", c.User, c.Password, c.Server, c.Port, c.Database) //https://godoc.org/github.com/lib/pq 参数请参阅
	case postgresql1:
		c.Source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=verify-full", c.User, c.Password, c.Server, c.Port, c.Database) //https://godoc.org/github.com/lib/pq 参数请参阅
	case postgresql2:
		c.Source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=verify-full", c.User, c.Password, c.Server, c.Port, c.Database) //https://godoc.org/github.com/lib/pq 参数请参阅
	default:
		log.Fatal("暂时没有设置该驱动")
	}
}
func NewDbConfig(config DbConfig) *DbConfig {
	c := DbConfig{
		DriverType: DriverType(config.DriverType),
		Server:     config.Server,
		Port:       config.Port,
		User:       config.User,
		Password:   config.Password,
		Database:   config.Database,
		Source:     config.Source,
		IsDebug:    config.IsDebug,
	}
	return &c
}
func NewDbConfigs(configs []DbConfig) *DbConfig {
	result := DbConfig{}
	for _, config := range configs {
		c := DbConfig{
			DriverType: DriverType(config.DriverType),
			Server:     config.Server,
			Port:       config.Port,
			User:       config.User,
			Password:   config.Password,
			Database:   config.Database,
			Source:     config.Source,
			IsDebug:    config.IsDebug,
		}
		c.init()
		result.IsDebug = config.IsDebug
		result.Sources = append(result.Sources, c.Source)
	}
	return &result
}

//QshXormNew 初始化一个新的Engine
func QshXormNew(driverName, dataSourceName string) (engine *xorm.Engine, err error) {
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return engine, err
	}
	if err = engine.Ping(); err != nil {
		return engine, err
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.Logger().ShowSQL()
	return engine, err
}

//QshXormNewEngine 新的engine
func (c *DbConfig) QshXormNewEngine() (engine *xorm.Engine, err error) {
	c.init()
	engine, err = xorm.NewEngine(c.DriverType.String(), c.Source)
	if err != nil {
		return engine, err
	}
	if err = engine.Ping(); err != nil {
		return engine, err
	}
	if c.IsDebug {
		engine.ShowExecTime(true)
		engine.ShowSQL(true)
	}
	engine.Logger().ShowSQL()
	return engine, err
}

//QshXormNewGroup 初始化一个新的Engine
func (c *DbConfig) QshXormNewGroup() (eg *xorm.EngineGroup, err error) {
	//默认随机策略
	//eg, err = xorm.NewEngineGroup(driverName, dataSourceNames)
	//随机策略
	//eg, err = xorm.NewEngineGroup(driverName, dataSourceNames, xorm.RandomPolicy())
	//权重随机策略（此时设置第一个的权重为2，第二个的权重为3）
	//eg, err = xorm.NewEngineGroup(driverName, dataSourceNames,xorm.WeightRandomPolicy([]int{2,3}))
	//轮询
	//eg, err = xorm.NewEngineGroup(driverName, dataSourceNames, xorm.RoundRobinPolicy())
	//轮询权重
	//eg, err = xorm.NewEngineGroup(driverName, dataSourceNames, xorm.WeightRandomPolicy([]int{2,3}))
	//最小连接数访问负载策略
	eg, err = xorm.NewEngineGroup(c.DriverType.String(), c.Sources, xorm.LeastConnPolicy())
	//自定义策略
	if err != nil {
		return eg, err
	}
	if err = eg.Ping(); err != nil {
		return eg, err
	}
	if c.IsDebug {
		eg.ShowSQL(true)      //显示sql
		eg.ShowExecTime(true) //显示sql执行时间
	}
	//eg.Logger().SetLevel(core.LOG_DEBUG) //控制台打印
	//eg.SetLogger()                       //日志保存
	//全局缓存(缓存的方式存放内存中，缓存struct的记录为1000条)
	/*//eg.SetDefaultCacher(xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000))
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	utype := models.SysUserType{}
	eg.MapCacher(&utype, cacher)
	eg.Exec("......")
	eg.ClearCache(new(models.SysUserType))//修改后清楚缓存
	*/
	return eg, err
}
