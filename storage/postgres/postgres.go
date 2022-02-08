package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"mutant-ms/settings"
	"mutant-ms/utils/logger"
)

type postgres struct {
	db *sql.DB
}

func NewPostgresStorage(settings settings.PostgresSettings) *postgres {
	log := logger.New("-")

	connectionURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.DBName,
		"disable",
	)

	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	return &postgres{
		db: db,
	}
}

func (p *postgres) QueryContext(ctx context.Context, query string) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query)
}

func (p *postgres) ExecContext(ctx context.Context, query string) (sql.Result, error) {
	return p.db.ExecContext(ctx, query)
}
