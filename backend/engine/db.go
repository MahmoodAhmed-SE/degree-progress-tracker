package engine

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgxDB *pgxpool.Pool

func config() *pgxpool.Config {
	const defaultMaxConnections = int32(10)
	const defaultMinConnections = int32(5)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = 30 * time.Minute
	const defaultHealthCheckPeriod = time.Minute
	const defaultAcquireTimeout = 5 * time.Second
	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}
	dbConfig.MaxConns = defaultMaxConnections
	dbConfig.MinConns = defaultMinConnections
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultAcquireTimeout
	return dbConfig
}

func InitDB() {
	connPool, err := pgxpool.NewWithConfig(context.Background(), config())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	PgxDB = connPool
}
