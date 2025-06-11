package veripass_test

import (
	"context"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
)

func TestMain(t *testing.T) {
	host := "http://localhost:5002/api/"

	for {
		if _, err := http.DefaultClient.Get(host + "ping"); err == nil {
			break
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
