package mzjgorm

import (
	"fmt"
	"log"
	"qshapi/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DriverType 数据库驱动
type DriverType int

const (
	mysql      DriverType = iota //0
	mssql                        //1
	oracle                       //2
	sqllite                      //3
	postgresql                   //4
)

func (d DriverType) String() string {
	switch d {
	case mysql:
		return "mysql" //go get github.com/jinzhu/gorm/dialects/mysql
	case mssql:
		return "mssql" //go get  github.com/jinzhu/gorm/dialects/mssql
	case oracle:
		return "oci8" //go get github.com/mattn/go-oci8
	case sqllite:
		return "sqlite3" //go get github.com/jinzhu/gorm/dialects/sqlite
	case postgresql:
		return "postgres" //go get github.com/jinzhu/gorm/dialects/postgres
	default:
		return "暂时没有设置该驱动"
	}
}

//DbConfig 数据库配置
type DbConfig struct {
	DriverType DriverType `json:"driverType"` //驱动类型（这个是我自定义的）
	Server     string     `json:"server"`     //服务器
	Port       int     `json:"port"`       //端口
	User       string     `json:"user"`       //用户名
	Password   string     `json:"password"`   //密码
	Database   string     `json:"database"`   //数据库
	Source     string     `json:"source"`     //完整连接（优先读取）
	IsDebug    bool       `json:"isDebug"`    //是否为调试模式
}

func NewDbConfig(config models.DbConfig) *DbConfig {
	c:=DbConfig{
		DriverType: DriverType(config.DriverType),
		Server: config.Server,
		Port: config.Port,
		User: config.User,
		Password: config.Password,
		Database: config.Database,
		Source: config.Source,
		IsDebug: config.IsDebug,
	}
	return &c
}
func (c *DbConfig) init() {
	if c.Source != "" { //如果有数据库链接了那么直接使用
		return
	}
	switch c.DriverType {
	case mysql:
		c.Source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	case mssql:
		c.Source = fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", c.Server, c.Port, c.Database, c.User, c.Password)
	case oracle:
		c.Source = fmt.Sprintf("%s/%s@%s:%d/%s", c.User, c.Password, c.Server, c.Port, c.Database)
	case sqllite:
		c.Source = "./foo.db" //设置默认的db
	case postgresql:
		//c.Source = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", c.User, c.Password, c.Server, c.Database)
		c.Source = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=verify-full", c.User, c.Password, c.Server, c.Port, c.Database) //https://godoc.org/github.com/lib/pq 参数请参阅
	default:
		log.Fatal("暂时没有设置该驱动")
	}
	fmt.Println(c.Source)
}

//New 创建新得gorm
func (c *DbConfig) New() (db *gorm.DB) {
	c.init()
	if c.IsDebug {
		fmt.Println(c.DriverType.String(), c.Source)
	}
	db, err := gorm.Open(c.DriverType.String(), c.Source)
	if err != nil {
		log.Fatal("数据库连接失败", err.Error())
	}
	if c.IsDebug {
		db.LogMode(true) //sql跟踪
	}
	db.SingularTable(true)       //代码结构体单复数和数据库表名单复数必须对应
	db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return db
}

/*
var (
	lock     *sync.Mutex = &sync.Mutex{}
	instanct *DbConfig
)

//GetInstance 单例默认初始化DbConfig
func GetInstance() *DbConfig {
	if instanct == nil {
		lock.Lock()
		defer lock.Unlock()
		if instanct == nil {
			instanct = &DbConfig{}
			fmt.Println(11)
		}
	}
	return instanct
}*/
func main() {
	c := DbConfig{0, "127.0.0.1", 3306, "root", "123456", "test", "", true}
	db := c.New()
	defer db.Close()
	type User struct {
		ID   int64
		Name string
	}
	user := User{}
	db.First(&user)
	fmt.Println(user)
}
