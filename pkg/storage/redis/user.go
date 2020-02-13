package redis

import (
	"errors"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/xuebing1110/notify/pkg/models"
)

const (
	USER_PREFIX = "user."
)

func (rs *RedisStorage) UpdateUser(uid string, props map[string]interface{}) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("HMSET", redigo.Args{}.Add(USER_PREFIX+uid).AddFlat(props)...)
	return err
}

func (rs *RedisStorage) AddUser(user *models.User) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("HMSET", redigo.Args{}.Add(USER_PREFIX+user.Id()).AddFlat(&user)...)
	return err
}

func (rs *RedisStorage) GetUser(uid string) (*models.User, error) {
	client := rs.Get()
	defer client.Close()

	values, err := redigo.Values(client.Do("HGETALL", USER_PREFIX+uid))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		err = errors.New("the refresh token is invalid or expired")
		return nil, err
	}

	u := new(models.User)
	if err = redigo.ScanStruct(values, u); err != nil {
		//err = errors.Wrapf(err, "parse \"%s\" result failed", key)
		return nil, err
	}
	return u, nil
}

func (rs *RedisStorage) Exist(uid string) bool {
	client := rs.Get()
	defer client.Close()

	count, err := redigo.Int(client.Do("EXISTS", USER_PREFIX+uid))
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}

	return false
}
