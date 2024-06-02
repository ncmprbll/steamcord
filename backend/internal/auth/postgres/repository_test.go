package postgres

import (
	"context"
	"main/backend/internal/models"
	"main/backend/internal/util"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

type result struct {
	insertID     int64
	rowsAffected int64
	err          error
}

func (r *result) LastInsertId() (int64, error) {
	return r.insertID, r.err
}

func (r *result) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}

func TestRegister(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()
	repository := New(sqlxDB)

	t.Run("Register", func(t *testing.T) {
		user := &models.User{
			Login:    "test_user",
			Email:    "test_user@example.com",
			Password: "test_user",
		}
		mock.ExpectExec(queryRegisterUser).WithArgs(user.Login, user.Email, user.Password).WillReturnResult(util.NewResult(0, 1))
		err := repository.Register(context.Background(), user)
		require.NoError(t, err)
	})
}
