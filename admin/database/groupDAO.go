package database

import (
	"context"
	"examples/admin/model"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var ErrDuplicateName = fmt.Errorf("duplicate name")

type GroupDAO interface {
	Create(ctx context.Context, group *model.Group) error
}

type groupDAO struct {
	pool *pgxpool.Pool
}

func NewGroupDAO(pool *pgxpool.Pool) GroupDAO {
	gd := groupDAO{
		pool: pool,
	}
	return &gd
}

func (dao *groupDAO) Create(ctx context.Context, group *model.Group) error {
	tx, err := dao.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction begin failed, %w", err)
	}

	err = tx.QueryRow(ctx, "insert into group_tbl(created_at, updated_at, name) values($1, $2, $3) returning id",
		group.CreatedAt, group.UpdatedAt, group.Name).Scan(&group.Id)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			zerolog.Ctx(ctx).Err(rollbackErr).Msg("failure to rollback a transaction")
		}
		if errorCode(err) == pgerrcode.UniqueViolation {
			return ErrDuplicateName
		}
		return fmt.Errorf("query failed, %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed, %w", err)
	}
	return nil
}
