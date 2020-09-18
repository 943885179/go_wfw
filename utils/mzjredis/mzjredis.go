package mzjredis

import (
	"encoding/json"
	"qshapi/models"
	"time"

	"github.com/go-redis/redis"
)

var options redis.Options

//QshRedisInit 初始化redis
func QshRedisInit(op models.RedisConfig) {
	options = redis.Options{
		Addr:     op.Addr,
		Password: op.Password,
		Network:  op.Network,
		DB:       op.DB,
	}
}

//Ping 测试
func Ping() (string, error) {
	client := redis.NewClient(&options)
	defer client.Close()
	pong, err := client.Ping().Result()
	return pong, err
}

//Set 写入
func Set(key string, value interface{}, expiration time.Duration) error {
	client := redis.NewClient(&options)
	defer client.Close()
	return client.Set(key, value, expiration).Err()
}

//SetEntity 写入实体（转json后写入）
func SetEntity(key string, entity interface{}, expiration time.Duration) error {
	client := redis.NewClient(&options)
	defer client.Close()
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return client.Set(key, value, expiration).Err()
}

//GetEntity 读取
func GetEntity(key string, resp interface{}) (err error) {
	client := redis.NewClient(&options)
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
func Exists(keys ...string) (int64, error) {
	client := redis.NewClient(&options)
	defer client.Close()
	return client.Exists(keys...).Result()
}

//Get 读取
func Get(key string) (value string, err error) {
	client := redis.NewClient(&options)
	defer client.Close()
	return client.Get(key).Result()
	/*if err == redis.Nil {

	}*/

}

//Del 删除
func Del(key string) error {
	client := redis.NewClient(&options)
	defer client.Close()
	return client.Del(key).Err()
}
