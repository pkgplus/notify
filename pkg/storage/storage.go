package storage

import (
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/pkg/wechat"
)

var GlobalStore Storage

type Storage interface {
	SaveSession(sess_3rd string, sessInfo *wechat.SessionResp) error
	QuerySession(sess_3rd string) (*wechat.SessionResp, error)

	UpsertUser(user models.User) error
	AddUser(user models.User) error
	Exist(uid string) bool

	AddEnergy(uid, energy string) error
	GetEnergyCount(uid string) int64
	PopEnergy(uid string) (string, error)
}
