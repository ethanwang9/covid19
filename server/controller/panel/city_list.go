package panel

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// name: 城市列表
// author: Ethan.Wang
// desc:

// GetCityList 获取城市列表
func GetCityList(ctx *gin.Context) {
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

	// 获取缓存数据
	cityList, err := global.REDIS.Get(ctx, "GovCityList").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库内容失败",
			Data:    nil,
		})
		return
	}
	var cityListData gov.CityList
	json.Unmarshal([]byte(cityList), &cityListData)

	// 清洗数据
	type resData struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}
	res := make([]resData, 0)
	for _, v := range cityListData.Data.List {
		res = append(res, resData{
			Value: v.Name,
			Label: v.Name,
		})
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    res,
	})
}
