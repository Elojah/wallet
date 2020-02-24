package store

import (
	"github.com/elojah/redis"
	"github.com/elojah/wallet/pkg/wallet"
)

var _ wallet.Store = (*Store)(nil)
var _ wallet.TxStore = (*Store)(nil)

// Store implements wallet and transactions stores.
type Store struct {
	*redis.Service
}

// NewStore returns a new wallet store.
func NewStore(s *redis.Service) *Store {
	return &Store{
		Service: s,
	}
}
