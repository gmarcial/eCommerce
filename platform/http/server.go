package http

import (
	"github.com/go-chi/chi"
	"github.com/sarulabs/di"
	"gmarcial/eCommerce/platform/configuration"
	"log"
	"net/http"
)

//ListenAndServe run the http server
func ListenAndServe(container di.Container, configuration *configuration.Configuration) {
	router := chi.NewRouter()

	MiddlewarePipeline(router)
	Routes(router, container)

	server := &http.Server{
		Addr:    configuration.HttpServerPort,
		Handler: router,
	}

	log.Print(server.ListenAndServe())
}
