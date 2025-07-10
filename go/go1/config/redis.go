package config

import (
	"github.com/go-redis/redis"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //端口
		Password: "",
		DB:       0, //默认数据库
	})
	//返回一个redis客户端
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
