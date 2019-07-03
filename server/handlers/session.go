package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/pkg/storage"
)

const (
	CONTEXT_OPENID_TAG = "OpenID"
	CONTEXT_UNION_TAG  = "UnionID"
)

func SessionCheck(ctx *gin.Context) {
	sid := ctx.Param("uid")
	if sid == "" {
		SendResponse(ctx, http.StatusUnauthorized, "get uid failed", "")
		return
	}

	resp, err := storage.GlobalStore.QuerySession(sid)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "session maybe expired,please login before sending a notice", err.Error())
		return
	}
	ctx.Set(CONTEXT_OPENID_TAG, resp.OpenID)
	ctx.Set(CONTEXT_UNION_TAG, resp.Unionid)

	ctx.Next()
}

func GetOpenID(ctx *gin.Context) string {
	return ctx.GetString(CONTEXT_OPENID_TAG)
}

func GetUnionID(ctx *gin.Context) string {
	return ctx.GetString(CONTEXT_UNION_TAG)
}

func getUID(ctx *gin.Context) (uid string) {
	// uid = ctx.GetString(CONTEXT_UNION_TAG)
	return ctx.GetString(CONTEXT_OPENID_TAG)
}
