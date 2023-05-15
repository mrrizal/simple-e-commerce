package configs

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(conf Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), conf.DBURI)
	if err != nil {
		log.Default().Fatal("cant connect to the database")
	}
	return pool, nil
}
