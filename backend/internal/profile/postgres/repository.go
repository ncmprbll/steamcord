package postgres

import (
	"context"
	"database/sql"
	"errors"
	"main/backend/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) *Repository {
	return &Repository{database: database}
}

func (s *Repository) Update(ctx context.Context, user *models.UserGeneralUpdate) error {
	const query = `
				UPDATE
					users
				SET
					avatar = CASE
						WHEN $1 = '' THEN avatar
						WHEN $1 <> '' THEN $1
					END,
					display_name = CASE
						WHEN $2 = '' THEN display_name
						WHEN $2 <> '' THEN $2
					END,
					about = $3
				WHERE id = $4;
				`
	_, err := s.database.ExecContext(ctx, query, user.Avatar, user.DisplayName, user.About, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) PasswordUpdate(ctx context.Context, user *models.UserPasswordUpdate) error {
	const query = `
				UPDATE
					users
				SET
					password = $1
				WHERE id = $2;
				`
	_, err := s.database.ExecContext(ctx, query, user.NewPassword, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) PrivacyUpdate(ctx context.Context, user *models.UserPrivacyUpdate) error {
	const query = `
				UPDATE
					users
				SET
					privacy = $1
				WHERE id = $2;
				`
	_, err := s.database.ExecContext(ctx, query, user.Privacy, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) DeleteAvatar(ctx context.Context, user *models.User) (string, error) {
	const query = `
				WITH u AS (
					SELECT avatar FROM users WHERE id = $1
				)
				UPDATE
					users
				SET
					avatar = ''
				WHERE id = $1
				RETURNING (SELECT avatar FROM u);
				`
	var avatar string
	if err := s.database.QueryRowxContext(ctx, query, user.UUID).Scan(&avatar); err != nil {
		return "", err
	}

	return avatar, nil
}

func (s *Repository) CreateComment(ctx context.Context, comment *models.Comment) error {
	const query = `
				INSERT INTO
					users_comments (profile_id, commentator, text)
				VALUES ($1, $2, $3);
				`
	_, err := s.database.ExecContext(ctx, query, comment.ProfileID, comment.Commentator, comment.Text)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) GetComments(ctx context.Context, user *models.User, pageLimit, pageOffset int) (*models.ProfileComments, error) {
	const queryTotal = `
						SELECT
							COUNT(*)
						FROM users_comments
						WHERE profile_id = $1
						`
	var total int
	if err := s.database.QueryRowxContext(ctx, queryTotal, user.UUID).Scan(&total); err != nil {
		return nil, err
	}

	const queryComments = `
						SELECT
							users_comments.id,
							commentator,
							avatar,
							display_name,
							text,
							users_comments.created_at
						FROM users_comments
							JOIN users ON users.id = users_comments.commentator
						WHERE profile_id = $1
						ORDER BY created_at DESC
						LIMIT $2 OFFSET $3;
						`
	rows, err := s.database.QueryxContext(ctx, queryComments, user.UUID, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.DisplayComment{}

	for rows.Next() {
		row := &models.DisplayComment{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return &models.ProfileComments{Comments: result, Total: total}, nil
}

func (s *Repository) DeleteComment(ctx context.Context, profileUser, requester *models.User, comment *models.Comment) (bool, error) {
	const query = `
				DELETE FROM
					users_comments
				WHERE id = $1 AND (profile_id = $2 OR commentator = $2);
				`
	result, err := s.database.ExecContext(ctx, query, comment.ID, requester.UUID)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected != 0, nil
}

func (s *Repository) IsFriend(ctx context.Context, user1 *models.User, user2 *models.User) (bool, error) {
	const queryIsFriend = `
						SELECT EXISTS ( SELECT * FROM users_friends WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1) )
						`
	var isFriend bool
	if err := s.database.QueryRowxContext(ctx, queryIsFriend, user1.UUID, user2.UUID).Scan(&isFriend); err != nil {
		return false, err
	}

	return isFriend, nil
}

func (s *Repository) HasIncomingInvite(ctx context.Context, user1 *models.User, user2 *models.User) (bool, error) {
	const queryIsFriend = `
						SELECT EXISTS ( SELECT * FROM users_friend_invites WHERE invitee = $1 AND inviter = $2 AND status = 'pending' )
						`
	var hasIncomingInvite bool
	if err := s.database.QueryRowxContext(ctx, queryIsFriend, user1.UUID, user2.UUID).Scan(&hasIncomingInvite); err != nil {
		return false, err
	}

	return hasIncomingInvite, nil
}

func (s *Repository) HasOutgoingInvite(ctx context.Context, user1 *models.User, user2 *models.User) (bool, error) {
	const queryIsFriend = `
						SELECT EXISTS ( SELECT * FROM users_friend_invites WHERE invitee = $1 AND inviter = $2 AND status = 'pending' )
						`
	var hasOutgoingInvite bool
	if err := s.database.QueryRowxContext(ctx, queryIsFriend, user2.UUID, user1.UUID).Scan(&hasOutgoingInvite); err != nil {
		return false, err
	}

	return hasOutgoingInvite, nil
}

func (s *Repository) FriendInvite(ctx context.Context, invitee *models.User, inviter *models.User) error {
	if invitee.UUID == inviter.UUID {
		return models.ErrCannotFriendSelf
	}

	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const queryIsFriend = `
						SELECT EXISTS ( SELECT * FROM users_friends WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1) )
						`
	var isFriend bool
	if err := tx.QueryRowxContext(ctx, queryIsFriend, invitee.UUID, inviter.UUID).Scan(&isFriend); err != nil {
		return err
	}
	if isFriend {
		return models.ErrAlreadyFriends
	}

	const queryLatestInvite = `
							SELECT
								status,
								created_at
							FROM users_friend_invites
							WHERE invitee = $1 AND inviter = $2
							ORDER BY created_at DESC;
							`
	invite := &models.Invite{}
	if err := tx.QueryRowxContext(ctx, queryLatestInvite, invitee.UUID, inviter.UUID).StructScan(invite); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if invite.Status == "pending" {
		return models.ErrFriendInvitePending
	} else if invite.Status == "rejected" && invite.CreatedAt.After(time.Now().Add(-models.FRIEND_REQUEST_REJECTED_WAIT_TIME)) {
		return models.ErrAlreadyTriedToFriendRejected
	}

	// const queryRemovePreviousInvites = `
	// 						DELETE FROM
	// 							users_friend_invites
	// 						WHERE invitee = $1 AND inviter = $2;
	// 						`
	// if _, err := tx.ExecContext(ctx, queryRemovePreviousInvites, invitee, inviter.UUID); err != nil {
	// 	return err
	// }

	const querySendInvite = `
				INSERT INTO
					users_friend_invites (invitee, inviter)
				VALUES ($1, $2);
				`
	if _, err := tx.ExecContext(ctx, querySendInvite, invitee.UUID, inviter.UUID); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Repository) HandleFriendInvite(ctx context.Context, invitee *models.User, inviter *models.User, status string) error {
	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const queryFriendInviteExists = `
						SELECT EXISTS ( SELECT * FROM users_friend_invites WHERE invitee = $1 AND inviter = $2 AND status = 'pending' )
						`
	var friendInviteExists bool
	if err := tx.QueryRowxContext(ctx, queryFriendInviteExists, invitee.UUID, inviter.UUID).Scan(&friendInviteExists); err != nil {
		return err
	}
	if !friendInviteExists {
		return models.ErrNoFriendRequest
	}

	const queryUpdateInvite = `
							WITH cte AS (
								SELECT
									id
								FROM
									users_friend_invites
								WHERE invitee = $1 AND inviter = $2 AND status = 'pending'
							)
							UPDATE
								users_friend_invites
							SET
								status = $3
							FROM cte
							WHERE cte.id = users_friend_invites.id
							`
	if _, err := tx.ExecContext(ctx, queryUpdateInvite, invitee.UUID, inviter.UUID, status); err != nil {
		return err
	}

	if status == "accepted" {
		const queryAddFriend = `
					INSERT INTO
						users_friends (user_id1, user_id2)
					VALUES ($1, $2);
					`
		if _, err := tx.ExecContext(ctx, queryAddFriend, invitee.UUID, inviter.UUID); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Repository) DeleteFriend(ctx context.Context, user1 *models.User, user2 *models.User) (bool, error) {
	const query = `
				DELETE FROM
					users_friends
				WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1);
				`
	result, err := s.database.ExecContext(ctx, query, user1.UUID, user2.UUID)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected != 0, nil
}

func (s *Repository) Search(ctx context.Context, name string, pageLimit, pageOffset int) ([]*models.User, error) {
	result := []*models.User{}

	if len(name) < models.MIN_NAME_SEARCH_LENGTH {
		return result, nil
	}

	const query = `
				SELECT
					id,
					avatar,
					display_name
				FROM users
				WHERE LOWER(display_name) LIKE '%' || LOWER($1) || '%'
				ORDER BY display_name
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, name, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.User{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetFriends(ctx context.Context, user *models.User, pageLimit, pageOffset int) ([]*models.User, error) {
	result := []*models.User{}

	const query = `
				SELECT
					user_id2 AS id,
					avatar,
					display_name
				FROM users_friends
					JOIN users ON users.id = users_friends.user_id2
				WHERE user_id1 = $1
				UNION
				SELECT
					user_id1 AS id,
					avatar,
					display_name
				FROM users_friends
					JOIN users ON users.id = users_friends.user_id1
				WHERE user_id2 = $1
				ORDER BY display_name
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.User{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetInvitesOutgoing(ctx context.Context, user *models.User, pageLimit, pageOffset int) ([]*models.User, error) {
	result := []*models.User{}

	const query = `
				SELECT
					invitee AS id,
					avatar,
					display_name
				FROM users_friend_invites
					JOIN users ON users.id = users_friend_invites.invitee
				WHERE inviter = $1 AND status = 'pending'
				ORDER BY users_friend_invites.created_at
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.User{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetInvitesIncoming(ctx context.Context, user *models.User, pageLimit, pageOffset int) ([]*models.User, error) {
	result := []*models.User{}

	const query = `
				SELECT
					inviter AS id,
					avatar,
					display_name
				FROM users_friend_invites
					JOIN users ON users.id = users_friend_invites.inviter
				WHERE invitee = $1 AND status = 'pending'
				ORDER BY users_friend_invites.created_at
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.User{}
		rows.Scan(&row.UUID, &row.Avatar, &row.DisplayName)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetGames(ctx context.Context, user *models.User, term string, pageLimit, pageOffset int) (*models.Games, error) {
	result := []*models.Product{}

	var total int
	const queryTotal = `
				SELECT
					COUNT(*)
				FROM users_games
				WHERE user_id = $1;
				`
	if err := s.database.QueryRowxContext(ctx, queryTotal, user.UUID).Scan(&total); err != nil {
		return nil, err
	}

	const query = `
				SELECT
					products.id,
					products.name,
					products_images.tier_background_img
				FROM users_games
					JOIN products ON users_games.product_id = products.id
					JOIN products_images ON users_games.product_id = products_images.product_id
				WHERE user_id = $1 AND LOWER(products.name) LIKE '%' || LOWER($2) || '%'
				ORDER BY users_games.created_at
				LIMIT $3 OFFSET $4;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID, term, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.Product{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return &models.Games{Games: result, Total: total}, nil
}
