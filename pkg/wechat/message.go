package wechat

import (
	"fmt"
	"strings"

	"github.com/xuebing1110/notify/pkg/storage"
	wxmodels "github.com/xuebing1110/notify/pkg/wechat/models"
)

func NoticeToTemplateMsg(n *wxmodels.Notice) (*wxmodels.TemplateMsg, error) {
	// if n.Page == "" {
	// 	n.Page = "/pages/index/index"
	// } else if strings.Index(n.Page, "/") < 0 {
	// 	n.Page = fmt.Sprintf("/pages/%s/%s", n.Page, n.Page)
	// }

	energy, err := storage.GlobalStore.PopEnergy(n.UserID)
	if err != nil {
		return nil, err
	}
	if n.Emphasis == "" {
		n.Emphasis = "1"
	}

	tmsg := wxmodels.NewTemplateMsg(n.UserID, n.Template, energy, n.Values)
	tmsg.SetEmphasis(n.Emphasis)
	tmsg.SetPage(tmsg.Page)
	return tmsg, nil
}
