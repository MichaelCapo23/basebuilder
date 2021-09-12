package auth

import (
	"context"

	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
)

type AuthService struct {
	db     *postgres.PsqlDB
	logger *logging.InternalLogger
}

const (
	loggerName string = "authServiceCtx"
)

func NewService(ctx context.Context, logger *logging.InternalLogger, db *postgres.PsqlDB) *AuthService {
	l := logging.FromContext(ctx).Named(loggerName)

	l.Infow("initializing auth service")

	return &AuthService{
		db:     db,
		logger: logger,
	}
}
