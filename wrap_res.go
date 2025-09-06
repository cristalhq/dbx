package dbx

import "database/sql"

func wrapRes(res sql.Result) Result {
	return &wrapResult{res: res}
}

type wrapResult struct {
	res sql.Result
}

func (res *wrapResult) LastInsertID() (int64, error) {
	n, err := res.res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (res *wrapResult) RowsAffected() (int64, error) {
	n, err := res.res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return n, nil
}
