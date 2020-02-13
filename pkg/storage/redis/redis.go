package redis

import (
	"os"
	"time"

	goredis "github.com/gomodule/redigo/redis"
	"github.com/xuebing1110/notify/pkg/storage"
)

type RedisStorage struct {
	*goredis.Pool
}

func init() {
	// RedisClient
	addr := os.Getenv("REDIS_ADDR")
	passwd := os.Getenv("REDIS_PWD")
	if addr == "" {
		addr = "localhost:6379"
	}

	dialFunc := func() (c goredis.Conn, err error) {
		c, err = goredis.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}

		if passwd != "" {
			if _, err := c.Do("AUTH", passwd); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", 0)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	pool := &goredis.Pool{
		MaxIdle:     10,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}

	c := pool.Get()
	defer c.Close()

	if c.Err() != nil {
		panic(c.Err())
	}

	// RedisStorage
	storage.GlobalStore = &RedisStorage{
		Pool: pool,
	}
}
