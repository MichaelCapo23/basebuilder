package router

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/jwtserver/cmd/gateway/middleware"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
	"github.com/MichaelCapo23/jwtserver/internal/user"
	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the server.
func AddRoutes(
	ctx context.Context,
	baseRouter *gin.Engine,
	fb *firebase.App,
	authService *auth.AuthService,
	parentService *user.UserService,
) {
	// use cors/logging/trace middleware on all routes
	baseRouter.Use(middleware.CORS())
	baseRouter.Use(middleware.LoggingMiddleware(ctx))
	baseRouter.Use(middleware.TraceMiddleware())

	// authorizedV1 := baseRouter.Group("/v1", middleware.AuthJWT(fb)) //fix AuthJWT
	authorizedV1 := baseRouter.Group("/v1")

	authorizedV1.GET("/profile", parentService.HandleGetProfile())
}
