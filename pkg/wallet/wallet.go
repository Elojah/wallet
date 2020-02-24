package wallet

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
)

// W represents a wallet at a specific time.
// Time is embed into ID.
type W struct {
	ID ksuid.KSUID

	Date time.Time

	Amount decimal.Decimal
}

// Filter object to fetch specific W.
type Filter struct {
	Date time.Time
}

// App application layer for W object.
type App interface {
	// Fetch recompute all Tx data from latest computed wallet
	Fetch(context.Context, Filter)
}

// Store storage layer for W object.
type Store interface {
	Insert(context.Context, W) error
	FetchLast(context.Context, time.Time) (W, error)
}
