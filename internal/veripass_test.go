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
	mockUser, mockAdmin := getUserAdminPair(t, dbClient, dexClient, true)

	studentClient := getClient(t, mockUser.Id+"@stu.manit.ac.in")
	adminClient := getClient(t, mockAdmin.Email)

	studentUserClient := veripassv1connect.NewUserServiceClient(studentClient, HOST)
	studentPassClient := veripassv1connect.NewPassServiceClient(studentClient, HOST)

	adminAdminClient := veripassv1connect.NewAdminServiceClient(adminClient, HOST)

	mockPass1 := veripassv1.Pass{
		UserId: mockUser.Id,
		Type:   veripassv1.Pass_PASS_TYPE_CLASS,
	}

	mockPass2 := veripassv1.Pass{
		UserId: mockUser.Id,
		Type:   veripassv1.Pass_PASS_TYPE_HOME,
	}

	publicKey, err := adminAdminClient.GetPublicKey(ctx, connect.NewRequest(&emptypb.Empty{}))
	attest(t, err)

	user, err := studentUserClient.GetUser(ctx, connect.NewRequest(&veripassv1.GetUserRequest{
		Id: nil,
	}))
	attest(t, err)
	if !proto.Equal(user.Msg, mockUser) {
		t.Fatal("User response not equal")
	}

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
	if passList1.Msg.NextPageToken.Seconds != pass.Msg.StartTime.Seconds {
		t.Fatalf("Expected %v, got %v", pass.Msg.StartTime, passList1.Msg.NextPageToken)
	}
	failIfNotEqualPass(t, passList1.Msg.Passes[0], pass2.Msg, publicKey.Msg.PublicKey)
	pageToken = passList1.Msg.NextPageToken

	passList2, err := studentPassClient.ListPassesByUser(ctx, connect.NewRequest(&veripassv1.ListPassesByUserRequest{
		Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED.Enum(),
		PageToken: pageToken,
		PageSize:  1,
	}))
	attest(t, err)
	if passList2.Msg.NextPageToken != nil {
		t.Fatal("Expected nil next page token")
	}
	failIfNotEqualPass(t, passList2.Msg.Passes[0], pass.Msg, publicKey.Msg.PublicKey)

	admin, err := adminAdminClient.GetAdmin(ctx, connect.NewRequest(&emptypb.Empty{}))
	attest(t, err)

	if !proto.Equal(admin.Msg, mockAdmin) {
		t.Fatalf("Expected %v, got %v", &mockAdmin, admin.Msg)
	}

	hostelPassList1, err := adminAdminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
		Hostel:     mockAdmin.Hostel,
		StartTime:  timestamppb.New(time.Unix(0, 0)),
		EndTime:    timestamppb.Now(),
		PassIsOpen: nil,
		Type:       veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
		PageSize:   1,
		PageToken:  timestamppb.Now(),
	}))
	attest(t, err)
	if (hostelPassList1.Msg.NextPageToken == nil) || hostelPassList1.Msg.NextPageToken.Seconds != pass.Msg.StartTime.Seconds {
		t.Fatalf("Expected %v, got %v", pass.Msg.StartTime, hostelPassList1.Msg.NextPageToken)
	}
	if hostelPassList1.Msg.Passes[0].StudentName != mockUser.Name {
		t.Fatalf("Expected student name %v, got %v", mockUser.Name, hostelPassList1.Msg.Passes[0].StudentName)
	}
	if hostelPassList1.Msg.Passes[0].StudentRoom != mockUser.Room {
		t.Fatalf("Expected student room %v, got %v", mockUser.Room, hostelPassList1.Msg.Passes[0].StudentRoom)
	}
	failIfNotEqualPass(t, hostelPassList1.Msg.Passes[0].Pass, pass2.Msg, publicKey.Msg.PublicKey)

	hostelPassList2, err := adminAdminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
		Hostel:     mockAdmin.Hostel,
		StartTime:  timestamppb.New(time.Unix(0, 0)),
		EndTime:    timestamppb.Now(),
		PassIsOpen: nil,
		Type:       veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
		PageSize:   1,
		PageToken:  hostelPassList1.Msg.NextPageToken,
	}))
	attest(t, err)
	if hostelPassList2.Msg.NextPageToken != nil {
		t.Fatal("Expected nil next page token")
	}
	if hostelPassList2.Msg.Passes[0].StudentName != mockUser.Name {
		t.Fatalf("Expected student name %v, got %v", mockUser.Name, hostelPassList2.Msg.Passes[0].StudentName)
	}
	if hostelPassList2.Msg.Passes[0].StudentRoom != mockUser.Room {
		t.Fatalf("Expected student room %v, got %v", mockUser.Room, hostelPassList2.Msg.Passes[0].StudentRoom)
	}
	failIfNotEqualPass(t, hostelPassList2.Msg.Passes[0].Pass, pass.Msg, publicKey.Msg.PublicKey)

	outCount, err := adminAdminClient.GetOutCountByHostel(ctx, connect.NewRequest(&veripassv1.GetOutCountByHostelRequest{
		Hostel:    mockAdmin.Hostel,
		StartTime: timestamppb.New(time.Unix(0, 0)),
		EndTime:   timestamppb.Now(),
		Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED,
	}))
	attest(t, err)
	if outCount.Msg.Out != 1 {
		t.Fatalf("Expected 1 out count, got:%v", outCount.Msg.Out)
	}
}

func attest(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func failIfNotEqualPass(t *testing.T, got *veripassv1.Pass, expected *veripassv1.Pass, publicKey ed25519.PublicKey) {
	t.Helper()
	verifyPass(t, got, publicKey)
	if got.Id != expected.Id ||
		got.UserId != expected.UserId ||
		got.Type != expected.Type ||
		math.Abs(float64(got.StartTime.Seconds-expected.StartTime.Seconds)) > 1 ||
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
	if passID != pass.Id {
		t.Fatalf("Expected pass ID %v, got %v", pass.Id, passID)
	}

	userID := string(signedQrCode[37:secondDelimemter])
	if userID != pass.UserId {
		t.Fatalf("Expected user ID %v, got %v", pass.UserId, userID)
	}
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
