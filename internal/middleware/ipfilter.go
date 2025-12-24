package veripass

import (
	"context"
	"fmt"
	"net"
	"strings"

	"connectrpc.com/connect"
	"github.com/chetan0402/veripass/internal/gen/veripass/v1/veripassv1connect"
)

// ALLOWED_IPv4_RANGE is CIDR notation of IPv4s allowed to access
const ALLOWED_IPv4_RANGE = "0.0.0.0/0"

// ALLOWED_IPv6_RANGE is CIDR notation of IPv6s allowed to access
const ALLOWED_IPv6_RANGE = "::/0"

var isIPProtected = map[string]bool{
	veripassv1connect.UserServiceEntryProcedure: true,
	veripassv1connect.UserServiceExitProcedure:  true,
}

// NewIpMiddleware returns a interceptor which checks if client's IP is
// within the allowed IP ranges and blocks the request if it isn't
func NewIpMiddleware() connect.UnaryInterceptorFunc {
	_, ipv4Net, err := net.ParseCIDR(ALLOWED_IPv4_RANGE)
	if err != nil {
		panic("Invalid IP range")
	}
	_, ipv6Net, err := net.ParseCIDR(ALLOWED_IPv6_RANGE)
	if err != nil {
		panic("Invalid IP range")
	}
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if isIPProtected[req.Spec().Procedure] {
				ip := getClientIP(req)
				if ip == nil {
					return nil, connect.NewError(connect.CodeInternal, nil)
				}
				if !ipv4Net.Contains(ip) && !ipv6Net.Contains(ip) {
					return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("IP not allowed: %v", net.IP(ip).String()))
				}
			}
			return next(ctx, req)
		}
	}
}

func getClientIP(req connect.AnyRequest) []byte {
	ipStr := req.Header().Get("X-Real-IP")
	if ipStr != "" {
		return net.ParseIP(ipStr)
	}
	ip, _, err := net.SplitHostPort(req.Peer().Addr)
	if err != nil {
		return nil
	}
	return net.ParseIP(strings.Split(ip, "%")[0])
}
