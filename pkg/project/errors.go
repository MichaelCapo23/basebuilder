package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	UnauthorizedRequest = &InternalError{
		Code:    "unauthorized_request",
		Message: "unauthorized request",
	}
	BadInput = &InternalError{
		Code:    "bad_input",
		Message: "bad input",
	}
	NotFound = &InternalError{
		Code:    "not_found",
		Message: "not found",
	}
	NotAllowed = &InternalError{
		Code:    "not_allowed",
		Message: "not allowed",
	}
	Conflict = &InternalError{
		Code:    "conflict",
		Message: "conflict",
	}
	UserExists = &InternalError{
		Code:    "user_exists",
		Message: "user already exists",
	}
)

type InternalError struct {
	Code    string
	Message string
	TraceID string
}

func (e *InternalError) Error() string {
	return e.Message
}

func ReturnGinError(c *gin.Context, code int, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    code,
		"message": err.Error(),
	})
}
