package Initialization

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/model"
	"go.uber.org/zap"
	"os"
	"strings"
)

// name: 初始化数据库
// author: Ethan.Wang
// desc:

// 初始化数据库
func initDatabase() {
	// 获取数据库内容
	_, err := model.SysApp.New(model.Sys{}).Get()
	if err != nil && err.Error() == "Error 1146: Table 'c19.c19_sys' doesn't exist" {
		// 写入 sql 文件
		if err := runSql(); err != nil {
			global.LOG.Error("初始化数据库#运行数据库文件失败", zap.Error(err))
			return
		}
	} else if err != nil {
		global.LOG.Error("初始化数据库#获取数据库失败", zap.Error(err))
		return
	}
}

// 执行sql文件
func runSql() error {
	sqls, err := os.ReadFile(global.DBFilePath)
	if err != nil {
		return err
	}
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		if sql == "" {
			continue
		}
		global.DB.Exec(sql)
	}
	return err
}
