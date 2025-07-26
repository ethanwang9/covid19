package auth

import (
	"encoding/json"
	"fmt"
	"github.com/ethanwang9/covid19/server/api"
	"github.com/ethanwang9/covid19/server/api/internet"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/model"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/gochat/offia"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// name: 微信登录操作
// author: Ethan.Wang
// desc:

// WxLogin 获取微信登录授权链接
func WxLogin(ctx *gin.Context) {
	// 获取参数
	redirectURL := ctx.PostForm("back")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"back":      redirectURL,
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

	// 处理回调地址
	redirectURL = url.QueryEscape(redirectURL)

	// 生成用户标识
	state := strings.ToUpper(utils.Algorithm.UUID())

	// 生成链接
	u := global.MP.OAuth2URL(offia.ScopeSnsapiUser, redirectURL, state)

	// 写入缓存
	cacheLogin, _ := json.Marshal(global.RedisLogin{IsLogin: false, UUID: state})
	err = global.REDIS.Set(ctx, "login#"+state, string(cacheLogin), 2*time.Minute).Err()
	if err != nil {
		global.LOG.Error("获取微信登录授权链接#写UUID到缓存失败", zap.Error(err))
		return
	}

	// 返回信息
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"code": state,
			"url":  u,
		},
	})
}

// WxToken 获取微信登录Token
func WxToken(ctx *gin.Context) {
	// 获取参数
	code := ctx.Query("code")
	state := ctx.Query("state")

	// 获得登录AccessToken
	token, err := global.MP.Code2OAuthToken(ctx, code)
	if err != nil {
		global.LOG.Error("获取微信登录Token#获得用户登录AccessToken失败", zap.Error(err), zap.Any("data", gin.H{
			"code":    code,
			"state":   state,
			"request": ctx.Request,
		}))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "接口请求失败",
			Data:    nil,
		})
		return
	}
	// 过滤用户IP
	userIp := ctx.ClientIP()

	// 获取用户信息
	wxUserInfo := new(offia.ResultOAuthUser)
	global.MP.Do(ctx, token.AccessToken, offia.GetOAuthUser(token.OpenID, wxUserInfo))

	// 判断用户是否存在
	isFind, err := model.UserApp.New(model.User{WxOpenid: wxUserInfo.OpenID}).HasUser()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "数据库操作失败",
			Data: gin.H{
				"res":   wxUserInfo,
				"state": state,
			},
		})
		return
	}

	// 用户不存在
	if !isFind {
		// 获取IP归属地
		ip, err := api.ApiApp.Internet.IP.New(internet.IP{Ip: userIp}).Query()
		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorByAPI,
				Message: "网络接口请求失败",
				Data:    nil,
			})
			return
		}

		// 判断用户权限
		tempUserCount, err := model.UserApp.New(model.User{}).GetUserCount()
		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySQL,
				Message: "获取数据库信息失败",
				Data:    nil,
			})
			return
		}
		userLevel := global.UserLevelByUser
		if tempUserCount == 0 {
			userLevel = global.UserLevelByAdmin
		}

		// 创建用户
		err = model.UserApp.New(model.User{
			Uid:      utils.Algorithm.UUID(),
			WxOpenid: wxUserInfo.OpenID,
			Avatar:   wxUserInfo.HeadImgURL,
			Nickname: wxUserInfo.Nickname,
			Level:    userLevel,
			Location: ip.Province,
			Base:     model.Base{},
		}).Add()

		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySQL,
				Message: "数据库操作失败",
				Data:    nil,
			})
			return
		}
	}

	// 获取用户数据库
	dbUserInfo, err := model.UserApp.New(model.User{WxOpenid: wxUserInfo.OpenID}).GetByWxOpenid()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "数据库操作失败",
			Data:    nil,
		})
		return
	}

	// 判断用户是否被禁用
	if dbUserInfo.Level == global.UserLevelByStop {
		temp := global.RedisLogin{IsStop: true}
		tempJsonStr, _ := json.Marshal(temp)
		err = global.REDIS.Set(ctx, fmt.Sprintf("login#%v", state), string(tempJsonStr), 3*time.Minute).Err()
		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySQL,
				Message: "数据库操作失败",
				Data:    nil,
			})
			return
		}

		ctx.String(http.StatusOK, "用户已被禁用，请联系管理员解决！")
		return
	}

	// 获取用户IP属地
	ip, err := api.ApiApp.Internet.IP.New(internet.IP{Ip: userIp}).Query()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "网络接口请求失败",
			Data:    nil,
		})
		return
	}

	// 写入数据库
	err = model.UserApp.New(model.User{
		Uid:      dbUserInfo.Uid,
		Location: ip.Province,
	}).UpdateByLocation()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "更新数据库失败",
			Data:    nil,
		})
		return
	}

	// 写入缓存
	redisName := fmt.Sprintf("login#%v", state)
	redisWrite := global.RedisLogin{
		UUID:     state,
		Uid:      dbUserInfo.Uid,
		Nickname: dbUserInfo.Nickname,
		Avatar:   dbUserInfo.Avatar,
		Level:    dbUserInfo.Level,
		Location: ip.Province,
		IsLogin:  true,
		IsToken:  false,
	}
	redisValue, _ := json.Marshal(redisWrite)
	err = global.REDIS.Set(ctx, redisName, string(redisValue), time.Minute*30).Err()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "数据库操作失败",
			Data:    nil,
		})
		return
	}

	// 返回信息
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta http-equiv="X-UA-Compatible"content="IE=edge"><meta name="viewport"content="width=device-width, initial-scale=1.0"><title>登录成功</title><style type="text/CSS">*{padding:0;margin:0}body{width:100vw;height:100vh;display:flex;justify-content:center;align-items:center;background-color:#010655}#app{background-color:#fff;border-radius:15px;width:80%;padding:20px;box-sizing:0 08px rgba(0,0,0,0.3)}h3{padding-bottom:20px}p{text-align:center;font-size:20px;letter-spacing:2px}</style></head><body><div id="app"><h3>登录成功</h3><p>微信扫码登录成功</p></div></body></html>`)
}

