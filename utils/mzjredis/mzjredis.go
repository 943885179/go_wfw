package mzjredis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Network  string `json:"network"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	IsDebug    bool     `json:"isDebug"`    //是否为调试模式
}
/*
func Init(op models.RedisConfig) *RedisConfig {
	cf:= RedisConfig{
		Addr:     op.Addr,
		Password: op.Password,
		Network:  op.Network,
		DB:       op.DB,
	}
	return &cf
}*/

func redisInit(op *RedisConfig)*redis.Options {
	option := redis.Options{
		Addr:     op.Addr,
		Password: op.Password,
		Network:  op.Network,
		DB:       op.DB,
	}
	return  &option
}
//Ping 测试
func (ops *RedisConfig) Ping() (string, error) {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	pong, err := client.Ping().Result()
	return pong, err
}

//Set 写入
func (ops *RedisConfig) Set(key string, value interface{}, expiration time.Duration) error {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	return client.Set(key, value, expiration).Err()
}

//SetEntity 写入实体（转json后写入）
func  (ops *RedisConfig)SetEntity(key string, entity interface{}, expiration time.Duration) error {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return client.Set(key, value, expiration).Err()
}

//GetEntity 读取
func  (ops *RedisConfig)GetEntity(key string, resp interface{}) (err error) {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	value, err := client.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), resp)
	/*if err == redis.Nil {

	}*/

}

//Exists 是否存在(存在返回1，不存在返回0)
func (ops *RedisConfig) Exists(keys ...string) (int64, error) {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	return client.Exists(keys...).Result()
}

//Get 读取
func (ops *RedisConfig) Get(key string) (value string, err error) {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	return client.Get(key).Result()
	/*if err == redis.Nil {

	}*/

}

//Del 删除
func  (ops *RedisConfig)Del(key string) error {
	client := redis.NewClient(redisInit(ops))
	defer client.Close()
	return client.Del(key).Err()
}
