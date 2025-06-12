package veripass_test

import (
	"net/http"
	"testing"
	"time"

	veripass "github.com/chetan0402/veripass/internal"
)

func TestMain(t *testing.T) {
	timeout := time.After(30 * time.Second)
	host := "http://localhost:8000"
	go veripass.Run("postgres://veripass:veripass@localhost:5432/veripass")

	for {
		if _, err := http.DefaultClient.Get(host + "/ping"); err == nil {
			break
		}
		select {
		case <-timeout:
			t.Fatal("Could not connect to server")
		default:
		}
	}
}
