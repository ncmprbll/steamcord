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
	var id int
	if err := s.database.QueryRowxContext(ctx, "SELECT id FROM users_roles WHERE name = $1;", user.Role).Scan(&id); err != nil {
		return nil, err
	}

	permissions := &models.Permissions{}
	if err := s.database.QueryRowxContext(ctx, "SELECT JSONB_AGG(permission) permissions FROM users_role_permissions WHERE role_id = $1;", id).Scan(permissions); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *Repository) GetUsers(ctx context.Context, term string) (*models.ManagementUsers, error) {
	const queryTotal = `
						SELECT
							COUNT(*)
						FROM users
						WHERE
							LOWER(id::TEXT) LIKE '%' || LOWER($1) || '%' OR
							LOWER(login) LIKE '%' || LOWER($1) || '%' OR
							LOWER(display_name) LIKE '%' || LOWER($1) || '%';
						`
	var total int
	if err := s.database.QueryRowxContext(ctx, queryTotal, term).Scan(&total); err != nil {
		return nil, err
	}

	const query = `
				SELECT
					*
				FROM users
				WHERE
					LOWER(id::TEXT) LIKE '%' || LOWER($1) || '%' OR
					LOWER(login) LIKE '%' || LOWER($1) || '%' OR
					LOWER(display_name) LIKE '%' || LOWER($1) || '%'
				ORDER BY created_at;
				`
	rows, err := s.database.QueryxContext(ctx, query, term)
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

func (s *Repository) GetRoles(ctx context.Context) (*models.Roles, error) {
	const query = `
				SELECT
					*
				FROM users_roles
				ORDER BY created_at;
				`
	rows, err := s.database.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := &models.Roles{}
	for rows.Next() {
		role := &models.Role{}
		rows.StructScan(role)
		*roles = append(*roles, role)
	}

	return roles, nil
}

func (s *Repository) CreateRole(ctx context.Context, role *models.Role) error {
	const query = `
				INSERT INTO
					users_roles (name)
				VALUES ($1);
				`
	if _, err := s.database.ExecContext(ctx, query, role.Name); err != nil {
		return err
	}
	return nil
}

func (s *Repository) DeleteRole(ctx context.Context, role *models.Role) (int64, error) {
	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	const querySelect = `
					SELECT
						name,
						can_delete
					FROM
						users_roles
					WHERE
						id = $1;
					`
	var (
		name      string
		canDelete bool
	)
	if err := tx.QueryRowxContext(ctx, querySelect, role.ID).Scan(&name, &canDelete); err != nil {
		return 0, err
	}

	if !canDelete {
		return 0, nil
	}

	// const queryUpdate = `
	// 			UPDATE
	// 				users
	// 			SET
	// 				role = 'user'
	// 			WHERE
	// 				role = $1;
	// 			`
	// if _, err := tx.ExecContext(ctx, queryUpdate, name); err != nil {
	// 	return 0, err
	// }

	const queryDelete = `
				DELETE FROM
					users_roles
				WHERE
					id = $1
				`
	result, err := tx.ExecContext(ctx, queryDelete, role.ID)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return affected, nil
}

func (s *Repository) GetRolePermissions(ctx context.Context) (*models.RolePermissions, error) {
	const queryAllPermissions = `
				SELECT
					name
				FROM permissions
				ORDER BY created_at;
				`
	rows, err := s.database.QueryxContext(ctx, queryAllPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []string{}
	for rows.Next() {
		var permission string
		rows.Scan(&permission)
		permissions = append(permissions, permission)
	}

	const queryRolePermissions = `
				SELECT
					name,
					permission
				FROM users_roles
					LEFT JOIN users_role_permissions ON role_id = id;
				`
	rows, err = s.database.QueryxContext(ctx, queryRolePermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := map[string][]string{}
	for rows.Next() {
		role := &models.ManagementRole{}
		rows.StructScan(role)
		if _, ok := roles[role.Name]; !ok {
			roles[role.Name] = make([]string, 0)
		}
		roles[role.Name] = append(roles[role.Name], role.Permission)
	}

	return &models.RolePermissions{Permissions: permissions, Roles: roles}, nil
}

func (s *Repository) AddPermission(ctx context.Context, role *models.Role, permissions *models.Permissions) error {
	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const queryAdd = `
				INSERT INTO
					users_role_permissions (role_id, permission)
				VALUES ($1, $2);
				`
	for _, p := range *permissions {
		if _, err := tx.ExecContext(ctx, queryAdd, role.ID, p); err != nil {
			return err
		}
	}

	const queryUpdate = `
				UPDATE
					users_roles
				SET
					updated_at = NOW()
				WHERE
					id = $1;
				`
	if _, err := tx.ExecContext(ctx, queryUpdate, role.ID); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Repository) DeletePermission(ctx context.Context, role *models.Role, permissions *models.Permissions) (int64, error) {
	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	acc := int64(0)
	const queryAdd = `
				DELETE FROM
					users_role_permissions
				WHERE role_id = $1 AND permission = $2;
				`
	for _, p := range *permissions {
		result, err := tx.ExecContext(ctx, queryAdd, role.ID, p)
		if err != nil {
			return 0, err
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}

		acc += affected
	}

	const queryUpdate = `
				UPDATE
					users_roles
				SET
					updated_at = NOW()
				WHERE
					id = $1;
				`
	if _, err := tx.ExecContext(ctx, queryUpdate, role.ID); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return acc, nil
}
