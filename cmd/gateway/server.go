package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/MichaelCapo23/jwtserver/cmd/gateway/router"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres/jwks"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres/user"
	"github.com/gin-gonic/gin"
)

type Api struct {
	addr   string
	router http.Handler
}

func New(ctx context.Context, db *postgres.PsqlDB, addr string, auth *auth.AuthService) *Api {
	gin.SetMode(gin.ReleaseMode)
	baseRouter := gin.New()

	UserStore := user.NewUserStore(db.WriterX)
	JWKSStore := jwks.NewJWKSStore(db.WriterX)

	router := &router.Router{
		UserStore: UserStore,
		JWKSStore: JWKSStore,
	}

	//add routes to baseRouter
	router.Routes(ctx, baseRouter)

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
