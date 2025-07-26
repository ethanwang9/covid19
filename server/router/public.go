package router

import (
	"github.com/ethanwang9/covid19/server/controller/public"
	"github.com/gin-gonic/gin"
)

// name: 公共查询
// author: Ethan.Wang
// desc: 前台数据返回

func Public(group *gin.RouterGroup) {
	r := group.Group("/public")
	{
		// 公开信息
		r.GET("/info", public.GetInfo)
	}
}
