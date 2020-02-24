package wallet

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
)

// Tx represents a transaction on a wallet.
// Time transaction is embed into ID.
type Tx struct {
	ID ksuid.KSUID

	WalletID ksuid.KSUID
	Date     time.Time

	Sum decimal.Decimal
}

// Filter object to fetch specific Txs.
type TxFilter struct {
	WalletID  ksuid.KSUID
	StartDate time.Time
	EndDate   time.Time
}

// TxStore storage layer for Tx object.
type TxStore interface {
	InsertTx(context.Context, Tx) error
	FetchManyTx(context.Context, Filter) ([]Tx, error)
}
