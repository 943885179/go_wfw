package v3

import (
	"time"
	//"github.com/micro/go-micro/v3"
)
//Service 服务
type Service struct {
	Ip string //Ip地址
	Port int //端口
	Version string //版本
	Name string //服务名称
	Describe string //叙述
	Etcd string //注入的etcd地址
}

const (
	interal time.Duration =time.Second*10 //重新注册时间
	ttl time.Duration=time.Second*time.Duration(30)//服务过期时间
)

func test()  {

}