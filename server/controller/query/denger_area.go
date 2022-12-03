package query

import (
	"encoding/json"
	"fmt"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// name: 风险地区
// author: Ethan.Wang
// desc:

// 用户归属地返回数据结构
type resData struct {
	UpdateTime string        `json:"update_time"`
	HCount     int           `json:"h_count"`
	HList      []resDataList `json:"h_list"`
	MCount     int           `json:"m_count"`
	MList      []resDataList `json:"m_list"`
	LCount     int           `json:"l_count"`
	LList      []resDataList `json:"l_list"`
}

type resDataList struct {
	Type       string   `json:"type"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	County     string   `json:"county"`
	AreaName   string   `json:"area_name"`
	Communitys []string `json:"communitys"`
}

type resDataList2 struct {
	Type       string            `json:"type"`
	Province   string            `json:"province"`
	City       string            `json:"city"`
	County     string            `json:"county"`
	AreaName   string            `json:"area_name"`
	Communitys []resDataListComm `json:"communitys"`
}

type resDataListComm struct {
	Level string `json:"level"`
	Info  string `json:"info"`
}

// GetDangerAreaLocation 获取归属地风险地区
func GetDangerAreaLocation(ctx *gin.Context) {
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

	// 获取用户UID
	token, err := core.JwtApp.Decode(ctx.GetHeader("Authorization")[7:])
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySystem,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 查询redis
	res, err := global.REDIS.Get(ctx, "login#"+token).Result()
	if err != nil {
		global.LOG.Warn("获取归属地风险地区#获取缓存信息失败", zap.Error(err))
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库信息失败",
			Data:    nil,
		})
		return
	}
	var userInfo global.RedisLogin
	json.Unmarshal([]byte(res), &userInfo)

	// 获取风险地区数据
	data, err := global.REDIS.Get(ctx, "GovDangerArea").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试！",
			Data:    nil,
		})
		return
	}
	var dangerData gov.DengerArea
	json.Unmarshal([]byte(data), &dangerData)

	// 过滤用户归属地数据
	h := make([]resDataList, 0)
	hc := 0
	m := make([]resDataList, 0)
	mc := 0
	l := make([]resDataList, 0)
	lc := 0
	for _, v := range dangerData.Data.Highlist {
		if v.Province == userInfo.Location {
			h = append(h, v)
			hc += len(v.Communitys)
		}
	}
	for _, v := range dangerData.Data.Middlelist {
		if v.Province == userInfo.Location {
			m = append(m, v)
			mc += len(v.Communitys)
		}
	}
	for _, v := range dangerData.Data.Lowlist {
		if v.Province == userInfo.Location {
			l = append(l, v)
			lc += len(v.Communitys)
		}
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: resData{
			UpdateTime: dangerData.Data.EndUpdateTime,
			HCount:     hc,
			HList:      h,
			MCount:     mc,
			MList:      m,
			LCount:     lc,
			LList:      l,
		},
	})
}

// QueryDangerArea 查询风险地区
func QueryDangerArea(ctx *gin.Context) {
	// 获取参数
	province := ctx.Query("province")
	county := ctx.Query("county")
	city := ctx.Query("city")
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"province":  province,
		"county":    county,
		"city":      city,
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

	// 获取风险地区数据
	redisData, err := global.REDIS.Get(ctx, "GovDangerArea").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试！",
			Data:    nil,
		})
		return
	}
	var dangerData gov.DengerArea
	json.Unmarshal([]byte(redisData), &dangerData)

	// 匹配数据
	var res resDataList2
	level := "n"
	for _, v := range dangerData.Data.Highlist {
		if v.AreaName == fmt.Sprintf("%v %v %v", province, county, city) {
			res = resDataList2(struct {
				Type       string
				Province   string
				City       string
				County     string
				AreaName   string
				Communitys []resDataListComm
			}{Type: v.Type, Province: v.Province, City: v.City, County: v.County, AreaName: v.AreaName, Communitys: nil})

			for _, v := range v.Communitys {
				res.Communitys = append(res.Communitys, resDataListComm{
					Level: "高风险",
					Info:  v,
				})
			}

			level = "h"
		}
	}
	for _, v := range dangerData.Data.Middlelist {
		if v.AreaName == fmt.Sprintf("%v %v %v", province, county, city) {
			if len(res.Province) == 0 {
				res = resDataList2(struct {
					Type       string
					Province   string
					City       string
					County     string
					AreaName   string
					Communitys []resDataListComm
				}{Type: v.Type, Province: v.Province, City: v.City, County: v.County, AreaName: v.AreaName, Communitys: nil})

				for _, v := range v.Communitys {
					res.Communitys = append(res.Communitys, resDataListComm{
						Level: "中风险",
						Info:  v,
					})
				}
				level = "m"
			} else {
				for _, v := range v.Communitys {
					res.Communitys = append(res.Communitys, resDataListComm{
						Level: "中风险",
						Info:  v,
					})
				}
			}
		}
	}
	for _, v := range dangerData.Data.Lowlist {
		if v.AreaName == fmt.Sprintf("%v %v %v", province, county, city) {
			if len(res.Province) == 0 {
				res = resDataList2(struct {
					Type       string
					Province   string
					City       string
					County     string
					AreaName   string
					Communitys []resDataListComm
				}{Type: v.Type, Province: v.Province, City: v.City, County: v.County, AreaName: v.AreaName, Communitys: nil})

				for _, v := range v.Communitys {
					res.Communitys = append(res.Communitys, resDataListComm{
						Level: "低风险",
						Info:  v,
					})
				}
				level = "l"
			} else {
				for _, v := range v.Communitys {
					res.Communitys = append(res.Communitys, resDataListComm{
						Level: "低风险",
						Info:  v,
					})
				}
			}
		}
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"info":        res,
			"level":       level,
			"update_time": dangerData.Data.EndUpdateTime,
		},
	})
}

// GetCity 获取城市列表
func GetCity(ctx *gin.Context) {
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

	// 获取风险地区数据
	redisData, err := global.REDIS.Get(ctx, "GovCityList").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试！",
			Data:    nil,
		})
		return
	}
	var cityList gov.CityList
	json.Unmarshal([]byte(redisData), &cityList)

	// 清洗数据
	type backData struct {
		Value    string     `json:"value"`
		Label    string     `json:"label"`
		Children []backData `json:"children"`
	}

	back := make([]backData, 0)
	for _, v := range cityList.Data.List {
		var temp backData
		temp.Label = v.Name
		temp.Value = v.Name

		for _, v2 := range v.Items {
			var temp2 backData
			temp2.Label = v2.Name
			temp2.Value = v2.Name

			for _, v3 := range v2.Items {
				var temp3 backData
				temp3.Label = v3.Name
				temp3.Value = v3.Name
				temp2.Children = append(temp2.Children, temp3)
			}
			temp.Children = append(temp.Children, temp2)
		}
		back = append(back, temp)
	}

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    back,
	})
}

// QueryDangerAreaAll 获取全国风险地区
func QueryDangerAreaAll(ctx *gin.Context) {
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

	// 获取风险地区数据
	redisData, err := global.REDIS.Get(ctx, "GovDangerArea").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据失败，请再次重试！",
			Data:    nil,
		})
		return
	}
	var dangerData gov.DengerArea
	json.Unmarshal([]byte(redisData), &dangerData)

	// 返回数据
	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    dangerData.Data,
	})
}
