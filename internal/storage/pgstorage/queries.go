package pgstorage

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

import "context"

type DB interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Queries struct {
	db DB
}

func NewQueries(db DB) *Queries {
	return &Queries{
		db: db,
	}
}
