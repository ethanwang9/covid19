package internet

import (
	"encoding/json"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// name: ip 接口
// author: Ethan.Wang
// desc:

type IP struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Isp      string `json:"isp"`
	Location string `json:"location"`
	Ip       string `json:"ip"`
}

// New 初始化
func (i *IP) New(ip IP) *IP {
	return &ip
}

// Query 查询IP地址
func (i *IP) Query() (IP, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"host":       "restapi.amap.com",
			"origin":     "https://widget.codelife.cc",
			"referer":    "https://widget.codelife.cc",
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
		}).
		SetQueryParams(map[string]string{
			"type": "4",
			"key":  "e1d86a0fa7f8d42d2add26e11a050e25",
			"ip":   i.Ip,
		}).
		Get("https://restapi.amap.com/v5/ip")

	if err != nil {
		global.LOG.Error("接口#网络接口-查询IP地址请求失败", zap.Error(err), zap.String("ip", i.Ip))
		return IP{}, err
	}

	// 解析数据
	var d IP
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
