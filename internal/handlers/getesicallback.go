package handlers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/neoevax/eveCal/internal/auth"
)

type GetEsiCallbackHandler struct {
}

func NewGetEsiCallbackHandler() *GetEsiCallbackHandler {
	return &GetEsiCallbackHandler{}
}

func (h *GetEsiCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	logger.Info("Code", slog.String("code", code))
	logger.Info("State", slog.String("state", state))
	auth.EveSSOAnswer(w, r)

	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
