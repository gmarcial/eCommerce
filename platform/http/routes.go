package http

import (
	"github.com/go-chi/chi"
	"github.com/sarulabs/di"
	"gmarcial/eCommerce/platform/http/resources/checkout"
)

//Routes define the routes per module
func Routes(router *chi.Mux, container di.Container) {
	checkout.Routes(router, container)
}
