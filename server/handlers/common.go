package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/models"
)

func SendResponse(ctx *gin.Context, code int, msg, detail string) {
	resp := &models.Response{
		Code:    code,
		Message: msg,
		Detail:  detail,
	}
	if resp.Code >= 400 {
		ctx.Abort()
	}
	ctx.JSON(resp.Code, resp)
}

func SendNormalResponse(ctx *gin.Context, data interface{}) {
	resp := &models.Response{
		Code:    200,
		Message: "OK",
		Detail:  "",
		Data:    data,
	}
	ctx.JSON(resp.Code, resp)
}
