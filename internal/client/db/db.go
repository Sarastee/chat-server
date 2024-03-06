package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Handler - function that is executed in transaction
type Handler func(ctx context.Context) error

// Client - DB Client
type Client interface {
	DB() DB
	Close() error
}

// TxManager - Transaction Manager
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query ...
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor interface
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecutor combines NamedExecutor and QueryExecutor
type SQLExecutor interface {
	QueryExecutor
	CopyExecutor
}

// QueryExecutor interface
type QueryExecutor interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// CopyExecutor interface
type CopyExecutor interface {
	CopyFromContext(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// Pinger interface to Ping DB
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB interface
type DB interface {
	SQLExecutor
	Transactor
	Pinger
	Close()
}
