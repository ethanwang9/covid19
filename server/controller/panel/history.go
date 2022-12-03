package panel

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	dxy2 "github.com/ethanwang9/covid19/server/api/dxy"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// name: 历史数据
// author: Ethan.Wang
// desc:

// GetHistoryCurrConf 获取历史现存确诊
func GetHistoryCurrConf(ctx *gin.Context) {
	// 获取参数
	province := ctx.Query("province")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"province":  province,
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

	// 获取统计数据
	dxy, err := global.REDIS.Get(ctx, "DXYAreaStat").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库内容失败",
			Data:    nil,
		})
		return
	}
	var dxyData []dxy2.AreaStat
	dxyHistoryLink := ""
	json.Unmarshal([]byte(dxy), &dxyData)
	for _, v := range dxyData {
		if v.ProvinceName == province {
			dxyHistoryLink = v.StatisticsData
		}
	}

	// 获取统计数据
	dxyHistory, err := api.ApiApp.DXY.History.Get(dxyHistoryLink)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "服务端接口数据获取失败",
			Data:    nil,
		})
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    dxyHistory.Data,
	})
}
