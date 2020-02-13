package redis

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

const (
	ENERGY_PREFIX = "energy."
)

var (
	SEVEN_DAY = 168 * time.Hour
)

func (rs *RedisStorage) AddEnergy(uid, energy string) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("RPUSH", ENERGY_PREFIX+uid, fmt.Sprintf("%s,%d", energy, time.Now().Unix()))
	if err != nil {
		return err
	}

	_, err = client.Do("EXPIRE", ENERGY_PREFIX+uid, SEVEN_DAY/time.Second)
	return err
}

func (rs *RedisStorage) GetEnergyCount(uid string) int64 {
	client := rs.Get()
	defer client.Close()

	count, _ := redigo.Int64(client.Do("LLEN", ENERGY_PREFIX+uid))
	return count
}

func (rs *RedisStorage) PopEnergy(uid string) (string, error) {
	var curtime int64 = time.Now().Unix()
	for {
		energy_ret, err := rs.popOneEnergy(uid)
		if err != nil {
			return "", err
		}
		energy_info := strings.SplitN(energy_ret, ",", 2)
		if len(energy_info) != 2 {
			return "", errors.New("text")
		}

		pushtime, err := strconv.Atoi(energy_info[1])
		if err != nil {
			log.Printf("convert to time failed:%s", err)
			continue
		}

		if curtime-int64(pushtime) < 604000 {
			return energy_info[0], nil
		}
	}

	return "", nil
}

func (rs *RedisStorage) popOneEnergy(uid string) (string, error) {
	client := rs.Get()
	defer client.Close()

	return redigo.String(client.Do("LPOP", ENERGY_PREFIX+uid))
}

func (rs *RedisStorage) ExpireEnergy(uid string) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("EXPIRE", ENERGY_PREFIX+uid, int64(SEVEN_DAY/time.Second))
	return err
}
