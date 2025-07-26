package sys

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/model"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// name: 系统信息
// author: Ethan.Wang
// desc: 系统信息控制器

// GetInfo 获取系统信息
func GetInfo(ctx *gin.Context) {
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

	// 获取数据库信息
	data, err := model.SysApp.New(model.Sys{}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败",
			Data:    nil,
		})
		return
	}

	// 返回信息
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"copyright": data.Copyright,
			"gov_no1":   data.GovNo1,
			"gov_no2":   data.GovNo2,
			"mp_url":    data.MpUrl,
			"mp_img":    data.MpImg,
			"mail":      data.Mail,
			"blog":      data.Blog,
		},
	})
	return
}

// SetInfo 设置系统信息
func SetInfo(ctx *gin.Context) {
	// 获取参数
	copyright := ctx.PostForm("copyright")
	govNo1 := ctx.PostForm("gov_no1")
	govNo2 := ctx.PostForm("gov_no2")
	mpUrl := ctx.PostForm("mp_url")
	mpImg := ctx.PostForm("mp_img")
	mail := ctx.PostForm("mail")
	blog := ctx.PostForm("blog")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"copyright": copyright,
		"gov_no1":   govNo1,
		"gov_no2":   govNo2,
		"mp_url":    mpUrl,
		"mp_img":    mpImg,
		"mail":      mail,
		"blog":      blog,
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

	// 清洗数据
	if copyright == "EMPTY" {
		copyright = ""
	}
	if govNo1 == "EMPTY" {
		govNo1 = ""
	}
	if govNo2 == "EMPTY" {
		govNo2 = ""
	}
	if mpUrl == "EMPTY" {
		mpUrl = ""
	}
	if mpImg == "EMPTY" {
		mpImg = ""
	}
	if mail == "EMPTY" {
		mail = ""
	}
	if blog == "EMPTY" {
		blog = ""
	}

	// 设置数据库数据
	err = model.SysApp.New(model.Sys{
		Copyright: copyright,
		GovNo1:    govNo1,
		GovNo2:    govNo2,
		MpUrl:     mpUrl,
		MpImg:     mpImg,
		Mail:      mail,
		Blog:      blog,
	}).Set()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "更新数据失败",
			Data:    nil,
		})
		return
	}

	// 返回信息
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    nil,
	})
	return
}
