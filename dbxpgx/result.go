package dbxpgx

import "github.com/jackc/pgx/v5/pgconn"

type wrapResult struct {
	res pgconn.CommandTag
}

func (r *wrapResult) LastInsertID() (int64, error) {
	return 0, nil
}

func (r *wrapResult) RowsAffected() (int64, error) {
	return r.res.RowsAffected(), nil
}
