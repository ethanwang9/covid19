package utils

import (
	"errors"
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"sort"
	"strconv"
	"time"
)

// name: 安全校验
// author: Ethan.Wang
// desc: 校验sign, 参数信息等

type Safe struct{}

var SafeApp = new(Safe)

// SafeVerify 验证签名
func (s *Safe) SafeVerify(data map[string]interface{}) error {
	if global.CONFIG.GetString("server.env") == "prod" {
		// 校验参数
		// 1. 参数是否为空
		for _, v := range data {
			if td := fmt.Sprintf("%v", v); len(td) == 0 {
				return errors.New("请求参数不能为空")
			}
		}

		// 校验请求时间 15秒内, 不超过3秒
		tTimeString := fmt.Sprintf("%v", data["timestamp"])
		tTime, _ := strconv.ParseInt(tTimeString, 10, 64)
		if tTime < time.Now().Unix()-15 && tTime > time.Now().Unix()+3 {
			return errors.New("请求超时")
		}

		// 校验签名
		sign1 := fmt.Sprintf("%v", data["sign"])
		sign2 := s.SignCreate(data)
		if sign1 != sign2 {
			return errors.New("签名验证失败")
		}
	}

	return nil
}

// SignCreate 生成签名
func (s *Safe) SignCreate(data map[string]interface{}) string {
	// 删除sign
	delete(data, "sign")

	// 获取keys
	keys := make([]string, 0)
	for k, _ := range data {
		keys = append(keys, k)
	}

	// keys排序
	sort.Strings(keys)

	// 构造sign
	var signString string
	for _, v := range keys {
		for k2, v2 := range data {
			if v == k2 {
				signString += v + fmt.Sprintf("%v", v2)
			}
		}
	}

	sign := Algorithm.MD5(signString + global.CONFIG.GetString("safe.sign"))

	// 加密
	return sign
}
