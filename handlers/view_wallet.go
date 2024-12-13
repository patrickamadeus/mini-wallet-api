package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewWallet(c *gin.Context) {
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

	if wallet.Status != "enabled" {
		utils.ErrorResponse(c, http.StatusForbidden, map[string][]string{"status": {"Wallet is not enabled"}})
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"wallet": wallet,
	})
}
