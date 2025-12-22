package adminservice

import (
	"context"
	"crypto/ed25519"
	"time"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/admin"
	"github.com/chetan0402/veripass/internal/ent/pass"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/services"
	passservice "github.com/chetan0402/veripass/internal/services/pass"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	client    *ent.Client
	publicKey ed25519.PublicKey
}

var _ veripassv1connect.AdminServiceHandler = (*AdminService)(nil)

func New(client *ent.Client, publicKey ed25519.PublicKey) *AdminService {
	return &AdminService{
		client:    client,
		publicKey: publicKey,
	}
}

// GetPublicKey implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetPublicKey(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[veripassv1.GetPublicKeyResponse], error) {
	return connect.NewResponse(&veripassv1.GetPublicKeyResponse{
		PublicKey: s.publicKey,
	}), nil
}

// GetAllPassesByHostel implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetAllPassesByHostel(ctx context.Context, r *connect.Request[veripassv1.GetAllPassesByHostelRequest]) (*connect.Response[veripassv1.GetAllPassesByHostelResponse], error) {
	var (
		_            = r.Msg.Hostel
		start_time   = r.Msg.StartTime.AsTime()
		end_time     = r.Msg.EndTime.AsTime()
		pass_is_open = r.Msg.PassIsOpen
		pass_type    = r.Msg.Type
		page_size    = int(r.Msg.PageSize)
		page_token   = r.Msg.PageToken
	)

	query := s.client.Pass.Query().Order(
		pass.ByID(sql.OrderDesc()),
	).Limit(page_size + 1)

	query = query.Where(
		pass.IDGTE(veripass.ToUUIDv7Nil(start_time)),
		pass.IDLTE(veripass.ToUUIDv7Max(end_time)),
		pass.IDLTE(veripass.ToUUIDv7Max(page_token.AsTime())),
	)

	if pass_is_open != nil {
		if *pass_is_open {
			query = query.Where(pass.EndTimeIsNil())
		} else {
			query = query.Where(pass.EndTimeNotNil())
		}
	}

	if pass_type != veripassv1.Pass_PASS_TYPE_UNSPECIFIED {
		query = query.Where(pass.TypeEQ(passservice.ProtoPassTypeToEnt(pass_type)))
	}

	// TODO - Filter by hostel

	passes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	response := &veripassv1.GetAllPassesByHostelResponse{}

	if len(passes) > page_size {
		response.NextPageToken = timestamppb.New(time.Unix(passes[len(passes)-1].ID.Time().UnixTime()))
	}

	for index, pass := range passes {
		if index == page_size {
			break
		}
		user, err := pass.QueryUser().Only(ctx)
		if err != nil {
			return nil, err
		}
		protoPass, err := passservice.ToProto(pass)
		if err != nil {
			return nil, err
		}
		response.Passes = append(response.Passes, &veripassv1.GetAllPassesByHostelResponse_InfoIncludedPass{
			Pass:        protoPass,
			StudentName: user.Name,
			StudentRoom: user.Room,
		})
	}

	return connect.NewResponse(response), nil
}

// GetAdmin implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetAdmin(ctx context.Context, r *connect.Request[emptypb.Empty]) (*connect.Response[veripassv1.Admin], error) {
	var (
		email = veripass.GetEmailFromCtx(ctx)
	)

	admin, err := s.client.Admin.Query().Where(admin.Email(email)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, err
	}

	return connect.NewResponse(toProto(admin)), nil
}

// GetOutCountByHostel implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetOutCountByHostel(ctx context.Context, req *connect.Request[veripassv1.GetOutCountByHostelRequest]) (*connect.Response[veripassv1.GetOutCountByHostelResponse], error) {
	var (
		_         = req.Msg.Hostel
		startTime = req.Msg.StartTime.AsTime()
		endTime   = req.Msg.EndTime.AsTime()
		passType  = req.Msg.Type
	)

	query := s.client.Pass.Query().Where(
		pass.IDGTE(veripass.ToUUIDv7Nil(startTime)),
		pass.IDLTE(veripass.ToUUIDv7Max(endTime)),
		pass.EndTimeIsNil(),
	)

	if passType != veripassv1.Pass_PASS_TYPE_UNSPECIFIED {
		query = query.Where(
			pass.TypeEQ(passservice.ProtoPassTypeToEnt(passType)),
		)
	}

	out, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&veripassv1.GetOutCountByHostelResponse{
		Out: int64(out),
	}), nil
}

func toProto(admin *ent.Admin) *veripassv1.Admin {
	return &veripassv1.Admin{
		Email:      admin.Email,
		Name:       admin.Name,
		Hostel:     admin.Hostel,
		CanAddPass: admin.CanAddPass,
	}
}
