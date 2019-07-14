package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func login() {
	// wechat login
	api.POST("/login", handlers.UserLogin)
}
