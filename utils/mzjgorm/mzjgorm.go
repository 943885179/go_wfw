package mzjgorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"

	//"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	//"gorm.io/datatypes" 让其支持json类型，sqlite, mysql, postgres支持
	//"gorm.io/hints"优化器索引注释等功能支持
	//"gorm.io/plugin/prometheus"使用Prometheus收集数据库状态
	//_ "gorm.io/driver/bigquery/driver" 配合"database/sql"使用
)

type DbType int

var dbList = map[string]*gorm.DB{} //数据库集合

const (
	DbType_Mysql      DbType = iota //Mysql
	DbType_Postgres                 //Postgres
	DbType_SqlServer                //SqlServer
	DbType_Sqlite                   //Sqlite
	DbType_Clickhouse               //云数据库ClickHouse
)

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Sqlite struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Sqlserver struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	Logger       bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Postgresql struct {
	Host                 string `mapstructure:"host" json:"host" yaml:"host"`
	Port                 string `mapstructure:"port" json:"port" yaml:"port"`
	Config               string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname               string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username             string `mapstructure:"username" json:"username" yaml:"username"`
	Password             string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns         int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns         int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	PreferSimpleProtocol bool   `mapstructure:"prefer-simple-protocol" json:"preferSimpleProtocol" yaml:"prefer-simple-protocol"`
	Logger               bool   `mapstructure:"logger" json:"logger" yaml:"logger"`
}

//DbConfig 数据库配置
type DbConfig struct {
	DbType     DbType     `json:"dbType"`      //驱动类型（这个是我自定义的） driverName
	Server     string     `json:"server"`      //服务器
	Port       int        `json:"port"`        //端口
	User       string     `json:"user"`        //用户名
	Password   string     `json:"password"`    //密码
	Database   string     `json:"database"`    //数据库
	Source     string     `json:"source"`      //完整连接（优先读取）
	IsDebug    bool       `json:"isDebug"`     //是否为调试模式
	ResolverDb ResolverDb `json:"resolver_db"` //读写分离设置
}
type ResolverDb struct {
	DbDsn    string   `json:"db_dsn"`   //主要连接数据库
	Sources  []string `json:"sources"`  //Sources
	Replicas []string `json:"replicas"` //replicas
}

/**
 * @Author mzj
 * @Description 创建一个db
 * @Date 上午 10:02 2020/10/29 0029
 * @Param
 * @return
 **/
func (c *DbConfig) New() (db *gorm.DB) {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	gormConfig := &gorm.Config{
		Logger:                                   newLogger, // logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,      //外键约束,默认开启
		DisableAutomaticPing:                     true,      //在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性
		//SkipDefaultTransaction:                   true, //禁用默认事务
		//DryRun: true, //缓存预编译语句
		/*NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},*/
		PrepareStmt: true, //带 PreparedStmt 的 SQL 生成器 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率

	}
	switch c.DbType {
	case DbType_Mysql:
		c.Source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
		//db, err = gorm.Open(mysql.Open(c.Source), gormConfig)
		for s, i := range dbList {
			if c.Source == s {
				return i
			}
		}
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       c.Source, // DSN data source name
			DefaultStringSize:         191,      // string 类型字段的默认长度
			DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,    // 根据版本自动配置,
		}), gormConfig)
		break
	case DbType_SqlServer:
		c.Source = fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s", c.Server, c.Port, c.Database, c.User, c.Password)
		db, err = gorm.Open(sqlserver.Open(c.Source), gormConfig)
		break
	//case ORACLE: //不支持
	//	c.Source = fmt.Sprintf("%s/%s@%s:%d/%s", c.User, c.Password, c.Server, c.Port, c.Database)
	case DbType_Postgres:
		//c.Source = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", c.User, c.Password, c.Server, c.Database)
		c.Source = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d  sslmode=disable TimeZone=Asia/Shanghai", c.Server, c.User, c.Password, c.Database, c.Port) //https://godoc.org/github.com/lib/pq 参数请参阅
		db, err = gorm.Open(postgres.Open(c.Source), gormConfig)
		break
		//这里这样导入sqllite会报错
		/*case DbType_Sqlite:
		c.Source = "gorm.db" //设置默认的db
		db, err = gorm.Open(sqlite.Open(c.Source), gormConfig)
		break*/
	/*case DbType_Clickhouse:
	c.Source = fmt.Sprintf("tcp://%s:%d?debug=true", c.Server, c.Port) //设置默认的db
	db, err = gorm.Open(clickhouse.Open(c.Source), gormConfig)*/
	default:
		log.Fatal("暂时没有设置该驱动")
	}
	if err != nil {
		log.Fatal(err)
	}
	//db = db.Session(&gorm.Session{PrepareStmt: true})
	// 下面这段是否有用不太清楚，不用了吧，v2找不到设置的连接数量了
	sqldb, _ := db.DB()
	sqldb.SetMaxOpenConns(100)            //设置数据库连接池最大连接数
	sqldb.SetMaxIdleConns(20)             //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqldb.SetConnMaxLifetime(time.Minute) //过期时间
	//defer sqldb.Close() 不能关闭，关闭后就查不到数据了
	if c.IsDebug {
		db = db.Debug()
	}
	dbList[c.Source] = db
	return db
}

