package user

import (
	"context"

	fbAuth "firebase.google.com/go/auth"
	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
)

type UserService struct {
	db          *postgres.PsqlDB
	logger      *logging.InternalLogger
	authService *auth.AuthService
}

type Claims struct {
	User          *models.UserProfile `json:"user,omitempty"`
	FirebaseToken *fbAuth.Token       `json:"firebase_token,omitempty"`
}

const (
	loggerName string = "UserServiceCtx"
)

func NewService(ctx context.Context, logger *logging.InternalLogger, db *postgres.PsqlDB, authService *auth.AuthService) *UserService {
	l := logging.FromContext(ctx).Named(loggerName)

	l.Infow("initializing parent service")

	return &UserService{
		db:          db,
		logger:      logger,
		authService: authService,
	}
}
