package Initialization

// name: 初始化项目
// author: Ethan.Wang
// desc:

// InitApp 初始化程序
func InitApp() {
	// 初始化配置
	initConfig()
	// 初始化数据库
	initDatabase()
	// 初始化接口
	initApi()
}
