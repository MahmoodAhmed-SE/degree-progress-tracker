package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB() {
	if newPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL")); err == nil {
		pool = newPool
	} else {
		log.Fatalf("Error while initiating db: %s", err.Error())
	}
}

func GetPool() *pgxpool.Pool {
	return pool
}
