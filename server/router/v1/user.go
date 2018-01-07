package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func user() {
	// user
	api.Post("/users", handlers.UserRegiste)

	// session
	api.Post("/login", handlers.UserLogin)
}
