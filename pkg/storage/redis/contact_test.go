package redis

import (
	"fmt"
	"github.com/xuebing1110/notify/pkg/models"
	"github.com/xuebing1110/notify/pkg/storage"
	"testing"
)

func TestContact(t *testing.T) {
	uid := "test"
	c := &models.MaskUserInfo{
		"1",
		"王重胜", "370201199410328213", "13355329830",
		&models.MaskStore{
			Area:    "李沧区",
			Address: "李沧区铜川路216号负一层超市",
			OrgCode: "91370213MA3DQENL2P",
			Name:    "丽达绿城店",
			ID:      132,
		},
	}

	db := storage.GlobalStore
	err := db.AddContact("test", c)
	if err != nil {
		t.Fatal(err)
	}

	// GET
	cnew, err := db.GetContact(uid, c.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err := compare(c, cnew); err != nil {
		t.Fatal(err)
	}

	// LIST
	cs, err := db.ListContacts(uid)
	if err != nil {
		t.Fatal(err)
	}

	if len(cs) != 1 {
		t.Fatalf("expect one contanct, but get %d", len(cs))
	}

	if err := compare(c, cs[0]); err != nil {
		t.Fatal(err)
	}

	// DELETE
	if err := db.DelContact(uid, c.Id); err != nil {
		t.Fatal(err)
	}
}

func compare(c, cnew *models.MaskUserInfo) error {
	if cnew.Name != c.Name || cnew.ExpectStore.ID != c.ExpectStore.ID {
		return fmt.Errorf("get contact failed: %+v", cnew)
	}

	return nil
}
