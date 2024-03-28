package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/neoevax/eveCal/internal/auth"
	"github.com/neoevax/eveCal/internal/db"
	"github.com/neoevax/eveCal/internal/handlers"

	m "github.com/neoevax/eveCal/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/jritsema/gotoolbox"
)

func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func main() {

	//load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db.Setup(logger)
	auth.SetupESI()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			m.TextHTMLMiddleware,
			m.CSPMiddleware,
		)

		//r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)

		r.Get("/", handlers.NewHomeHandler(handlers.GetHomeHandlerParams{
			UserStore: db.DB,
		}).ServeHTTP)

		r.Get("/api/esi/callback", handlers.NewGetEsiCallbackHandler().ServeHTTP)

		r.Get("/auth/esi/login", handlers.NewGetEsiAuthHandler().ServeHTTP)
	})

	//exit process immediately upon sigterm
	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	port := gotoolbox.GetEnvWithDefault("PORT", "8080")
	url := gotoolbox.GetEnvWithDefault("LISTEN_URL", "localhost")

	srv := &http.Server{
		Addr:    url + ":" + port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", port))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")

}
