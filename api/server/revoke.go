package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/switch-bit/tenprotocol-server/tenp"
)

func (s Server) Revoke(ctx context.Context, in *tenp.RevokeRequest) (*empty.Empty, error) {
	return nil, nil
}
