package http

import (
	"net/http"

	"github.com/Cwby333/testapp/pkg/api/v1"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func NewServer(handler Handler) *http.Server {
	strict := api.NewStrictHandler(&handler, nil)

	handler.RegisterRoutes(strict)

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	mwValidator := nethttpmiddleware.OapiRequestValidator(swagger)

	hand := api.HandlerWithOptions(strict, api.StdHTTPServerOptions{
		BaseURL: ":8885",
		BaseRouter: handler.mux,
		Middlewares: []api.MiddlewareFunc{mwValidator},
	})

	server := &http.Server{
		Handler: hand,
		Addr: ":8885",
	}

	return server
}