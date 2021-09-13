package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project"
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

func (s *UserStore) GetUserByID(ctx context.Context, externalID string) (*models.UserProfile, error) {
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
