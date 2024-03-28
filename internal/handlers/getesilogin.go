package handlers

import (
	"net/http"

	"github.com/neoevax/eveCal/internal/auth"
)

type GetEsiAuthHandler struct {
}

func NewGetEsiAuthHandler() *GetEsiAuthHandler {
	return &GetEsiAuthHandler{}
}

func (h *GetEsiAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth.EveSSO(w, r)

	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
