package main

import (
	"os"

	"github.com/xuebing1110/notify/server/app"

	_ "github.com/xuebing1110/notify/pkg/storage/redis"
	_ "github.com/xuebing1110/notify/server/router/v2"
)

func main() {
	// http server

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	app.GetApp().Run(addr)
}
