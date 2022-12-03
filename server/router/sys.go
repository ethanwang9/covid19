package router

import (
	"github.com/ethanwang9/covid19/server/controller/sys"
	"github.com/ethanwang9/covid19/server/middleware"
	"github.com/gin-gonic/gin"
)

// name: 系统路由
// author: Ethan.Wang
// desc:

func Sys(group *gin.RouterGroup) {
	r := group.Group("/sys")
	r.Use(middleware.Auth())
	{
		// 获取公众号配置
		r.GET("/mp", sys.GetMp)
		// 修改公众号配置
		r.POST("/mp", sys.PostMp)
		// 获取系统信息
		r.GET("/info", sys.GetInfo)
		// 修改
		r.POST("/info", sys.SetInfo)
	}
}
