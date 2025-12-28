// Package passservice implements PassService
package passservice

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/admin"
	"github.com/chetan0402/veripass/internal/ent/pass"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/services"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// PassService represents PassService defined in proto
type PassService struct {
	client *ent.Client
}

var _ veripassv1connect.PassServiceHandler = (*PassService)(nil)
var passPrivateKey ed25519.PrivateKey

// New returns an instance of PassService and sets unexported fields
func New(client *ent.Client, privateKey ed25519.PrivateKey) *PassService {
	passPrivateKey = privateKey
	return &PassService{
		client: client,
	}
}

// CreateManualPass implements veripassv1connect.PassServiceHandler
func (s *PassService) CreateManualPass(ctx context.Context, r *connect.Request[veripassv1.CreateManualPassRequest]) (*connect.Response[veripassv1.Pass], error) {
	var (
		adminEmail = veripass.GetEmailFromCtx(ctx)
		userId     = r.Msg.UserId
		passType   = ProtoPassTypeToEnt(r.Msg.Type)
	)

	admin, err := s.client.Admin.Query().Where(admin.Email(adminEmail)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}
	if !admin.CanAddPass {
		return nil, connect.NewError(connect.CodePermissionDenied, err)
	}

	passId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	timeNow := time.Unix(passId.Time().UnixTime())

	if err := s.client.Pass.Create().SetID(passId).SetUserID(userId).SetType(passType).Exec(ctx); err != nil {
		return nil, err
	}

	return connect.NewResponse(&veripassv1.Pass{
		Id:        passId.String(),
		UserId:    userId,
		Type:      r.Msg.Type,
		StartTime: timestamppb.New(timeNow),
	}), nil
}

// GetLatestPassByUser implements veripassv1connect.PassServiceHandler
func (s *PassService) GetLatestPassByUser(ctx context.Context, r *connect.Request[emptypb.Empty]) (*connect.Response[veripassv1.Pass], error) {
	var (
		userId = veripass.GetUsernamefromCtx(ctx)
	)

	entPass, err := s.client.Pass.Query().Where(
		pass.UserID(userId),
	).Order(
		pass.ByID(sql.OrderDesc()),
	).First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	protoPass, err := ToProto(entPass)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(protoPass), nil
}

// GetPass implements veripassv1connect.PassServiceHandler
func (s *PassService) GetPass(ctx context.Context, r *connect.Request[veripassv1.GetPassRequest]) (*connect.Response[veripassv1.Pass], error) {
	var (
		passId = r.Msg.Id
	)

	passUUID, err := uuid.Parse(passId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	pass, err := s.client.Pass.Get(ctx, passUUID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	protoPass, err := ToProto(pass)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(protoPass), nil
}

// ListPassesByUser implements veripassv1connect.PassServiceHandler
//
// Cursor based pagination is used on ID (UUIDv7)
func (s *PassService) ListPassesByUser(ctx context.Context, r *connect.Request[veripassv1.ListPassesByUserRequest]) (*connect.Response[veripassv1.ListPassesByUserResponse], error) {
	var (
		userId    = veripass.GetUsernamefromCtx(ctx)
		pageToken = r.Msg.PageToken
		pageSize  = r.Msg.PageSize
		passType  = r.Msg.Type
		startTime = r.Msg.StartTime
		endTime   = r.Msg.EndTime
	)

	if pageToken == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid page token"))
	}

	query := s.client.Pass.Query().
		Order(pass.ByID(sql.OrderDesc())).
		Where(
			pass.UserID(userId),
		).Limit(int(pageSize) + 1)

	if passType != nil && *passType != veripassv1.Pass_PASS_TYPE_UNSPECIFIED {
		query = query.Where(
			pass.TypeEQ(ProtoPassTypeToEnt(*passType)),
		)
	}

	if startTime != nil {
		query = query.Where(
			pass.IDGTE(veripass.ToUUIDv7Nil(startTime.AsTime())),
		)
	}

	if endTime != nil {
		query = query.Where(
			pass.IDLTE(veripass.ToUUIDv7Max(endTime.AsTime())),
		)
	}

	passes, err := query.Where(
		pass.IDLTE(veripass.ToUUIDv7Max(pageToken.AsTime())),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	response := &veripassv1.ListPassesByUserResponse{}

	if len(passes) == int(pageSize)+1 {
		response.NextPageToken = timestamppb.New(time.Unix(passes[len(passes)-1].ID.Time().UnixTime()))
	}

	for index, pass := range passes {
		if index == int(pageSize) {
			break
		}
		protoPass, err := ToProto(pass)
		if err != nil {
			return nil, err
		}
		response.Passes = append(response.Passes, protoPass)
	}

	return connect.NewResponse(response), nil
}

// ToProto returns Pass entity
func ToProto(entPass *ent.Pass) (*veripassv1.Pass, error) {
	passType := veripassv1.Pass_PASS_TYPE_UNSPECIFIED

	switch entPass.Type {
	case pass.TypeClass:
		passType = veripassv1.Pass_PASS_TYPE_CLASS
	case pass.TypeMarket:
		passType = veripassv1.Pass_PASS_TYPE_MARKET
	case pass.TypeHome:
		passType = veripassv1.Pass_PASS_TYPE_HOME
	case pass.TypeEvent:
		passType = veripassv1.Pass_PASS_TYPE_EVENT
	}

	qrCode := entPass.ID.String() + "|" + entPass.UserID

	signature := ed25519.Sign(passPrivateKey, []byte(qrCode))

	signedQrCode := base64.StdEncoding.EncodeToString(append([]byte(qrCode+"|"), signature...))

	protoPass := &veripassv1.Pass{
		Id:        entPass.ID.String(),
		UserId:    entPass.UserID,
		Type:      passType,
		StartTime: timestamppb.New(time.Unix(entPass.ID.Time().UnixTime())),
		QrCode:    signedQrCode,
	}

	if !entPass.EndTime.IsZero() {
		protoPass.EndTime = timestamppb.New(entPass.EndTime)
	}

	return protoPass, nil
}

// ProtoPassTypeToEnt returns corresponding proto pass type or unspecified
func ProtoPassTypeToEnt(passType veripassv1.Pass_PassType) pass.Type {
	switch passType {
	case veripassv1.Pass_PASS_TYPE_CLASS:
		return pass.TypeClass
	case veripassv1.Pass_PASS_TYPE_MARKET:
		return pass.TypeMarket
	case veripassv1.Pass_PASS_TYPE_HOME:
		return pass.TypeHome
	case veripassv1.Pass_PASS_TYPE_EVENT:
		return pass.TypeEvent
	default:
		return pass.TypeUnspecified
	}
}
