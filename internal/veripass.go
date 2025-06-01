package veripass

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
	err := http.ListenAndServe("0.0.0.0:8000", h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		panic(err)
	}
}
