package public

import (
	"github.com/ethanwang9/covid19/server/global"
	"github.com/ethanwang9/covid19/server/model"
	"github.com/ethanwang9/covid19/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// name: 公开信息
// author: Ethan.Wang
// desc:

func GetInfo(ctx *gin.Context) {
	// 获取参数
	timestamp := ctx.Query("timestamp")
	sign := ctx.Query("sign")

	// 校验签名
	err := utils.SafeApp.SafeVerify(gin.H{
		"timestamp": timestamp,
		"sign":      sign,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorByNotTrue,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 获取数据库信息
	info, err := model.SysApp.New(model.Sys{}).Get()
	if err != nil {
		ctx.JSON(http.StatusOK, global.MsgBack{
			Code:    global.CodeErrorBySQL,
			Message: "获取数据库内容失败",
			Data:    nil,
		})
		return
	}

	info.Id = ""

	ctx.JSON(http.StatusOK, global.MsgBack{
		Code:    global.CodeSuccess,
		Message: "success",
		Data:    info,
	})
}
