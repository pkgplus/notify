package handlers

import (
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify/pkg/models"
)

func SendResponse(ctx context.Context, code int, msg, detail string) {
	resp := &models.Response{
		code,
		msg,
		detail,
	}
	ctx.StatusCode(resp.Code)
	ctx.JSON(resp)
}
