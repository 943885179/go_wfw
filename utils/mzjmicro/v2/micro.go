package v2

import (
	"context"
	"fmt"
	"github.com/micro/go-plugins/logger/logrus/v2"
	logruss "github.com/sirupsen/logrus"
	"qshapi/utils/mzjlog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/registry/nacos/v2"

	"github.com/micro/go-plugins/broker/kafka/v2"
)

//Service 服务
type Service struct {
	Ip       string `json:"ip"`       //Ip地址
	Port     int    `json:"port"`     //端口
	Version  string `json:"version"`  //版本
	Name     string `json:"name"`     //服务名称
	Describe string `json:"describe"` //叙述
	Etcd     string `json:"etcd"`     //注入的etcd地址
	NacOs    string `json:"nac_os"`
	Consul   string `json:"consul"`
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
	interal         time.Duration = time.Second * 10                //重新注册时间
	ttl             time.Duration = time.Second * time.Duration(30) //服务过期时间
	logpath         string        = "logs"                          //日志存储位置
	logfile         string        = "wfw.log"                       //日志文件
	logMaxAge       time.Duration = time.Hour * 24 * 3              //日志自动回收时间
	logRotationTime time.Duration = time.Hour                       //日志切割时间
)

func init() {
	hooks := logruss.LevelHooks{}
	hooks.Add(mzjlog.GetHook(logpath, logfile, logMaxAge, logRotationTime))
	logger.DefaultLogger = logrus.NewLogger(logrus.WithJSONFormatter(&logruss.JSONFormatter{}), logrus.WithLevelHooks(hooks)).Fields(map[string]interface{}{
		"版本": "v1",
		//"k2": 123456,
	})

}

func (s *Service) NewGinWeb(g *gin.Engine) web.Service {
	//regs:=consul.NewRegistry(registry.Addrs(s.Etcd))
	sv := web.NewService(
		web.Name(s.Name),
		web.Version(s.Version),
		//web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)),
		web.RegisterInterval(interal), //间隔多久再次注册服务
		web.RegisterTTL(ttl),          //注册服务的过期时间
	//web.Registry(reg),
	)
	if s.Port > 0 { //设置了特定的端口和地址
		sv.Init(web.Address(fmt.Sprintf("%s:%d", s.Ip, s.Port)))
	}
	sv.Handle("/", g)
	if len(s.Etcd) > 0 {
		reg := etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d", s.Name, s.Ip, s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	return sv
}
func (s *Service) NewWeb() web.Service {
	sv := web.NewService(
		web.Name(s.Name),
		web.Version(s.Version),
		//web.Address(fmt.Sprintf("%s:%d",s.Ip,s.Port)),
		web.RegisterInterval(interal), //间隔多久再次注册服务
		web.RegisterTTL(ttl),          //注册服务的过期时间
	)
	if s.Port > 0 { //设置了特定的端口和地址
		sv.Init(web.Address(fmt.Sprintf("%s:%d", s.Ip, s.Port)))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d", s.Name, s.Ip, s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	if len(s.Etcd) > 0 {
		reg := etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	if len(s.NacOs) > 0 {
		reg := nacos.NewRegistry(registry.Addrs(s.NacOs))
		sv.Init(web.Registry(reg))
	}
	if len(s.Consul) > 0 {
		reg := consul.NewRegistry(registry.Addrs(s.Consul))
		sv.Init(web.Registry(reg))
	}
	return sv
}

//logWrapper 日志记录
func logWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc { //请求服务前先记录日志
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		str := fmt.Sprintf("服务：%s\t全称：%s \t方法：%s \t头部：%s \t请求体：%v \n", req.Endpoint(), req.Service(), req.Method(), req.Header(), req.Body())
		//logger.DefaultLogger = logrus.NewLogger(logger.WithOutput(os.Stdout))
		//logger.DefaultLogger = logrus.NewLogger(logger.WithLevel(logger.DebugLevel))
		logger.Logf(logger.InfoLevel, str)
		//todo: 记录日志
		return handlerFunc(ctx, req, rsp)
	}
}
func roleWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc { //请求服务前先记录日志
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//fmt.Printf("[%v] 服务请求:服务：%s\t全称：%s \t方法：%s \t头部：%s \t请求体：%s \n", time.Now(), req.Endpoint(), req.Service(), req.Method(), req.Header(), req.Body())
		//todo: 判断权限是否足够
		//return errors.New("权限不足")
		return handlerFunc(ctx, req, rsp)
		//return errors.New("我是错误")
	}
}
func (s *Service) NewSrv() micro.Service {
	//reg:=consul.NewRegistry(registry.Addrs(s.Etcd))
	sv := micro.NewService(
		micro.Name(s.Name),
		micro.Version(s.Version),

		micro.RegisterInterval(interal), //间隔多久再次注册服务
		micro.RegisterTTL(ttl),          //注册服务的过期时间
		micro.Transport(grpc.NewTransport()),
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
		micro.WrapHandler(logWrapper),
		micro.WrapHandler(roleWrapper),
	)
	if s.Port > 0 { //设置了特定的端口和地址
		sv.Init(micro.Address(fmt.Sprintf("%s:%d", s.Ip, s.Port)))
	}
	// 服务注册
	if len(s.Etcd) > 0 {
		reg := etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(micro.Registry(reg))
	}
	if len(s.NacOs) > 0 {
		reg := nacos.NewRegistry(registry.Addrs(s.NacOs))
		sv.Init(micro.Registry(reg))

	}
	if len(s.Consul) > 0 {
		reg := consul.NewRegistry(registry.Addrs(s.Consul))
		sv.Init(micro.Registry(reg))
	}
	//Mq
	b := kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"127.0.0.1:9092"}
	})
	b.Init()
	b.Connect()
	sv.Init(micro.Broker(b))

	sv.Init(micro.AfterStart(func() error {
		fmt.Printf("启动服务成功:%s,地址为:%s:%d", s.Name, s.Ip, s.Port)
		fmt.Println(s.Describe)
		return nil
	}))
	sv.Init()
	return sv
}

