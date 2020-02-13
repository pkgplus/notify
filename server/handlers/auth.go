package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CONTEXT_OPENID_TAG = "OpenID"
	CONTEXT_UNION_TAG  = "UnionID"
)

func AuthN(ctx *gin.Context) {
	uid, err := getUidFromJwt(ctx)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "用户未登录，请重新登录", "")
		ctx.Abort()
		return
	}

	//u, err := storage.GlobalStore.GetUser(uid)
	//if err != nil {
	//	SendResponse(ctx, http.StatusUnauthorized, "read user info failed", err.Error())
	//	return
	//}

	ctx.Set(CONTEXT_UNION_TAG, uid)
	//ctx.Set(CONTEXT_OPENID_TAG, u.OpenId)
	//ctx.Set(CONTEXT_UNION_TAG, u.UnionId)

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

	log.Printf("%+v", claims)
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
	if ctx.Param("uid") != "" {
		return ctx.Param("uid")
	}
	return ctx.GetString(CONTEXT_UNION_TAG)
}
