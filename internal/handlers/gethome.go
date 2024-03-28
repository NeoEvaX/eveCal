package handlers

import (
	"net/http"

	"github.com/neoevax/eveCal/internal/db"
	"github.com/neoevax/eveCal/internal/templates"
)

type HomeHandler struct {
	userStore *db.Queries
}

type GetHomeHandlerParams struct {
	UserStore *db.Queries
}

func NewHomeHandler(params GetHomeHandlerParams) *HomeHandler {
	return &HomeHandler{
		userStore: params.UserStore,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//_, claims, _ := jwtauth.FromContext(r.Context())

	//_, ok := claims["email"].(string)

	// if !ok {
	// 	err := templates.GuestIndex().Render(r.Context(), w)

	// 	if err != nil {
	// 		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	return
	// }

	err := templates.Index().Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
