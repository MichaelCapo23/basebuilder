package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/MichaelCapo23/jwtserver/cmd/gateway"
	"github.com/MichaelCapo23/jwtserver/internal/auth"
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

	logger := logging.NewLogger(viper.GetBool("AIRVET_DEBUG"))
	defer logger.Sync()

	logger.Infow("initializing admin-api")

	ctx = logging.WithLogger(ctx, logger)

	db := postgres.NewDBFromSql(viper.GetString("PG_READER_URI"))

	var (
		authService = auth.NewService(ctx, db)
	)

	server := gateway.New(ctx, db, *addr, authService)

	logger.Infow("starting server")

	server.Serve(ctx)
	done()
}
