package Initialization

import (
	"context"
	"encoding/json"
	"github.com/ethanwang9/covid19/server/api"
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
	"time"
)

// name: 初始化接口
// author: Ethan.Wang
// desc:

// 初始化接口
func initApi() {
	// 初始化腾讯-城市列表 [出行政策附属接口]
	tencentCityList()
	// 初始化国家卫健委-风险地区
	govDangerArea()
	// 初始化国家卫健委-城市列表
	govCityList()
	// 初始化丁香园-确诊数据
	dxyAreaStat()
}

// 初始化腾讯-城市列表 [出行政策附属接口]
func tencentCityList() {
	// 判断是否有 TencentCityList
	res, err := global.REDIS.Get(context.Background(), "TencentCityList").Result()
	if err != nil && err.Error() != "redis: nil" {
		global.LOG.Error("初始化腾讯-城市列表 [出行政策附属接口]#获取redis数据失败",
			zap.String("key", "TencentCityList"),
			zap.String("error", err.Error()),
		)
		return
	}
	if len(res) != 0 {
		return
	}

	// 获取数据
	data, err := api.ApiApp.Tencent.CityList.Get()
	if err != nil {
		global.LOG.Error("初始化腾讯-城市列表 [出行政策附属接口]#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "TencentCityList", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("初始化腾讯-城市列表 [出行政策附属接口]#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// 初始化国家卫健委-风险地区
func govDangerArea() {
	// 判断是否有 GovDangerArea
	res, err := global.REDIS.Get(context.Background(), "GovDangerArea").Result()
	if err != nil && err.Error() != "redis: nil" {
		global.LOG.Error("初始化国家卫健委-风险地区#获取redis数据失败",
			zap.String("key", "GovDangerArea"),
			zap.String("error", err.Error()),
		)
		return
	}
	if len(res) != 0 {
		return
	}

	// 获取数据
	data, err := api.ApiApp.Gov.DengerArea.Get()
	if err != nil {
		global.LOG.Error("初始化国家卫健委-风险地区#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "GovDangerArea", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("初始化国家卫健委-风险地区#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// 初始化国家卫健委-城市列表
func govCityList() {
	// 判断是否有 GovCityList
	res, err := global.REDIS.Get(context.Background(), "GovCityList").Result()
	if err != nil && err.Error() != "redis: nil" {
		global.LOG.Error("初始化国家卫健委-城市列表#获取redis数据失败",
			zap.String("key", "GovCityList"),
			zap.String("error", err.Error()),
		)
		return
	}
	if len(res) != 0 {
		return
	}

	// 获取数据
	data, err := api.ApiApp.Gov.CityList.Get()
	if err != nil {
		global.LOG.Error("初始化国家卫健委-城市列表#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "GovCityList", string(writeStr), 24*time.Hour).Err()
	if err != nil {
		global.LOG.Error("初始化国家卫健委-城市列表#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}

// 初始化丁香园-确诊数据
func dxyAreaStat() {
	// 判断是否有 DXYAreaStat
	res, err := global.REDIS.Get(context.Background(), "DXYAreaStat").Result()
	if err != nil && err.Error() != "redis: nil" {
		global.LOG.Error("初始化丁香园-确诊数据#获取redis数据失败",
			zap.String("key", "DXYAreaStat"),
			zap.String("error", err.Error()),
		)
		return
	}
	if len(res) != 0 {
		return
	}

	// 获取数据
	data, err := api.ApiApp.DXY.AreaStat.Get()
	if err != nil {
		global.LOG.Error("初始化丁香园-确诊数据#获取网络请求失败", zap.String("error", err.Error()))
		return
	}

	// 写入内容
	writeStr, _ := json.Marshal(data)
	err = global.REDIS.Set(context.Background(), "DXYAreaStat", string(writeStr), 12*time.Hour).Err()
	if err != nil {
		global.LOG.Error("初始化丁香园-确诊数据#写入数据到Redis中失败", zap.String("error", err.Error()))
		return
	}
}
