package dto

import (
	"time"

	oulid "github.com/oklog/ulid"
	"github.com/shopspring/decimal"

	"github.com/elojah/wallet/pkg/errors"
	"github.com/elojah/wallet/pkg/ulid"
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
	if oulid.Timestamp(req.Date) >= oulid.MaxTime() {
		return errors.ErrInvalidField{Field: "date", Value: req.Date.String()}
	}
	// # No check on date
	// # TODO it may be smart to disabled transactions too far in the past to avoid overomputation
	return nil
}
