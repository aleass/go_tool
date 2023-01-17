package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func GenRedis() *redis.Client {
	Redis = getRedisClient()
	if _, err := Redis.Ping().Result(); err != nil {
		panic("redis初始化打开失败:" + err.Error())
	}
	return Redis
}

/**
 * 获取 redis 客户端
 */
func getRedisClient() *redis.Client {
	redisConfig := Configer.Redis
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
		Network:  "tcp",
		PoolSize: 50,
	})
}
