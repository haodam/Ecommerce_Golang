package main

import (
	"Ecommerce_Golang/component/appctx"
	"Ecommerce_Golang/module/account/transport/ginaccount"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	dsn := "host=localhost user=root password=ecommerce123 dbname=ecommercegolang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)

	db = db.Debug()
	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	//r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	account := v1.Group("/accounts")

	account.POST("", ginaccount.CreateAccount(appContext))

	r.Run()
}
