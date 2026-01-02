package veripass_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	veripass "github.com/chetan0402/veripass/internal"
	"github.com/chetan0402/veripass/internal/ent"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	dex "github.com/dexidp/dex/api/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

// HOST is URL of running backend
const HOST = "http://localhost:8000"

// VERIPASS_HASH is bcrypt hash of the password "veripass"
const VERIPASS_HASH = "$2y$10$6JmUm254bTxZSICSXtFWZ.joFmhkGZuHIMoXx6S7aHi2krB/BpsUm"

// CONFIG is veripass configuration for running server
var CONFIG = veripass.Config{
	DatabaseUrl:    "postgres://veripass:veripass@localhost:5432/veripass",
	OAuthServer:    "http://localhost:1433/dex",
	ClientID:       "veripass",
	ClientSecret:   "veripass",
	RedirectionURI: HOST + "/callback",
}

// TODO - test passClient create manual after create admin entity
func TestMain(t *testing.T) {
	ctx := t.Context()
	timeout := time.After(30 * time.Second)
	go veripass.Run(&CONFIG)

	for {
		if _, err := http.DefaultClient.Get(HOST + "/ping"); err == nil {
			break
		}
		select {
		case <-timeout:
			t.Fatal("Could not connect to server")
		default:
		}
	}

	db, err := sql.Open("pgx", CONFIG.DatabaseUrl)
	attest(t, err)

	dbClient := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))

	conn, err := grpc.NewClient("127.0.0.1:5557", grpc.WithTransportCredentials(insecure.NewCredentials()))
	attest(t, err)

	dexClient := dex.NewDexClient(conn)
	t.Run("Happy path", func(t *testing.T) {
		student, admin := getUserAdminPair(t, dbClient, dexClient, true)

		sc := getClient(t, student.Id+"@stu.manit.ac.in")
		ac := getClient(t, admin.Email)

		studentUserClient := veripassv1connect.NewUserServiceClient(sc, HOST)
		studentPassClient := veripassv1connect.NewPassServiceClient(sc, HOST)

		adminAdminClient := veripassv1connect.NewAdminServiceClient(ac, HOST)

		mockPass1 := veripassv1.Pass{
			UserId: student.Id,
			Type:   veripassv1.Pass_PASS_TYPE_CLASS,
		}

		mockPass2 := veripassv1.Pass{
			UserId: student.Id,
			Type:   veripassv1.Pass_PASS_TYPE_HOME,
		}

		publicKey, err := adminAdminClient.GetPublicKey(ctx, connect.NewRequest(&emptypb.Empty{}))
		attest(t, err)

		user, err := studentUserClient.GetUser(ctx, connect.NewRequest(&veripassv1.GetUserRequest{
			Id: nil,
		}))
		attest(t, err)
		failIfNotEqual(t, user.Msg, student)

		exitResponse, err := studentUserClient.Exit(ctx, connect.NewRequest(&veripassv1.ExitRequest{
			Type: veripassv1.ExitRequest_ExitType(mockPass1.Type),
		}))
		attest(t, err)

		mockPass1.Id = exitResponse.Msg.PassId
		mockPass1.StartTime = timestamppb.Now()

		_, err = studentUserClient.Entry(ctx, connect.NewRequest(&veripassv1.EntryRequest{
			PassId: mockPass1.Id,
		}))
		attest(t, err)

		mockPass1.EndTime = timestamppb.Now()

		exitResponse2, err := studentUserClient.Exit(ctx, connect.NewRequest(&veripassv1.ExitRequest{
			Type: veripassv1.ExitRequest_ExitType(mockPass2.Type),
		}))
		attest(t, err)

		mockPass2.Id = exitResponse2.Msg.PassId
		mockPass2.StartTime = timestamppb.Now()

		pass, err := studentPassClient.GetPass(ctx, connect.NewRequest(&veripassv1.GetPassRequest{
			PassId: mockPass1.Id,
		}))
		attest(t, err)

		failIfNotEqualPass(t, pass.Msg, &mockPass1, publicKey.Msg.PublicKey)

		pass2, err := studentPassClient.GetPass(ctx, connect.NewRequest(&veripassv1.GetPassRequest{
			PassId: mockPass2.Id,
		}))
		attest(t, err)

		failIfNotEqualPass(t, pass2.Msg, &mockPass2, publicKey.Msg.PublicKey)

		latestPass, err := studentPassClient.GetLatestPassByUser(ctx, connect.NewRequest(&emptypb.Empty{}))
		attest(t, err)

		failIfNotEqualPass(t, latestPass.Msg, &mockPass2, publicKey.Msg.PublicKey)

		pageToken := timestamppb.Now()

		passList1, err := studentPassClient.ListPassesByUser(ctx, connect.NewRequest(&veripassv1.ListPassesByUserRequest{
			Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED.Enum(),
			PageToken: pageToken,
			PageSize:  1,
		}))
		attest(t, err)
		failIfNotEqual(t, passList1.Msg.NextPageToken.Seconds, pass.Msg.StartTime.Seconds)
		failIfNotEqualPass(t, passList1.Msg.Passes[0], pass2.Msg, publicKey.Msg.PublicKey)
		pageToken = passList1.Msg.NextPageToken

		passList2, err := studentPassClient.ListPassesByUser(ctx, connect.NewRequest(&veripassv1.ListPassesByUserRequest{
			Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED.Enum(),
			PageToken: pageToken,
			PageSize:  1,
		}))
		attest(t, err)
		failIfNotEqual(t, passList2.Msg.NextPageToken, nil)
		failIfNotEqualPass(t, passList2.Msg.Passes[0], pass.Msg, publicKey.Msg.PublicKey)

		a, err := adminAdminClient.GetAdmin(ctx, connect.NewRequest(&emptypb.Empty{}))
		attest(t, err)
		failIfNotEqual(t, a.Msg, admin)

		hostelPassList1, err := adminAdminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
			Hostel:     admin.Hostel,
			StartTime:  timestamppb.New(time.Unix(0, 0)),
			EndTime:    timestamppb.Now(),
			PassIsOpen: nil,
			Type:       veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
			PageSize:   1,
			PageToken:  timestamppb.Now(),
		}))
		attest(t, err)
		failIfNotEqual(t, hostelPassList1.Msg.NextPageToken.Seconds, pass.Msg.StartTime.Seconds)
		failIfNotEqual(t, hostelPassList1.Msg.Passes[0].StudentName, student.Name)
		failIfNotEqual(t, hostelPassList1.Msg.Passes[0].StudentRoom, student.Room)
		failIfNotEqualPass(t, hostelPassList1.Msg.Passes[0].Pass, pass2.Msg, publicKey.Msg.PublicKey)

		hostelPassList2, err := adminAdminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
			Hostel:     admin.Hostel,
			StartTime:  timestamppb.New(time.Unix(0, 0)),
			EndTime:    timestamppb.Now(),
			PassIsOpen: nil,
			Type:       veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
			PageSize:   1,
			PageToken:  hostelPassList1.Msg.NextPageToken,
		}))
		attest(t, err)
		failIfNotEqual(t, hostelPassList2.Msg.NextPageToken, nil)
		failIfNotEqual(t, hostelPassList2.Msg.Passes[0].StudentName, student.Name)
		failIfNotEqual(t, hostelPassList2.Msg.Passes[0].StudentRoom, student.Room)
		failIfNotEqualPass(t, hostelPassList2.Msg.Passes[0].Pass, pass.Msg, publicKey.Msg.PublicKey)

		outCount, err := adminAdminClient.GetOutCountByHostel(ctx, connect.NewRequest(&veripassv1.GetOutCountByHostelRequest{
			Hostel:    admin.Hostel,
			StartTime: timestamppb.New(time.Unix(0, 0)),
			EndTime:   timestamppb.Now(),
			Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
		}))
		attest(t, err)
		failIfNotEqual(t, outCount.Msg.Out, 1)
	})
}

