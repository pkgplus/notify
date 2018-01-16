package redis

import (
	"github.com/xuebing1110/notify/pkg/wechat/models"
)

const (
	SESS_PREFIX = "sess."
)

func (rs *RedisStorage) SaveSession(sess_3rd string, sessInfo *models.SessionResp) error {
	ret := rs.HMSet(SESS_PREFIX+sess_3rd, sessInfo.Convert2Map())
	return ret.Err()
}

func (rs *RedisStorage) QuerySession(sess_3rd string) (*models.SessionResp, error) {
	ret := rs.HGetAll(SESS_PREFIX + sess_3rd)
	if ret.Err() != nil {
		return nil, ret.Err()
	}
	return models.NewSessionResp(ret.Val())
}
