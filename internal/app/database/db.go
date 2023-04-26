package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func InitDatabase() (*Database, error) {
	db, err := sqlx.Open("postgres", "user=foo dbname=bar sslmode=disable")

	if err != nil {
		return nil, errors.New("failed conencting to db")
	}

	return &Database{db}, nil
}

func (db *Database) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return db.MustBegin().MustExec(query, args...)
}

func (db *Database) SelectContext(ctx context.Context, q sqlx.QueryerContext, dest interface{}, query string, args ...interface{}) error {
	return db.SelectContext(ctx, q, dest, query, args...)
}
