package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/shenghui0779/gochat/offia"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// name: 程序变量
// author: Ethan.Wang
// desc: 程序运行主要变量

var (
	CONFIG *viper.Viper
	LOG    *zap.Logger
	DB     *gorm.DB
	MP     *offia.Offia
	CRON   *cron.Cron
	REDIS  *redis.Client
)