func (c *DbConfig) IsErrRecordNotFound(err error) bool {
	return gorm.ErrRecordNotFound == err
}

func (c *DbConfig) Resolver() (db *gorm.DB) { //第一个为主写,其他为辅读
	var err error
	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true, //外键约束,默认开启
	}
	switch c.DbType {
	case DbType_Mysql:
		db, err = gorm.Open(mysql.Open(c.ResolverDb.DbDsn), gormConfig)
	/*case DbType_SqlServer:
		db, err = gorm.Open(sqlserver.Open(c.ResolverDb.DbDsn), gormConfig)
	//case ORACLE: //不支持
	//	c.Source = fmt.Sprintf("%s/%s@%s:%d/%s", c.User, c.Password, c.Server, c.Port, c.Database)
	case DbType_Sqlite:
		c.Source = "gorm.db" //设置默认的db
		db, err = gorm.Open(sqlite.Open(c.ResolverDb.DbDsn), gormConfig)
	case DbType_Postgres:
		db, err = gorm.Open(postgres.Open(c.ResolverDb.DbDsn), gormConfig)
	case DbType_Clickhouse:
		db, err = gorm.Open(clickhouse.Open(c.ResolverDb.DbDsn), gormConfig)*/
	default:
		log.Fatal("暂时没有设置该驱动")
	}
	if err != nil {
		log.Fatal(err)
	}
	/*db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  c.ResolverDb.Sources,
		Replicas: c.ResolverDb.Replicas,
		Policy:   dbresolver.DBResolver{}, //负载均衡策略 RandomPolicy:随机
	}).Register(dbresolver.Config{
		Replicas: []gorm.Dialector{},
	}))*/
	return db
}

/*
type IDb interface {
	Open()
}
type myMysql struct{}

func (*myMysql) Open() {
	//dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	//	DefaultStringSize: 256, // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
	//	DisableDatetimePrecision: true, // disable datetime precision support, which not supported before MySQL 5.6
	//	DontSupportRenameIndex: true, // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
	//	DontSupportRenameColumn: true, // use change when rename column, rename rename not supported before MySQL 8, MariaDB
	//	SkipInitializeWithVersion: false, // smart configure based on used version
	//}), &gorm.Config{})
	gorm.Open(mysql.Open(""), &gorm.Config{})
}

type mySqlite struct{}

func (*mySqlite) Open() {
	//dsn="gorm.db"
	gorm.Open(sqlite.Open(""), &gorm.Config{})
}

type myPostgres struct{}

func (*myPostgres) Open() {
	//dsn := "user=gorm password=gorm DB.name=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN:                  "user=gorm password=gorm DB.name=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
	//	PreferSimpleProtocol: true,                                                                                    // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	//}), &gorm.Config{})
	gorm.Open(postgres.Open(""), &gorm.Config{})
}

type mySqlserver struct{}

func (*mySqlserver) Open() {
	//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	gorm.Open(sqlserver.Open(""), &gorm.Config{})
}

type myClickhouse struct{}

func (*myClickhouse) Open() {
	//dsn := "tcp://localhost:9000?debug=true"
	_, _ = gorm.Open(clickhouse.Open("db1_dsn"), &gorm.Config{})
}
*/
