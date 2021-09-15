package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUserByExternalID(ctx context.Context, externalID string) (*models.User, error) {
	var user models.User

	q := `
	SELECT
		id,
		email,
		first_name,
		last_name,
		email_verified,
		banned,
		deleted,
		created_at,
		updated_at
	FROM 
		"user"
	WHERE
		external_id = $1`

	if err := s.db.Get(&user, q, externalID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, project.NotFound
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	q := `
	SELECT
		id,
		email,
		first_name,
		last_name,
		email_verified,
		banned,
		deleted,
		created_at,
		updated_at
	FROM 
		"user"
	WHERE
		id = $1`

	if err := s.db.Get(&user, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, project.NotFound
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) SetUser(ctx context.Context, user models.User) error {
	q := `
	INSERT INTO
		"user"
	(id, external_id, email, first_name, last_name)
		VALUES
	(:id, :external_id, :email, :first_name, :last_name)
		ON CONFLICT
	(external_id)
		DO UPDATE SET
	(id, external_id, email, first_name, last_name, email_verified, banned, deleted, updated_at)
		=
	(EXCLUDED.id, EXCLUDED.external_id, EXCLUDED.email, EXCLUDED.first_name, EXCLUDED.last_name, EXCLUDED.email_verified, EXCLUDED.banned, EXCLUDED.deleted, EXCLUDED.updated_at)
	`

	rows, err := s.db.NamedExec(q, user)
	if err != nil {
		return project.Conflict
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == int64(0) {
		return project.Conflict
	}

	return nil
}
