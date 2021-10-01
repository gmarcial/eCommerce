package checkout

import (
	"github.com/go-chi/chi"
	"github.com/sarulabs/di"
	"gmarcial/eCommerce/platform/http/resources/checkout/handlers"
)

//Routes define the routes of the checkout module
func Routes(router *chi.Mux, container di.Container) {
	router.Post("/checkout", handlers.HandleMakeCart(container))
}