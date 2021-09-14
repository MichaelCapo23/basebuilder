package router

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/basebuilder/cmd/gateway/middleware"
	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/internal/user"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the server.
func AddRoutes(
	ctx context.Context,
	logger *logging.InternalLogger,
	baseRouter *gin.Engine,
	fb *firebase.App,
	authService *auth.AuthService,
	userService *user.UserService,
) {
	// use cors/logging/trace middleware on all routes
	baseRouter.Use(middleware.CORS())
	baseRouter.Use(middleware.LoggingMiddleware(logger))
	baseRouter.Use(middleware.TraceMiddleware(logger))

	//unauthorized routes
	// unauthorizedV1 := baseRouter.Group("/")

	//authorized routes
	authorizedV1 := baseRouter.Group("/v1", middleware.AuthJWT(fb, logger, authService))
	authorizedV1.GET("/profile", userService.HandleGetUserProfile())
	authorizedV1.POST("/signup", userService.HandleSignUp()) //use after FE creates fb account and passes token/params to create user locally
}
