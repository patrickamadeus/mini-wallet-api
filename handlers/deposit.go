package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DepositMoney(c *gin.Context) {
	var request struct {
		Amount      float64 `form:"amount" binding:"required"`
		ReferenceID string  `form:"reference_id" binding:"required"`
	}

	var errors = make(map[string][]string)

	if err := c.ShouldBind(&request); err != nil {
		errors["amount"] = append(errors["amount"], "Amount is required")
		errors["reference_id"] = append(errors["reference_id"], "Reference ID is required")
		utils.ErrorResponse(c, http.StatusBadRequest, errors)
		return
	}

	customerXID, _ := c.Get("customer_xid")
	wallet, exists := models.Wallets[customerXID.(string)]
	if !exists {
		errors["customer_xid"] = append(errors["customer_xid"], "Wallet not found")
		utils.ErrorResponse(c, http.StatusNotFound, errors)
		return
	}

	if wallet.Status != "enabled" {
		errors["status"] = append(errors["status"], "Wallet is not enabled")
		utils.ErrorResponse(c, http.StatusBadRequest, errors)
		return
	}

	if _, exists := models.Transactions[request.ReferenceID]; exists {
		errors["reference_id"] = append(errors["reference_id"], "Reference ID already used")
		utils.ErrorResponse(c, http.StatusBadRequest, errors)
		return
	}

	newTransaction := &models.Transaction{
		ID:           uuid.NewString(),
		Status:       "success",
		TransactedAt: utils.CurrentTime(),
		Type:         "deposit",
		Amount:       request.Amount,
		ReferenceID:  request.ReferenceID,
	}

	models.Transactions[wallet.OwnedBy] = append(models.Transactions[wallet.OwnedBy], newTransaction)

	go func() {
		time.Sleep(3 * time.Second)
		wallet.Balance += request.Amount
	}()

	depositResponse := models.DepositResponse{
		ID:          newTransaction.ID,
		Status:      newTransaction.Status,
		DepositedBy: wallet.ID,
		DepositedAt: newTransaction.TransactedAt,
		Amount:      newTransaction.Amount,
		ReferenceID: newTransaction.ReferenceID,
	}

	utils.SuccessResponse(c, http.StatusCreated, gin.H{
		"deposit": depositResponse,
	})
}
