package handlers

import (
	"encoding/json"
	"github.com/sarulabs/di"
	applicationCheckout "gmarcial/eCommerce/core/checkout/application"
	"net/http"
)

//HandleMakeCart the entry-port responsible per expose the construction of make cart in http interface
func HandleMakeCart(container di.Container) http.HandlerFunc {

	requestContainer, _ := container.SubContainer()
	makeCartUseCase := requestContainer.Get("makeCartUseCase").(*applicationCheckout.MakeCartUseCase)

	return func(writer http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var selectedProducts applicationCheckout.SelectedProducts
		err := decoder.Decode(&selectedProducts)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		shoppingCart, err := makeCartUseCase.Execute(&selectedProducts)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(shoppingCart)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
