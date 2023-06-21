package grpc

import (
	"context"
	"crypto/tls"
	"examples/admin/database"
	"examples/admin/generated"
	"fmt"
	"github.com/iyarkov/foundation/auth"
	foundationgrpc "github.com/iyarkov/foundation/grpc"
	"github.com/iyarkov/foundation/support"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
)

type Configuration struct {
	Port uint16
}

type Module struct {
}

func InitGrpc(ctx context.Context, cfg *Configuration, authCfg *auth.Configuration, dbModule *database.Module, tlsCfg *tls.Config) error {
	if cfg.Port == 0 {
		cfg.Port = 8443
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsCfg)),
		grpc.ChainUnaryInterceptor(
			foundationgrpc.ServerContextId,
			foundationgrpc.ServerConnectionInfo,
			foundationgrpc.NewServerAuthInterceptor(ctx, authCfg),
		),
	)

	gm := groupsServer{
		groupDAO: dbModule.GroupDAO,
	}
	generated.RegisterGroupsServer(grpcServer, &gm)
	log := zerolog.Ctx(ctx)
	log.Info().Msgf("starting grpc server on %d", cfg.Port)

	support.OnSigTerm(func(signal os.Signal) {
		log.Info().Msg("Shutting down gRPC server")
		grpcServer.GracefulStop()
		log.Info().Msgf("gRPC server stopped")
	})

	return grpcServer.Serve(lis)
}
