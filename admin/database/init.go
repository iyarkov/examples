package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/iyarkov/foundation/config"
	"github.com/iyarkov/foundation/schema"
	"github.com/iyarkov/foundation/support"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"os"
)

type Module struct {
	connectionPool *pgxpool.Pool

	GroupDAO GroupDAO
}

func InitDb(ctx context.Context, cfg *config.DbConfig) (*Module, error) {
	log := zerolog.Ctx(ctx)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password.Value(), cfg.DbName)
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	defer support.CloseWithWarning(ctx, db, "Failed to close the DB after schema updated")
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	log.Info().Msgf("connected to DB on %d:%d", cfg.Host, cfg.Port)

	_, newDbVersion, schemaErr := schema.Update(ctx, db, changeset)
	if schemaErr != nil {
		return nil, fmt.Errorf("failed to upgrade schema: %w", schemaErr)
	}

	validationMessages, validationErr := schema.Validate(ctx, db, expectedSchema, newDbVersion == changeset[len(changeset)-1].Version)
	if validationErr != nil {
		return nil, fmt.Errorf("failed to validate schema: %w", validationErr)
	}
	if len(validationMessages) != 0 {
		log.Error().Msg("DB Schema validation failed:")
		log.Error().Msg("======================================================")
		for _, msg := range validationMessages {
			log.Error().Msg(msg)
		}
		log.Error().Msg("======================================================")
		return nil, fmt.Errorf("db schema validation failed")
	}
	log.Info().Msgf("DB Ready, schema version: %s", newDbVersion)

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db pool: %w", err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db pool: %w", err)
	}

	support.OnSigTerm(func(signal os.Signal) {
		log.Info().Msg("Closing DB connection pool")
		pool.Close()
		log.Info().Msg("DB connection pool closed")
	})

	module := Module{
		connectionPool: pool,
		GroupDAO:       NewGroupDAO(pool),
	}
	return &module, nil
}