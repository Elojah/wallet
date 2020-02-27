package main

import (
	"encoding/json"
	"net/http"

	werrors "github.com/elojah/wallet/pkg/errors"
	"github.com/elojah/wallet/pkg/ulid"
	"github.com/elojah/wallet/pkg/wallet"
	"github.com/elojah/wallet/pkg/wallet/dto"
	oulid "github.com/oklog/ulid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (h handler) PostTx(w http.ResponseWriter, r *http.Request) {

	// TODO this should be done on handler side, not inside controller
	switch r.Method {
	case http.MethodPost:
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	logger := log.With().Str("route", r.RequestURI).Str("method", r.Method).Logger()

	// #Read payload
	var payload dto.PostTxReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "failed to read payload", http.StatusBadRequest)
		return
	}
	if err := payload.Check(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Create transaction
	if err := h.Tx.CreateTx(ctx, wallet.Tx{
		ID:       ulid.NewTimeID(oulid.Timestamp(payload.Date)),
		WalletID: ulid.MustParse(payload.WalletID),
		Sum:      payload.Sum,
	}); err != nil {
		switch errors.Cause(err).(type) {
		case werrors.ErrNotFound:
			http.Error(w, "wallet doesn't exist", http.StatusNotFound)
			return
		default:
			logger.Error().Err(err).Msg("failed to create transaction")
			http.Error(w, "failed to create transaction", http.StatusInternalServerError)
			return
		}
	}

	// TODO invalidate wallets future to this transaction

	w.WriteHeader(http.StatusOK)
}
