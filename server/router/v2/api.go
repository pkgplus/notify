package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/server/app"
)

var api *gin.RouterGroup

func init() {
	api = app.GetApp().Group("/api/v2/notify")

	// registe path: /users
	user()

	// registe path: /login
	login()
}
