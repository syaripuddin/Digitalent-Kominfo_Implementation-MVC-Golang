package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/syaripuddin/Digitalent-Kominfo_Implementation-MVC-Golang/app/controller"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/v1/account/add", controller.CreateAccount)
	router.POST("/api/v1/login", controller.Login)
	router.GET("/api/v1/account", middleware.Auth, controller.GetAccount)
	router.POST("/api/v1/transfer", middleware.Auth, controller.Transfer)
	router.POST("/api/v1/withdraw", middleware.Auth, controller.Withdraw)
	router.POST("/api/v1/deposit", middleware.Auth, controller.Deposit)
	router.Run(":8081")
}
