package database

import (
	"errors"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName = "pgx"

func New(config *Config) (*sqlx.DB, error) {
	if config == nil {
		return nil, errors.New("db: config is nil")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.SSLMode,
	)

	db, err := sqlx.Connect(driverName, dsn)
	if err != nil {
		slog.Error("db: failed to connect to database:", slog.Any("error", err))
		return nil, err
	}

	return db, nil
}
