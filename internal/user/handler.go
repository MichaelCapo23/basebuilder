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
		var createUserOpts models.CreateUser
		ctx := c.Request.Context()

		// create logger
		internalLogger := logging.NewLogger(false)
		logger := logging.FromContext(ctx).Named("HandleSignUp")
		internalLogger.Logger = logger

		if err := c.Bind(&createUserOpts); err != nil {
			s.logger.ErrorCtx(ctx, "parse error", "err", err)
			c.JSON(http.StatusBadRequest, project.BadInput)
			return
		}

		uid, err := s.SignUpUser(ctx, createUserOpts)
		if err != nil {
			s.logger.ErrorCtx(ctx, "error signing up user", "err", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		store := user.NewUserStore(s.db.ReaderX)
		userProfile, err := store.GetUserByID(ctx, *uid)

		c.JSON(http.StatusOK, userProfile)
	}
}

// func (s *UserService) HandleLogin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		logger := logging.FromContext(ctx).Named("HandleLogin")

// 		claims := auth.FromContext(ctx)
// 		if claims == nil {
// 			logger.Errorw("missing claims")
// 			c.JSON(http.StatusForbidden, project.NotAllowed)
// 			return
// 		}

// 	}
// }
