package handlers

import (
	"log/slog"
	"net/http"

	"github.com/antihax/goesi"
	"github.com/neoevax/eveCal/internal/auth"
	"github.com/neoevax/eveCal/internal/db"
	"github.com/neoevax/eveCal/internal/session"
	"github.com/neoevax/eveCal/internal/templates"
	"golang.org/x/oauth2"
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

		token := session.Scs.Get(r.Context(), "token").(oauth2.Token)

		ctx := auth.GetTokenContext(&token)
		info, _, err := auth.EsiClient.ESI.CalendarApi.GetCharactersCharacterIdCalendar(ctx, v.CharacterID, nil)
		if err != nil {
			slog.Error("Error getting calendar events", slog.Any("Error", err))
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		slog.Info("Info", slog.Any("Info", info))
		for _, event := range info {
			event, _, err := auth.EsiClient.ESI.CalendarApi.GetCharactersCharacterIdCalendarEventId(ctx, v.CharacterID, event.EventId, nil)
			if err != nil {
				slog.Error("Error getting calendar events", slog.Any("Error", err))
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			slog.Info("Event", slog.Any("Event", event))
		}
		templates.Index(v.CharacterName).Render(r.Context(), w)
	} else {
		slog.Info("No Character")
		templates.GuestIndex().Render(r.Context(), w)
	}
}
