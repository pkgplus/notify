package wechat

import (
	"fmt"
	"os"

	"github.com/esap/wechat"
	"github.com/esap/wechat/util"
	// "github.com/xuebing1110/notify/pkg/models"
)

const (
	FMT_URL_JSCODE2SESSION = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

var (
	WxServer *wechat.Server
)

func init() {
	appid := os.Getenv("WX_APPID")
	secret := os.Getenv("WX_SECRET")
	if appid == "" || secret == "" {
		panic("can't get WX_APPID and WX_SECRET from env")
	}

	WxServer = wechat.New("", appid, secret, "")
	WxServer.MsgUrl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
}

func Jscode2Session(jscode string) (sessResp *SessionResp, err error) {
	req_url := fmt.Sprintf(FMT_URL_JSCODE2SESSION,
		WxServer.AppId,
		WxServer.Secret,
		jscode)

	sessResp = new(SessionResp)
	err = util.GetJson(req_url, sessResp)
	return
}
