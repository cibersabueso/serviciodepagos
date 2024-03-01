// internal/database/db.go
package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	dbURL := "postgres://postgres:calasiki@database-1.ct0o62yqab33.us-east-1.rds.amazonaws.com:5432/pruebas"
	return pgxpool.Connect(context.Background(), dbURL)
}
