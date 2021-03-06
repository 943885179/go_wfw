package mzjdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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

//New 创建新得gorm
func (c *DbConfig) New() *sql.DB {
	c.init()
	if c.IsDebug {
		fmt.Println(c.DriverType.String(), c.Source)
	}
	db, err := sql.Open(c.DriverType.String(), c.Source)
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	return db
}

func FindStr(db *sql.DB, sqlStr string, args ...interface{}) (string, error) {
	maps, err := QueryMap(db, sqlStr, args)
	if err != nil || len(maps) == 0 {
		return "", err
	}
	jsonData, err := json.Marshal(maps)
	if err != nil {
		log.Fatal(fmt.Sprintf("Map转错误:%s", err.Error()))
	}
	return string(jsonData), nil
}

func Find(db *sql.DB, sqlStr string, data interface{}, args ...interface{}) error {
	str, err := FindStr(db, sqlStr, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), data)
}

func FirstStr(db *sql.DB, sqlStr string, args ...interface{}) (string, error) {
	maps, err := QueryMap(db, sqlStr, args)
	if err != nil || len(maps) == 0 {
		return "", err
	}
	jsonData, err := json.Marshal(maps[0])
	if err != nil {
		log.Fatal(fmt.Sprintf("Map转错误:%s", err.Error()))
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
func First(db *sql.DB, sqlStr string, data interface{}, args ...interface{}) error {
	str, err := FirstStr(db, sqlStr, args...)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), data)
}
func QueryMap(db *sql.DB, sqlStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlStr)
	if err != nil {
		//log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		//log.Fatal(err.Error())
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	//遍历每一行
	for rows.Next() {
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				//TODO:编码需要确认
				v = val
				//先确定编码，如果编码是gb2312使用下面的代码
				/*c, ok := val.(string)
				if ok {
					//v, _ = iconv.ConvertString(c, "utf-8", "GB2312")
					v, _ = iconv.ConvertString(c, "GB2312", "utf-8")
				} else {
					v = val
				}*/
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

//QueryOne 只能查询一条数据的一个值
func QueryOne(db *sql.DB, sqlStr string, data interface{}) error {
	defer db.Close()
	var dataCopy interface{}
	fmt.Println(db)
	defer db.Close()
	row := db.QueryRow(sqlStr)
	err := row.Scan(&dataCopy)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return err
	}
	if dataCopy != nil {
		err = db.QueryRow(sqlStr).Scan(data)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	return nil
}
func main() {
	c := DbConfig{0, "127.0.0.1", 3306, "root", "123456", "test", "", true}
	db := c.New()
	defer db.Close()
	result, err := db.Exec("insert into user(name)values('weixiao1')")
	fmt.Println(result, err)
}
