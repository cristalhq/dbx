package dbx

import (
	"context"
	"database/sql"
)

// WrapConn ...
func WrapConn(conn *sql.Conn) Conn {
	return &wrapConn{conn: conn}
}

type wrapConn struct {
	conn *sql.Conn
}

func (c *wrapConn) Ping(ctx context.Context) error {
	return c.conn.PingContext(ctx)
}

func (c *wrapConn) BeginTx(ctx context.Context, opts TxOptions) (Tx, error) {
	tx, err := c.conn.BeginTx(ctx, convTxOptions(opts))
	if err != nil {
		return nil, err
	}
	return WrapTx(tx), nil
}

func (c *wrapConn) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	res, err := c.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return wrapRes(res), nil
}

func (c *wrapConn) Query(ctx context.Context, query string, args ...any) (Rows, error) {
	rows, err := c.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (c *wrapConn) QueryRow(ctx context.Context, query string, args ...any) Row {
	return c.conn.QueryRowContext(ctx, query, args...)
}

func (c *wrapConn) Close() error {
	return c.Close()
}
