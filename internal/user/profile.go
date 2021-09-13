package user

import (
	"net/http"

	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/gin-gonic/gin"
)

func (s *UserService) HandleGetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		logger := logging.FromContext(ctx).Named("HandleUpdateProfile")

		claims := auth.FromContext(ctx)
		if claims == nil {
			logger.Errorw("missing claims")
			c.Writer.WriteHeader(http.StatusForbidden)
			return
		}

		c.JSON(http.StatusOK, claims.User)
	}
}
