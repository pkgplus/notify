package models

type MpUser struct {
	OpenId     string `json:"openid"`     // 用户OpenId
	UnionId    string `json:"unionid"`    // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	NickName   string `json:"nickname"`   // 用户的昵称
	Mobile     string `json:"mobile"`     // mobile
	Sex        int    `json:"sex"`        // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City       string `json:"city"`       // 用户所在城市
	Country    string `json:"country"`    // 用户所在国家
	Province   string `json:"province"`   //用户所在省份
	HeadimgUrl string `json:"headimgurl"` //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。

	CreateTime int64 `json:"createTime"`
}
