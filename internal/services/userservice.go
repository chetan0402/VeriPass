package userservice

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/pass"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	client *ent.Client
}

var _ veripassv1connect.UserServiceHandler = (*UserService)(nil)

func New(client *ent.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) Entry(ctx context.Context, r *connect.Request[veripassv1.EntryRequest]) (*connect.Response[emptypb.Empty], error) {
	var (
		passId = r.Msg.PassId
	)

	passUUID, err := uuid.Parse(passId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if err := s.client.Pass.UpdateOneID(passUUID).SetEndTime(time.Now()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *UserService) Exit(ctx context.Context, r *connect.Request[veripassv1.ExitRequest]) (*connect.Response[veripassv1.ExitResponse], error) {
	var (
		id       = r.Msg.Id
		passType = pass.TypeUnspecified
	)

	switch r.Msg.Type {
	case veripassv1.ExitRequest_EXIT_TYPE_CLASS:
		passType = pass.TypeClass
	case veripassv1.ExitRequest_EXIT_TYPE_MARKET:
		passType = pass.TypeMarket
	case veripassv1.ExitRequest_EXIT_TYPE_HOME:
		passType = pass.TypeHome
	case veripassv1.ExitRequest_EXIT_TYPE_EVENT:
		passType = pass.TypeEvent
	}

	var passId = uuid.New()

	// TODO - after OAuth is done, use tokens instead of user id.
	if _, err := s.client.User.Get(ctx, id); err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	if err := s.client.Pass.Create().
		SetID(passId).
		SetUserID(id).
		SetType(passType).
		SetStartTime(time.Now()).
		Exec(ctx); err != nil {
		return nil, err
	}

	return connect.NewResponse(&veripassv1.ExitResponse{
		PassId: passId.String(),
	}), nil
}

func (s *UserService) GetUser(ctx context.Context, r *connect.Request[veripassv1.GetUserRequest]) (*connect.Response[veripassv1.User], error) {
	var (
		id = r.Msg.Id
	)

	user, err := s.client.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	return connect.NewResponse(&veripassv1.User{
		Id:     user.ID,
		Name:   user.Name,
		Room:   user.Room,
		Hostel: user.Hostel,
		Phone:  user.Phone,
	}), nil
}
