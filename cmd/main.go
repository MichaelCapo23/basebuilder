package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/MichaelCapo23/jwtserver/cmd/gateway"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/MichaelCapo23/jwtserver/pkg/repository/postgres"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("AIRVET_PORT", "3000")
	viper.SetDefault("PG_READER_URI", "postgres://root:postgres@localhost/localauth?sslmode=disable")
	viper.AutomaticEnv()
}

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logger := logging.NewLogger(viper.GetBool("AIRVET_DEBUG"))
	defer logger.Sync()

	logger.Infow("initializing admin-api")

	ctx = logging.WithLogger(ctx, logger)

	db := postgres.NewDBFromSql(viper.GetString("PG_READER_URI"))
	port := viper.GetInt("AIRVET_PORT")

	api := gateway.New() //fill out

	logger.Infow("starting server")

	api.Serve(ctx)
	done()
}
