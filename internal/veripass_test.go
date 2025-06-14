package veripass_test

import (
	"context"
	"database/sql"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	veripass "github.com/chetan0402/veripass/internal"
	"github.com/chetan0402/veripass/internal/ent"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	"google.golang.org/protobuf/proto"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

func TestMain(t *testing.T) {
	timeout := time.After(30 * time.Second)
	host := "http://localhost:8000"
	dbUrl := "postgres://veripass:veripass@localhost:5432/veripass"
	go veripass.Run(dbUrl)

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

	db, err := sql.Open("pgx", dbUrl)
	attest(t, err)

	dbClient := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	mockUser := veripassv1.User{
		Id:     "test_id",
		Name:   "test_name",
		Room:   "test_room",
		Hostel: "test_hostel",
		Phone:  "1234567890",
	}

	if err := dbClient.User.Create().
		SetID(mockUser.Id).
		SetName(mockUser.Name).
		SetRoom(mockUser.Room).
		SetHostel(mockUser.Hostel).
		SetPhone(mockUser.Phone).
		Exec(ctx); err != nil {
		if !ent.IsConstraintError(err) {
			t.Fatal(err)
		}
	}

	user, err := client.GetUser(ctx, connect.NewRequest(&veripassv1.GetUserRequest{
		Id: mockUser.Id,
	}))
	attest(t, err)
	if !proto.Equal(user.Msg, &mockUser) {
		t.Fatal("User response not equal")
	}

	exitResponse, err := client.Exit(ctx, connect.NewRequest(&veripassv1.ExitRequest{
		Id:   mockUser.Id,
		Type: veripassv1.ExitRequest_EXIT_TYPE_CLASS,
	}))
	attest(t, err)

	_, err = client.Entry(ctx, connect.NewRequest(&veripassv1.EntryRequest{
		PassId: exitResponse.Msg.PassId,
	}))
	attest(t, err)
}

func attest(t *testing.T, err error) {
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}
