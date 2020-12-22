package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"

	//"github.com/bsm/sarama-cluster" //会报错，需要修改把 sarama 版本改成 v1.24.1 就可以用啦, github.com/Shopify/sarama v1.24.1
	"os"
	"os/signal"
	"qshapi/utils/mzjkafka/al/configs"
)

var cfg *configs.MqConfig

//var consumer *cluster.Consumer
var consumer *sarama.Consumer
var sig chan os.Signal

//需要创建kafka.json
func init() {
	fmt.Println("init kafka consumer, it may take a few seconds...")

	var err error

	//cfg := &configs.MqConfig{}
	//configs.LoadJsonConfig(cfg, "kafka.json")
	cfg = &configs.MqConfig{
		[]string{"topic1"},
		[]string{"127.0.0.1:9092"},
		"group1",
	}

	clusterCfg := cluster.NewConfig()

	clusterCfg.Consumer.Return.Errors = true
	clusterCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	clusterCfg.Group.Return.Notifications = true

	clusterCfg.Version = sarama.V0_10_2_1
	if err = clusterCfg.Validate(); err != nil {
		msg := fmt.Sprintf("Kafka consumer config invalidate. config: %v. err: %v", *clusterCfg, err)
		fmt.Println(msg)
		panic(msg)
	}

	consumer, err = cluster.NewConsumer(cfg.Servers, cfg.ConsumerId, cfg.Topics, clusterCfg)
	if err != nil {
		msg := fmt.Sprintf("Create kafka consumer error: %v. config: %v", err, clusterCfg)
		fmt.Println(msg)
		panic(msg)
	}
	sig = make(chan os.Signal, 1)

}

func Start() {
	go consume()
}

func consume() {
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s Timestamp:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value), msg.Timestamp)
				consumer.MarkOffset(msg, "")
			}
		case err, more := <-consumer.Errors():
			if more {
				fmt.Println("Kafka consumer error: %v", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				fmt.Println("Kafka consumer rebalance: %v", ntf)
			}
		case <-sig:
			fmt.Errorf("Stop consumer server...")
			consumer.Close()
			return
		}
	}

}

func Stop(s os.Signal) {
	fmt.Println("Recived kafka consumer stop signal...")
	sig <- s
	fmt.Println("kafka consumer stopped!!!")
}

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	Start()

	select {
	case s := <-signals:
		Stop(s)
	}
}
