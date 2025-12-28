// Package userservice implements UserService
package userservice

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/pass"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/services"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserService implements UserService defined in proto
type UserService struct {
	client *ent.Client
}

var _ veripassv1connect.UserServiceHandler = (*UserService)(nil)

// New returns an instance of UserService and sets unexported fields
func New(client *ent.Client) *UserService {
	return &UserService{
		client: client,
	}
}

// Entry implements veripass veripassv1connect.UserServiceHandler
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

// Exit implements veripassv1connect.UserServiceHandler
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

	var passId, err = uuid.NewV7()
	if err != nil {
		return nil, err
	}

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
		Exec(ctx); err != nil {
		return nil, err
	}

	return connect.NewResponse(&veripassv1.ExitResponse{
		PassId: passId.String(),
	}), nil
}

// GetPhoto implements veripassv1connect.UserServiceHandler
func (s *UserService) GetPhoto(context.Context, *connect.Request[veripassv1.GetPhotoRequest]) (*connect.Response[veripassv1.GetPhotoResponse], error) {
	return nil, connect.NewError(connect.CodeInternal, nil)
}

// GetUser implements veripassv1connect.UserServiceHandler
func (s *UserService) GetUser(ctx context.Context, r *connect.Request[veripassv1.GetUserRequest]) (*connect.Response[veripassv1.User], error) {
	var (
		id = r.Msg.GetId()
	)

	if id == "" {
		id = veripass.GetUsernamefromCtx(ctx)
	}

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
