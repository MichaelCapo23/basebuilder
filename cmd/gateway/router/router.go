package router

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/jwtserver/cmd/gateway/middleware"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the server.
func AddRoutes(
	ctx context.Context,
	baseRouter *gin.Engine,
	fb *firebase.App,
	authService *auth.AuthService) {

	// use cors/logging/trace middleware on all routes
	baseRouter.Use(middleware.CORS())
	baseRouter.Use(middleware.LoggingMiddleware(ctx))
	baseRouter.Use(middleware.TraceMiddleware())

	baseRouter.Use(middleware.AuthJWT(fb))

}
