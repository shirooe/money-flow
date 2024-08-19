package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	SQLExecer

	Pooler
	Pinger
	Closer
}

type Pooler interface {
	Pool() *pgxpool.Pool
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type Closer interface {
	Close() error
}

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, query Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, query Query, ards ...interface{}) error
	ScanRowContext(ctx context.Context, dest interface{}, query Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query Query, args ...interface{}) pgx.Row
}
