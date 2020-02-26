package dto

import (
	"time"

	"github.com/elojah/wallet/pkg/errors"
	"github.com/elojah/wallet/pkg/ulid"
	"github.com/shopspring/decimal"
)

// PostTxReq request format for /POST tx
type PostTxReq struct {
	WalletID string    `json:"wallet_id"`
	Sum      string    `json:"sum"`
	Date     time.Time `json:"date"`
}

// Check checks params validity.
func (req PostTxReq) Check() error {
	if _, err := ulid.Parse(req.WalletID); err != nil {
		return errors.ErrInvalidField{Field: "wallet_id", Value: req.WalletID}
	}
	if _, err := decimal.NewFromString(req.Sum); err != nil {
		return errors.ErrInvalidField{Field: "sum", Value: req.Sum}
	}
	return nil
}
