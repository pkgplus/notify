package mask

import (
	"fmt"
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/pkg/storage"
	"time"
)

func SaveContacts(uid string, mui *models.MaskUserInfo) error {
	if mui.Id == "" {
		mui.Id = fmt.Sprintf("%d", time.Now().Unix())
	}

	return storage.GlobalStore.AddContact(uid, mui)
}

func ListContacts(uid string) ([]*models.MaskUserInfo, error) {
	return storage.GlobalStore.ListContacts(uid)
}

func GetContact(user, contactUser string) (*models.MaskUserInfo, error) {
	return storage.GlobalStore.GetContact(user, contactUser)
}

func DeleteContact(user, contactUser string) error {
	return storage.GlobalStore.DelContact(user, contactUser)
}
