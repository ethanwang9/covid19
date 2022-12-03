package main

import (
	"fmt"
	"github.com/ethanwang9/covid19/server/Initialization"
	"github.com/ethanwang9/covid19/server/core"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/router"
	"go.uber.org/zap"
)

// name: 主程序
// author: Ethan.Wang
// desc: 程序运行主要入口

func main() {
	// 初始化配置文件
	global.CONFIG = core.InitViper()
	// 初始化日志
	global.LOG = core.InitZap()
	// 初始化数据库
	global.DB = core.InitGorm()
	// 初始化Redis
	global.REDIS = core.InitRedis()
	defer global.REDIS.Close()
	// 初始化微信公众号配置
	global.MP = core.InitMP()
	// 初始化定时任务
	global.CRON = core.InitCron()
	defer global.CRON.Stop()
	// 初始化路由
	r := router.Init()

	// 启动定时任务
	core.CronRun()

	// 初始化程序
	Initialization.InitApp()

	// 运行日志
	global.LOG.Info("系统已开始运行", zap.String("port", global.CONFIG.GetString("server.port")))

	// 运行
	r.Run(fmt.Sprintf(":%v", global.CONFIG.GetString("server.port")))
}
