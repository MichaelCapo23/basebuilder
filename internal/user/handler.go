package user

import (
	"net/http"

	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/pkg/models"
	"github.com/MichaelCapo23/basebuilder/pkg/project"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres/user"
	"github.com/gin-gonic/gin"
)

func (s *UserService) HandleGetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		logger := logging.FromContext(ctx).Named("HandleUpdateProfile")

		claims := auth.FromContext(ctx)
		if claims == nil {
			logger.Errorw("missing claims")
			c.JSON(http.StatusForbidden, project.NotAllowed)
			return
		}

		c.JSON(http.StatusOK, claims.User)
	}
}

func (s *UserService) HandleSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userProfile models.User
		ctx := c.Request.Context()

		claims := auth.FromContext(ctx)

		// create logger
		internalLogger := logging.NewLogger(false)
		logger := logging.FromContext(ctx).Named("HandleSignUp")
		internalLogger.Logger = logger

		if err := c.Bind(&userProfile); err != nil {
			s.logger.ErrorCtx(ctx, "parse error", "err", err)
			c.JSON(http.StatusBadRequest, project.BadInput)
			return
		}

		id, err := s.SignUpUser(ctx, userProfile, claims)
		if err != nil {
			s.logger.ErrorCtx(ctx, "error signing up user", "err", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		store := user.NewUserStore(s.db.ReaderX)
		user, err := store.GetUserByID(ctx, id)
		if err != nil {
			s.logger.ErrorCtx(ctx, "error getting new user", "err", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, *user)
	}
}
