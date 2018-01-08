package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify/pkg/storage"
)

const (
	CONTEXT_OPENID_TAG = "OpenID"
	CONTEXT_UNION_TAG  = "UnionID"
)

func SessionCheck(ctx context.Context) {
	sid := ctx.Params().Get("uid")
	if sid == "" {
		SendResponse(ctx, http.StatusUnauthorized, "get uid failed", "")
		return
	}

	resp, err := storage.GlobalStore.QuerySession(sid)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "session maybe expired,please login before sending a notice", err.Error())
		return
	}
	ctx.Values().Set(CONTEXT_OPENID_TAG, resp.OpenID)
	ctx.Values().Set(CONTEXT_UNION_TAG, resp.Unionid)

	ctx.Next()
}

func GetOpenID(ctx context.Context) string {
	return ctx.Values().GetString(CONTEXT_OPENID_TAG)
}

func GetUnionID(ctx context.Context) string {
	return ctx.Values().GetString(CONTEXT_UNION_TAG)
}

func getUID(ctx context.Context) (uid string) {
	// uid = ctx.Values().GetString(CONTEXT_UNION_TAG)
	return ctx.Values().GetString(CONTEXT_OPENID_TAG)
}