func (s *Service) KafkaMQ() {
	fmt.Println("消息中间件")
	b := kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"127.0.0.1:9092"}
	})
	b.Init()
	b.Connect()
	micro.Broker(b)
	broker.Publish("go.micro.topic.test", &broker.Message{
		Header: map[string]string{
			"AAA":   "BBBBB",
			"CCCCC": "DDDDDD",
		},
		Body: []byte("消息内容"),
	})
	fmt.Println("消息中间件接受")
	broker.Subscribe("go.micro.topic.test", func(event broker.Event) error {
		fmt.Println("啊实打实的撒")
		fmt.Println(event.Topic())
		x := event.Message()
		fmt.Print(x.Header)
		fmt.Print(string(x.Body))

		return nil
	})
}

/*
func (s *Service) NewGrpcSrc() service.Service {
	sv:=grpc.NewService()
	sv.Init()
	return sv
}*/
func (s *Service) NewRoundWeb() web.Service {
	//reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
	sv := web.NewService(
		web.RegisterInterval(interal), //间隔多久再次注册服务
		web.RegisterTTL(ttl),          //注册服务的过期时间
	//web.Registry(reg),
	)
	if len(s.Etcd) > 0 {
		reg := etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(web.Registry(reg))
	}
	if len(s.NacOs) > 0 {
		reg := nacos.NewRegistry(registry.Addrs(s.NacOs))
		sv.Init(web.Registry(reg))
	}
	if len(s.Consul) > 0 {
		reg := consul.NewRegistry(registry.Addrs(s.Consul))
		sv.Init(web.Registry(reg))
	}
	sv.Init(web.AfterStart(func() error {
		fmt.Printf("启动服务成功:%v", sv.Options())
		fmt.Println(s.Describe)
		return nil
	}))
	return sv
}
func (s *Service) NewRoundSrv() micro.Service {
	//reg:=etcd.NewRegistry(registry.Addrs(s.Etcd))
	sv := micro.NewService(
		micro.RegisterInterval(interal), //间隔多久再次注册服务
		micro.RegisterTTL(ttl),          //注册服务的过期时间
	//micro.Registry(reg),
	)
	if len(s.Etcd) > 0 {
		reg := etcd.NewRegistry(registry.Addrs(s.Etcd))
		sv.Init(micro.Registry(reg))
	}
	if len(s.NacOs) > 0 {
		reg := nacos.NewRegistry(registry.Addrs(s.NacOs))
		sv.Init(micro.Registry(reg))
	}
	if len(s.Consul) > 0 {
		reg := consul.NewRegistry(registry.Addrs(s.Consul))
		sv.Init(micro.Registry(reg))
	}
	sv.Init(micro.AfterStart(func() error {
		fmt.Printf("启动服务成功:%v", sv.Options())
		fmt.Println(s.Describe)
		return nil
	}))
	return sv
}
