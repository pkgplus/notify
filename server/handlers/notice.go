package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/wechat"
	"github.com/xuebing1110/notify/pkg/wechat/models"
)

func SendNotice(ctx *gin.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get uid failed", "")
		return
	}

	// create
	n := new(models.Notice)
	err := ctx.BindJSON(n)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse to notice failed", err.Error())
		return
	}
	n.UserID = uid

	// send
	err = wechat.SendNotice(n)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "send message failed", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "OK", "")
}
