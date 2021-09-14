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

func (s *UserStore) GetUserByExternalID(ctx context.Context, externalID string) (*models.UserProfile, error) {
	var user models.UserProfile

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

func (s *UserStore) GetUserByID(ctx context.Context, id uuid.UUID) (*models.UserProfile, error) {
	var user models.UserProfile

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

func (s *UserStore) Create(ctx context.Context, opts models.CreateUserRepo) error {
	q := `
	INSERT INTO 
		"user"
	(id, external_id, email, first_name, last_name) 
		VALUES 
	(:id, :external_id, :email, :first_name, :last_name)`

	if _, err := s.db.NamedExecContext(ctx, q, opts); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return project.NotFound
		}
		return err
	}

	return nil
}
