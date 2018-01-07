package redis

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	ENERGY_PREFIX = "energy."
)

var (
	SEVEN_DAY = 168 * time.Hour
)

func (rs *RedisStorage) AddEnergy(uid, energy string) error {
	ret := rs.RPush(ENERGY_PREFIX+uid, fmt.Sprintf("%s,%d", energy, time.Now().Unix()))
	if ret.Err() != nil {
		return ret.Err()
	}

	rs.ExpireAt(ENERGY_PREFIX+uid, time.Now().Add(SEVEN_DAY))
	return nil
}

func (rs *RedisStorage) GetEnergyCount(uid string) int64 {
	ret := rs.LLen(ENERGY_PREFIX + uid)
	return ret.Val()
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
	ret := rs.LPop(ENERGY_PREFIX + uid)
	return ret.Result()
}

func (rs *RedisStorage) ExpireEnergy(uid string) error {
	ret := rs.ExpireAt(ENERGY_PREFIX+uid, time.Now().Add(SEVEN_DAY))
	return ret.Err()
}
