package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

type KafkaConfig struct {
	Topics        []string `json:"topics"`
	Servers       []string `json:"servers"`
	ConsumerGroup string   `json:"consumer_group"`
}

//SyncProducerString 同步模式生产者
func (c *KafkaConfig) SyncProducerString(topic, key, value string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true ///是否开启消息发送成功后通知 successes channel
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机分区器
	cli, err := sarama.NewClient(c.Servers, config)
	if err != nil {
		return err
	}
	defer cli.Close()
	producer, err := sarama.NewSyncProducerFromClient(cli)
	if err != nil {
		return err
	}
	defer producer.Close()
	//partition, offset, err := producer.SendMessage()
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	})
	fmt.Println(partition, offset, key, value)
	return err
}

//SyncProducerString 异步模式生产者 　异步模式，顾名思义就是produce一个message之后不等待发送完成返回；这样调用者可以继续做其他的工作。
func (c *KafkaConfig) ASyncProducerString(topic, key, value string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true ///是否开启消息发送成功后通知 successes channel 打开了Return.Successes配置，则上述代码段等同于同步方式
	//config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机分区器
	cli, err := sarama.NewClient(c.Servers, config)
	if err != nil {
		return err
	}
	defer cli.Close()
	producer, err := sarama.NewAsyncProducerFromClient(cli)
	if err != nil {
		return err
	}
	defer producer.Close()
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.StringEncoder(value)}
	select {
	case m := <-producer.Successes():
		fmt.Println("成功", m.Value)
		return nil
	case err = <-producer.Errors():
		fmt.Println("写入错误", err)
		return err
	}
	return nil
}

//ProducerString 发布普通消息
func (c *KafkaConfig) ProducerString(topic, key, content string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	//broker:=sarama.NewBroker("")
	msg := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(content),
	}
	//sarama.NewSyncProducer() //同步发送者
	//sarama.NewAsyncProducer() //异步发送者
	client, err := sarama.NewSyncProducer(c.Servers, config)
	if err != nil {
		//log.Fatalln("kafka连接失败", err.Error()) //一般情况是端口不存在或者说docker的kafka配置错误导致宿主机无法访问
		return err
	}
	defer client.Close()
	//pid, offset, err := client.SendMessage(&msg) //pid，offset
	if _, _, err := client.SendMessage(&msg); err != nil {
		//log.Fatalln("kafka消息推送失败", err.Error())
		return err
	}
	return nil
}

//ProducerString 发布普通消息
func (c *KafkaConfig) ProducerJosn(topic, key string, msgEntity interface{}) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	bt, _ := json.Marshal(msgEntity)
	msg := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(bt),
	}
	client, err := sarama.NewSyncProducer(c.Servers, config)
	if err != nil {
		return err
	}
	defer client.Close()
	if _, _, err := client.SendMessage(&msg); err != nil {
		//log.Fatalln("kafka消息推送失败", err.Error())
		return err
	}
	return nil
}

//ProducerByte 发布字节
func (c *KafkaConfig) ProducerByte(topic, key, content string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	msg := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(content),
	}
	client, err := sarama.NewSyncProducer(c.Servers, config)
	if err != nil {
		return err
	}
	defer client.Close()
	if _, _, err := client.SendMessage(&msg); err != nil {
		return err
	}
	return nil
}

//Consumer 消息订阅
func (c *KafkaConfig) Consumer(topic string) {
	consumer, err := sarama.NewConsumer(c.Servers, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer consumer.Close()
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for p := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(p), sarama.OffsetNewest) //这个很关键，如果要读取旧数据的话需要Old
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		/*for msg := range pc.Messages() {
			fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}*/
		defer pc.AsyncClose()
		for {
			select {
			case msg := <-pc.Messages():
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			case err := <-pc.Errors():
				fmt.Println("错误", err)
			}
		}
	}
}
