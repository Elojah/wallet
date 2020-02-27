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

func (h handler) PostWalletHistory(w http.ResponseWriter, r *http.Request) {

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
	var payload dto.PostHistoryReq
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "failed to read payload", http.StatusBadRequest)
		return
	}
	if err := payload.Check(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Fetch wallet history
	ws, err := h.Wallet.ComputeAndFetch(ctx, wallet.Filter{
		ID:        ulid.MustParse(payload.WalletID),
		StartDate: payload.StartDate,
		EndDate:   payload.EndDate,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch wallet history")
		http.Error(w, "failed to fetch wallet history", http.StatusInternalServerError)
	}

	resp := dto.PostHistoryResp{
		History: make([]dto.Wallet, 0, len(ws)),
	}
	// force UTC location display in response
	// it's ok to ignore error here, UTC is valid
	loc, _ := time.LoadLocation("UTC")
	for _, w := range ws {
		resp.History = append(resp.History, dto.Wallet{
			Amount: w.Amount,
			Date:   time.Unix(w.Timestamp, 0).In(loc),
		})
	}

	// #Respond
	raw, err := json.Marshal(resp)
	if err != nil {
		logger.Error().Err(err).Msg("failed to respond")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(raw)
}
