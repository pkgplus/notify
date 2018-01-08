package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func login() {
	api.Post("/login", handlers.UserLogin)
}
