package app

import (
	"context"
	"time"

	"github.com/elojah/wallet/pkg/wallet"
)

// App wallet logic implementation.
type App struct {
	wallet.Store
	wallet.TxStore
}

// ComputeAndFetch logic implementation.
func (a App) ComputeAndFetch(ctx context.Context, filter wallet.Filter) ([]wallet.W, error) {

	// #Fetch last before date computed wallet
	w, err := a.Fetch(ctx, wallet.Filter{
		ID:             filter.ID,
		LastBeforeDate: filter.StartDate,
	})
	if err != nil {
		return nil, err
	}

	// #Fetch all transactions between last computed wallet and end date
	txs, err := a.FetchManyTx(ctx, wallet.TxFilter{
		StartDate: time.Unix(w.Timestamp, 0),
		EndDate:   filter.EndDate,
	})
	if err != nil {
		return nil, err
	}

	var result []wallet.W
	for _, tx := range txs {
		if time.Unix(tx.Timestamp, 0).After(filter.StartDate) {
			result = append(result, w)
		}

	}

	return []wallet.W{}, nil
}
