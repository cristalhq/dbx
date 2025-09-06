package dbx

import (
	"context"
)

// Client ...
type Client interface {
	Conn

	// Acquire ...
	Acquire(ctx context.Context) (Conn, error)

	// Release ...
	Release(Conn)
}

// Conn ...
type Conn interface {
	// Ping ...
	Ping(ctx context.Context) error

	// BeginTx ...
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)

	// Exec ...
	Exec(ctx context.Context, query string, args ...any) (Result, error)

	// Query ...
	Query(ctx context.Context, query string, args ...any) (Rows, error)

	// QueryRow ...
	QueryRow(ctx context.Context, query string, args ...any) Row

	// Close ...
	Close() error
}

// Result ...
type Result interface {
	// LastInsertID ...
	LastInsertID() (int64, error)

	// RowsAffected ...
	RowsAffected() (int64, error)
}

// Rows ...
type Rows interface {
	Row

	// Next ...
	Next() bool

	// Close ...
	Close() error

	// Err ...
	Err() error
}

// Row ...
type Row interface {
	// Scan ...
	Scan(dst ...any) error
}

// Tx ...
type Tx interface {
	// BeginTx ...
	BeginTx(ctx context.Context, opts TxOptions) (Tx, error)

	// Commit ...
	Commit(ctx context.Context) error

	// Rollback ...
	Rollback(ctx context.Context) error

	// Exec ...
	Exec(ctx context.Context, query string, args ...any) (Result, error)

	// Query ...
	Query(ctx context.Context, query string, args ...any) (Rows, error)

	// QueryRow ...
	QueryRow(ctx context.Context, query string, args ...any) Row
}

// TxOptions ...
type TxOptions struct{}
