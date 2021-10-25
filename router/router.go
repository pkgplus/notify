package router

import (
	_ "github.com/pkgplus/notify/docs"
	_ "github.com/pkgplus/notify/router/api/v1"
	"github.com/pkgplus/notify/router/app"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	app.Get().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Run(addr ...string) error {
	return app.Get().Run(addr...)
}
