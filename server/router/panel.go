package router

import (
	"github.com/ethanwang9/covid19/server/controller/panel"
	"github.com/ethanwang9/covid19/server/middleware"
	"github.com/gin-gonic/gin"
)

// name: 可视化面板数据
// author: Ethan.Wang
// desc:

func Panel(group *gin.RouterGroup) {
	r := group.Group("/panel")
	r.Use(middleware.Auth())
	{
		// 获取统计数据
		r.GET("/area_stat", panel.GetAreaStat)
		// 获取统计详细数据
		r.GET("/area_stat/details", panel.GetAreaStatDetails)
		// 获取风险地区
		r.GET("/area_danger", panel.GetAreaDanger)
		// 获取热点消息
		r.GET("/hot_message", panel.GetHotMessage)
		// 获取历史现存确诊
		r.GET("/history/currentConfirmed", panel.GetHistoryCurrConf)
		// 获取地图
		r.GET("/geo", panel.GetGeo)
		// 获取地图数据
		r.GET("/geo/data", panel.GetGeoData)
		// 获取城市列表
		r.GET("/city_list", panel.GetCityList)
	}
}
