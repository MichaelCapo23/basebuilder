package user

import (
	"context"
	"fmt"
	"log"

	fbAuth "firebase.google.com/go/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project"
)

func (s *UserService) SignUpUser(ctx context.Context, opts models.CreateUser) (*string, error) {
	fb := s.authService.GetFirebase()
	client, err := fb.Auth(context.Background())

	params := (&fbAuth.UserToCreate{}).
		Email(opts.Email).
		EmailVerified(false).
		Password(opts.Password).
		DisplayName(fmt.Sprintf("%s %s", opts.FirstName, opts.LastName)).
		Disabled(false)

	if opts.Phone != nil {
		params.PhoneNumber(*opts.Phone)
	}

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		s.logger.ErrorCtx(ctx, "error creating user", "err", err)
		return nil, project.Conflict
	}
	log.Printf("Successfully created user: %v\n", u)

	return &u.UID, nil
}
