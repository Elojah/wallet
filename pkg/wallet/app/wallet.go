package app

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

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
		WalletID:  filter.ID,
		StartDate: time.Unix(w.Timestamp, 0),
		EndDate:   filter.EndDate,
	})
	if err != nil {
		return nil, err
	}

	result := wallet.GenerateHourRange(filter.StartDate, filter.EndDate)

	// For each result wallet
	current, err := decimal.NewFromString(w.Amount)
	if err != nil {
		return nil, err
	}
	for i, r := range result {

		// Run transactions in time order
		var j int
		var tx wallet.Tx
		for j, tx = range txs {

			// exit condition, next transactions will be sum for next wallet result
			if int64(tx.ID.Timestamp()) > r.Timestamp {
				// Here is a trick to not remove this transaction of txs slice
				j--
				break
			}

			// boiler, parse decimal string to correctly sum them
			sum, err := decimal.NewFromString(tx.Sum)
			if err != nil {
				return nil, err
			}
			current = current.Add(sum)
		}

		// Remove transactions already added
		if j < len(txs) {
			txs = txs[j+1:]
		}
		result[i].Amount = current.String()

		// TODO Insert wallet to recompute once
	}

	return result, nil
}
