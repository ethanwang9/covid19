package query

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	internet "github.com/ethanwang9/covid19/server/api/tencent"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// name: 出行政策
// author: Ethan.Wang
// desc: 出行政策控制器

// GetTravel 获取出行政策
func GetTravel(ctx *gin.Context) {
	// 获取参数
	from := ctx.Query("from")
	to := ctx.Query("to")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"from":      from,
		"to":        to,
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

	// 请求接口
	data, err := api.ApiApp.Tencent.Travel.Get(from, to)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "接口请求失败，请稍后重试！",
			Data:    nil,
		})
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    data.Rsp,
	})
}

// GetTravelCityList 获取出行政策城市列表
func GetTravelCityList(ctx *gin.Context) {
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

	// 查询redis
	data, err := global.REDIS.Get(ctx, "TencentCityList").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "查询失败，请再次重试",
			Data:    nil,
		})
		return
	}

	// 数据清洗
	var list internet.CityList
	json.Unmarshal([]byte(data), &list)

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    list.Rsp.CityList,
	})
}
