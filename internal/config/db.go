package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	if dsn == "" {
        return ErrNoDBURL
    }
    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        return err
    }
    if err := pool.Ping(context.Background()); err != nil {
        return err
    }
    DB = pool
	return nil
}

var ErrNoDBURL = &ConfigError{"DATABASE_URL environment variable not set"}

type ConfigError struct {
    Msg string
}

func (e *ConfigError) Error() string {
    return e.Msg
}