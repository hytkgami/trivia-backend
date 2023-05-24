package infrastructure

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func NewDB(ctx context.Context) (*sqlx.DB, error) {
	opt := "?sslmode=disable"
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	if os.Getenv("APP_ENV") == "development" {
		dsn += opt
	}
	return sqlx.ConnectContext(ctx, "postgres", dsn)
}
