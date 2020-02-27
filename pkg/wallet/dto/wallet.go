package dto

import (
	"time"

	"github.com/elojah/wallet/pkg/errors"
	"github.com/oklog/ulid"
)

// PostResp response format for POST /wallet
type PostResp struct {
	ID string `json:"id"`
}

// PostHistoryReq request format for POST /wallet-history
type PostHistoryReq struct {
	WalletID  string    `json:"wallet_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Check checks params validity.
func (req PostHistoryReq) Check() error {
	if _, err := ulid.Parse(req.WalletID); err != nil {
		return errors.ErrInvalidField{Field: "wallet_id", Value: req.WalletID}
	}
	return nil
}

// Wallet represents wallet domain type for API.
type Wallet struct {
	Date   time.Time `json:"date"`
	Amount string    `json:"amount"`
}

// PostHistoryResp response format for POST /wallet-history
type PostHistoryResp struct {
	History []Wallet `json:"history"`
}
