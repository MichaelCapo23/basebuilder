package auth

import (
	"context"

	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres"
)

type AuthService struct {
	db *postgres.PsqlDB
}

const (
	loggerName string = "auth"
)

func NewService(ctx context.Context, db *postgres.PsqlDB) *AuthService {
	l := logging.FromContext(ctx).Named(loggerName)

	l.Infow("initializing auth service")

	return &AuthService{
		db: db,
	}
}
