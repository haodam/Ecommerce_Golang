package main

import (
	"Ecommerce_Golang/component/appctx"
	"Ecommerce_Golang/db"
	"Ecommerce_Golang/server"
)

func main() {
	DB := db.InitDB()

	appContext := appctx.NewAppContext(DB)

	server.StartServer(appContext)
}
