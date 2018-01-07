package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify/pkg/storage"
	"github.com/xuebing1110/notify/pkg/wechat"
)

type Notice struct {
	UserID   string   `json:"touser"`
	Template string   `json:"template_id"`
	Emphasis string   `json:"emphasis"`
	Page     string   `json:"page"`
	Values   []string `json:"values"`
}

func SendNotice(ctx context.Context) {
	uid := getUID(ctx)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "get uid failed", "")
		return
	}

	notice := new(Notice)
	err := ctx.ReadJSON(notice)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse to notice failed", err.Error())
		return
	}
	notice.UserID = uid
	if notice.Page == "" {
		notice.Page = "/pages/index/index"
	} else if strings.Index(notice.Page, "/") >= 0 {
		notice.Page = fmt.Sprintf("/pages/%s/%s", notice.Page, notice.Page)
	}

	energy, err := storage.GlobalStore.PopEnergy(uid)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "pop energy failed", err.Error())
		return
	}

	err = wechat.SendMsg(wechat.NewTemplateMsg(notice.UserID, notice.Template, energy, notice.Values))
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "send message failed", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "OK", "")
}
