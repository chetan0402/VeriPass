package adminservice

import (
	"context"

	"connectrpc.com/connect"
	"entgo.io/ent/dialect/sql"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/admin"
	"github.com/chetan0402/veripass/internal/ent/pass"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	passservice "github.com/chetan0402/veripass/internal/services/pass"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	client *ent.Client
}

// GetAllPassesByHostel implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetAllPassesByHostel(ctx context.Context, r *connect.Request[veripassv1.GetAllPassesByHostelRequest]) (*connect.Response[veripassv1.GetAllPassesByHostelResponse], error) {
	var (
		_            = r.Msg.Hostel
		start_time   = r.Msg.StartTime.AsTime()
		pass_is_open = r.Msg.PassIsOpen
		pass_type    = r.Msg.Type
		page_size    = int(r.Msg.PageSize)
		page_token   = r.Msg.PageToken
	)

	query := s.client.Pass.Query().Order(
		pass.ByStartTime(sql.OrderDesc()),
	).Limit(page_size + 1)

	query = query.Where(
		pass.StartTimeGTE(start_time),
		pass.StartTimeLTE(page_token.AsTime()),
	)

	if pass_is_open {
		query = query.Where(pass.EndTimeIsNil())
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
		response.NextPageToken = timestamppb.New(passes[len(passes)-1].StartTime)
	}

	for index, pass := range passes {
		if index == int(page_size) {
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

var _ veripassv1connect.AdminServiceHandler = (*AdminService)(nil)

func New(client *ent.Client) *AdminService {
	return &AdminService{
		client: client,
	}
}

// GetAdmin implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetAdmin(ctx context.Context, r *connect.Request[veripassv1.GetAdminRequest]) (*connect.Response[veripassv1.Admin], error) {
	var (
		email = r.Msg.Email
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

func toProto(admin *ent.Admin) *veripassv1.Admin {
	return &veripassv1.Admin{
		Email:      admin.Email,
		Name:       admin.Name,
		Hostel:     admin.Hostel,
		CanAddPass: admin.CanAddPass,
	}
}
