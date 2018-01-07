package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func notice() {
	api.Post("/session/:sid/notice", handlers.SessionCheck, handlers.SendNotice)
}
