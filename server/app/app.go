package app

import (
	"github.com/gin-gonic/gin"
)

var (
	app = gin.New()
)

func GetApp() *gin.Engine {
	return app
}
