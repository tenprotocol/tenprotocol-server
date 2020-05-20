package server

import (
	"context"
	"github.com/gocql/gocql"
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
	b := s.cassandra.NewBatch(ctx, gocql.LoggedBatch)
	//
	//clientID := ctx.Value(clientIDContext).(string)
	//
	//now := time.Now().Unix()
	//
	//for _, key := range in.Keys {
	//	b.Query("INSERT INTO tenp.diagnosed (key_type, key, ts, reporting_client_id)" +
	//		"VALUES (?, ?, ?, ?)" +
	//		"IF NOT EXIST", in.KeyType, key, now, clientID)
	//}

	return &empty.Empty{}, s.cassandra.ExecuteBatch(b)
}
