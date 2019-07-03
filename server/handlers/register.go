package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/pkg/storage"
)

func UserRegiste(ctx *gin.Context) {
	user := new(models.User)

	// request
	err := ctx.BindJSON(user)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "Parse json to User failed", err.Error())
		return
	}

	// save
	err = storage.GlobalStore.AddUser(*user)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "save user failed", err.Error())
		return
	}

	SendResponse(ctx, http.StatusCreated, "OK", "")
}
