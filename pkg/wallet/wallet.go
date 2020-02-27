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
	StartDate      time.Time
	EndDate        time.Time
}

// App application layer for W object.
type App interface {
	Store
	// Fetch recompute all Tx data from latest computed wallet
	ComputeAndFetch(context.Context, Filter) ([]W, error)
}

// Store storage layer for W object.
type Store interface {
	Insert(context.Context, W) error
	Fetch(context.Context, Filter) (W, error)
	Remove(context.Context, Filter) error
}

// GenerateHourRange generates a slice of wallets at hours end (e.g: 10:00, 22:00) from a date range
// !!! Amounts are always empty and need to be populated.
func GenerateHourRange(start time.Time, end time.Time) []W {

	var result []W

	// First we generate first rounded hour equal or superior to start
	var ts int64
	if start.Minute() == 0 {
		ts = start.Unix()
	} else {
		ts = start.Add(time.Duration(60-start.Minute()) * time.Minute).Unix()
	}

	// While we didn't reach end date
	for {
		// exit condition
		if ts > end.Unix() {
			return result
		}

		// append current result timestamp
		result = append(result, W{Timestamp: ts})

		// add 1 hours (3600sec) to timestamp
		ts += 3600
	}
}
