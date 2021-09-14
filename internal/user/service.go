package user

import (
	"context"

	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
)

type UserService struct {
	db          *postgres.PsqlDB
	logger      *logging.InternalLogger
	authService *auth.AuthService
}

const (
	loggerName string = "UserServiceCtx"
)

func NewService(ctx context.Context, logger *logging.InternalLogger, db *postgres.PsqlDB, authService *auth.AuthService) *UserService {
	l := logging.FromContext(ctx).Named(loggerName)

	l.Infow("initializing user service")

	return &UserService{
		db:          db,
		logger:      logger,
		authService: authService,
	}
}
