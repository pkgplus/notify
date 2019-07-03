package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func login() {
	api.POST("/login", handlers.UserLogin)
}
