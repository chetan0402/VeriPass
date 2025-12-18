package veripass_test

import (
	"context"
	"crypto/ed25519"
	"database/sql"
	"encoding/base64"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	veripass "github.com/chetan0402/veripass/internal"
	"github.com/chetan0402/veripass/internal/ent"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	"golang.org/x/net/publicsuffix"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/jackc/pgx/v5/stdlib" // Import the pgx driver for PostgreSQL
)

// TODO - test passClient create manual after create admin entity
func TestMain(t *testing.T) {
	timeout := time.After(30 * time.Second)
	host := "http://localhost:8000"
	dbUrl := "postgres://veripass:veripass@localhost:5432/veripass"
	go veripass.Run(&veripass.Config{
		DatabaseUrl:    dbUrl,
		OAuthServer:    "http://localhost:1433/dex",
		ClientID:       "veripass",
		ClientSecret:   "veripass",
		RedirectionURI: "http://localhost:8000/callback",
	})

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

	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	attest(t, err)

	jar.SetCookies(&url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
	}, []*http.Cookie{getToken(t)})
	client := &http.Client{Jar: jar}

	userClient := veripassv1connect.NewUserServiceClient(client, host)
	passClient := veripassv1connect.NewPassServiceClient(client, host)
	adminClient := veripassv1connect.NewAdminServiceClient(client, host)
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

	mockPass1 := veripassv1.Pass{
		UserId: mockUser.Id,
		Type:   veripassv1.Pass_PASS_TYPE_CLASS,
	}

	mockPass2 := veripassv1.Pass{
		UserId: mockUser.Id,
		Type:   veripassv1.Pass_PASS_TYPE_HOME,
	}

	mockAdmin := veripassv1.Admin{
		Email:      "test_email",
		Name:       "test_name",
		Hostel:     "test_hostel",
		CanAddPass: true,
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

	if err := dbClient.Admin.Create().
		SetEmail(mockAdmin.Email).
		SetName(mockAdmin.Name).
		SetHostel(mockAdmin.Hostel).
		SetCanAddPass(mockAdmin.CanAddPass).
		Exec(ctx); err != nil {
		if !ent.IsConstraintError(err) {
			t.Fatal(err)
		}
	}

	publicKey, err := adminClient.GetPublicKey(ctx, connect.NewRequest(&emptypb.Empty{}))
	attest(t, err)

	user, err := userClient.GetUser(ctx, connect.NewRequest(&veripassv1.GetUserRequest{
		Id: mockUser.Id,
	}))
	attest(t, err)
	if !proto.Equal(user.Msg, &mockUser) {
		t.Fatal("User response not equal")
	}

	exitResponse, err := userClient.Exit(ctx, connect.NewRequest(&veripassv1.ExitRequest{
		Id:   mockPass1.UserId,
		Type: veripassv1.ExitRequest_ExitType(mockPass1.Type),
	}))
	attest(t, err)

	mockPass1.Id = exitResponse.Msg.PassId
	mockPass1.StartTime = timestamppb.Now()

	_, err = userClient.Entry(ctx, connect.NewRequest(&veripassv1.EntryRequest{
		PassId: mockPass1.Id,
	}))
	attest(t, err)

	mockPass1.EndTime = timestamppb.Now()

	exitResponse2, err := userClient.Exit(ctx, connect.NewRequest(&veripassv1.ExitRequest{
		Id:   mockPass2.UserId,
		Type: veripassv1.ExitRequest_ExitType(mockPass2.Type),
	}))
	attest(t, err)

	mockPass2.Id = exitResponse2.Msg.PassId
	mockPass2.StartTime = timestamppb.Now()

	pass, err := passClient.GetPass(ctx, connect.NewRequest(&veripassv1.GetPassRequest{
		Id: mockPass1.Id,
	}))
	attest(t, err)

	failIfNotEqualPass(t, pass.Msg, &mockPass1, publicKey.Msg.PublicKey)

	pass2, err := passClient.GetPass(ctx, connect.NewRequest(&veripassv1.GetPassRequest{
		Id: mockPass2.Id,
	}))
	attest(t, err)

	failIfNotEqualPass(t, pass2.Msg, &mockPass2, publicKey.Msg.PublicKey)

	latestPass, err := passClient.GetLatestPassByUser(ctx, connect.NewRequest(&veripassv1.GetLatestPassByUserRequest{
		UserId: mockUser.Id,
	}))
	attest(t, err)

	failIfNotEqualPass(t, latestPass.Msg, &mockPass2, publicKey.Msg.PublicKey)

	pageToken := timestamppb.Now()

	passList1, err := passClient.ListPassesByUser(ctx, connect.NewRequest(&veripassv1.ListPassesByUserRequest{
		UserId:    mockUser.Id,
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

	passList2, err := passClient.ListPassesByUser(ctx, connect.NewRequest(&veripassv1.ListPassesByUserRequest{
		UserId:    mockUser.Id,
		Type:      veripassv1.Pass_PASS_TYPE_UNSPECIFIED.Enum(),
		PageToken: pageToken,
		PageSize:  1,
	}))
	attest(t, err)
	if passList2.Msg.NextPageToken != nil {
		t.Fatal("Expected nil next page token")
	}
	failIfNotEqualPass(t, passList2.Msg.Passes[0], pass.Msg, publicKey.Msg.PublicKey)

	admin, err := adminClient.GetAdmin(ctx, connect.NewRequest(&veripassv1.GetAdminRequest{
		Email: mockAdmin.Email,
	}))
	attest(t, err)

	if !proto.Equal(admin.Msg, &mockAdmin) {
		t.Fatalf("Expected %v, got %v", &mockAdmin, admin.Msg)
	}

	hostelPassList1, err := adminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
		Hostel:     "H mock",
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

	hostelPassList2, err := adminClient.GetAllPassesByHostel(ctx, connect.NewRequest(&veripassv1.GetAllPassesByHostelRequest{
		Hostel:     "H mock",
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

	outCount, err := adminClient.GetOutCountByHostel(ctx, connect.NewRequest(&veripassv1.GetOutCountByHostelRequest{
		Hostel:    "H mock",
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
	if err != nil {
		t.Helper()
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

	encodedSignedQrCode := pass.QrCode

	signedQrCodeBytes, err := base64.StdEncoding.DecodeString(encodedSignedQrCode)
	attest(t, err)

	signedQrCode := string(signedQrCodeBytes)

	pass_id, useridAndSignature, ok := strings.Cut(signedQrCode, "|")
	if !ok {
		t.Fatal("Invalid signed QR code format")
	}

	userid, signature, ok := strings.Cut(useridAndSignature, "|")
	if !ok {
		t.Fatal("Invalid signed QR code format")
	}

	message := pass_id + "|" + userid

	if !ed25519.Verify(publicKey, []byte(message), []byte(signature)) {
		t.Fatal("Invalid signature")
	}

	if pass_id != pass.Id {
		t.Fatalf("Expected pass ID %v, got %v", pass.Id, pass_id)
	}

	if userid != pass.UserId {
		t.Fatalf("Expected user ID %v, got %v", pass.UserId, userid)
	}
}

func getToken(t *testing.T) *http.Cookie {
	jar, err := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get("http://localhost:1433/dex/auth?client_id=veripass&redirect_uri=http://localhost:8000/callback&response_type=code&scope=openid")
	attest(t, err)

	dexLoginFormURL := resp.Request.URL.String()

	formData := url.Values{}
	formData.Set("login", "student@example.com")
	formData.Set("password", "veripass")

	resp, err = client.PostForm(dexLoginFormURL, formData)
	attest(t, err)

	formData = url.Values{}
	formData.Set("req", resp.Request.URL.Query().Get("req"))
	formData.Set("approval", "approve")

	resp, err = client.PostForm(resp.Request.URL.String(), formData)
	attest(t, err)

	for _, c := range jar.Cookies(&url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
	}) {
		if c.Name == "token" {
			return c
		}
	}

	t.Fatalf("Couldn't get token from dex")
	return nil
}
