package server

import (
	"Ecommerce_Golang/component/appctx"
	"Ecommerce_Golang/module/account/transport/ginaccount"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer(appContext appctx.AppContext) {
	fmt.Println("API server is running !!!")

	r := gin.Default()
	//r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome my API ecommerce",
			"status":  1,
		})
	})

	v1 := r.Group("/v1")
	account := v1.Group("/account")

	account.POST("/create", ginaccount.CreateAccount(appContext))
	account.DELETE("/delete", ginaccount.DeleteAccount(appContext))
	account.GET("/list", ginaccount.ListAccount(appContext))
	account.GET("/find", ginaccount.FindAccount(appContext))

	r.Run()
}
