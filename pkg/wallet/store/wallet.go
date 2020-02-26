package store

import (
	"context"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"

	werrors "github.com/elojah/wallet/pkg/errors"
	"github.com/elojah/wallet/pkg/wallet"
)

const (
	walletKey = "wallet:"
)

// Insert wallet redis implementation.
func (s *Store) Insert(ctx context.Context, w wallet.W) error {
	raw, err := w.Marshal()
	if err != nil {
		return errors.Wrapf(err, "insert wallet %s", w.ID.String())
	}
	return errors.Wrapf(s.ZAddNX(
		walletKey+w.ID.String(),
		redis.Z{
			Score:  float64(w.Timestamp),
			Member: raw,
		},
	).Err(), "insert wallet %s", w.ID.String())
}

// Fetch wallet redis implementation.
func (s *Store) Fetch(ctx context.Context, filter wallet.Filter) (wallet.W, error) {
	vals, err := s.ZRevRangeByScore(
		walletKey+filter.ID.String(),
		redis.ZRangeBy{
			Count: 1,
			Min:   "-inf",
			Max:   strconv.FormatInt(filter.LastBeforeDate.Unix(), 10),
		},
	).Result()
	if err != nil {
		return wallet.W{}, errors.Wrapf(err, "fetch wallet %s", filter.ID)
	}
	if len(vals) == 0 {
		return wallet.W{}, errors.Wrapf(
			werrors.ErrNotFound{Store: walletKey, Index: filter.ID.String()},
			"fetch wallet %s", filter.ID)
	}

	var result wallet.W
	if err := result.Unmarshal([]byte(vals[0])); err != nil {
		return wallet.W{}, errors.Wrapf(err, "fetch wallet %s", filter.ID.String())
	}
	return result, nil
}

// Remove wallet redis implementation.
func (s *Store) Remove(ctx context.Context, filter wallet.Filter) error {
	return errors.Wrapf(s.ZRemRangeByScore(
		walletKey+filter.ID.String(),
		strconv.FormatInt(filter.StartDate.Unix(), 10),
		"+inf",
	).Err(), "remove wallet %s from TS %d", filter.ID.String(), filter.StartDate.Unix())
}
