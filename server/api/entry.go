package api

import (
	"github.com/ethanwang9/covid19/server/api/ali"
	"github.com/ethanwang9/covid19/server/api/dxy"
	"github.com/ethanwang9/covid19/server/api/gov"
	"github.com/ethanwang9/covid19/server/api/internet"
	"github.com/ethanwang9/covid19/server/api/tencent"
)

// name: 接口入口
// author: Ethan.Wang
// desc:

type Api struct {
	// 网络接口
	Internet internet.Internet
	// 腾讯接口
	Tencent tencent.Tencent
	// 国家卫健委
	Gov gov.Gov
	// 丁香园
	DXY dxy.DXY
	// 阿里Geo地图
	Ali ali.Ali
}

var ApiApp = new(Api)
