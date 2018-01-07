package handlers

import (
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify/pkg/storage"
	"net/http"
)

func AddEnery(ctx context.Context) {
	eneryMap := make(map[string]string)
	err := ctx.ReadJSON(&eneryMap)
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
		SendResponse(ctx, http.StatusInternalServerError, "get openid failed", "")
		return
	}

	err = storage.GlobalStore.AddEnergy(uid, energy)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "add energy failed", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "OK", "")
}

func EneryCount(ctx context.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get openid failed", "")
		return
	}

	// count := storage.GlobalStore.GetEnergyCount(uid)
	// if err != nil {
	//  SendResponse(ctx, http.StatusInternalServerError, "get energy count failed", err.Error())
	//  return
	// }

	ctx.JSON(map[string]int64{"count": storage.GlobalStore.GetEnergyCount(uid)})
}

func PopEnergy(ctx context.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get openid failed", "")
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
		ctx.JSON(map[string]string{"enery": energy})
	}
}
