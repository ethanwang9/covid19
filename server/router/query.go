package router

import (
	"github.com/ethanwang9/covid19/server/controller/query"
	"github.com/ethanwang9/covid19/server/middleware"
	"github.com/gin-gonic/gin"
)

// name: 主要业务
// author: Ethan.Wang
// desc: 主要业务路由,查询系统各种数据

func Query(group *gin.RouterGroup) {
	r := group.Group("/query")
	r.Use(middleware.Auth())
	{
		// 获取出行政策
		r.GET("/travel", query.GetTravel)
		// 获取出行政策城市列表
		r.GET("/travel/city_list", query.GetTravelCityList)
		// 获取归属地风险地区
		r.GET("/denger_area/location", query.GetDangerAreaLocation)
		// 获取城市列表
		r.GET("/danger_area/city", query.GetCity)
		// 查询风险地区
		r.GET("/danger_area/query", query.QueryDangerArea)
		// 获取全国风险地区
		r.GET("/danger_area/all", query.QueryDangerAreaAll)
	}
}
