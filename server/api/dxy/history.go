package dxy

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

// name: 获取历史数据
// author: Ethan.Wang
// desc:

type History struct {
	Code string `json:"code"`
	Data []struct {
		ConfirmedCount        int `json:"confirmedCount"`
		ConfirmedIncr         int `json:"confirmedIncr"`
		CuredCount            int `json:"curedCount"`
		CuredIncr             int `json:"curedIncr"`
		CurrentConfirmedCount int `json:"currentConfirmedCount"`
		CurrentConfirmedIncr  int `json:"currentConfirmedIncr"`
		DateId                int `json:"dateId"`
		DeadCount             int `json:"deadCount"`
		DeadIncr              int `json:"deadIncr"`
		HighDangerCount       int `json:"highDangerCount"`
		MidDangerCount        int `json:"midDangerCount"`
		SuspectedCount        int `json:"suspectedCount"`
		SuspectedCountIncr    int `json:"suspectedCountIncr"`
	} `json:"data"`
}

// Get 获取数据
func (h History) Get(url string) (History, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"referer":    "https://ncov.dxy.cn/ncovh5/view/pneumonia?share=0&source=appshare",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
		}).
		Get(url)

	if err != nil {
		global.LOG.Error("接口#丁香园-历史数据请求失败", zap.Error(err))
		return History{}, err
	}

	// 解析数据
	var data History
	json.Unmarshal(resp.Body(), &data)

	return data, nil
}
