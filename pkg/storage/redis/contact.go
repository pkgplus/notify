package redis

import (
	"encoding/json"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/xuebing1110/notify/pkg/models"
)

//AddContact(uid string, contact *models.MaskUserInfo) error
//ListContacts(uid string) ([]*models.MaskUserInfo, error)
//UpdateContact(uid string, contact *models.MaskUserInfo) error
//DelContact(uid, contact_id string) error

const (
	CONTACT_PREFIX = "mask.contact."
)

func (rs *RedisStorage) AddContact(uid string, contact *models.MaskUserInfo) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("HSET", CONTACT_PREFIX+uid, contact.Id, string(contact.Bytes()))
	if err != nil {
		return err
	}

	return nil
}

func (rs *RedisStorage) GetContact(uid, cid string) (*models.MaskUserInfo, error) {
	client := rs.Get()
	defer client.Close()

	bytes, err := redigo.Bytes(client.Do("HGET", CONTACT_PREFIX+uid, cid))
	if err != nil {
		return nil, err
	}

	mui := new(models.MaskUserInfo)
	if err := json.Unmarshal(bytes, mui); err != nil {
		return nil, err
	}
	return mui, nil
}

func (rs *RedisStorage) ListContacts(uid string) (muis []*models.MaskUserInfo, err error) {
	client := rs.Get()
	defer client.Close()

	values, err := redigo.Values(client.Do("HGETALL", CONTACT_PREFIX+uid))
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return
	}

	muis = make([]*models.MaskUserInfo, 0, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		//cid := string(values[i.([]byte))
		mui := new(models.MaskUserInfo)
		bytes := values[i+1].([]byte)
		if err := json.Unmarshal(bytes, mui); err != nil {
			return muis, err
		}
		muis = append(muis, mui)
	}

	//resp := make([][]byte, 0)
	//if err = redigo.ScanStruct(values, resp); err != nil {
	//	return nil, err
	//}
	//
	//for _, item := range resp {
	//	log.Printf("%s", item)
	//}

	return
}

func (rs *RedisStorage) UpdateContact(uid string, contact *models.MaskUserInfo) error {
	return rs.AddContact(uid, contact)
}

func (rs *RedisStorage) DelContact(uid, contact_id string) error {
	client := rs.Get()
	defer client.Close()

	_, err := client.Do("HDEL", CONTACT_PREFIX+uid, contact_id)
	if err != nil {
		return err
	}

	return nil
}
