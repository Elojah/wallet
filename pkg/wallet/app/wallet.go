package app

import (
	"context"

	"github.com/elojah/wallet/pkg/wallet"
)

// App wallet logic implementation.
type App struct {
	wallet.Store
}

// ComputeAndFetch logic implementation.
func (a App) ComputeAndFetch(context.Context, wallet.Filter) ([]wallet.W, error) {
	return []wallet.W{}, nil
}
