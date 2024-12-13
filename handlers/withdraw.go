package handlers

import (
	"mini-wallet-api/models"
	"mini-wallet-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func WithdrawMoney(c *gin.Context) {
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

	if request.Amount > wallet.Balance {
		errors["amount"] = append(errors["amount"], "Insufficient funds")
		utils.ErrorResponse(c, http.StatusPaymentRequired, errors)
		return
	}

	for _, transactions := range models.Transactions {
		for _, transaction := range transactions {
			if transaction.ReferenceID == request.ReferenceID {
				errors["reference_id"] = append(errors["reference_id"], "Reference ID already used")
				utils.ErrorResponse(c, http.StatusBadRequest, errors)
				return
			}
		}
	}

	newTransaction := &models.Transaction{
		ID:           uuid.NewString(),
		Status:       "success",
		TransactedAt: time.Now(),
		Type:         "withdrawal",
		Amount:       request.Amount,
		ReferenceID:  request.ReferenceID,
	}

	models.Transactions[wallet.OwnedBy] = append(models.Transactions[wallet.OwnedBy], newTransaction)

	wallet.Balance -= request.Amount

	withdrawResponse := models.WithdrawResponse{
		ID:          newTransaction.ID,
		Status:      newTransaction.Status,
		WithdrawnBy: wallet.ID,
		WithdrawnAt: newTransaction.TransactedAt,
		Amount:      newTransaction.Amount,
		ReferenceID: newTransaction.ReferenceID,
	}

	utils.SuccessResponse(c, http.StatusCreated, gin.H{
		"withdrawal": withdrawResponse,
	})
}
