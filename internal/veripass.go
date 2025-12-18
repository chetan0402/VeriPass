package veripass

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/middleware"
	adminservice "github.com/chetan0402/veripass/internal/services/admin"
	passservice "github.com/chetan0402/veripass/internal/services/pass"
	userservice "github.com/chetan0402/veripass/internal/services/user"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/oauth2"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

type Config struct {
	DatabaseUrl    string
	OAuthServer    string
	ClientID       string
	ClientSecret   string
	RedirectionURI string
}

func Run(config *Config) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	interceptor := connect.WithInterceptors(
		veripass.NewIpMiddleware(),
	)

	provider, err := oidc.NewProvider(ctx, config.OAuthServer)
	if err != nil {
		log.Fatal(err)
	}

	oauth2config := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectionURI,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	mux.HandleFunc("GET /callback", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: oauth2config.ClientID})

		oauth2Token, err := oauth2config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if _, err := verifier.Verify(ctx, rawIDToken); err != nil {
			http.Error(w, "Bad credentials", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    rawIDToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
		if r.URL.Query().Get("state") == "admin" {
			http.Redirect(w, r, "/admin", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})
	mux.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("redirect") {
			http.Error(w, "No redirect param set.", http.StatusBadRequest)
			return
		}

		redirect := r.URL.Query().Get("redirect")

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now(),
			MaxAge:   -1,
		})
		http.Redirect(w, r, redirect, http.StatusFound)
	})
	mux.Handle(veripassv1connect.NewUserServiceHandler(userservice.New(client), interceptor))
	mux.Handle(veripassv1connect.NewPassServiceHandler(passservice.New(client, privateKey), interceptor))
	mux.Handle(veripassv1connect.NewAdminServiceHandler(adminservice.New(client, publicKey), interceptor))
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", h2c.NewHandler(mux, &http2.Server{})))
}
