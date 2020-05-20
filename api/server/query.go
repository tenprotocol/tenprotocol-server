package server

import (
	"context"
	"github.com/switch-bit/tenprotocol-server/tenp"
)

func (s Server) Query(ctx context.Context, in *tenp.QueryRequest) (*tenp.QueryResponse, error) {
	out := &tenp.QueryResponse{}
	return out, nil
}
