package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/xuebing1110/notify/server/app"
	"github.com/xuebing1110/notify/server/handlers"
)

func init() {
	api := app.GetApp().Group("/api/v2/notify")
	{
		// 登录
		api.POST("/login-by-miniprogram", handlers.LoginByMiniprogram)

		// 用户信息
		maskBase := api.Group("/mask")
		{
			maskBase.GET("/captcha", handlers.GetCaptcha)
			maskBase.GET("/stores", handlers.ListMaskStores)
		}

		maskWithAuth := maskBase.Group("", handlers.AuthN)
		userMask := maskBase.Group("/mask/users/:uid")
		for _, mask := range []*gin.RouterGroup{maskWithAuth, userMask} {
			mask.POST("/contacts", handlers.AddContact)
			mask.GET("/contacts", handlers.ListContacts)
			mask.GET("/contacts/:cid", handlers.GetContact)
			mask.PUT("/contacts/:cid", handlers.AddContact)
			mask.DELETE("/contacts/:cid", handlers.DelContact)
			mask.POST("/reserve", handlers.ReserveMask)
		}
	}
}
