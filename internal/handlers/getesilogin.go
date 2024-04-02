package handlers

import (
	"net/http"

	"github.com/neoevax/eveCal/internal/auth"
	"github.com/neoevax/eveCal/internal/session"
)

type GetEsiAuthHandler struct {
}

func NewGetEsiAuthHandler() *GetEsiAuthHandler {
	return &GetEsiAuthHandler{}
}

func (h *GetEsiAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth.EveSSO(w, r, session.Scs)
}
