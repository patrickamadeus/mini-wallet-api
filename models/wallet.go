package models

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	OwnedBy   string    `json:"owned_by"`
	Status    string    `json:"status"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

var Wallets = make(map[string]*Wallet)
