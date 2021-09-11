package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"google.golang.org/api/option"
)

func NewFirebase(ctx context.Context, dbCredsFile string) *firebase.App {
	opt := option.WithCredentialsFile(dbCredsFile)
	l := logging.FromContext(ctx)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		l.Fatalw("failed to initialize firebase", "err", err)
	}
	return app
}

func GenerateToken(ctx context.Context, fbApiKey string, fbAuth *auth.Client, uid string) (string, error) {
	idToken, err := fbAuth.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}

	return idToken, nil
}
