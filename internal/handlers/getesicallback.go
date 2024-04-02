package handlers

import (
	"log/slog"
	"net/http"

	"github.com/neoevax/eveCal/internal/auth"
	"github.com/neoevax/eveCal/internal/session"
)

type GetEsiCallbackHandler struct {
}

func NewGetEsiCallbackHandler() *GetEsiCallbackHandler {
	return &GetEsiCallbackHandler{}
}

func (h *GetEsiCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	slog.Info("Code", slog.String("code", code))
	slog.Info("State", slog.String("state", state))
	auth.EveSSOAnswer(w, r, session.Scs)
}
