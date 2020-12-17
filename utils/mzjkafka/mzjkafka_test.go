package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

type TestDto struct {
	Name string
}

func TestProducer(t *testing.T) {
	cc := KafkaConfig{
		Servers: []string{"127.0.0.1:9092"},
	}
	if err := cc.ProducerString("sun", "", "你好"); err != nil {
		t.Error(err)
	}
	if err := cc.ProducerByte("sun", "", "你好啊啊啊"); err != nil {
		t.Error(err)
	}
	d := TestDto{
		Name: "测试",
	}
	if err := cc.ProducerJosn("sun", "", d); err != nil {
		t.Error(err)
	}
	t.Log("kafka消息推送成功")
}

func TestConsumer(t *testing.T) {
	cc := KafkaConfig{
		Servers: []string{"127.0.0.1:9092"},
	}
	cc.Consumer("bbb")
	fmt.Scanln()
}

func TestProducerFor1(t *testing.T) {
	cc := KafkaConfig{
		Servers: []string{"127.0.0.1:9092"},
	}
	var i = 0
	for {
		i++
		cc.ProducerString("bbb", strconv.Itoa(i), "hello"+strconv.Itoa(int(time.Now().UnixNano())))
		time.Sleep(time.Second)
	}
}
func TestProducerFor2(t *testing.T) {
	cc := KafkaConfig{
		Servers: []string{"127.0.0.1:9092"},
	}
	var i = 0
	for {
		i++
		cc.ProducerString("bbb", "copy"+strconv.Itoa(i), "hello"+strconv.Itoa(int(time.Now().UnixNano())))
		time.Sleep(time.Second)
	}
}
