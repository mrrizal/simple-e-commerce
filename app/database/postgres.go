// implement of database interface, in this case i'm using postgres
package database

import (
	"context"
	"e-commerce-api/app/configs"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transaction struct {
	tx pgx.Tx
}

func NewTransaction(tx pgx.Tx) Transaction {
	return &transaction{tx: tx}
}

func (this *transaction) Rollback(ctx context.Context) error {
	return this.tx.Rollback(ctx)
}

func (this *transaction) BulkInsert(ctx context.Context, tableName string, columns []string, rows [][]any) (int, error) {
	copyFrom := pgx.CopyFromRows(rows)
	temp, err := this.tx.CopyFrom(ctx, pgx.Identifier{tableName}, columns, copyFrom)
	return int(temp), err
}

func (this *transaction) Commit(ctx context.Context) error {
	return this.tx.Commit(ctx)
}

type postgresDB struct {
	pool *pgxpool.Pool
}

func (this *postgresDB) Close() {
	this.pool.Close()
}

func (this *postgresDB) Begin(ctx context.Context) (Transaction, error) {
	pgTx, err := this.pool.Begin(ctx)
	tx := NewTransaction(pgTx)
	return tx, err
}

func (this *postgresDB) QueryRow(ctx context.Context, sql string, args ...any) Row {
	return this.pool.QueryRow(ctx, sql, args...)
}

func (this *postgresDB) Query(ctx context.Context, sql string, args ...any) (Rows, error) {
	return this.pool.Query(ctx, sql, args...)
}

func NewDB(conf configs.Config) (DB, error) {
	pool, err := pgxpool.New(context.Background(), conf.DBURI)
	db := &postgresDB{pool: pool}
	if err != nil {
		log.Default().Fatal("cant connect to the database")
	}
	return db, nil
}
