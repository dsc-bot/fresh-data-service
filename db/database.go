package db

import (
	"context"
	"fmt"
	"os"

	"github.com/dsc-bot/fresh-data-service/config"
	"github.com/dsc-bot/fresh-data-service/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	pool, err := pgxpool.New(context.Background(), config.Conf.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to acquire database connection: %v\n", err)
		os.Exit(1)
	}
	conn.Release()

	Pool = pool

	utils.Logger.Debug("Connected to the database")
}
