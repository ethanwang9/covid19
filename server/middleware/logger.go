package middleware

import (
	"bytes"
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
	"time"
)

// name: 中间件-处理日志
// author: Ethan.Wang, GVA [github.com/flipped-aurora/gin-vue-admin]
// desc: 中间件处理日志，记录http请求, 修改自 GVA

// LogLayout 日志layout
type LogLayout struct {
	Method string              // 请求类型
	Path   string              // 访问路径
	Query  string              // Get请求参数
	Header map[string][]string // 请求头
	Body   string              // 请求body参数
	IP     string              // ip地址
	Error  string              // 错误
	Cost   time.Duration       // 花费时间
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 访问路径
		path := c.Request.URL.Path
		// Get请求参数
		query := c.Request.URL.RawQuery
		// 请求body参数
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// 请求类型
		method := c.Request.Method

		c.Next()

		// 花费时间
		cost := time.Since(start)
		// 状态码
		status := strconv.Itoa(c.Writer.Status())

		// 写入日志
		global.LOG.Info(
			fmt.Sprintf("%s %s", status, method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Any("header", c.Request.Header),
			zap.ByteString("body", body),
			zap.String("ip", c.ClientIP()),
			zap.String("error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)

	}
}
