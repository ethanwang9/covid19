package panel

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	"github.com/ethanwang9/covid19/server/api/ali"
	dxy2 "github.com/ethanwang9/covid19/server/api/dxy"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// name: Geo 地图
// author: Ethan.Wang
// desc:

type resGeoStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	H     int    `json:"h"`
	M     int    `json:"m"`
	L     int    `json:"l"`
}

// GetGeo 获取地图
func GetGeo(ctx *gin.Context) {
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
	if len(code) == 0 {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByAPI,
			Message: "参数 province 输入不正确",
			Data:    nil,
		})
		return
	}

	// 获取缓存数据
	var geo ali.Geo
	geoCaData, err := global.REDIS.Get(ctx, "GeoMap#"+code).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			// 获取接口消息
			geo, err = api.ApiApp.Ali.Geo.Get(code)
			if err != nil {
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorByAPI,
					Message: "接口请求失败",
					Data:    nil,
				})
				return
			}
			geoStr, _ := json.Marshal(geo)
			err := global.REDIS.Set(ctx, "GeoMap#"+code, string(geoStr), time.Hour*24).Err()
			if err != nil {
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorBySQL,
					Message: "数据库存储数据发送错误",
					Data:    nil,
				})
				return
			}
		} else {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySQL,
				Message: "获取数据缓存失败",
				Data:    nil,
			})
			return
		}
	} else {
		json.Unmarshal([]byte(geoCaData), &geo)
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    geo,
	})
}

// GetGeoData 获取地图数据
func GetGeoData(ctx *gin.Context) {
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
	resData := make([]resGeoStruct, 0)

	// 获取风险地区
	danger, err := global.REDIS.Get(ctx, "GovDangerArea").Result()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取缓存数据库内容失败",
			Data:    nil,
		})
		return
	}
	var dangerData gov.DengerArea
	json.Unmarshal([]byte(danger), &dangerData)

	for _, v := range dangerData.Data.Highlist {
		if v.Province == province {
			tempName := ""

			if v.Province == v.City {
				tempName = v.County
			} else {
				tempName = v.City
			}

			tempFind := false
			for index, v2 := range resData {
				if v2.Name == tempName {
					tempFind = true
					resData[index].H = v2.H + len(v.Communitys)
				}
			}
			if !tempFind {
				resData = append(resData, resGeoStruct{
					Name:  tempName,
					Value: 0,
					H:     len(v.Communitys),
					M:     0,
					L:     0,
				})
			}

		}
	}
	for _, v := range dangerData.Data.Middlelist {
		if v.Province == province {
			tempName := ""

			if v.Province == v.City {
				tempName = v.County
			} else {
				tempName = v.City
			}

			for index, v2 := range resData {
				if v2.Name == tempName {
					resData[index].M = v2.M + len(v.Communitys)
				}
			}

		}
	}
	for _, v := range dangerData.Data.Lowlist {
		if v.Province == province {
			tempName := ""

			if v.Province == v.City {
				tempName = v.County
			} else {
				tempName = v.City
			}

			for index, v2 := range resData {
				if v2.Name == tempName {
					resData[index].L = v2.L + len(v.Communitys)
				}
			}

		}
	}

	// 确诊信息
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
	json.Unmarshal([]byte(dxy), &dxyData)

	for _, v := range dxyData {
		if v.ProvinceName == province {
			for _, v2 := range v.Cities {
				isNotFind := true
				cityName := changeCityName(province, v2.CityName)
				for index, v3 := range resData {
					if cityName == v3.Name {
						if v2.CurrentConfirmedCountStr == "-" {
							resData[index].Value = -99999999
						} else {
							resData[index].Value = v2.CurrentConfirmedCount
						}
						isNotFind = false
					}
				}

				if isNotFind {
					resData = append(resData, resGeoStruct{
						Name:  cityName,
						Value: v2.CurrentConfirmedCount,
						H:     0,
						M:     0,
						L:     0,
					})
				}
			}
		}
	}

	// 获取城市中心
	var center []float64
	var geo ali.Geo
	geoCaData, err := global.REDIS.Get(ctx, "GeoMap#100000").Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			// 获取接口消息
			geo, err = api.ApiApp.Ali.Geo.Get("100000")
			if err != nil {
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorByAPI,
					Message: "接口请求失败",
					Data:    nil,
				})
				return
			}
			geoStr, _ := json.Marshal(geo)
			err := global.REDIS.Set(ctx, "GeoMap#100000", string(geoStr), time.Hour*24).Err()
			if err != nil {
				ctx.JSON(http.StatusOK, global.MsgBack{
					Code:    global.CodeErrorBySQL,
					Message: "数据库存储数据发送错误",
					Data:    nil,
				})
				return
			}
		} else {
			ctx.JSON(http.StatusOK, global.MsgBack{
				Code:    global.CodeErrorBySQL,
				Message: "获取数据缓存失败",
				Data:    nil,
			})
			return
		}
	} else {
		json.Unmarshal([]byte(geoCaData), &geo)
	}
	for _, v := range geo.Features {
		if v.Properties.Name == province {
			center = v.Properties.Center
		}
	}

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data: gin.H{
			"value":  resData,
			"center": center,
		},
	})
}

