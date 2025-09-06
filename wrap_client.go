package dbx

import (
	"context"
	"database/sql"
)

func WrapClient(db *sql.DB) Client {
	return &wrapClient{db: db}
}

type wrapClient struct {
	db *sql.DB
}

func (c *wrapClient) Acquire(ctx context.Context) (Conn, error) {
	conn, err := c.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return WrapConn(conn), nil
}

func (c *wrapClient) Release(conn Conn) {
	conn.Close()
}

func (c *wrapClient) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *wrapClient) BeginTx(ctx context.Context, opts TxOptions) (Tx, error) {
	tx, err := c.db.BeginTx(ctx, convTxOptions(opts))
	if err != nil {
		return nil, err
	}
	return WrapTx(tx), nil
}

func (c *wrapClient) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	res, err := c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return wrapRes(res), nil
}

func (c *wrapClient) Query(ctx context.Context, query string, args ...any) (Rows, error) {
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c *wrapClient) QueryRow(ctx context.Context, query string, args ...any) Row {
	return c.db.QueryRowContext(ctx, query, args...)
}

func (c *wrapClient) Close() error {
	return c.db.Close()
}
