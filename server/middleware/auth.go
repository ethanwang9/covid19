package middleware

import (
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// name: 中间件-用户认证
// author: Ethan.Wang
// desc:

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 验证参数
		if len(ctx.GetHeader("Authorization")) == 0 {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySystem,
				Message: "认证失败",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		// 获取用户token
		token, err := core.JwtApp.Decode(ctx.GetHeader("Authorization")[7:])
		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorByAuth,
				Message: err.Error(),
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		// 检查用户当前登录状态
		redisUserData, err := global.REDIS.Get(ctx, "login#"+token).Result()
		if err != nil {
			if err.Error() == "redis: nil" {
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorByAuth,
					Message: "token 已过期",
					Data:    nil,
				})
				ctx.Abort()
				return
			} else {
				global.LOG.Warn("中间件-Auth#获取缓存信息失败", zap.Error(err))
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorBySQL,
					Message: "获取缓存数据库信息失败",
					Data:    nil,
				})
				ctx.Abort()
				return
			}
		}

		// 用户心跳存活设置
		err = global.REDIS.Set(ctx, "login#"+token, redisUserData, time.Minute*30).Err()
		if err != nil {
			global.LOG.Warn("中间件-Auth#用户心跳存活设置失败", zap.Error(err))
		}
	}
}
