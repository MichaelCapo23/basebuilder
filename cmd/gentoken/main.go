package main

import (
	"context"
	"flag"
	"log"

	"github.com/MichaelCapo23/jwtserver/pkg/firebase"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/spf13/viper"
)

var DefaultUID string

func init() {
	viper.AutomaticEnv()
	flag.StringVar(&DefaultUID, "uid", "WSHIIlUNgYQ0ZJuZGBByrlzulTB3", "firebase UID flag")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	internalLogger := logging.NewLogger(false)
	ctx = logging.WithLogger(ctx, internalLogger)

	logger := logging.FromContext(ctx)
	internalLogger.Logger = logger

	fbConfigFile := viper.GetString("FIREBASE_CONFIG_FILE")
	fbApiKey := viper.GetString("FB_API_KEY")
	fb := firebase.NewFirebase(ctx, fbConfigFile)

	client, err := fb.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	t, _ := firebase.GenerateToken(ctx, fbApiKey, client, DefaultUID)
	internalLogger.Infow("success", "firebase_token", t)
}
