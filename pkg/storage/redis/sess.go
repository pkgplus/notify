package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	SESS_PREFIX = "sess."
)

func (rs *RedisStorage) SaveSession(sess_3rd string, user_id string) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("SET", SESS_PREFIX+sess_3rd, user_id, 24*time.Hour/time.Second)
	return err
}

func (rs *RedisStorage) QuerySession(sess_3rd string) (string, error) {
	client := rs.Get()
	defer client.Close()

	return redis.String(client.Do("GET", SESS_PREFIX+sess_3rd))
}
