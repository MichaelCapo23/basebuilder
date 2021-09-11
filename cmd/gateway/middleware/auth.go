package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/jwtserver/pkg/project"
	"github.com/gin-gonic/gin"
)

const (
	CtxUserID  = "user_id"
	authHeader = "Authorization"
	fbIdToken  = "FIREBASE_ID_TOKEN"
)

func AuthJWT(fbApp *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(authHeader)
		token := strings.Replace(authHeader, "Bearer", "", 1)
		client, err := fbApp.Auth(context.Background())
		IDToken, err := client.VerifyIDToken(c, token)
		if err != nil {
			project.ReturnGinError(c, http.StatusUnauthorized, project.UnauthorizedRequest)
		}

		// pass IDToken to gin context
		c.Set(fbIdToken, IDToken)
		c.Next()
	}
}
