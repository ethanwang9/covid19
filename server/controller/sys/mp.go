package sys

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// name: 公众号
// author: Ethan.Wang
// desc: 公众号控制器

// GetMp 获取公众号配置
func GetMp(ctx *gin.Context) {
	// 获取参数
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

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
		global.LOG.Warn("获取公众号配置#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}
	var d global.RedisLogin
	json.Unmarshal([]byte(res), &d)

	// 判断用户权限
	if d.Level != global.UserLevelByAdmin {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByLevel,
			Message: "用户权限不足",
			Data:    nil,
		})
		return
	}

	// 返回信息
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"appid":  global.CONFIG.GetString("mp.appid"),
			"secret": global.CONFIG.GetString("mp.secret"),
		},
	})
}

// PostMp 修改公众号配置
func PostMp(ctx *gin.Context) {
	// 获取参数
	appid := ctx.PostForm("appid")
	secret := ctx.PostForm("secret")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"appid":     appid,
		"secret":    secret,
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
		global.LOG.Warn("获取公众号配置#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}
	var d global.RedisLogin
	json.Unmarshal([]byte(res), &d)

	// 判断用户权限
	if d.Level != global.UserLevelByAdmin {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByLevel,
			Message: "用户权限不足",
			Data:    nil,
		})
		return
	}

	// 修改信息
	global.CONFIG.Set("mp.appid", appid)
	global.CONFIG.Set("mp.secret", secret)
	err = global.CONFIG.WriteConfig()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySystem,
			Message: "保存配置失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 返回消息
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "修改成功",
		Data:    nil,
	})
	return
}
