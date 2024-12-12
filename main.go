// main.go
package main

import (
	"github.com/gin-gonic/gin"

	"mini-wallet-api/handlers"
	"mini-wallet-api/middleware"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/init", handlers.InitializeAccount)
	r.POST("/api/v1/wallet", middleware.AuthRequired(), handlers.EnableWallet)

	r.Run(":8080")
}