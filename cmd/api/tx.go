package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h handler) PostTx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logger := log.With().Str("route", r.RequestURI).Str("method", r.Method).Logger()

}
