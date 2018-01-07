package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func energy() {
	api.Post("/session/:sid/energy", handlers.SessionCheck, handlers.AddEnery)
	api.Get("/session/:sid/energy/count", handlers.SessionCheck, handlers.EneryCount)
	api.Get("/session/:sid/energy/one", handlers.SessionCheck, handlers.PopEnergy)
}
