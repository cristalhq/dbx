package dbxpgx

import (
	"context"

	"github.com/cristalhq/dbx"
	"github.com/jackc/pgx/v5"
)

type wrapTx struct {
	tx pgx.Tx
}

func (t *wrapTx) BeginTx(ctx context.Context, opts dbx.TxOptions) (dbx.Tx, error) {
	tx, err := t.tx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &wrapTx{tx: tx}, nil
}

func (t *wrapTx) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *wrapTx) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

func (t *wrapTx) Exec(ctx context.Context, query string, args ...any) (dbx.Result, error) {
	_, err := t.tx.Exec(ctx, query, args...)
	return nil, err
}

func (t *wrapTx) Query(ctx context.Context, query string, args ...any) (dbx.Rows, error) {
	rows, err := t.tx.Query(ctx, query, args...)
	return &wrapRows{rows: rows}, err
}

func (t *wrapTx) QueryRow(ctx context.Context, query string, args ...any) dbx.Row {
	return t.tx.QueryRow(ctx, query, args...)
}

func convTxOptions(opts dbx.TxOptions) pgx.TxOptions {
	// TODO: impl conversion
	return pgx.TxOptions{}
}
