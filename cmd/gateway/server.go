package gateway

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go"
	"github.com/MichaelCapo23/basebuilder/cmd/gateway/router"
	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/internal/user"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
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
	router.AddRoutes(ctx, logger, baseRouter, fb, authService, userService)

	return &Api{
		addr:   addr,
		router: baseRouter,
	}
}

func (a *Api) Serve(ctx context.Context) {
	server := &http.Server{
		Addr:    a.addr,
		Handler: a.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
