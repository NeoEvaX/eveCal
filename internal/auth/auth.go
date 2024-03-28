package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"net/http"
	"os"

	"github.com/antihax/goesi"
)

var (
	ESI              *goesi.APIClient
	SSOAuthenticator *goesi.SSOAuthenticator
	scopes           []string
)

func SetupESI() {
	// create ESI client
	httpClient := &http.Client{}
	// call Status endpoint
	scopes = []string{"esi-calendar.respond_calendar_events.v1", "esi-calendar.read_calendar_events.v1"}
	ESI = goesi.NewAPIClient(httpClient, "EveCal (ian.kremer@gmail.com, @neoevax on Discord")
	SSOAuthenticator = goesi.NewSSOAuthenticator(httpClient, os.Getenv("CLIENT_ID"), os.Getenv("SECRET_KEY"), "http://localhost:3000/api/esi/callback", scopes)
}

func EveSSO(w http.ResponseWriter, r *http.Request) (int, error) {

	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Save the state on the session
	// s.Values["state"] = state
	// err := s.Save(r, w)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// Generate the SSO URL with the state string
	url := SSOAuthenticator.AuthorizeURL(state, true, scopes)

	// Send the user to the URL
	http.Redirect(w, r, url, http.StatusFound)
	return http.StatusMovedPermanently, nil
}

func EveSSOAnswer(w http.ResponseWriter, r *http.Request) (int, error) {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// get our code and state
	code := r.FormValue("code")
	//state := r.FormValue("state")

	// Verify the state matches our randomly generated string from earlier.
	// if s.Values["state"] != state {
	// 	return http.StatusInternalServerError, errors.New("Invalid State.")
	// }

	// Exchange the code for an Access and Refresh token.
	token, err := SSOAuthenticator.TokenExchange(code)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	logger.Info("Token", slog.String("token", token.AccessToken))

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc := SSOAuthenticator.TokenSource(token)

	// Assign an auth context to the calls
	//auth := context.WithValue(context.TODO(), goesi.ContextOAuth2, tokSrc.Token)

	// Verify the client (returns clientID)
	v, _ := SSOAuthenticator.Verify(tokSrc)
	logger.Info("CharacterName", slog.String("CharacterName", v.CharacterName))
	logger.Info("CharacterOwnerHash", slog.String("CharacterOwnerHash", v.CharacterOwnerHash))
	logger.Info("ExpiresOn", slog.String("ExpiresOn", v.ExpiresOn))
	logger.Info("Scopes", slog.String("Scopes", v.Scopes))
	logger.Info("TokenType", slog.String("TokenType", v.TokenType))
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// Save the verification structure on the session for quick access.
	// s.Values["character"] = v
	// err = s.Save(r, w)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }

	// Redirect to the front page for now.
	http.Redirect(w, r, "/", http.StatusFound)
	return http.StatusMovedPermanently, nil
}
