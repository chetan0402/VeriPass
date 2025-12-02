package veripass

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/middleware"
	adminservice "github.com/chetan0402/veripass/internal/services/admin"
	passservice "github.com/chetan0402/veripass/internal/services/pass"
	userservice "github.com/chetan0402/veripass/internal/services/user"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

func Run(databaseUrl string) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", databaseUrl)
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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	mux.Handle(veripassv1connect.NewUserServiceHandler(userservice.New(client), interceptor))
	mux.Handle(veripassv1connect.NewPassServiceHandler(passservice.New(client, privateKey), interceptor))
	mux.Handle(veripassv1connect.NewAdminServiceHandler(adminservice.New(client, publicKey), interceptor))
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", h2c.NewHandler(mux, &http2.Server{})))
}
