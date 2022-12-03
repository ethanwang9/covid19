package dxy

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"strings"
)

// name: 确诊信息
// author: Ethan.Wang
// desc:

type AreaStat struct {
	ProvinceName          string `json:"provinceName"`
	ProvinceShortName     string `json:"provinceShortName"`
	CurrentConfirmedCount int    `json:"currentConfirmedCount"`
	ConfirmedCount        int    `json:"confirmedCount"`
	SuspectedCount        int    `json:"suspectedCount"`
	CuredCount            int    `json:"curedCount"`
	DeadCount             int    `json:"deadCount"`
	Comment               string `json:"comment"`
	LocationId            int    `json:"locationId"`
	StatisticsData        string `json:"statisticsData"`
	HighDangerCount       int    `json:"highDangerCount"`
	MidDangerCount        int    `json:"midDangerCount"`
	DetectOrgCount        int    `json:"detectOrgCount"`
	VaccinationOrgCount   int    `json:"vaccinationOrgCount"`
	Cities                []struct {
		CityName                 string `json:"cityName"`
		CurrentConfirmedCount    int    `json:"currentConfirmedCount"`
		ConfirmedCount           int    `json:"confirmedCount"`
		SuspectedCount           int    `json:"suspectedCount"`
		CuredCount               int    `json:"curedCount"`
		DeadCount                int    `json:"deadCount"`
		HighDangerCount          int    `json:"highDangerCount"`
		MidDangerCount           int    `json:"midDangerCount"`
		LocationId               int    `json:"locationId"`
		CurrentConfirmedCountStr string `json:"currentConfirmedCountStr"`
	} `json:"cities"`
	DangerAreas []struct {
		CityName    string `json:"cityName"`
		AreaName    string `json:"areaName"`
		DangerLevel int    `json:"dangerLevel"`
	} `json:"dangerAreas"`
}

// Get 获取确诊信息
func (a *AreaStat) Get() ([]AreaStat, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"referer":    "https://ncov.dxy.cn/ncovh5/view/pneumonia?share=0&source=appshare",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
		}).
		Get("https://ncov.dxy.cn/ncovh5/view/pneumonia?share=0&source=appshare")

	if err != nil {
		global.LOG.Error("接口#丁香园-确诊信息请求失败", zap.Error(err))
		return []AreaStat{}, err
	}

	// 解析数据
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		global.LOG.Error("接口#丁香园-确诊信息#解析为DOM数据失败", zap.Error(err))
		return []AreaStat{}, err
	}

	var data []AreaStat
	dom.Find("script[id='getAreaStat']").Each(func(i int, selection *goquery.Selection) {
		tempStr := selection.Text()
		json.Unmarshal([]byte(tempStr[27:len(tempStr)-11]), &data)
	})

	return data, nil
}
