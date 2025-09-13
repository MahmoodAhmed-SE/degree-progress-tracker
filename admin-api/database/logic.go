package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgxPool *pgxpool.Pool

func InitDB() {
	var err error
	PgxPool, err = pgxpool.New(context.Background(), os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatalf("Error while connecting to db: %s", err.Error())
	}
}
