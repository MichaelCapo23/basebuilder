package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/MichaelCapo23/jwtserver/pkg/project"
	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"google.golang.org/api/option"
)

func NewFirebase(ctx context.Context, dbCredsFile string) *firebase.App {
	opt := option.WithCredentialsFile(dbCredsFile)
	l := logging.FromContext(ctx)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		l.Fatalw("failed to initialize firebase", "err", err)
	}
	return app
}

// GenerateToken will generate a firebase jwt token based off of the uid
func GenerateToken(ctx context.Context, fbApiKey string, fbAuth *auth.Client, uid string) (string, error) {
	customToken, err := fbAuth.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}
	reqByte, _ := json.Marshal(struct {
		Token             string `json:"token"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}{
		customToken,
		true,
	})
	reqBody := bytes.NewReader(reqByte)
	uri := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s", fbApiKey)
	resp, err := http.DefaultClient.Post(uri, "application/json", reqBody)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	var respJSON map[string]interface{}
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		return "", err
	}

	jwt, ok := respJSON["idToken"]
	if !ok {
		return "", fmt.Errorf("idToken %w", project.NotFound)
	}

	return fmt.Sprintf("%s", jwt), nil
}
