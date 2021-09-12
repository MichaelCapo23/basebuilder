package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/jwtserver/cmd/gateway/router"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
	"github.com/MichaelCapo23/jwtserver/internal/user"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres"
	"github.com/gin-gonic/gin"
)

type Api struct {
	addr   string
	router http.Handler
}

func New(ctx context.Context,
	logger *logging.InternalLogger,
	db *postgres.PsqlDB,
	fb *firebase.App,
	addr string,
	authService *auth.AuthService,
	userService *user.UserService,
) *Api {
	gin.SetMode(gin.ReleaseMode)
	baseRouter := gin.New()

	//add routes to baseRouter
	router.AddRoutes(ctx, baseRouter, fb, authService, userService)

	return &Api{
		addr:   addr,
		router: baseRouter,
	}
}

func (a *Api) Serve(ctx context.Context) error {
	server := &http.Server{
		Addr:    a.addr,
		Handler: a.router,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
