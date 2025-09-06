package dbx

import (
	"context"
	"database/sql"
)

// WrapTx ...
func WrapTx(tx *sql.Tx) Tx {
	return &wrapTx{tx: tx}
}

type wrapTx struct {
	tx *sql.Tx
}

func (tx *wrapTx) BeginTx(ctx context.Context, opts TxOptions) (Tx, error) {
	return nil, nil
}

func (tx *wrapTx) Commit(context.Context) error {
	return tx.tx.Commit()
}

func (tx *wrapTx) Rollback(context.Context) error {
	return tx.tx.Rollback()
}

func (tx *wrapTx) Exec(ctx context.Context, query string, args ...any) (Result, error) {
	res, err := tx.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return wrapRes(res), nil
}

func (tx *wrapTx) Query(ctx context.Context, query string, args ...any) (Rows, error) {
	rows, err := tx.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (tx *wrapTx) QueryRow(ctx context.Context, query string, args ...any) Row {
	return tx.tx.QueryRowContext(ctx, query, args...)
}

func convTxOptions(opts TxOptions) *sql.TxOptions {
	// TODO: impl conversion
	return &sql.TxOptions{}
}
