package app

import (
	"context"
	"time"

	"github.com/elojah/wallet/pkg/wallet"
)

// TxApp application implementation.
type TxApp struct {
	wallet.TxStore
	wallet.Store
}

// CreateTx logic implementation.
func (a TxApp) CreateTx(ctx context.Context, tx wallet.Tx) error {

	// # Insert transaction
	if err := a.InsertTx(ctx, tx); err != nil {
		return err
	}

	// # Remove all wallet savepoints after this transaction
	// Any transaction invalidates posterior wallet data.
	if err := a.Remove(ctx, wallet.Filter{
		ID:        tx.WalletID,
		StartDate: time.Unix(int64(tx.ID.Time()), 0),
	}); err != nil {
		return err
	}

	return nil
}
