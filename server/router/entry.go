package router

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/middleware"
	"github.com/gin-gonic/gin"
)

// name: 路由
// author: Ethan.Wang
// desc: 路由总控制器

// Init 初始化路由
func Init() *gin.Engine {
	// gin 模式
	if env := global.CONFIG.GetString("server.env"); env == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 路由
	router := gin.New()

	// 开放temp目录
	router.Static("assets", "./assets")

	// 中间件
	router.Use(middleware.GinLogger())
	router.Use(middleware.GinRecovery(true))

	// 版本 v1
	v1 := router.Group("/v1")
	{
		// 身份认证
		Auth(v1)
		// 用户
		User(v1)
		// 系统
		Sys(v1)
		// 查询业务
		Query(v1)
		// 可视化面板数据
		Panel(v1)
		// 公共查询
		Public(v1)

	}

	return router
}
