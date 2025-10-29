package veripass

import (
	"context"
	"fmt"
	"net"
	"strings"

	"connectrpc.com/connect"
)

const ALLOWED_IPv4_RANGE = "0.0.0.0/0"
const ALLOWED_IPv6_RANGE = "::/0"

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
			ip := getClientIP(req)
			if ip == nil {
				return nil, connect.NewError(connect.CodeInternal, nil)
			}
			if !ipv4Net.Contains(ip) && !ipv6Net.Contains(ip) {
				return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("IP not allowed: %v", net.IP(ip).String()))
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
