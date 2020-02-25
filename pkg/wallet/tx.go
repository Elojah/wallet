package wallet

import (
	"context"
	"time"

	"github.com/elojah/wallet/pkg/ulid"
)

// Filter object to fetch specific Txs.
type TxFilter struct {
	WalletID  ulid.ID
	StartDate time.Time
	EndDate   time.Time
}

// TxStore storage layer for Tx object.
type TxStore interface {
	InsertTx(context.Context, Tx) error
	FetchManyTx(context.Context, TxFilter) ([]Tx, error)
}
