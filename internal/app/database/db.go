package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
}

func InitDatabase() (*Database, error) {
	// // docker
	db, err := sqlx.Open("postgres", "postgres://gouser:password@go-database:5432/godatabase?sslmode=disable")

	// // localhost
	// db, err := sqlx.Open("postgres", "postgres://gouser:password@localhost:5433/godatabase?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (db *Database) SelectxInContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	q, args, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	q = db.Rebind(q)
	return db.SelectContext(ctx, dest, q, args...)
}

func (db *Database) NamedExecxContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	r, err := tx.NamedExecContext(ctx, query, arg)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db *Database) GetTxContext(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
