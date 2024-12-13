package models

import "time"

type Wallet struct {
	ID        string     `json:"id"`
	OwnedBy   string     `json:"owned_by"`
	Status    string     `json:"status"` // initialized / enabled / disabled
	Balance   float64    `json:"balance"`
	CreatedAt time.Time  `json:"created_at"`
	EnabledAt *time.Time `json:"enabled_at,omitempty"`
}

var Wallets = make(map[string]*Wallet)
