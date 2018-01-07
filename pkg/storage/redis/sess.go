package redis

import (
	"github.com/xuebing1110/fortify/pkg/wechat"
)

const (
	SESS_PREFIX = "sess."
)

func (rs *RedisStorage) SaveSession(sess_3rd string, sessInfo *wechat.SessionResp) error {
	ret := rs.HMSet(SESS_PREFIX+sess_3rd, sessInfo.Convert2Map())
	return ret.Err()
}

func (rs *RedisStorage) QuerySession(sess_3rd string) (*wechat.SessionResp, error) {
	ret := rs.HGetAll(SESS_PREFIX + sess_3rd)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return wechat.NewSessionResp(ret.Val())
}