// WxStatus 获取扫码登录状态
func WxStatus(ctx *gin.Context) {
	// 获取参数
	query := ctx.PostForm("query")
	timestamp := ctx.PostForm("timestamp")
	sign := ctx.PostForm("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"query":     query,
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

	// 查询 Redis
	res, err := global.REDIS.Get(ctx, "login#"+query).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取登录状态失败",
			Data:    nil,
		})
		return
	}

	// 获取redis信息
	var d global.RedisLogin
	json.Unmarshal([]byte(res), &d)

	// 判断用户是否被禁用
	if d.IsStop {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAuth,
			Message: "用户已被禁用",
			Data:    nil,
		})
		return
	}

	// 生成token
	token := ""
	if d.IsLogin == true && d.IsToken == false {
		token, err = core.JwtApp.Generate(d.UUID)
		if err != nil {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySystem,
				Message: "生成登录凭证失败",
				Data:    nil,
			})
			return
		} else {
			d.IsToken = true
			dStr, _ := json.Marshal(d)
			global.REDIS.Set(ctx, "login#"+query, dStr, 2*time.Hour)
		}
	} else if d.IsLogin && d.IsToken {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySystem,
			Message: "Token已生成,请求已拒绝",
			Data:    nil,
		})
		return
	}

	// 返回
	if d.IsLogin {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeSuccess,
			Message: "请求成功",
			Data: gin.H{
				"status":   d.IsLogin,
				"token":    token,
				"uid":      d.Uid,
				"nickname": d.Nickname,
				"avatar":   d.Avatar,
				"location": d.Location,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeSuccess,
			Message: "请求成功",
			Data: gin.H{
				"status": d.IsLogin,
			},
		})
	}

}
