package store

import (
	"context"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	"github.com/elojah/wallet/pkg/wallet"
)

const (
	txKey = "tx:"
)

// InsertTx transaction redis implementation.
func (s *Store) InsertTx(ctx context.Context, tx wallet.Tx) error {
	raw, err := tx.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert transaction %s", tx.ID.String())
	}
	return errors.Wrapf(s.ZAddNX(
		txKey+tx.WalletID.String(),
		redis.Z{
			Score:  float64(tx.ID.Time()),
			Member: raw,
		},
	).Err(), "insert transaction %s", tx.ID.String())
}

// FetchManyTx transaction redis implementation.
func (s *Store) FetchManyTx(ctx context.Context, filter wallet.TxFilter) ([]wallet.Tx, error) {
	vals, err := s.ZRevRangeByScore(
		txKey+filter.WalletID.String(),
		redis.ZRangeBy{
			Min: strconv.FormatInt(filter.StartDate.Unix(), 10),
			Max: strconv.FormatInt(filter.EndDate.Unix(), 10),
		},
	).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "fetch transactions for wallet %s", filter.WalletID)
	}
	if len(vals) == 0 {
		return []wallet.Tx{}, nil
	}

	result := make([]wallet.Tx, len(vals))
	for i, val := range vals {
		if err := result[i].Unmarshal([]byte(val)); err != nil {
			return nil, errors.Wrapf(err, "fetch transactions for wallet %s", filter.WalletID)
		}
	}
	return result, nil
}
