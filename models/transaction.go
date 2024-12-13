package models

import "time"

type Transaction struct {
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	TransactedAt time.Time `json:"transacted_at"`
	Type         string    `json:"type"` // "deposit" or "withdrawal"
	Amount       float64   `json:"amount"`
	ReferenceID  string    `json:"reference_id"`
}

var Transactions = make(map[string][]*Transaction)

type DepositResponse struct {
	ID          string    `json:"id"`
	DepositedBy string    `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}

type WithdrawResponse struct {
	ID          string    `json:"id"`
	WithdrawnBy string    `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}
