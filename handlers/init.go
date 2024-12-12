package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitializeAccount(c *gin.Context) {
	var request struct {
		CustomerXID string `form:"customer_xid" binding:"required"`
	}

	var errors = make(map[string][]string)

	if err := c.ShouldBind(&request); err != nil {
		errors["customer_xid"] = append(errors["customer_xid"], "Missing data for required field.")
	}

	if _, exists := models.Wallets[request.CustomerXID]; exists {
		errors["customer_xid"] = append(errors["customer_xid"], "Account already initialized")
	}

	if len(errors) > 0 {
		utils.ErrorResponse(c, http.StatusBadRequest, errors)
		return
	}

	newWallet := &models.Wallet{
		ID:        uuid.NewString(),
		OwnedBy:   request.CustomerXID,
		Status:    "initialized",
		Balance:   0.0,
		CreatedAt: utils.CurrentTime(),
	}

	models.Wallets[request.CustomerXID] = newWallet

	token, err := utils.GenerateJWT(request.CustomerXID)
	if err != nil {
		errors["token"] = append(errors["token"], "Failed to generate token")
		utils.ErrorResponse(c, http.StatusInternalServerError, errors)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, gin.H{
		"token": token,
	})
}
