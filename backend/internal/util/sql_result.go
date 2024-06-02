package util

import "database/sql/driver"

type result struct {
	insertID     int64
	rowsAffected int64
	err          error
}

func NewResult(lastInsertID int64, rowsAffected int64) driver.Result {
	return &result{
		insertID:     lastInsertID,
		rowsAffected: rowsAffected,
	}
}

func NewErrorResult(err error) driver.Result {
	return &result{
		err: err,
	}
}

func (r *result) LastInsertId() (int64, error) {
	return r.insertID, r.err
}

func (r *result) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}