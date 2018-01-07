package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/xuebing1110/notify/server/app"

	_ "github.com/xuebing1110/notify/pkg/storage/redis"
	_ "github.com/xuebing1110/notify/server/router/v1"
)

func main() {
	// http server
	irisApp := app.GetIrisApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	irisApp.Run(iris.Addr(":" + port))
}
