package dbxpgx

import "github.com/jackc/pgx/v5"

type wrapRows struct {
	rows pgx.Rows
}

func (r *wrapRows) Scan(dst ...any) error {
	return r.rows.Scan(dst...)
}

func (r *wrapRows) Next() bool {
	return r.rows.Next()
}

func (r *wrapRows) Close() error {
	r.rows.Close()
	return nil
}

func (r *wrapRows) Err() error {
	return r.rows.Err()
}
