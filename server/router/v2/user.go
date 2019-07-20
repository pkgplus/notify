package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/server/handlers"
)

func user() {
	// user register
	api.POST("/users", handlers.UserRegister)

	// user group
	uGroup := api.Group("/users/:uid")

	// user route
	for _, g := range []*gin.RouterGroup{api, uGroup} {
		// energy
		g.POST("/energy", handlers.SessionCheck, handlers.AddEnery)
		g.GET("/energy/count", handlers.SessionCheck, handlers.EneryCount)
		g.GET("/energy/one", handlers.SessionCheck, handlers.PopEnergy)

		//// notice
		//g.POST("/notice", handlers.SessionCheck, handlers.SendNotice)
	}
}
