package init

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/micro/go-micro/v2/config"
	"qshapi/models"
	"qshapi/utils/jsonos"
	"sync"
)

var defaultPath = "config.yaml" //"E:/go/qshapi/config.configs" 如果测试用go run 启动最好指定固定路径

func InitByJson(path string, resp interface{}) error {
	if path == "" {
		return errors.New("请输入配置文件")
	}
	return jsonos.JSONReadEntity(path, resp)
}

func InitByMicroConfig(path string, resp interface{}) error { //测试发现只支持json风格 要v2，v3取消了LoadFile "github.com/micro/go-micro/v2/config"
	if path == "" {
		return errors.New("请输入配置文件")
	}
	var once sync.Once
	once.Do(func() {
		config.LoadFile(path)
	})
	return json.Unmarshal(config.Bytes(), resp)
}

//MicroConfig 读取config
func MicroConfig(resp interface{}) error { //测试发现只支持json风格 要v2，v3取消了LoadFile "github.com/micro/go-micro/v2/config"
	return json.Unmarshal(config.Bytes(), resp)
}
func Default(resp interface{}) error {
	return InitByMicroConfig(defaultPath, resp)
}

/*
func InitByMicroConfig(path string,resp interface{})error  {//测试发现只支持json风格 v3"github.com/micro/go-micro/v3/config" "github.com/micro/go-micro/v3/config/source/file"
	if path=="" {
		return errors.New("请输入配置文件")
	}
	configs, _ := config.NewConfig()
	configs.Load(file.NewSource(file.WithPath(path)))
	fmt.Println(string(configs.Bytes()))
	return jsonos.Unmarshal(configs.Bytes(),resp)
}*/

func initByIni(path string) error { //github.com/Unknwon/goconfig 是能够读取ini风格的配置文件，但是好像加注释有点问题
	//go get gopkg.in/ini.v1 这个也可以，还可以保存文件
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		return err
	}
	//读取单个值
	value, err := cfg.GetValue("mysql", "username")
	fmt.Println("读取到的值为", value)
	//读取整个区
	fmt.Println(cfg.GetSection("mysql"))
	return nil
}

func main() {
	/*configs:= models.APIConfig{}
	InitByJson("config.jsonos",&configs)
	fmt.Println(configs.DbConfig.Port)*/
	conf := models.APIConfig{}
	InitByMicroConfig("config.configs", &conf)
	fmt.Println(conf.DbConfig.Port)
	//InitByIni("test.configs")
}
