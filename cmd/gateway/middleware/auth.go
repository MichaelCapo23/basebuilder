package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/project"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/gin-gonic/gin"
)

const (
	CtxUserID  = "user_id"
	authHeader = "Authorization"
	fbIdToken  = "FIREBASE_ID_TOKEN"
)

func AuthJWT(fbApp *firebase.App, internalLogger *logging.InternalLogger, authService *auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		logger := logging.FromContext(ctx).Named("authMiddleware")

		authHeader := c.Request.Header.Get(authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		client, err := fbApp.Auth(context.Background())
		IDToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			project.ReturnGinError(c, http.StatusUnauthorized, project.UnauthorizedRequest)
			c.Abort()
			return
		}

		claims, err := authService.GetClaims(ctx, IDToken)
		if err != nil {
			project.ReturnGinError(c, http.StatusUnauthorized, err)
		}

		// add user id to logger
		logger = logger.With("claims", claims)
		internalLogger.Logger = logger

		ctx = logging.WithLogger(ctx, internalLogger)
		c.Request = c.Request.Clone(auth.WithClaims(ctx, claims))

		c.Next()
	}
}
