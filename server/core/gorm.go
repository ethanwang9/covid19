package core

import (
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// name: 数据库配置
// author: Ethan.Wang
// desc: 使用GORM对数据库进行连接

// InitGorm 初始化 GORM
func InitGorm() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			global.CONFIG.GetString("database.username"),
			global.CONFIG.GetString("database.password"),
			global.CONFIG.GetString("database.host"),
			global.CONFIG.GetString("database.port"),
			global.CONFIG.GetString("database.dbname"),
		),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "c19_",
			SingularTable: true, //禁用表名复数
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Millisecond,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Silent,
			},
		),
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		global.LOG.Panic("数据库MySQL-初始化连接失败", zap.String("error", err.Error()))
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(global.CONFIG.GetInt("database.max_idle_conns"))
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(global.CONFIG.GetInt("database.max_open_conns"))
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		global.LOG.Panic("数据库MySQL-连接池初始化失败", zap.String("error", err.Error()))
	}

	return db
}
