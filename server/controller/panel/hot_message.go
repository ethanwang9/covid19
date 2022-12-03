package panel

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// name: 热点消息
// author: Ethan.Wang
// desc:

// GetHotMessage 获取热点消息
func GetHotMessage(ctx *gin.Context) {
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

	// 省名转为城市代码
	code := ""
	redi, err := global.REDIS.Get(ctx, "GovCityList").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据缓存失败",
			Data:    nil,
		})
		return
	}
	var rediData gov.CityList
	json.Unmarshal([]byte(redi), &rediData)
	for _, v := range rediData.Data.List {
		if province == v.Name {
			code = v.Code
		}
	}

	// 获取接口消息
	res, err := api.ApiApp.Tencent.HotMessage.Get(code)
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "接口请求失败",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    res.Rsp.HotnewsRsp.Contents,
	})
}
