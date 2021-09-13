package auth

import (
	"context"

	"firebase.google.com/go/auth"
	fbAuth "firebase.google.com/go/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/firebase"

	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres/user"
)

type ContextKey string

var (
	ClaimsKey  ContextKey = "ClaimsKey"
	loggerName string     = "authServiceCtx"
)

type AuthService struct {
	db     *postgres.PsqlDB
	logger *logging.InternalLogger
}

type Claims struct {
	User          *models.UserProfile `json:"user,omitempty"`
	FirebaseToken *fbAuth.Token       `json:"firebase_token,omitempty"`
}

func NewService(ctx context.Context, logger *logging.InternalLogger, db *postgres.PsqlDB) *AuthService {
	l := logging.FromContext(ctx).Named(loggerName)

	l.Infow("initializing auth service")

	return &AuthService{
		db:     db,
		logger: logger,
	}
}

func WithClaims(ctx context.Context, claims *Claims) context.Context {
	ctx = context.WithValue(ctx, ClaimsKey, claims)
	return ctx
}

func FromContext(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(ClaimsKey).(*Claims); ok {
		return claims
	}
	return nil
}

func (s *AuthService) GetClaims(ctx context.Context, IDToken *auth.Token) (*Claims, error) {
	claims := &Claims{
		FirebaseToken: IDToken,
	}

	externalID := firebase.GetUID(IDToken)

	store := user.NewUserStore(s.db.ReaderX)
	user, err := store.GetClaims(ctx, externalID)
	if err != nil {
		return nil, err
	}

	claims.User = user

	return claims, nil
}
