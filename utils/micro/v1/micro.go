package v1

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/service/grpc"
)

func NewSrv() micro.Service {
	//consulReg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	etcdReg := etcd.NewRegistry(registry.Addrs("106.12.72.181:23791"))
	myservice := grpc.NewService( //原来是micro.NewService还支持http等其他访问方式，但是grpc这种方法只支持grpc访问，所以需要创建网关让其支持http访问
		micro.Name("api.xiahualou.com.test"),
		micro.Address(":8001"),
		micro.Registry(etcdReg),
		//micro.Registry(consulReg),
	)
	return myservice
}
