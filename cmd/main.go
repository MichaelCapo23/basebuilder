package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/MichaelCapo23/jwtserver/cmd/gateway"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
	"github.com/MichaelCapo23/jwtserver/internal/user"
	"github.com/MichaelCapo23/jwtserver/pkg/firebase"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("PG_READER_URI", "postgres://root:postgres@localhost/localauth?sslmode=disable")
	viper.AutomaticEnv()
}

func main() {
	fs := flag.NewFlagSet("jwtserver", flag.ExitOnError)
	addr := fs.String("addr", ":3000", "server address")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	fbConfigFile := viper.GetString("FIREBASE_CONFIG_FILE")
	fb := firebase.NewFirebase(ctx, fbConfigFile)

	internalLogger := logging.NewLogger(false)
	ctx = logging.WithLogger(ctx, internalLogger)

	logger := logging.FromContext(ctx)
	internalLogger.Logger = logger

	db := postgres.NewDBFromSql(viper.GetString("PG_WRITER_URI"))

	var (
		authService = auth.NewService(ctx, internalLogger, db)
		userService = user.NewService(ctx, internalLogger, db, authService)
	)

	server := gateway.New(ctx, internalLogger, db, fb, *addr, authService, userService)
	logger.Infow("starting server")

	server.Serve(ctx)
	done()
}
