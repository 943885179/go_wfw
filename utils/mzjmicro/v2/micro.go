package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
	//"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/web"
	"time"
)
//Service 服务
type Service struct {
	Ip string `json:"ip"` //Ip地址
	Port int        `json:"port"`//端口
	Version string  `json:"version"`//版本
	Name string     `json:"name"`//服务名称
	Describe string `json:"describe"`//叙述
	Etcd string `json:"etcd"`//注入的etcd地址
}
/*
func NewService(sv models.Service)*Service{
	s:=Service{
		Ip:sv.Ip,
		Port:sv.Port,
		Version	:sv.Version,
		Name:sv.Name,
		Describe:sv.Describe,
		Etcd:sv.Etcd,
	}
	return &s
}*/

const (
	interal time.Duration =time.Second*10 //重新注册时间
	ttl time.Duration=time.Second*time.Duration(30)//服务过期时间
)

func (s *Service) NewGinWeb(g *gin.Engine) web.Service {
	//regs:=consul.NewRegistry(registry.Addrs(s.Etcd))
	sv:=web.NewService(
		web.Name(s.Name),
		web.Version(s.Version),
		//web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)),
		web.RegisterInterval(interal),//间隔多久再次注册服务
		web.RegisterTTL(ttl),//注册服务的过期时间
		//web.Registry(reg),
		)
	if s.Port>0  {//设置了特定的端口和地址
		sv.Init(web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)))
	}
	sv.Handle("/",g)
	if len(s.Etcd)>0 {
		reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d",s.Name,s.Ip,s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	return sv
}
func (s *Service) NewWeb() web.Service {
	sv:=web.NewService(
		web.Name(s.Name),
		web.Version(s.Version),
		//web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)),
		web.RegisterInterval(interal),//间隔多久再次注册服务
		web.RegisterTTL(ttl),//注册服务的过期时间
		)
	if s.Port>0 {//设置了特定的端口和地址
		sv.Init(web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d",s.Name,s.Ip,s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	if len(s.Etcd)>0 {
		reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	return sv
}
func (s *Service) NewSrv() micro.Service  {
	//reg:=consul.NewRegistry(registry.Addrs(s.Etcd))
	sv:=micro.NewService(
		micro.Name(s.Name),
		micro.Version(s.Version),

		micro.RegisterInterval(interal),//间隔多久再次注册服务
		micro.RegisterTTL(ttl),//注册服务的过期时间
		micro.Transport(grpc.NewTransport()),
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
		)
	if s.Port>0 {//设置了特定的端口和地址
		sv.Init(micro.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)))
	}
	if len(s.Etcd)>0 {
		reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(micro.Registry(reg))
	}
	sv.Init(micro.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d",s.Name,s.Ip,s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	sv.Init()
	return sv
}

/*
func (s *Service) NewGrpcSrc() service.Service {
	sv:=grpc.NewService()
	sv.Init()
	return sv
}*/
func (s *Service) NewRoundWeb() web.Service  {
	//reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
	sv:=web.NewService(
		web.RegisterInterval(interal),//间隔多久再次注册服务
		web.RegisterTTL(ttl),//注册服务的过期时间
		//web.Registry(reg),
		)
	if len(s.Etcd)>0 {
		reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%v",sv.Options())
		fmt.Println(s.Describe)
		return nil
	}))
	return  sv
}
func (s *Service) NewRoundSrv() micro.Service {
	//reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
	sv:=micro.NewService(
		micro.RegisterInterval(interal),//间隔多久再次注册服务
		micro.RegisterTTL(ttl),//注册服务的过期时间
		//micro.Registry(reg),
		)
	if len(s.Etcd)>0 {
		reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(micro.Registry(reg))
	}
	sv.Init(micro.AfterStart(func() error {
		fmt.Printf("启动服务成功:%v",sv.Options())
		fmt.Println(s.Describe)
		return nil
	}))
	return sv
}