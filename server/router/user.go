package router

import (
	"github.com/ethanwang9/covid19/server/controller/user"
	"github.com/ethanwang9/covid19/server/middleware"
	"github.com/gin-gonic/gin"
)

// name: 用户
// author: Ethan.Wang
// desc: 用户路由

func User(group *gin.RouterGroup) {
	r := group.Group("/user")
	r.Use(middleware.Auth())
	{
		// 获取用户信息 - user
		r.POST("/info/base", user.GetBaseInfo)
		// 获取用户列表 - admin
		r.GET("/list", user.GetUserList)
		// 获取用户信息#复合条件查询 - admin
		r.GET("/query", user.QueryUser)
		// 更新用户权限 - admin
		r.POST("/update/level", user.UpdateUserLevel)
	}
}
