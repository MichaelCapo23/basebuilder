package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/MichaelCapo23/basebuilder/cmd/gateway"
	"github.com/MichaelCapo23/basebuilder/internal/auth"
	"github.com/MichaelCapo23/basebuilder/internal/user"
	"github.com/MichaelCapo23/basebuilder/pkg/firebase"
	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/MichaelCapo23/basebuilder/pkg/repository/postgres"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("PG_READER_URI", "postgres://root:postgres@localhost/localauth?sslmode=disable")
	viper.AutomaticEnv()
}

func main() {
	fs := flag.NewFlagSet("basebuilder", flag.ExitOnError)
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
		authService = auth.NewService(ctx, internalLogger, db, fb)
		userService = user.NewService(ctx, internalLogger, db, authService)
	)

	server := gateway.New(ctx, internalLogger, db, fb, *addr, authService, userService)
	logger.Infow("starting server")

	server.Serve(ctx)
	done()
}
