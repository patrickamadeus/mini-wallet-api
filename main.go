// main.go
package main

import (
	"github.com/gin-gonic/gin"

	"mini-wallet-api/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/init", handlers.InitializeAccount)

	r.Run(":8080")
}