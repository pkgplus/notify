package models

import (
	"fmt"
	"reflect"
	"strings"
)

//const (
//	USER_FIELD_UNIONID  = "unionid"
//	USER_FIELD_OPENID   = "openid"
//	USER_FIELD_NICKNAME = "nickName"
//	USER_FIELD_GENDER   = "gender"
//	USER_FIELD_PROVINCE = "province"
//	USER_FIELD_CITY     = "city"
//	USER_FIELD_COUNTRY  = "country"
//	USER_FIELD_SUBTIME  = "subTime"
//)

type User struct {
	OpenId   string `json:"openid"`
	UnionId  string `json:"unionid"`
	NickName string `json:"nickName"`
	Gender   int    `json:"gender"`
	Province string `json:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
	SubTime  int64  `json:"subTime"`
}

func (u *User) Get(field string) interface{} {
	values := reflect.ValueOf(u)
	types := reflect.TypeOf(u)
	for i := 0; i < types.NumField(); i++ {
		tag := string(types.Field(i).Tag)
		if strings.Contains(tag, fmt.Sprintf(`json:"%s"`, field)) {
			value := values.Field(i)
			switch value.Kind() {
			case reflect.String:
				return value.String()
			case reflect.Int, reflect.Int64:
				return value.Int()
			case reflect.Interface:
				return value.Interface()
			default:
				return "<" + value.Type().String() + " Value>"
			}
		}
	}

	return nil
}

func (u *User) Id() string {
	return u.OpenId
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"openid":   u.OpenId,
		"unionid":  u.UnionId,
		"nickName": u.NickName,
		"gender":   u.Gender,
		"province": u.Province,
		"city":     u.City,
		"country":  u.Country,
		"subTime":  u.SubTime,
	}
}
