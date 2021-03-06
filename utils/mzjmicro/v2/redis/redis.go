package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/redis/v2"
	"strconv"
	"time"
)

//MICRO_BROKER=redis;MICRO_BROKER_ADDRESS=127.0.0.1:6379

//在ridis cli里面监听 subscribe topic
//redis手动添加: publish topic value

func main() {
	broker := redis.NewBroker(func(op *broker.Options) {
		op.Addrs = []string{"127.0.0.1:6379"}
	})
	broker.Init()
	s := micro.NewService(
		micro.Name("kafkatTest"),
		micro.Broker(broker),
	)
	s.Init(micro.AfterStart(func() error {
		brk := s.Options().Broker
		if err := brk.Connect(); err != nil {
			fmt.Println(err)
		}
		//go sub(brk)
		//time.Sleep(time.Second * 5) //先让订阅者创建，否则在创建途中无法读取数据
		go pub(brk)
		return nil
	}))
	s.Run()
}

var (
	topic = "www"
)

func pub(brk broker.Broker) {
	for i := 0; i < 100; i++ {
		err := brk.Publish(topic, &broker.Message{
			Header: map[string]string{"id": strconv.Itoa(i)},
			Body:   []byte("你好啊"),
		}, func(op *broker.PublishOptions) {
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}

func sub(brk broker.Broker) {
	brk.Subscribe(topic, func(e broker.Event) error {
		fmt.Println(e.Message().Body, e.Message().Header)
		return nil
	})
}
