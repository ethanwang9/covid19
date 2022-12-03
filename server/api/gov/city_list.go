package gov

import (
	"encoding/json"
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"strings"
	"time"
)

// name: 城市列表
// author: Ethan.Wang
// desc:

type CityList struct {
	Data struct {
		List []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Key   string `json:"key"`
			Items []struct {
				Code  string `json:"code"`
				Name  string `json:"name"`
				Key   string `json:"key"`
				Items []struct {
					Code  string        `json:"code"`
					Name  string        `json:"name"`
					Key   string        `json:"key"`
					Items []interface{} `json:"items"`
				} `json:"items"`
			} `json:"items"`
		} `json:"list"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CityListReq struct {
	Flag            string `json:"flag"`
	Key             string `json:"key"`
	AppId           string `json:"appId"`
	PaasHeader      string `json:"paasHeader"`
	TimestampHeader string `json:"timestampHeader"`
	NonceHeader     string `json:"nonceHeader"`
	SignatureHeader string `json:"signatureHeader"`
}

// Get 获取城市列表
func (t *CityList) Get() (CityList, error) {
	// 准备参数
	ms := time.Now().Unix()
	xWifSignature := strings.ToUpper(utils.Algorithm.Sha256(fmt.Sprintf("%v%v%v", ms, "fTN2pfuisxTavbTuYVSsNJHetwq5bJvCQkjjtiLM2dCratiA", ms)))
	nonceHeader := "123456789abcdefg"
	signatureHeader := strings.ToUpper(utils.Algorithm.Sha256(fmt.Sprintf("%v%v%v%v", ms, "23y0ufFl5YxIyGrI8hWRUZmKkvtSjLQA", nonceHeader, ms)))

	// 发起请求
	client := resty.New()
	resp, err := client.R().
		SetHeaders(map[string]string{
			"host":            "bmfw.www.gov.cn",
			"origin":          "http://bmfw.www.gov.cn",
			"referer":         "http://bmfw.www.gov.cn/yqfxdjcx/index.html",
			"content-type":    "application/json",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
			"x-wif-nonce":     "QkjjtiLM2dCratiA",
			"x-wif-paasid":    "smt-application",
			"x-wif-signature": xWifSignature,
			"x-wif-timestamp": fmt.Sprintf("%v", ms),
		}).
		SetBody(CityListReq{
			Flag:            "11",
			Key:             "243D215B2CA449ECABF1E6C93B7D973C",
			AppId:           "NcApplication",
			PaasHeader:      "zdww",
			TimestampHeader: fmt.Sprintf("%v", ms),
			NonceHeader:     nonceHeader,
			SignatureHeader: signatureHeader,
		}).
		Post("http://bmfw.www.gov.cn/bjww/interface/interfaceJson")

	if err != nil {
		global.LOG.Error("接口#gov-获取城市列表请求失败", zap.Error(err))
		return CityList{}, err
	}

	// 解析数据
	var d CityList
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
