package core

import (
	"context"
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// name: redis
// author: Ethan.Wang
// desc:

// InitRedis 初始化 Redis
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", global.CONFIG.GetString("redis.host"), global.CONFIG.GetString("redis.port")),
		Password: global.CONFIG.GetString("redis.password"),
		DB:       global.CONFIG.GetInt("redis.db"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Panic("redis初始化失败", zap.Error(err))
	}

	return client
}
