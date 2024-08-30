package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/antihax/goesi"
	"github.com/jackc/pgx/v5"
	"github.com/neoevax/eveCal/internal/db"
	"golang.org/x/oauth2"
)

var (
	EsiClient        *goesi.APIClient
	SSOAuthenticator *goesi.SSOAuthenticator
	scopes           []string
	database         *db.Queries
)

type AuthHandlerParams struct {
	DataBase *db.Queries
}

func Setup(params AuthHandlerParams) {
	database = params.DataBase
	// create ESI client
	httpClient := &http.Client{}
	// call Status endpoint
	scopes = []string{"esi-calendar.respond_calendar_events.v1", "esi-calendar.read_calendar_events.v1"}
	EsiClient = goesi.NewAPIClient(httpClient, "EveCal (ian.kremer@gmail.com, @neoevax on Discord")
	SSOAuthenticator = goesi.NewSSOAuthenticator(httpClient, os.Getenv("CLIENT_ID"), os.Getenv("SECRET_KEY"), "http://localhost:3000/api/esi/callback", scopes)
}

func EveSSO(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) (int, error) {
	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	s.Put(r.Context(), "state", state)

	// Generate the SSO URL with the state string
	url := SSOAuthenticator.AuthorizeURL(state, true, scopes)

	// Send the user to the URL
	http.Redirect(w, r, url, http.StatusFound)
	return http.StatusMovedPermanently, nil
}

func EveSSOAnswer(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) (int, error) {
	// get our code and state
	code := r.FormValue("code")
	state := r.FormValue("state")

	// Verify the state matches our randomly generated string from earlier.
	if s.Get(r.Context(), "state") != state {
		return http.StatusInternalServerError, errors.New("invalid scopestate")
	}

	// Exchange the code for an Access and Refresh token.
	token, err := SSOAuthenticator.TokenExchange(code)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	slog.Info("Token Info", slog.Any("Token", token))

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc := SSOAuthenticator.TokenSource(token)

	// Assign an auth context to the calls
	s.Put(r.Context(), "token", &token)

	// Verify the client (returns clientID)
	v, err := SSOAuthenticator.Verify(tokSrc)
	ctx := context.Background()
	slog.Info("Checking to see if user already exists")
	user, err := database.GetUser(ctx, v.CharacterOwnerHash)
	if err == pgx.ErrNoRows {
		slog.Info("User does not exist, creating")
		createErr := database.CreateUser(ctx, v.CharacterOwnerHash)
		if createErr != nil {
			slog.Error("User Create Error", slog.Any("Error", createErr))
			return 0, createErr
		}
		user, err = database.GetUser(ctx, v.CharacterOwnerHash)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	slog.Info("User", slog.Any("User", user))
	// characters, err := database.GetUserCharacters(ctx, user.Characterownerhash)

	// slog.Info("characters", slog.Any("characters", characters))
	slog.Info("Character Info",
		slog.Any("CharacterInfo", v),
		slog.String("CharacterName", v.CharacterName),
		slog.String("CharacterOwnerHash", v.CharacterOwnerHash),
		slog.String("ExpiresOn", v.ExpiresOn),
		slog.String("Scopes", v.Scopes),
		slog.String("TokenType", v.TokenType))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Save the verification structure on the session for quick access.
	s.Put(r.Context(), "character", v)

	slog.Info("CharactersHas", slog.Any("charaterHas", v.CharacterOwnerHash))
	// Redirect to the front page for now.
	http.Redirect(w, r, "/", http.StatusFound)
	return http.StatusMovedPermanently, nil
}

func GetTokenContext(oauth2Token *oauth2.Token) context.Context {
	tokSrc := SSOAuthenticator.TokenSource(oauth2Token)
	ctx := context.WithValue(context.Background(), goesi.ContextOAuth2, tokSrc)
	return ctx
}
