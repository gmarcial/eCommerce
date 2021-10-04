package handlers

import (
	"encoding/json"
	applicationCheckout "gmarcial/eCommerce/core/checkout/application"
	"gmarcial/eCommerce/core/checkout/application/model"
	"io"
)

//decode the payload of endpoint to mount the shopping carts
func decode(body io.ReadCloser) (*applicationCheckout.SelectedProducts, error) {
	decoder := json.NewDecoder(body)
	selectedProducts := make([]*model.SelectedProduct, 0)

	var id uint32
	var quantity uint32

	for  {
		token, err := decoder.Token()
		if err != nil && err != io.EOF {
			return nil, err
		} else if  err == io.EOF {
			break
		}

		stringValue, isString := token.(string)
		if isString && stringValue == "id" {
			token, err := decoder.Token()
			if err != nil && err != io.EOF {
				return nil, err
			} else if  err == io.EOF {
				break
			}

			id = getUInt32Value(token)
		} else if isString && stringValue == "quantity" {
			token, err := decoder.Token()
			if err != nil && err != io.EOF {
				return nil, err
			} else if  err == io.EOF {
				break
			}
			
			quantity = getUInt32Value(token)
		}

		if id > 0 && quantity > 0 {
			selectedProduct := &model.SelectedProduct{
				ID:       id,
				Quantity: quantity,
			}
			selectedProducts = append(selectedProducts, selectedProduct)

			id = 0
			quantity = 0
		}
	}

	return &applicationCheckout.SelectedProducts{Products: selectedProducts}, nil
}

func getUInt32Value(token json.Token) uint32 {
	float64Value := token.(float64)
	uint32Value := uint32(float64Value)
	if uint32Value > 0 && float64(uint32Value) == float64Value {
		return uint32Value
	}

	return 0
}