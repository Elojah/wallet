package wallet

import (
	"context"
	"time"

	"github.com/elojah/wallet/pkg/ulid"
)

// Filter object to fetch specific W.
type Filter struct {
	ID             ulid.ID
	LastBeforeDate time.Time
	StartRange     time.Time
}

// App application layer for W object.
type App interface {
	Store
	// Fetch recompute all Tx data from latest computed wallet
	ComputeAndFetch(context.Context, Filter) (W, error)
}

// Store storage layer for W object.
type Store interface {
	Insert(context.Context, W) error
	Fetch(context.Context, Filter) (W, error)
	Remove(context.Context, Filter) error
}
