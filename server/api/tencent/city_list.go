package tencent

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// name: 城市列表
// author: Ethan.Wang
// desc:

type CityList struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Rsp  struct {
		CityList []struct {
			Children []struct {
				Children []interface{} `json:"children"`
				CityCode string        `json:"cityCode"`
				Label    string        `json:"label"`
			} `json:"children"`
			CityCode string `json:"cityCode"`
			Label    string `json:"label"`
		} `json:"cityList"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"rsp"`
	ThRequestId string `json:"th-request-id"`
}

// 请求结构体
type cityListReq struct {
	Request struct {
		Req struct {
		} `json:"req"`
	} `json:"request"`
	Service string `json:"service"`
	Func    string `json:"func"`
	Context struct {
		UserId string `json:"userId"`
	} `json:"context"`
}

// Get 获取城市列表
func (c *CityList) Get() (CityList, error) {
	// 构造请求
	req := cityListReq{
		Service: "TravelPrevention",
		Func:    "getCityList",
		Context: struct {
			UserId string `json:"userId"`
		}(struct{ UserId string }{UserId: utils.Algorithm.UUID2()}),
	}

	// 发起请求
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"content-type": "application/json",
			"origin":       "https://feiyan.wecity.qq.com",
			"referer":      "https://feiyan.wecity.qq.com",
			"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
		}).
		SetBody(req).
		Post("https://wechat.wecity.qq.com/trpcapi/TravelPrevention/getCityList")

	if err != nil {
		global.LOG.Error("接口#腾讯-城市列表 [出行政策附属接口]-获取城市列表请求失败", zap.Error(err))
		return CityList{}, err
	}

	// 解析数据
	var d CityList
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
