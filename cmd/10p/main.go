package main

import (
	"context"
	"github.com/switch-bit/tenprotocol-server/api/server"
	"github.com/switch-bit/tenprotocol-server/swagger"
	"github.com/switch-bit/tenprotocol-server/tenp"
	"github.com/switch-bit/tenprotocol-server/version"
	"github.com/switch-bit/orlop"
	"google.golang.org/grpc"
)

func main() {
	orlop.Run(version.Name, func(ctx context.Context, cfg *server.Config) error {
		s, err := server.NewServer(ctx, cfg)
		if err != nil {
			return err
		}
		return orlop.Serve(ctx, version.Name,
			orlop.WithServerConfig(cfg.Server),
			orlop.WithTLSProvider(orlop.NewVaultTLSProvider(cfg.Vault)),
			orlop.WithSwagger(swagger.FS),
			orlop.WithGRPCServices(func(ctx context.Context, g *grpc.Server) {
				tenp.RegisterThreatExposureNotificationServiceServer(g, s)
			}),
			orlop.WithGateway(
				tenp.RegisterThreatExposureNotificationServiceHandler),
		)
	}, &server.Config{})
}
