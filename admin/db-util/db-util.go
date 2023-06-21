package main

import (
	"context"
	"examples/admin/generated"
	"fmt"
	"github.com/google/uuid"
	"github.com/iyarkov/foundation/config"
	"github.com/iyarkov/foundation/logger"
	"github.com/iyarkov/foundation/support"
	"github.com/rs/zerolog"
	"os"
)

type Configuration struct {
	Manifest support.Manifest
	Log      logger.Configuration
	Db       config.DbConfig
}

func main() {
	cfg := Configuration{}
	if err := config.Read(&cfg); err != nil {
		fmt.Printf("failed to read initial configuration %v", err)
		os.Exit(1)
	}

	cfg.Manifest.Name = "Admin DB Util"
	cfg.Manifest.Version = fmt.Sprintf("%s.%s", generated.Version, generated.BuildNumber)
	support.AppManifest = cfg.Manifest

	logger.InitLogger(cfg.Log)
	ctx := logger.WithContextIdAndLogger(context.Background(), uuid.New().String())
	log := zerolog.Ctx(ctx)
	log.Info().Msgf("configuration: %+v", cfg)
}
