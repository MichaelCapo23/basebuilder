package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CtxUserID string = "user_id"

func Auth(next http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// write auth middleware
	}
}
