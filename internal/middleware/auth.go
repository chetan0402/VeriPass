package veripass

import (
	"context"
	"errors"
	"net/http"

	"connectrpc.com/connect"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
	veripass "github.com/chetan0402/veripass/internal/services"
	"github.com/coreos/go-oidc/v3/oidc"
)

var isAuthProtected = map[string]bool{
	veripassv1connect.AdminServiceGetAdminProcedure:             true,
	veripassv1connect.AdminServiceGetAllPassesByHostelProcedure: true,
	veripassv1connect.AdminServiceGetOutCountByHostelProcedure:  true,

	veripassv1connect.PassServiceCreateManualPassProcedure:    true,
	veripassv1connect.PassServiceGetPassProcedure:             true,
	veripassv1connect.PassServiceGetLatestPassByUserProcedure: true,
	veripassv1connect.PassServiceListPassesByUserProcedure:    true,

	veripassv1connect.UserServiceEntryProcedure:    true,
	veripassv1connect.UserServiceExitProcedure:     true,
	veripassv1connect.UserServiceGetPhotoProcedure: true,
	veripassv1connect.UserServiceGetUserProcedure:  true,
}

func NewAuthMiddleware(verifier *oidc.IDTokenVerifier) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if isAuthProtected[req.Spec().Procedure] {
				cookies, err := http.ParseCookie(req.Header().Get("Cookie"))
				if err != nil {
					return nil, connect.NewError(connect.CodeInvalidArgument, err)
				}

				for _, c := range cookies {
					if c.Name == "token" {
						token, err := verifier.Verify(ctx, c.Value)
						if err != nil {
							return nil, connect.NewError(connect.CodeUnauthenticated, err)
						}

						var claims struct {
							Email string `json:"email,omitempty"`
							Name  string `json:"name,omitempty"`
						}

						if err := token.Claims(&claims); err != nil {
							return nil, connect.NewError(connect.CodeInternal, errors.New("token doesn't contain required claims"))
						}

						ctx = veripass.GetCtxWithEmail(ctx, claims.Email)
						ctx = veripass.GetCtxWithUsername(ctx, claims.Name)

						// TODO: check student or admin
						return next(ctx, req)
					}
				}

				return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("no token cookie found"))
			}
			return next(ctx, req)
		}
	}
}
