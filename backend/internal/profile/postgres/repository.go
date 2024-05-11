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
	_, err := s.database.ExecContext(ctx, query, comment.ProfileID, comment.Commnetator, comment.Text)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) GetComments(ctx context.Context, user *models.User, pageLimit, pageOffset int) ([]*models.Comment, error) {
	const query = `
				SELECT
					commentator,
					text,
					created_at
				FROM users_comments
				WHERE profile_id = $1
				ORDER BY created_at DESC
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.Comment{}

	for rows.Next() {
		row := &models.Comment{}
		rows.StructScan(&row)
		result = append(result, row)
	}

	return result, nil
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

