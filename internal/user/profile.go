package user

import (
	"net/http"
	"time"

	"github.com/MichaelCapo23/jwtserver/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (p *UserService) HandleGetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := models.UserProfile{
			ID:            uuid.New(),
			FirstName:     "GET ME",
			LastName:      "DATA ALREADY",
			Email:         "michael@airvet.com",
			EmailVerified: false,
			Banned:        false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// get gin context
		// name the logging from ctx
		// get claims for authService
		// get profile
		// check errors

		c.JSON(http.StatusOK, p)
	}
}
