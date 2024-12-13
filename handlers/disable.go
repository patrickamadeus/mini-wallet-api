package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DisableWallet(c *gin.Context) {
	walletID := c.GetString("customer_xid")
	wallet, exists := models.Wallets[walletID]
	if !exists {
		utils.ErrorResponse(c, http.StatusOK, map[string][]string{
			"wallet": {"Wallet not found"},
		})
		return
	}

	var requestBody struct {
		IsDisabled bool `form:"is_disabled"`
	}

	if err := c.ShouldBind(&requestBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, map[string][]string{
			"request": {"Invalid request body"},
		})
		return
	}

	if requestBody.IsDisabled && wallet.Status != "disabled" {
		wallet.Status = "disabled"

		utils.SuccessResponse(c, http.StatusOK, gin.H{
			"wallet": gin.H{
				"id":          wallet.ID,
				"owned_by":    wallet.OwnedBy,
				"status":      wallet.Status,
				"disabled_at": time.Now(),
				"balance":     wallet.Balance,
			},
		})
		return
	}

	utils.ErrorResponse(c, http.StatusOK, map[string][]string{
		"wallet": {"Wallet is already disabled or the request is invalid"},
	})
}
