package handlers

import (
	"mini-wallet-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewTransactions(c *gin.Context) {
	walletID := c.GetString("customer_xid")
	wallet, exists := models.Wallets[walletID]
	if !exists || wallet.Status != "enabled" {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"data": gin.H{
				"error": "Wallet disabled",
			},
		})
		return
	}

	var transactions []models.Transaction
	for _, transaction := range models.Transactions[walletID] {
		transactions = append(transactions, *transaction) // Dereference pointer to value
	}

	if len(models.Transactions[walletID]) == 0 {
		transactions = []models.Transaction{}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"transactions": transactions,
		},
	})
}
