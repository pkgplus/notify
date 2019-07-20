package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	UserID  string `json:"userId"`
}

var (
	signingKey = []byte("notify_wx_2019")
)

func UserLogin(ctx *gin.Context) {
	lr := new(LoginReq)

	// request
	err := ctx.BindJSON(lr)
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

	// user
	user_id := sessRet.OpenID

	// create session
	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience:  "wx",
		ExpiresAt: now.Add(time.Hour * 24 * 30).Unix(),
		Issuer:    "notify",
		Subject:   user_id,
		IssuedAt:  now.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sess_3rd, err := token.SignedString(signingKey)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "sign token failed", err.Error())
		return
	}

	// storage
	//err = storage.GlobalStore.SaveSession(sess_3rd, user_id)
	//if err != nil {
	//	SendResponse(ctx, http.StatusInternalServerError, "save session failed", err.Error())
	//	return
	//}

	SendNormalResponse(ctx, &LoginResp{Session: sess_3rd, UserID: user_id})
}

// Deprecated
func SessCheck(ctx *gin.Context) {
	sess := ctx.Param("sess")

	resp, err := storage.GlobalStore.QuerySession(sess)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "check session failed", err.Error())
		return
	}

	SendNormalResponse(ctx, resp)
}
