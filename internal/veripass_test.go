package veripass_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	veripass "github.com/chetan0402/veripass/internal"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
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

	client := veripassv1connect.NewUserServiceClient(&http.Client{}, host)
	ctx := context.Background()

	_, err := client.GetUser(ctx, connect.NewRequest(&veripassv1.GetUserRequest{
		Id: "test_id",
	}))
	if connect.CodeOf(err) != connect.CodeNotFound {
		t.Log(err)
		t.FailNow()
	}
}
