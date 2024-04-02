package handlers

import (
	"log/slog"
	"net/http"

	"github.com/antihax/goesi"
	"github.com/neoevax/eveCal/internal/db"
	"github.com/neoevax/eveCal/internal/session"
	"github.com/neoevax/eveCal/internal/templates"
)

type HomeHandler struct {
	UserStore *db.Queries
}

type GetHomeHandlerParams struct {
	UserStore *db.Queries
}

func NewHomeHandler(params HomeHandler) *HomeHandler {
	return &HomeHandler{
		UserStore: params.UserStore,
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

	if session.Scs.Exists(r.Context(), "character") {
		v := session.Scs.Get(r.Context(), "character").(goesi.VerifyResponse)

		slog.Info("Character Name", slog.String("Character Name", v.CharacterName))
		slog.Info("Expiration", slog.String("Expiration", v.ExpiresOn))
		templates.Index(v.CharacterName).Render(r.Context(), w)
	} else {
		slog.Info("No Character")
		templates.GuestIndex().Render(r.Context(), w)
	}

	// if err != nil {
	// 	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	return
	// }
}
