package Initialization

import (
	"context"
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
	"time"
)

// name: 初始化-配置文件
// author: Ethan.Wang
// desc:

// 初始化配置
func initConfig() {
	// 初始化微信公众号AccessToken
	mpAccessToken()
}

// 初始化微信公众号AccessToken
func mpAccessToken() {
	// 判断 Redis 中是否有 WxMpAccessToken
	res, err := global.REDIS.Get(context.Background(), "WxMpAccessToken").Result()
	if err != nil && err.Error() != "redis: nil" {
		global.LOG.Error("初始化微信公众号AccessToken#获取redis数据失败",
			zap.String("key", "wx_mp_accesstoken"),
			zap.String("error", err.Error()),
		)
		return
	}
	if len(res) != 0 {
		return
	}

	// 获取微信公众号AccessToken
	token, err := global.MP.AccessToken(context.Background())
	if err != nil {
		global.LOG.Error("初始化微信公众号AccessToken#获取微信公众号AccessToken失败", zap.String("error", err.Error()))
		return
	}

	// 写入微信公众号AccessToken到Redis中失败
	err = global.REDIS.Set(context.Background(), "WxMpAccessToken", token.Token, 2*time.Hour).Err()
	if err != nil {
		global.LOG.Error("初始化微信公众号AccessToken#写入微信公众号AccessToken到Redis中失败", zap.String("error", err.Error()))
		return
	}
}
