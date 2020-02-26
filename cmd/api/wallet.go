package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/wallet/pkg/ulid"
	"github.com/elojah/wallet/pkg/wallet"
	"github.com/elojah/wallet/pkg/wallet/dto"
)

func (h handler) PostWallet(w http.ResponseWriter, r *http.Request) {

	// TODO this should be done on handler side, not inside controller
	switch r.Method {
	case http.MethodPost:
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	logger := log.With().Str("route", r.RequestURI).Str("method", r.Method).Logger()

	// #Create transaction
	id := ulid.NewID()
	if err := h.Wallet.Insert(ctx, wallet.W{
		ID:        id,
		Timestamp: time.Now().Unix(),
		Amount:    "0",
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create wallet")
		http.Error(w, "failed to create wallet", http.StatusInternalServerError)
		return
	}

	// #Respond
	raw, err := json.Marshal(dto.PostResp{ID: id.String()})
	if err != nil {
		logger.Error().Err(err).Msg("failed to respond")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(raw)
}
