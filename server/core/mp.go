package core

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/shenghui0779/gochat"
	"github.com/shenghui0779/gochat/offia"
)

// name: 微信公众号SDK
// author: Ethan.Wang
// desc:

// InitMP 初始化微信公众号配置
func InitMP() *offia.Offia {
	mp := gochat.NewOffia(
		global.CONFIG.GetString("mp.appid"),
		global.CONFIG.GetString("mp.secret"),
	)
	return mp
}
