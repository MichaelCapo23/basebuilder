package gateway

import (
	"context"
	"net/http"
)

type Api struct {
	router http.Handler
}

func New() *Api {

}

func (a *Api) Serve(ctx context.Context) error {

}
