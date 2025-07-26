package user

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// name: 获取用户信息
// auth: Ethan.Wang
// desc:

// GetBaseInfo 获取用户基本信息
func GetBaseInfo(ctx *gin.Context) {
	// 获取参数
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"timestamp": timestamp,
		"sign":      sign,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByNotTrue,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 获取用户UID
	token, err := core.JwtApp.Decode(ctx.GetHeader("Authorization")[7:])
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySystem,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 查询redis
	res, err := global.REDIS.Get(ctx, "login#"+token).Result()
	if err != nil {
		global.LOG.Warn("获取用户基本信息#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}

	var d global.RedisLogin
	json.Unmarshal([]byte(res), &d)

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"uid":      d.Uid,
			"nickname": d.Nickname,
			"avatar":   d.Avatar,
		},
	})
}
