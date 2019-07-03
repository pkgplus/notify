package redis

import (
	"github.com/xuebing1110/notify/pkg/storage"
	"testing"
)

func TestSaveSession(t *testing.T) {
	err := storage.GlobalStore.SaveSession("1", "a")
	if err != nil {
		t.Fatal(err)
	}
}
