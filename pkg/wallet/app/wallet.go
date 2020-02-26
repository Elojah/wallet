package app

import (
	"context"

	"github.com/elojah/wallet/pkg/wallet"
)

// App wallet logic implementation.
type App struct{}

// Fetch logic implementation.
func (a App) Fetch(context.Context, wallet.Filter) (wallet.W, error) {
	return wallet.W{}, nil
}
