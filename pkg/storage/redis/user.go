package redis

import (
	"github.com/xuebing1110/notify/pkg/models"
)

const (
	USER_PREFIX = "user."
)

func (rs *RedisStorage) UpsertUser(user models.User) error {
	return rs.AddUser(user)
}

func (rs *RedisStorage) AddUser(user models.User) error {
	ret := rs.HMSet(USER_PREFIX+user.ID(), map[string]interface{}(user))
	return ret.Err()
}

func (rs *RedisStorage) Exist(uid string) bool {
	ret := rs.HGet(USER_PREFIX+uid, models.USER_FIELD_SUBTIME)
	if ret.Err() != nil {
		return false
	}

	_, err := ret.Int64()
	if err != nil {
		return false
	}

	return true
}
