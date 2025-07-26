package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/ethanwang9/covid19/server/global"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"strings"
)

// name: 算法
// author: Ethan.Wang
// desc:

type algorithm struct{}

var Algorithm = new(algorithm)

// UUID 生成 UUID
func (a *algorithm) UUID() (u string) {
	u = uuid.NewV4().String()
	u = strings.Replace(u, "-", "", -1)
	return u
}

// UUID2 生成 UUID 带 -
func (a *algorithm) UUID2() (u string) {
	u = uuid.NewV4().String()
	return u
}

// Base64Encode Base64 编码
func (a *algorithm) Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Base64Decode Base64 解码
func (a *algorithm) Base64Decode(dst string) (src []byte, err error) {
	src, err = base64.StdEncoding.DecodeString(dst)
	if err != nil {
		global.LOG.Warn(
			"工具类-base64解码失败",
			zap.String("dst", dst),
			zap.String("error", err.Error()),
		)
		return nil, err
	}
	return src, nil
}

// MD5 MD5加密
func (a *algorithm) MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 Sha1加密
func (a *algorithm) Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// Sha256 Sha256加密
func (a *algorithm) Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