func attest(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func failIfNotEqual[T comparable](t *testing.T, got T, expected T) {
	t.Helper()
	failed := false
	switch any(got).(type) {
	case proto.Message:
		failed = !proto.Equal(any(got).(proto.Message), any(expected).(proto.Message))
	default:
		failed = (got != expected)
	}

	if failed {
		t.Fatalf("Expected: %v\nGot:%v", expected, got)
	}
}

func failIfNotEqualPass(t *testing.T, got *veripassv1.Pass, expected *veripassv1.Pass, publicKey ed25519.PublicKey) {
	t.Helper()
	verifyPass(t, got, publicKey)
	failIfNotEqual(t, got.Id, expected.Id)
	failIfNotEqual(t, got.UserId, expected.UserId)
	failIfNotEqual(t, got.Type, expected.Type)
	if math.Abs(float64(got.StartTime.Seconds-expected.StartTime.Seconds)) > 1 ||
		(expected.EndTime != nil && math.Abs(float64(got.EndTime.Seconds-expected.EndTime.Seconds)) > 1) {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func verifyPass(t *testing.T, pass *veripassv1.Pass, publicKey ed25519.PublicKey) {
	t.Helper()
	signedQrCode, err := base64.StdEncoding.DecodeString(pass.QrCode)
	attest(t, err)

	secondDelimemter := -1
	for i := 38; i < len(signedQrCode); i++ {
		if signedQrCode[i] == '|' {
			secondDelimemter = i
			break
		}
	}
	if secondDelimemter == -1 {
		t.Fatal("Invalid QR code format")
	}

	if !ed25519.Verify(publicKey, signedQrCode[:secondDelimemter], signedQrCode[secondDelimemter+1:]) {
		t.Log(string(signedQrCode))
		t.Fatalf("Invalid signature")
	}

	passID := string(signedQrCode[:36])
	failIfNotEqual(t, passID, pass.Id)

	userID := string(signedQrCode[37:secondDelimemter])
	failIfNotEqual(t, userID, pass.UserId)
}

func getClient(t *testing.T, email string) *http.Client {
	jar, err := cookiejar.New(nil)
	attest(t, err)
	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get("http://localhost:1433/dex/auth?client_id=veripass&redirect_uri=http://localhost:8000/callback&response_type=code&scope=openid%20profile%20email")
	attest(t, err)
	if resp.StatusCode >= 400 {
		t.Fatal(resp)
	}

	formData := url.Values{}
	formData.Set("login", email)
	formData.Set("password", "veripass")

	resp, err = client.PostForm(resp.Request.URL.String(), formData)
	attest(t, err)
	if resp.StatusCode >= 400 {
		t.Fatal(resp)
	}

	formData = url.Values{}
	formData.Set("req", resp.Request.URL.Query().Get("req"))
	formData.Set("approval", "approve")

	_, err = client.PostForm(resp.Request.URL.String(), formData)
	attest(t, err)

	for _, c := range jar.Cookies(&url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
	}) {
		if c.Name == "token" {
			return client
		}
	}

	t.Fatalf("Couldn't get token from dex")
	return nil
}

func getUserAdminPair(t *testing.T, db *ent.Client, dexClient dex.DexClient, canAdminAddPass bool) (*veripassv1.User, *veripassv1.Admin) {
	u := &veripassv1.User{
		Id:     rand.Text(),
		Name:   rand.Text(),
		Room:   rand.Text(),
		Hostel: rand.Text(),
		Phone:  rand.Text(),
	}
	a := &veripassv1.Admin{
		Email:      rand.Text() + "@manit.ac.in",
		Name:       rand.Text(),
		Hostel:     u.Hostel,
		CanAddPass: canAdminAddPass,
	}

	_, err := db.User.Create().
		SetID(u.Id).
		SetName(u.Name).
		SetRoom(u.Room).
		SetHostel(u.Hostel).
		SetPhone(u.Phone).
		Save(t.Context())
	attest(t, err)

	_, err = db.Admin.Create().
		SetEmail(a.Email).
		SetName(a.Name).
		SetHostel(a.Hostel).
		SetCanAddPass(a.CanAddPass).
		Save(t.Context())
	attest(t, err)

	_, err = dexClient.CreatePassword(t.Context(), &dex.CreatePasswordReq{
		Password: &dex.Password{
			Email:    u.Id + "@stu.manit.ac.in",
			Hash:     []byte(VERIPASS_HASH),
			Username: u.Id,
			UserId:   u.Id,
		},
	})
	attest(t, err)

	_, err = dexClient.CreatePassword(t.Context(), &dex.CreatePasswordReq{
		Password: &dex.Password{
			Email:    a.Email,
			Hash:     []byte(VERIPASS_HASH),
			Username: a.Name,
			UserId:   a.Name,
		},
	})
	attest(t, err)
	return u, a
}
