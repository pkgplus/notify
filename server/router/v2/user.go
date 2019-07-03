package v1

import (
	"github.com/xuebing1110/notify/server/handlers"
)

func user() {
	// user register
	api.POST("/users", handlers.UserRegiste)

	// energy
	api.POST("/users/:uid/energy", handlers.SessionCheck, handlers.AddEnery)
	api.GET("/users/:uid/energy/count", handlers.SessionCheck, handlers.EneryCount)
	api.GET("/users/:uid/energy/one", handlers.SessionCheck, handlers.PopEnergy)

	// notice
	api.POST("/users/:uid/notice", handlers.SessionCheck, handlers.SendNotice)
}
