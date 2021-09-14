package user

import (
	"context"

	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres/user"
	"github.com/google/uuid"
)

func (s *UserService) SignUpUser(ctx context.Context, opts models.CreateUser, claims *auth.Claims) (uuid.UUID, error) {
	store := user.NewUserStore(s.db.ReaderX)

	//create new user id
	id := uuid.New()

	// map handler model to repo layer model
	repoOpts := models.CreateUserRepo{
		ID:         id,
		ExternalID: claims.FirebaseToken.UID,
		Email:      opts.Email,
		FirstName:  opts.FirstName,
		LastName:   opts.LastName,
	}

	// add user to internal db
	err := store.Create(ctx, repoOpts)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
