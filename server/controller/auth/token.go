package auth

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetToken 微信消息验证
func GetToken(ctx *gin.Context) {
	//获取参数
	signature := ctx.Query("signature")
	timestamp := ctx.Query("timestamp")
	nonce := ctx.Query("nonce")
	echostr := ctx.Query("echostr")

	if global.MP.VerifyEventSign(signature, timestamp, nonce, global.CONFIG.GetString("mp.token")) {
		ctx.String(http.StatusOK, echostr)
	}
}
