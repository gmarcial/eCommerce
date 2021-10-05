package handlers

import (
	"encoding/json"
	"github.com/sarulabs/di"
	applicationCheckout "gmarcial/eCommerce/core/checkout/application"
	"go.uber.org/zap"
	"net/http"
)

//HandleMakeCart the entry-port responsible per expose the construction of make cart in http interface
func HandleMakeCart(container di.Container) http.HandlerFunc {

	requestContainer, _ := container.SubContainer()
	logger := requestContainer.Get("logger").(*zap.SugaredLogger).Named("HandleMakeCart")
	makeCartUseCase := requestContainer.Get("makeCartUseCase").(*applicationCheckout.MakeCartUseCase)

	return func(writer http.ResponseWriter, request *http.Request) {
		selectedProducts, err := decode(request.Body)
		if err != nil {
			logger.Errorw("an error occurred in process of the decode the request body",
				"error", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Infow("received a make cart request.",
			"request", selectedProducts)

		shoppingCart, err := makeCartUseCase.Execute(selectedProducts)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err = encoder.Encode(shoppingCart)
		if err != nil {
			logger.Errorw("an error occurred in process of the encode the response",
				"error", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
