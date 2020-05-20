package server

import (
	"context"
	//"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/switch-bit/tenprotocol-server/tenp"
)

/*
CREATE TABLE IF NOT EXISTS tenp.diagnosed (
	key_type varchar,
    key varchar,
    ts timestamp,
	reporting_client_id varchar,
    PRIMARY KEY((key_type, key), ts)
);
*/

func (s Server) Upload(ctx context.Context, in *tenp.UploadRequest) (*empty.Empty, error) {
	return nil, nil
}
