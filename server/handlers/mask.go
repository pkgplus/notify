package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/service/mask"
	"net/http"
)

func ListMaskStores(ctx *gin.Context) {
	SendNormalResponse(ctx, mask.ListMaskStores())
}

func GetMaskStore(ctx *gin.Context) {
	sid := ctx.Param("sid")
	ms, err := mask.GetMaskStoreInventory(ctx)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "查询剩余量失败", err.Error())
		return
	}

	var maskStore *models.MaskStore
	for _, m := range ms {
		if fmt.Sprintf("%d", m.ID) == sid {
			maskStore = m
			break
		}
	}
	SendNormalResponse(ctx, maskStore)
}

func AddContact(ctx *gin.Context) {
	uid := GetUID(ctx)
	c := new(models.MaskUserInfo)
	err := ctx.BindJSON(c)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "请求用户信息格式错误", err.Error())
		return
	}

	if c.Name == "" || c.CardNo == "" || c.Mobile == "" {
		SendResponse(ctx, http.StatusBadRequest, "填写不能为空", "")
		return
	}

	if err := mask.SaveContacts(uid, c); err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "保存用户信息失败", err.Error())
		return
	}

	SendNormalResponse(ctx, nil)
}

func ListContacts(ctx *gin.Context) {
	uid := GetUID(ctx)

	ret, err := mask.ListContacts(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "查询联系人列表失败", err.Error())
		return
	}

	SendNormalResponse(ctx, ret)
}

func GetContact(ctx *gin.Context) {
	uid := GetUID(ctx)
	cid := ctx.Param("cid")

	ret, err := mask.GetContact(uid, cid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "查询联系人失败", err.Error())
		return
	}

	SendNormalResponse(ctx, ret)
}

func DelContact(ctx *gin.Context) {
	uid := GetUID(ctx)
	cid := ctx.Param("cid")

	err := mask.DeleteContact(uid, cid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "删除联系人失败", err.Error())
		return
	}

	SendNormalResponse(ctx, nil)
}

func GetCaptcha(ctx *gin.Context) {
	c, err := mask.GetCaptcha(ctx.Request.Context())
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "获取验证码失败", err.Error())
		return
	}

	SendNormalResponse(ctx, c)
}

func ReserveMask(ctx *gin.Context) {
	uid := GetUID(ctx)

	req := new(models.ReserveRequest)
	if err := ctx.BindJSON(req); err != nil {
		SendResponse(ctx, http.StatusBadRequest, "请求格式错误", err.Error())
		return
	}

	if req.Contact == "" {
		SendResponse(ctx, http.StatusBadRequest, "缺少联系人信息", "")
		return
	}

	err := mask.DealReserveRequest(ctx, uid, req)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}
	//SendNormalResponse(ctx, "预约成功，请等待短信通知")
	SendResponse(ctx, http.StatusOK, "预约成功，请等待短信通知", "")
}
