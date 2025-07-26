package core

import (
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/spf13/viper"
)

// name: 读取配置
// author: Ethan.Wang
// desc: 读取配置文件数据

// InitViper 初始化 Viper
func InitViper() *viper.Viper {
	// 判断是否有配置文件
	if f, _ := utils.PathExists(fmt.Sprintf("%v%v.%v", global.ConfigPath, global.ConfigFileName, global.ConfigFileSuffix)); !f {
		panic("初始化 [Viper] 失败, 配置文件config/config.ini找不到")
	}

	v := viper.New()
	// 设置配置文件
	v.SetConfigName(global.ConfigFileName)
	v.SetConfigType(global.ConfigFileSuffix)
	v.AddConfigPath(global.ConfigPath)
	// 读取文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("初始化 [Viper] 出现错误, Error: %s\n", err))
	}
	// 监听文件
	v.WatchConfig()
	// 文件改变
	//v.OnConfigChange(func(e fsnotify.Event) {
	//fmt.Println("[Viper] 文件改变", e.Name)
	//})

	return v
}
