package main

import (
	"context"
	"examples/admin/database"
	"examples/admin/generated"
	"examples/admin/grpc"
	"fmt"
	"github.com/google/uuid"
	"github.com/iyarkov/foundation/auth"
	"github.com/iyarkov/foundation/config"
	"github.com/iyarkov/foundation/logger"
	"github.com/iyarkov/foundation/support"
	"github.com/iyarkov/foundation/telemetry"
	"github.com/iyarkov/foundation/tls"
	"github.com/rs/zerolog"
	"os"
)

type Configuration struct {
	Auth      auth.Configuration
	Manifest  support.Manifest
	Log       logger.Configuration
	GRPC      grpc.Configuration
	Db        config.DbConfig
	TLS       tls.Configuration
	Telemetry telemetry.Configuration
}

func main() {
	cfg := Configuration{}
	if err := config.Read(&cfg); err != nil {
		fmt.Printf("failed to read initial configuration %v", err)
		os.Exit(1)
	}

	cfg.Manifest.Name = "Admin"
	cfg.Manifest.Version = fmt.Sprintf("%s.%s", generated.Version, generated.BuildNumber)
	support.AppManifest = cfg.Manifest

	logger.InitLogger(&cfg.Log)
	ctx := logger.WithContextIdAndLogger(context.Background(), uuid.New().String())
	log := zerolog.Ctx(ctx)
	log.Info().Msgf("configuration: %+v", cfg)

	cfg.Telemetry.Mode = "docker"
	telemetry.InitTelemetry(ctx, &cfg.Telemetry)

	tlsConfig, err := cfg.TLS.NewCryptoTlsConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("TLS initialization failed")
	}

	databaseModule, err := database.InitDb(ctx, &cfg.Db)
	if err != nil {
		log.Fatal().Err(err).Msg("database initialization failed")
	}

	if err := grpc.InitGrpc(ctx, &cfg.GRPC, &cfg.Auth, databaseModule, tlsConfig); err != nil {
		log.Fatal().Err(err).Msg("gRPC initialization failed")
	}

}
