package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EnableWallet(c *gin.Context) {
	customerXID, exists := c.Get("customer_xid")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, map[string][]string{"customer_xid": {"Customer ID not found"}})
		return
	}

	customerID := customerXID.(string)

	wallet, exists := models.Wallets[customerID]
	if !exists {
		utils.ErrorResponse(c, http.StatusNotFound, map[string][]string{"wallet": {"Wallet not found"}})
		return
	}

	if wallet.Status == "enabled" {
		utils.ErrorResponse(c, http.StatusConflict, map[string][]string{"status": {"Already enabled"}})
		return
	}

	currentTime := time.Now()
	wallet.Status = "enabled"
	wallet.EnabledAt = &currentTime

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"wallet": wallet,
	})
}
