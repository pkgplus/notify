package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/pkg/storage"
	"github.com/xuebing1110/notify/pkg/wechat"
)

type LoginReq struct {
	Code string `json:"code"`
}

type LoginResp struct {
	*models.Response
	Session string `json:"session"`
}

func UserLogin(ctx context.Context) {
	lr := new(LoginReq)

	// request
	err := ctx.ReadJSON(lr)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "Parse to json failed", err.Error())
		return
	}
	if lr.Code == "" {
		SendResponse(ctx, http.StatusBadRequest, "code is required", "")
		return
	}

	// openid
	sessRet, err := wechat.Jscode2Session(lr.Code)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "jscode2session failed", err.Error())
		return
	}

	// create session
	sess_3rd := sessRet.OpenID
	// sess_3rd, err := user.GetRandomID(16)
	// if err != nil {
	//  SendResponse(ctx, http.StatusInternalServerError, "create 3rd_sess failed", err.Error())
	//  return
	// }

	// storage
	err = storage.GlobalStore.SaveSession(sess_3rd, sessRet)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "save sess_3rd and sessinfo failed", err.Error())
		return
	}

	SendNormalResponse(ctx, &LoginResp{Session: sess_3rd})
}

func SessCheck(ctx context.Context) {
	sess := ctx.Params().Get("sess")

	resp, err := storage.GlobalStore.QuerySession(sess)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "check session failed", err.Error())
		return
	}

	SendNormalResponse(ctx, resp)
}
