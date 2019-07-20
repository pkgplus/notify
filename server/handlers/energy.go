package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/storage"
	"net/http"
)

func AddEnery(ctx *gin.Context) {
	eneryMap := make(map[string]string)
	err := ctx.BindJSON(&eneryMap)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse json failed", err.Error())
		return
	}

	energy, ok := eneryMap["energy"]
	if !ok {
		SendResponse(ctx, http.StatusBadRequest, "energy is required", "")
		return
	}

	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get uid failed", "")
		return
	}

	err = storage.GlobalStore.AddEnergy(uid, energy)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "add energy failed", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "OK", "")
}

func EneryCount(ctx *gin.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get uid failed", "")
		return
	}

	// count := storage.GlobalStore.GetEnergyCount(uid)
	// if err != nil {
	//  SendResponse(ctx, http.StatusInternalServerError, "get energy count failed", err.Error())
	//  return
	// }

	data := map[string]int64{"count": storage.GlobalStore.GetEnergyCount(uid)}
	SendNormalResponse(ctx, data)
}

func PopEnergy(ctx *gin.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get uid failed", "")
		return
	}

	energy, err := storage.GlobalStore.PopEnergy(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "pop one energy failed", err.Error())
		return
	}

	if energy == "" {
		SendResponse(ctx, http.StatusBadRequest, "no energy to pop", "")
	} else {
		data := map[string]string{"enery": energy}
		SendNormalResponse(ctx, data)
	}
}
