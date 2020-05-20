package server

import (
	"context"
	"github.com/switch-bit/orlop"
	orlop_cassandra "github.com/switch-bit/orlop-cassandra"
	"github.com/switch-bit/orlop/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

var (
	errNotAuthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
	errNotAuthorized = status.Error(codes.PermissionDenied, "permission denied")
	clientIDContext = struct{}{}
)

// Config provide configuration options for the server
type Config struct {
	Server    orlop.ServerConfig
	Vault     orlop.VaultConfig
	Cassandra orlop_cassandra.Config
	TTL       time.Duration
}

// Server provides an implementation for the server
type Server struct {
	options   *Config
	cassandra *orlop_cassandra.Connection
}

// NewServer returns a new server given the options
func NewServer(ctx context.Context, opts *Config) (*Server, error) {
	cassandra, err := orlop_cassandra.Connect(opts.Cassandra, opts.Vault)
	if err != nil {
		return nil, err
	}

	s := &Server{
		options: opts,
		cassandra: cassandra,
	}
	return s, nil
}

/*
CREATE TABLE IF NOT EXISTS tenp.api_keys (
    api_key varchar,
    client_id varchar,
    PRIMARY KEY(api_key)
);
*/

func (s *Server) Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errNotAuthenticated
	}

	// Get the API key from Authentication header, erroring out if not found
	var apiKey string

	if m := md.Get("Authorization"); len(m) > 0 {
		prefix := "apikey "
		if strings.HasPrefix(strings.ToLower(m[0]), prefix) {
			apiKey = m[0][len(prefix):]
			log.Trace("found Authorization header", apiKey)
		}
	}

	if len(apiKey) == 0 {
		return nil, errNotAuthenticated
	}

	// Get the client ID (assuming one exists) for the API key, and error out if not found
	var clientID string
	err := s.cassandra.Query(ctx, `SELECT client_id FROM api_keys WHERE api_key = ? LIMIT 1`, apiKey).Scan(&clientID)
	if err != nil {
		return nil, errNotAuthorized
	}

	// Add to the context
	ctx = context.WithValue(ctx, clientIDContext, clientID)

	return handler(ctx, req)
}
