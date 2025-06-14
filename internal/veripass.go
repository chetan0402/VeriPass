package veripass

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	userservice "github.com/chetan0402/veripass/internal/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

func Run(databaseUrl string) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	mux.Handle(veripassv1connect.NewUserServiceHandler(userservice.New(client)))
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", h2c.NewHandler(mux, &http2.Server{})))
}
