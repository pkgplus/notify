package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func user() {
	// user register
	api.Post("/users", handlers.UserRegiste)

	// energy
	api.Post("/users/:uid/energy", handlers.SessionCheck, handlers.AddEnery)
	api.Get("/users/:uid/energy/count", handlers.SessionCheck, handlers.EneryCount)
	api.Get("/users/:uid/energy/one", handlers.SessionCheck, handlers.PopEnergy)

	// notice
	api.Post("/users/:uid/notice", handlers.SessionCheck, handlers.SendNotice)
}
