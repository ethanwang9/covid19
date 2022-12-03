package global

// name: 程序常量
// author: Ethan.Wang
// desc: 程序运行主要常量

// 配置文件

const (
	// ConfigPath 配置文件目录
	ConfigPath = "./config/"
	// ConfigFileName 配置文件名
	ConfigFileName = "config"
	// ConfigFileSuffix 配置文件后缀
	ConfigFileSuffix = "ini"
	// DBFilePath 数据库文件位置
	DBFilePath = "./config/c19.sql"
)

// 缓存

const (
	// TempPath 缓存文件路径
	TempPath = "./temp/"
)

// 接口消息状态 - 成功

const (
	// CodeSuccess 请求成功
	CodeSuccess = 200
)

// 接口消息状态 - 失败

const (
	// CodeErrorBySystem 系统处理消息失败
	CodeErrorBySystem = 301
	// CodeErrorByAPI 接口请求失败
	CodeErrorByAPI = 302
	// CodeErrorBySQL 数据库操作错误
	CodeErrorBySQL = 303
	// CodeErrorByNotTrue 请求不合规
	CodeErrorByNotTrue = 304
	// CodeErrorByAuth 认证失败
	CodeErrorByAuth = 305
	// CodeErrorByLevel 权限不足
	CodeErrorByLevel = 306
)

// 用户权限等级

const (
	// UserLevelByStop 用户已被禁用
	UserLevelByStop = -99
	// UserLevelByUser 用户
	UserLevelByUser = 1
	// UserLevelByAdmin 管理员
	UserLevelByAdmin = 99
)
