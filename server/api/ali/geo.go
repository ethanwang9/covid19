package ali

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// name: Geo 地图
// author: Ethan.Wang
// desc:

type Geo struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Adcode      interface{} `json:"adcode"`
			Name        string      `json:"name"`
			Center      []float64   `json:"center,omitempty"`
			Centroid    []float64   `json:"centroid,omitempty"`
			ChildrenNum int         `json:"childrenNum,omitempty"`
			Level       string      `json:"level,omitempty"`
			Parent      struct {
				Adcode int `json:"adcode"`
			} `json:"parent,omitempty"`
			SubFeatureIndex int    `json:"subFeatureIndex,omitempty"`
			Acroutes        []int  `json:"acroutes,omitempty"`
			Adchar          string `json:"adchar,omitempty"`
		} `json:"properties"`
		Geometry struct {
			Type        string            `json:"type"`
			Coordinates [][][]interface{} `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

// Get 获取信息
func (g *Geo) Get(province string) (Geo, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
		}).
		Get("http://geo.datav.aliyun.com/areas_v3/bound/" + province + "_full.json")

	if err != nil {
		global.LOG.Error("接口#阿里-Geo地图请求失败", zap.Error(err))
		return Geo{}, err
	}

	var d Geo
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
