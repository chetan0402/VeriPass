package adminservice

import (
	"context"

	"connectrpc.com/connect"
	"github.com/chetan0402/veripass/internal/ent"
	"github.com/chetan0402/veripass/internal/ent/admin"
	veripassv1 "github.com/chetan0402/veripass/internal/gen/veripass/v1"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
)

type AdminService struct {
	client *ent.Client
}

// GetAllPassesByHostel implements veripassv1connect.AdminServiceHandler.
func (s *AdminService) GetAllPassesByHostel(context.Context, *connect.Request[veripassv1.GetAllPassesByHostelRequest]) (*connect.Response[veripassv1.GetAllPassesByHostelResponse], error) {
	panic("unimplemented")
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
