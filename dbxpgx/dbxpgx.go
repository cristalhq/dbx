package dbxpgx

import (
	"context"
	"fmt"

	"github.com/cristalhq/dbx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string

	// TODO: add more fields supported by [pgxpool.ParseConfig]
}

func (cfg Config) AsConnStr() string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
	return dsn
}

func New(ctx context.Context, cfg Config) (dbx.Client, error) {
	connStr := cfg.AsConnStr()

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}
	return WrapPool(pool), nil
}

func WrapPool(pool *pgxpool.Pool) dbx.Client {
	return &wrapClient{pool: pool}
}

type wrapClient struct {
	pool *pgxpool.Pool
}

func (c *wrapClient) Acquire(ctx context.Context) (dbx.Conn, error) {
	conn, err := c.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	return &wrapConn{conn: conn}, nil
}

func (c *wrapClient) Release(conn dbx.Conn) {
	// TODO: think about other types under conn
	conn.(*wrapConn).conn.Release()
}

func (c *wrapClient) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx)
}

func (c *wrapClient) BeginTx(ctx context.Context, opts dbx.TxOptions) (dbx.Tx, error) {
	tx, err := c.pool.BeginTx(ctx, convTxOptions(opts))
	if err != nil {
		return nil, err
	}
	return &wrapTx{tx: tx}, nil
}

func (c *wrapClient) Exec(ctx context.Context, query string, args ...any) (dbx.Result, error) {
	res, err := c.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &wrapResult{res: res}, nil
}

func (c *wrapClient) Query(ctx context.Context, query string, args ...any) (dbx.Rows, error) {
	rows, err := c.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &wrapRows{rows: rows}, nil
}

func (c *wrapClient) QueryRow(ctx context.Context, query string, args ...any) dbx.Row {
	return c.pool.QueryRow(ctx, query, args...)
}

func (c *wrapClient) Close() error {
	c.pool.Close()
	return nil
}

type wrapConn struct {
	conn *pgxpool.Conn
}

func (c *wrapConn) Ping(ctx context.Context) error {
	return c.conn.Ping(ctx)
}

func (c *wrapConn) BeginTx(ctx context.Context, opts dbx.TxOptions) (dbx.Tx, error) {
	tx, err := c.conn.BeginTx(ctx, convTxOptions(opts))
	if err != nil {
		return nil, err
	}
	return &wrapTx{tx: tx}, err
}

func (c *wrapConn) Exec(ctx context.Context, query string, args ...any) (dbx.Result, error) {
	res, err := c.conn.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &wrapResult{res: res}, nil
}

func (c *wrapConn) Query(ctx context.Context, query string, args ...any) (dbx.Rows, error) {
	rows, err := c.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &wrapRows{rows: rows}, err
}

func (c *wrapConn) QueryRow(ctx context.Context, query string, args ...any) dbx.Row {
	return c.conn.QueryRow(ctx, query, args...)
}

func (c *wrapConn) Close() error {
	return c.Close()
}
