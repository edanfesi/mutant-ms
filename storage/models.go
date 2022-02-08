package storage

import (
	"context"
	"database/sql"
)

type PostgresDrivers interface {
	QueryContext(ctx context.Context, sql string) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string) (sql.Result, error)
}
