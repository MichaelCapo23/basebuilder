package user

import (
	"context"
	"time"

	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres/user"
	"github.com/google/uuid"
)

func (s *UserService) SignUpUser(ctx context.Context, userProfile models.User, claims *auth.Claims) (uuid.UUID, error) {
	store := user.NewUserStore(s.db.ReaderX)

	//create new user id
	id := uuid.New()
	userProfile.ID = id
	userProfile.ExternalID = &claims.FirebaseToken.UID
	userProfile.UpdatedAt = time.Now()

	// add user to internal db
	err := store.SetUser(ctx, userProfile)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
