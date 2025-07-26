package internal

import (
	"context"
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
	"time"
)

// name: 定时任务
// author: Ethan.Wang
// desc: 定时任务具体执行

type cronToDo struct{}

var CronToDo = new(cronToDo)

// New 初始化实例
func (d *cronToDo) New() *cronToDo {
	return &cronToDo{}
}

// Run 运行
func (d *cronToDo) Run() {
	global.CRON.Start()
}

// TODO 任务列表
func (d *cronToDo) TODO() *cronToDo {
	// 获取微信公众号AccessToken
	global.CRON.AddFunc("0 0 */1 * * ?", d.GetMpAccessToken)

	// 获取 腾讯-城市列表 [出行政策附属接口]
	global.CRON.AddFunc("0 0 10 * * ?", d.GetTencentCityList)

	// 获取 国家卫健委-风险地区
	global.CRON.AddFunc("0 0 8,9,11,13,15,18,20,21 * * ?", d.GetGovDangerArea)

	// 获取 国家卫健委-城市列表
	global.CRON.AddFunc("0 0 9 * * ?", d.GetGovCityList)

	// 获取 丁香园-确诊数据
	global.CRON.AddFunc("38 24 7,9,11,13,18,19,21,22 * * ? ", d.GetDXYAreaStat)

	return d
}

// GetMpAccessToken 获取微信公众号AccessToken
func (d *cronToDo) GetMpAccessToken() {
	// 获取微信公众号AccessToken
	token, err := global.MP.AccessToken(context.Background())
	if err != nil {
		global.LOG.Error("定时任务#获取微信公众号AccessToken失败", zap.String("error", err.Error()))
		return
	}

	// 写入微信公众号AccessToken到Redis中失败
	err = global.REDIS.Set(context.Background(), "WxMpAccessToken", token.Token, 2*time.Hour).Err()
	if err != nil {
		global.LOG.Error("定时任务#写入微信公众号AccessToken到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// GetTencentCityList 获取 腾讯-城市列表 [出行政策附属接口]
func (d *cronToDo) GetTencentCityList() {
	// 获取数据
	data, err := api.ApiApp.Tencent.CityList.Get()
	if err != nil {
		global.LOG.Error("定时任务#腾讯-城市列表 [出行政策附属接口]-获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "TencentCityList", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("定时任务#腾讯-城市列表 [出行政策附属接口]-写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// GetGovDangerArea 获取 国家卫健委-风险地区
func (d *cronToDo) GetGovDangerArea() {
	// 获取数据
	data, err := api.ApiApp.Gov.DengerArea.Get()
	if err != nil {
		global.LOG.Error("定时任务#国家卫健委-风险地区#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "GovDangerArea", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("定时任务#国家卫健委-风险地区#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// GetGovCityList 获取 国家卫健委-城市列表
func (d *cronToDo) GetGovCityList() {
	// 获取数据
	data, err := api.ApiApp.Gov.CityList.Get()
	if err != nil {
		global.LOG.Error("定时任务#国家卫健委-城市列表#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "GovCityList", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("定时任务#国家卫健委-城市列表#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// GetDXYAreaStat 获取 丁香园-确诊数据
func (d *cronToDo) GetDXYAreaStat() {
	// 获取数据
	data, err := api.ApiApp.DXY.AreaStat.Get()
	if err != nil {
		global.LOG.Error("定时任务#丁香园-确诊数据#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "DXYAreaStat", string(writeStr), 12*time.Hour).Err()
	if err != nil {
		global.LOG.Error("定时任务#丁香园-确诊数据#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}
