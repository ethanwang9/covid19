package router

import (
	"github.com/ethanwang9/covid19/server/controller/auth"
	"github.com/gin-gonic/gin"
)

// name: 身份认证
// author: Ethan.Wang
// desc: 身份认证路由

func Auth(group *gin.RouterGroup) {
	r := group.Group("/auth")
	{
		// 微信验证消息事件签名
		r.GET("/wx/token", auth.GetToken)
		// 获取微信登录链接
		r.POST("/wx/login", auth.WxLogin)
		// 获取微信登录token
		r.GET("/wx/login/token", auth.WxToken)
		// 获取扫码登录状态
		r.POST("/wx/login/status", auth.WxStatus)
	}
}
