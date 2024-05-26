package postgres

import (
	"context"
	"main/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) *Repository {
	return &Repository{database: database}
}

func (s *Repository) GetPermissions(ctx context.Context, user *models.User) (*models.Permissions, error) {
	permissions := &models.Permissions{}
	err := s.database.QueryRowxContext(ctx, "SELECT JSONB_AGG(permission) permissions FROM users_role_permissions WHERE role = $1;", user.Role).Scan(permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *Repository) GetUsers(ctx context.Context) (*models.ManagementUsers, error) {
	const queryTotal = `
						SELECT
							COUNT(*)
						FROM users;
						`
	var total int
	if err := s.database.QueryRowxContext(ctx, queryTotal).Scan(&total); err != nil {
		return nil, err
	}

	const query = `
				SELECT
					*
				FROM users
				ORDER BY created_at;
				`
	rows, err := s.database.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.User{}

	for rows.Next() {
		row := &models.User{}
		rows.StructScan(row)
		result = append(result, row)
	}

	const queryRoles = `
				SELECT
					name
				FROM users_roles;
				`
	rows, err = s.database.QueryxContext(ctx, queryRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []string{}

	for rows.Next() {
		var role string
		rows.Scan(&role)
		roles = append(roles, role)
	}

	const queryCurrencies = `
					SELECT
						code
					FROM currencies;
					`
	rows, err = s.database.QueryxContext(ctx, queryCurrencies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currencies := []string{}

	for rows.Next() {
		var currency string
		rows.Scan(&currency)
		currencies = append(currencies, currency)
	}

	return &models.ManagementUsers{Users: result, Total: total, Roles: roles, Currencies: currencies}, nil
}

func (s *Repository) UpdateUser(ctx context.Context, user *models.User) error {
	const query = `
				UPDATE
					users
				SET
					avatar = CASE
						WHEN $1 = '' THEN avatar
						WHEN $1 = 'remove' THEN ''
						WHEN $1 <> '' THEN avatar
					END,
					display_name = CASE
						WHEN $2 = '' THEN display_name
						WHEN $2 <> '' THEN $2
					END,
					privacy = CASE
						WHEN $3 = '' THEN privacy
						WHEN $3 <> '' THEN $3
					END,
					currency_code = CASE
						WHEN $4 = '' THEN currency_code
						WHEN $4 <> '' THEN $4
					END,
					password = CASE
						WHEN $5 = '' THEN password
						WHEN $5 <> '' THEN $5
					END,
					role = CASE
						WHEN $6 = '' THEN role
						WHEN $6 <> '' THEN $6
					END,
					banned = COALESCE($7, banned)
				WHERE id = $8;
				`
	_, err := s.database.ExecContext(ctx, query, user.Avatar, user.DisplayName, user.Privacy, user.CurrencyCode, user.Password, user.Role, user.Banned, user.UUID)
	if err != nil {
		return err
	}

	return nil
}
