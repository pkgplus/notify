package wechat

import (
	"fmt"
	"log"
	"os"

	"github.com/esap/wechat"
	"github.com/esap/wechat/util"

	// "github.com/xuebing1110/notify/pkg/models"
	wxmodels "github.com/xuebing1110/notify/pkg/wechat/models"
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

func SendMsg(t *wxmodels.TemplateMsg) error {
	// log.Println("prepare send message...")
	wxErr := WxServer.SendMsg(t)
	if wxErr.Error() != nil {
		log.Println(wxErr.ErrMsg)
	}

	return wxErr.Error()
}

func SendNotice(n *wxmodels.Notice) error {
	// log.Println("prepare send message...")
	// log.Printf("get noitce %+v", n)
	msg, err := NoticeToTemplateMsg(n)
	if err != nil {
		return err
	}

	log.Printf("send message %+v", msg)
	wxErr := WxServer.SendMsg(msg)
	if wxErr.Error() != nil {
		log.Println(wxErr.ErrMsg)
	}

	return wxErr.Error()
}

func Jscode2Session(jscode string) (sessResp *wxmodels.SessionResp, err error) {
	req_url := fmt.Sprintf(FMT_URL_JSCODE2SESSION,
		WxServer.AppId,
		WxServer.Secret,
		jscode)

	sessResp = new(wxmodels.SessionResp)
	err = util.GetJson(req_url, sessResp)
	return
}
