package panel

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
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

// GetAreaDanger 获取风险地区
func GetAreaDanger(ctx *gin.Context) {
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
		if v.Province == province {
			h = append(h, v)
			hc += len(v.Communitys)
		}
	}
	for _, v := range dangerData.Data.Middlelist {
		if v.Province == province {
			m = append(m, v)
			mc += len(v.Communitys)
		}
	}
	for _, v := range dangerData.Data.Lowlist {
		if v.Province == province {
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
