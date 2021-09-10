package router

import (
	"context"

	"github.com/MichaelCapo23/jwtserver/cmd/gateway/middleware"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres/jwks"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserStore *user.UserStore
	JWKSStore *jwks.JWKSStore
}

// Routes sets up the routes for the server.
func (r *Router) Routes(ctx context.Context, baseRouter *gin.Engine) {

	baseRouter.Use(middleware.CORS())
	baseRouter.Use(middleware.TraceMiddleware())
	baseRouter.Use(middleware.LoggingMiddleware(ctx))

	//use auth middleware on all routes except the jwt/jwks routes

}
