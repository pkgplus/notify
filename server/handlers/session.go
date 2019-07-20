package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/storage"
)

const (
	CONTEXT_OPENID_TAG = "OpenID"
	CONTEXT_UNION_TAG  = "UnionID"
)

func SessionCheck(ctx *gin.Context) {
	var err error
	uid := ctx.Param("uid")
	if uid == "" {
		uid, err = getUidFromJwt(ctx)
		if err != nil {
			SendResponse(ctx, http.StatusUnauthorized, "get uid failed", "")
			return
		}
	}

	//uid, err := storage.GlobalStore.QuerySession(sid)
	//if err != nil {
	//	SendResponse(ctx, http.StatusUnauthorized, "session maybe expired,please login before sending a notice", err.Error())
	//	return
	//}

	u, err := storage.GlobalStore.GetUser(uid)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "read user info failed", err.Error())
		return
	}

	ctx.Set(CONTEXT_OPENID_TAG, u.OpenId)
	ctx.Set(CONTEXT_UNION_TAG, u.UnionId)

	ctx.Next()
}

func getUidFromJwt(ctx *gin.Context) (uid string, err error) {
	ss := getToken(ctx)
	if ss == "" {
		return "", fmt.Errorf("token not found")
	}

	token, err := jwt.ParseWithClaims(ss, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return "", errors.Wrap(err, "parse jwt failed")
	}
	claims := token.Claims.(*jwt.StandardClaims)
	err = claims.Valid()
	if err != nil {
		return "", errors.Wrap(err, "jwt invalid")
	}

	return claims.Subject, nil
}

func getToken(ctx *gin.Context) string {
	auth_info := ctx.GetHeader("Authorization")
	if auth_info != "" {
		if strings.HasPrefix(auth_info, "Bearer ") {
			return strings.TrimPrefix(auth_info, "Bearer ")
		}
	}
	return ""
}

//func GetOpenID(ctx *gin.Context) string {
//	return ctx.GetString(CONTEXT_OPENID_TAG)
//}
//
//func GetUnionID(ctx *gin.Context) string {
//	return ctx.GetString(CONTEXT_UNION_TAG)
//}

func GetUID(ctx *gin.Context) string {
	return getUID(ctx)
}

func getUID(ctx *gin.Context) (uid string) {
	// uid = ctx.GetString(CONTEXT_UNION_TAG)
	return ctx.GetString(CONTEXT_OPENID_TAG)
}
