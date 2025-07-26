package user

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/model"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// name: 管理用户信息
// author: Ethan.Wang
// desc:

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	// 获取参数
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"limit":     limit,
		"page":      page,
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

	// 获取数据库信息数据
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	offsetN := limitN * (pageN - 1)
	users, err := model.UserApp.New(model.User{}).Get(limitN, offsetN)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 获取用户数量
	total, err := model.UserApp.New(model.User{}).GetCount()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 清洗数据
	type user struct {
		Uid       string `gorm:"column:uid" json:"uid"`
		Avatar    string `gorm:"column:avatar" json:"avatar"`
		Nickname  string `gorm:"column:nickname" json:"nickname"`
		Level     string `gorm:"column:level" json:"level"`
		Location  string `gorm:"column:location" json:"location"`
		CreatedAt string `gorm:"column:create_at" json:"create_at"`
		UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	}
	users2 := make([]user, 0)
	for _, v := range users {
		tempLevel := ""

		switch v.Level {
		case -99:
			tempLevel = "stop"
		case 1:
			tempLevel = "user"
		case 99:
			tempLevel = "admin"
		}

		users2 = append(users2, user{
			Uid:       v.Uid,
			Avatar:    v.Avatar,
			Nickname:  v.Nickname,
			Level:     tempLevel,
			Location:  v.Location,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"list":  users2,
			"total": total,
		},
	})
}

// QueryUser 复合查询用户信息
func QueryUser(ctx *gin.Context) {
	// 获取参数
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	uid := ctx.Query("uid")
	nickname := ctx.Query("nickname")
	level := ctx.Query("level")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"limit":     limit,
		"page":      page,
		"uid":       uid,
		"nickname":  nickname,
		"level":     level,
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

	var levelInt = 0
	switch level {
	case "admin":
		levelInt = global.UserLevelByAdmin
	case "user":
		levelInt = global.UserLevelByUser
	case "stop":
		levelInt = global.UserLevelByStop
	}

	// 获取数据库信息数据
	pageN, _ := strconv.Atoi(page)
	limitN, _ := strconv.Atoi(limit)
	offsetN := limitN * (pageN - 1)
	users, total, err := model.UserApp.New(model.User{
		Uid:      uid,
		Nickname: nickname,
		Level:    levelInt,
	}).GetUserByMore(limitN, offsetN)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 清洗数据
	type user struct {
		Uid       string `gorm:"column:uid" json:"uid"`
		Avatar    string `gorm:"column:avatar" json:"avatar"`
		Nickname  string `gorm:"column:nickname" json:"nickname"`
		Level     string `gorm:"column:level" json:"level"`
		Location  string `gorm:"column:location" json:"location"`
		CreatedAt string `gorm:"column:create_at" json:"create_at"`
		UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
	}
	users2 := make([]user, 0)
	for _, v := range users {
		tempLevel := ""

		switch v.Level {
		case global.UserLevelByStop:
			tempLevel = "stop"
		case global.UserLevelByUser:
			tempLevel = "user"
		case global.UserLevelByAdmin:
			tempLevel = "admin"
		}

		users2 = append(users2, user{
			Uid:       v.Uid,
			Avatar:    v.Avatar,
			Nickname:  v.Nickname,
			Level:     tempLevel,
			Location:  v.Location,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"list":  users2,
			"total": total,
		},
	})
}

// UpdateUserLevel 更新用户权限
func UpdateUserLevel(ctx *gin.Context) {
	// 获取参数
	uid := ctx.PostForm("uid")
	level := ctx.PostForm("level")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"uid":       uid,
		"level":     level,
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

	var levelInt = global.UserLevelByStop
	switch level {
	case "admin":
		levelInt = global.UserLevelByAdmin
	case "user":
		levelInt = global.UserLevelByUser
	case "stop":
		levelInt = global.UserLevelByStop
	}

	// 更新数据库
	err = model.UserApp.New(model.User{
		Uid:   uid,
		Level: levelInt,
	}).UpdateByLevel()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    nil,
	})
}
