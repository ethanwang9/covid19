package tencent

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// name: 热点消息
// author: Ethan.Wang
// desc:

type HotMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Rsp  struct {
		ContentMap struct {
		} `json:"contentMap"`
		HotnewsRsp struct {
			Code     int `json:"code"`
			Contents []struct {
				Area       string `json:"area"`
				AreaCode   string `json:"areaCode"`
				CityCode   string `json:"cityCode"`
				ContentTag string `json:"contentTag"`
				Desc       string `json:"desc"`
				From       string `json:"from"`
				FromDesc   string `json:"fromDesc"`
				Garea      string `json:"garea"`
				Id         int    `json:"id"`
				ImgUrl     string `json:"imgUrl"`
				JumpLink   struct {
					AppId  string `json:"appId"`
					AppVer string `json:"appVer"`
					Type   int    `json:"type"`
					Url    string `json:"url"`
				} `json:"jumpLink"`
				PublicTime    string        `json:"publicTime"`
				Tags          []interface{} `json:"tags"`
				TimeShaftDesc string        `json:"timeShaftDesc"`
				Title         string        `json:"title"`
			} `json:"contents"`
			Message string `json:"message"`
			Result  struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			} `json:"result"`
			TotalCnt  int    `json:"totalCnt"`
			UiMessage string `json:"uiMessage"`
		} `json:"hotnewsRsp"`
		Result struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"result"`
	} `json:"rsp"`
	ThRequestId string `json:"th-request-id"`
}

type HotMessageReq struct {
	Request struct {
		Req struct {
			AreaCode  string `json:"areaCode"`
			QueryList []struct {
			} `json:"queryList"`
			HotnewsReq struct {
				ReqType      int    `json:"reqType"`
				LocationCode string `json:"locationCode"`
				Offset       int    `json:"offset"`
				Limit        int    `json:"limit"`
				Tab          string `json:"tab"`
			} `json:"hotnewsReq"`
		} `json:"req"`
	} `json:"request"`
	Service string `json:"service"`
	Func    string `json:"func"`
	Context struct {
		UserId string `json:"userId"`
	} `json:"context"`
}

// Get 获取热点消息
//
// value:
// 		areaCode<string> 省级地区代码
func (h *HotMessage) Get(areaCode string) (HotMessage, error) {
	// 构造请求
	req := HotMessageReq{
		Request: struct {
			Req struct {
				AreaCode   string     `json:"areaCode"`
				QueryList  []struct{} `json:"queryList"`
				HotnewsReq struct {
					ReqType      int    `json:"reqType"`
					LocationCode string `json:"locationCode"`
					Offset       int    `json:"offset"`
					Limit        int    `json:"limit"`
					Tab          string `json:"tab"`
				} `json:"hotnewsReq"`
			} `json:"req"`
		}{
			Req: struct {
				AreaCode   string     `json:"areaCode"`
				QueryList  []struct{} `json:"queryList"`
				HotnewsReq struct {
					ReqType      int    `json:"reqType"`
					LocationCode string `json:"locationCode"`
					Offset       int    `json:"offset"`
					Limit        int    `json:"limit"`
					Tab          string `json:"tab"`
				} `json:"hotnewsReq"`
			}(struct {
				AreaCode   string
				QueryList  []struct{}
				HotnewsReq struct {
					ReqType      int
					LocationCode string
					Offset       int
					Limit        int
					Tab          string
				}
			}{
				AreaCode: areaCode,
				HotnewsReq: struct {
					ReqType      int
					LocationCode string
					Offset       int
					Limit        int
					Tab          string
				}{
					ReqType:      1,
					LocationCode: areaCode,
					Offset:       3,
					Limit:        5,
					Tab:          "shishitongbao",
				},
			}),
		},
		Service: "THPneumoniaOuterService",
		Func:    "getTopicContent",
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
		Post("https://wechat.wecity.qq.com/trpcapi/THPneumoniaOuterService/getTopicContent")

	if err != nil {
		global.LOG.Error("接口#腾讯-热点消息-获取热点消息请求失败", zap.Error(err))
		return HotMessage{}, err
	}

	// 解析数据
	var d HotMessage
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
