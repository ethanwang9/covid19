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

// name: 风险地区
// author: Ethan.Wang
// desc:

type DengerArea struct {
	Data struct {
		EndUpdateTime string `json:"end_update_time"`
		Hcount        int    `json:"hcount"`
		Mcount        int    `json:"mcount"`
		Lcount        int    `json:"lcount"`
		Highlist      []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communitys []string `json:"communitys"`
		} `json:"highlist"`
		Middlelist []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communitys []string `json:"communitys"`
		} `json:"middlelist"`
		Lowlist []struct {
			Type       string   `json:"type"`
			Province   string   `json:"province"`
			City       string   `json:"city"`
			County     string   `json:"county"`
			AreaName   string   `json:"area_name"`
			Communitys []string `json:"communitys"`
		} `json:"lowlist"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DengerAreaReq struct {
	Key             string `json:"key"`
	AppId           string `json:"appId"`
	PaasHeader      string `json:"paasHeader"`
	TimestampHeader string `json:"timestampHeader"`
	NonceHeader     string `json:"nonceHeader"`
	SignatureHeader string `json:"signatureHeader"`
}

// Get 获取出行政策
func (t *DengerArea) Get() (DengerArea, error) {
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
			"referer":         "http://bmfw.www.gov.cn/yqfxdjcx/risk.html",
			"content-type":    "application/json",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.56",
			"x-wif-nonce":     "QkjjtiLM2dCratiA",
			"x-wif-paasid":    "smt-application",
			"x-wif-signature": xWifSignature,
			"x-wif-timestamp": fmt.Sprintf("%v", ms),
		}).
		SetBody(DengerAreaReq{
			Key:             "3C502C97ABDA40D0A60FBEE50FAAD1DA",
			AppId:           "NcApplication",
			PaasHeader:      "zdww",
			TimestampHeader: fmt.Sprintf("%v", ms),
			NonceHeader:     nonceHeader,
			SignatureHeader: signatureHeader,
		}).
		Post("http://bmfw.www.gov.cn/bjww/interface/interfaceJson")

	if err != nil {
		global.LOG.Error("接口#gov-获取出行政策请求失败", zap.Error(err))
		return DengerArea{}, err
	}

	// 解析数据
	var d DengerArea
	json.Unmarshal(resp.Body(), &d)

	return d, nil
}
