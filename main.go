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
	r.GET("/api/v1/wallet", middleware.AuthRequired(), handlers.ViewWallet)
	r.GET("/api/v1/wallet/transactions", middleware.AuthRequired(), handlers.ViewTransactions)
	r.POST("/api/v1/wallet/deposits", middleware.AuthRequired(), handlers.DepositMoney)
	r.POST("/api/v1/wallet/withdrawals", middleware.AuthRequired(), handlers.DepositMoney)
	r.PATCH("/api/v1/wallet", middleware.AuthRequired(), handlers.DisableWallet)

	r.Run(":8080")
}
