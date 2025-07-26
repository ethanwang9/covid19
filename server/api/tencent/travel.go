package tencent

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// name: 出行政策
// author: Ethan.Wang
// desc:

type Travel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Rsp  struct {
		Code int `json:"code"`
		From struct {
			CityCode          string `json:"cityCode"`
			CityName          string `json:"cityName"`
			Code              int    `json:"code"`
			DiseaseControlTel string `json:"diseaseControlTel"`
			HighInDesc        string `json:"highInDesc"`
			IsInUpdate        bool   `json:"isInUpdate"`
			IsOutUpdate       bool   `json:"isOutUpdate"`
			Labels            []struct {
				Colour   string `json:"colour"`
				IsJump   bool   `json:"isJump"`
				IsShow   bool   `json:"isShow"`
				JumpLink struct {
					AppId  string `json:"appId"`
					AppVer string `json:"appVer"`
					Type   int    `json:"type"`
					Url    string `json:"url"`
				} `json:"jumpLink"`
				Label string `json:"label"`
			} `json:"labels"`
			LowInDesc    string `json:"lowInDesc"`
			MedianInDesc string `json:"medianInDesc"`
			Message      string `json:"message"`
			News         struct {
				Context  string `json:"context"`
				IsShow   bool   `json:"isShow"`
				JumpLink struct {
					AppId  string `json:"appId"`
					AppVer string `json:"appVer"`
					Type   int    `json:"type"`
					Url    string `json:"url"`
				} `json:"jumpLink"`
			} `json:"news"`
			OutDesc       string `json:"outDesc"`
			ProvinceName  string `json:"provinceName"`
			RiskLevel     int    `json:"riskLevel"`
			SourceContext string `json:"sourceContext"`
			SourceLink    struct {
				AppId  string `json:"appId"`
				AppVer string `json:"appVer"`
				Type   int    `json:"type"`
				Url    string `json:"url"`
			} `json:"sourceLink"`
		} `json:"from"`
		ImportantDict struct {
			IsShow bool   `json:"isShow"`
			Type   string `json:"type"`
			Value  string `json:"value"`
		} `json:"importantDict"`
		ImportantNotice string `json:"importantNotice"`
		InKeywordDict   struct {
			IsShow bool   `json:"isShow"`
			Type   string `json:"type"`
			Value  string `json:"value"`
		} `json:"inKeywordDict"`
		Message        string `json:"message"`
		OutKeywordDict struct {
			IsShow bool   `json:"isShow"`
			Type   string `json:"type"`
			Value  string `json:"value"`
		} `json:"outKeywordDict"`
		To struct {
			CityCode          string `json:"cityCode"`
			CityName          string `json:"cityName"`
			Code              int    `json:"code"`
			DiseaseControlTel string `json:"diseaseControlTel"`
			HighInDesc        string `json:"highInDesc"`
			IsInUpdate        bool   `json:"isInUpdate"`
			IsOutUpdate       bool   `json:"isOutUpdate"`
			Labels            []struct {
				Colour   string `json:"colour"`
				IsJump   bool   `json:"isJump"`
				IsShow   bool   `json:"isShow"`
				JumpLink struct {
					AppId  string `json:"appId"`
					AppVer string `json:"appVer"`
					Type   int    `json:"type"`
					Url    string `json:"url"`
				} `json:"jumpLink"`
				Label string `json:"label"`
			} `json:"labels"`
			LowInDesc    string `json:"lowInDesc"`
			MedianInDesc string `json:"medianInDesc"`
			Message      string `json:"message"`
			News         struct {
				Context  string `json:"context"`
				IsShow   bool   `json:"isShow"`
				JumpLink struct {
					AppId  string `json:"appId"`
					AppVer string `json:"appVer"`
					Type   int    `json:"type"`
					Url    string `json:"url"`
				} `json:"jumpLink"`
			} `json:"news"`
			OutDesc       string `json:"outDesc"`
			ProvinceName  string `json:"provinceName"`
			RiskLevel     int    `json:"riskLevel"`
			SourceContext string `json:"sourceContext"`
			SourceLink    struct {
				AppId  string `json:"appId"`
				AppVer string `json:"appVer"`
				Type   int    `json:"type"`
				Url    string `json:"url"`
			} `json:"sourceLink"`
		} `json:"to"`
	} `json:"rsp"`
	ThRequestId string `json:"th-request-id"`
}

type TravelReq struct {
	Request struct {
		Req struct {
			FromCityCode string `json:"fromCityCode"`
			ToCityCode   string `json:"toCityCode"`
		} `json:"req"`
	} `json:"request"`
	Service string `json:"service"`
	Func    string `json:"func"`
	Context struct {
		UserId string `json:"userId"`
	} `json:"context"`
}

// Get 获取城市列表
func (t *Travel) Get(FromCityCode, ToCityCode string) (Travel, error) {
	// 构造请求
	req := TravelReq{
		Request: struct {
			Req struct {
				FromCityCode string `json:"fromCityCode"`
				ToCityCode   string `json:"toCityCode"`
			} `json:"req"`
		}{Req: struct {
			FromCityCode string `json:"fromCityCode"`
			ToCityCode   string `json:"toCityCode"`
		}{
			FromCityCode: FromCityCode,
			ToCityCode:   ToCityCode,
		}},
		Service: "TravelPrevention",
		Func:    "getTravelPrevent",
		Context: struct {
			UserId string `json:"userId"`
		}{UserId: utils.Algorithm.UUID2()},
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
		Post("https://wechat.wecity.qq.com/trpcapi/TravelPrevention/getTravelPrevent")

	if err != nil {
		global.LOG.Error("接口#腾讯-获取出行政策请求失败", zap.Error(err))
		return Travel{}, err
	}

	// 解析数据
	var d Travel
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
