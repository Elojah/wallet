package main

import (
	"context"
	"net/http"

	"github.com/elojah/wallet/pkg/wallet"
	"github.com/rs/zerolog/log"
)

type handler struct {
	srv *http.Server

	Wallet wallet.App
	Tx     wallet.TxApp
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/transaction", h.PostTx)
	mux.HandleFunc("/wallet", h.PostWallet)
	mux.HandleFunc("/wallet-history", h.PostWalletHistory)

	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() {
		if err := h.srv.ListenAndServeTLS(c.Cert, c.Key); err != nil {
			log.Error().Err(err).Msg("failed to start server")
		}
	}()
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
