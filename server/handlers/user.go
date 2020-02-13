package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/models"
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

func LoginByMiniprogram(ctx *gin.Context) {
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

	if sessRet.OpenID == "" || sessRet.Unionid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "登录失败", "get openID")
		return
	}

	// user
	log.Printf("%+v", sessRet)
	user_id := sessRet.Unionid

	// create session
	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience:  "miniprogram",
		ExpiresAt: now.Add(time.Hour * 24 * 30).Unix(),
		Issuer:    "notify",
		Subject:   user_id,
		IssuedAt:  now.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "sign token failed", err.Error())
		return
	}

	SendNormalResponse(ctx, token)
}
