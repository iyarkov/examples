package database

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func errorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
