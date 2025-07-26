package panel

import (
	"encoding/json"
	"fmt"
	"github.com/ethanwang9/covid19/server/api/dxy"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// name: 统计数据
// author: Ethan.Wang
// desc:

// GetAreaStat 获取统计数据
func GetAreaStat(ctx *gin.Context) {
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

	// 获取redis数据
	res, err := global.REDIS.Get(ctx, "DXYAreaStat").Result()
	if err != nil {
		global.LOG.Warn("获取用户基本信息#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}
	var coreData []dxy.AreaStat
	json.Unmarshal([]byte(res), &coreData)

	// 清洗数据
	type resData struct {
		CurrentConfirmedCount int `json:"currentConfirmedCount"`
		ConfirmedCount        int `json:"confirmedCount"`
		DeadCount             int `json:"deadCount"`
		CuredCount            int `json:"curedCount"`
	}
	var data resData
	for _, v := range coreData {
		if v.ProvinceName == province {
			data = resData{
				CurrentConfirmedCount: v.CurrentConfirmedCount,
				ConfirmedCount:        v.ConfirmedCount,
				DeadCount:             v.DeadCount,
				CuredCount:            v.CuredCount,
			}
		}
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// GetAreaStatDetails 获取统计详细数据
func GetAreaStatDetails(ctx *gin.Context) {
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

	// 获取redis数据
	res, err := global.REDIS.Get(ctx, "DXYAreaStat").Result()
	if err != nil {
		global.LOG.Warn("获取用户基本信息#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}
	var coreData []dxy.AreaStat
	json.Unmarshal([]byte(res), &coreData)

	// 清洗数据
	data := make([][]string, 0)
	for _, v := range coreData {
		if v.ProvinceName == province {
			for _, v2 := range v.Cities {
				temp := []string{
					v2.CityName,
					fmt.Sprintf("%v", v2.CurrentConfirmedCount),
					fmt.Sprintf("%v", v2.ConfirmedCount),
					fmt.Sprintf("%v", v2.DeadCount),
					fmt.Sprintf("%v", v2.CuredCount),
				}
				data = append(data, temp)
			}
		}
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    data,
	})
}