// 替换城市名称
// desc: 丁香园省份城市名称为简写，需要特殊处理，不然echarts地图上无法显示正确的信息
func changeCityName(province, cityName string) string {
	switch province {
	case "广西壮族自治区":
		fallthrough
	case "广东省":
		fallthrough
	case "江西省":
		fallthrough
	case "山东省":
		fallthrough
	case "河南省":
		fallthrough
	case "宁夏回族自治区":
		fallthrough
	case "福建省":
		fallthrough
	case "安徽省":
		fallthrough
	case "浙江省":
		fallthrough
	case "江苏省":
		fallthrough
	case "辽宁省":
		fallthrough
	case "陕西省":
		fallthrough
	case "山西省":
		fallthrough
	case "河北省":
		cityName += "市"
	case "西藏自治区":
		if cityName != "阿里地区" {
			cityName += "市"
		}
	case "内蒙古自治区":
		if cityName != "阿拉善盟" && cityName != "锡林郭勒盟" && cityName != "锡林郭勒盟" && cityName != "乌海市" && cityName != "兴安盟" {
			cityName += "市"
		}
	case "吉林省":
		if cityName == "延边" {
			cityName = "延边朝鲜族自治州"
		} else if cityName != "吉林市" {
			cityName += "市"
		}
	case "黑龙江省":
		if cityName == "大兴安岭" {
			cityName += "地区"
		} else {
			cityName += "市"
		}
	case "湖北省":
		if cityName == "恩施州" {
			cityName = "恩施土家族苗族自治州"
		} else if cityName != "神农架林区" {
			cityName += "市"
		}
	case "湖南省":
		if cityName == "湘西自治州" {
			cityName = "湘西土家族苗族自治州"
		} else {
			cityName += "市"
		}
	case "海南省":
		if cityName == "临高" || cityName == "澄迈" || cityName == "定安" || cityName == "屯昌" {
			cityName += "县"
		} else if cityName == "乐东" || cityName == "陵水" || cityName == "昌江" || cityName == "保亭" || cityName == "琼中" || cityName == "白沙" {
			cityName += "黎族自治县"
		} else {
			cityName += "市"
		}
	case "重庆市":
		if cityName == "酉阳县" {
			cityName = "酉阳土家族苗族自治县"
		} else if cityName == "秀山县" {
			cityName = "秀山土家族苗族自治县"
		} else if cityName == "彭水县" {
			cityName = "彭水苗族土家族自治县"
		} else if cityName == "石柱县" {
			cityName = "石柱土家族自治县"
		}
	case "四川省":
		if cityName == "凉山州" {
			cityName = "凉山彝族自治州"
		} else if cityName == "甘孜州" {
			cityName = "甘孜藏族自治州"
		} else if cityName == "阿坝州" {
			cityName = "阿坝霾族羌族自治州"
		} else {
			cityName += "市"
		}
	case "贵州省":
		if cityName == "黔西南州" {
			cityName = "黔西南布依族苗族自治州"
		} else if cityName == "黔南州" {
			cityName = "黔南布依族苗族自治州"
		} else if cityName == "黔东南州" {
			cityName = "黔东南苗族侗族自治州"
		} else {
			cityName += "市"
		}
	case "云南省":
		if cityName == "德宏州" {
			cityName = "德宏傣族景颇族白治州"
		} else if cityName == "大理州" {
			cityName = "大理白族自治州"
		} else if cityName == "迪庆州" {
			cityName = "迪庆藏族自治州"
		} else if cityName == "楚雄州" {
			cityName = "楚雄彝族自治州"
		} else if cityName == "西双版纳" {
			cityName = "西双版纳傣族自治州"
		} else if cityName == "红河州" {
			cityName = "红河哈尼族彝族自治州"
		} else if cityName == "文山州" {
			cityName = "文山壮族苗族自治州"
		} else {
			cityName += "市"
		}
	case "甘肃省":
		if cityName == "临夏" {
			cityName = "临夏回族自治州"
		} else if cityName == "甘南" {
			cityName = "甘南藏族自治州"
		} else {
			cityName += "市"
		}
	case "青海省":
		if cityName == "海西州" {
			cityName = "海西蒙古族藏族自治州"
		} else if cityName == "玉树" {
			cityName = "玉树藏族自治州"
		} else if cityName == "海北州" {
			cityName = "海北藏族自治州"
		} else if cityName == "海南州" {
			cityName = "海南藏族自治州"
		} else if cityName == "黄南" {
			cityName = "黄南藏族自治州"
		} else if cityName == "果洛" {
			cityName = "果洛藏族自治州"
		} else {
			cityName += "市"
		}
	case "新疆维吾尔自治区":
		if cityName == "巴州" {
			cityName = "巴音郭楞蒙古自治州"
		} else if cityName == "伊犁州" {
			cityName = "伊犁哈萨克自治州"
		} else if cityName == "昌吉州" {
			cityName = "昌吉回族自治州"
		} else if cityName == "克孜勒苏" {
			cityName = "克孜勒苏柯尔克孜自治州"
		} else if cityName == "博尔塔拉" {
			cityName = "博尔塔拉蒙古自治州"
		} else if cityName != "喀什地区" && cityName != "和田地区" && cityName != "塔城地区" && cityName != "阿克苏地区" {
			cityName += "市"
		}
	}

	return cityName
}
