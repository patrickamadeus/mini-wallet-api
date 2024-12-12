package models

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	OwnedBy   string    `json:"owned_by"`
	Status    string    `json:"status"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	EnabledAt *time.Time `json:"enabled_at,omitempty"` // New field to track when the wallet was enabled
}

var Wallets = make(map[string]*Wallet)
