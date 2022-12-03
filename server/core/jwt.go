package core

import (
	"errors"
	"fmt"
	"github.com/ethanwang9/covid19/server/global"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

// name: JWT 无状态认证
// author: Ethan.Wang
// desc:

type jwtToken struct{}

var JwtApp = new(jwtToken)

// Generate 生成
func (j *jwtToken) Generate(uuid string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"token": uuid,
	}).SignedString([]byte(global.CONFIG.GetString("safe.jwt")))
	if err != nil {
		global.LOG.Error("jwt-生成失败", zap.String("error", err.Error()))
		return "", err
	}

	return token, nil
}

// Decode 解密
func (j *jwtToken) Decode(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			global.LOG.Error("jwt-加密失败", zap.Any("error", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(global.CONFIG.GetString("safe.jwt")), nil
	})

	// 错误处理
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 不是 token
				return "", errors.New("token 已过期")
			} else {
				// 无法处理token
				return "", errors.New("token 出现错误")
			}
		}
	}

	// 返回机密结果
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 解密成功
		return claims["token"].(string), nil
	} else {
		return "", errors.New("token 处理失败")
	}
}
