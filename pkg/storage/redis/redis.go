package redis

import (
	"os"

	goredis "github.com/go-redis/redis"
	"github.com/xuebing1110/notify/pkg/storage"
)

type RedisStorage struct {
	*goredis.Client
}

func init() {
	// RedisClient
	addr := os.Getenv("REDIS_ADDR")
	passwd := os.Getenv("REDIS_PASSWD")
	if addr == "" {
		addr = "localhost:6379"
	}
	rc := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	// RedisStorage

	storage.GlobalStore = &RedisStorage{
		Client: rc,
	}
}
