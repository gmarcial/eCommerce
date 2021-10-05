package handlers

import (
	"encoding/json"
	"errors"
	applicationCheckout "gmarcial/eCommerce/core/checkout/application"
	"gmarcial/eCommerce/core/checkout/application/model"
	"io"
)

const (
	leftCurlyBracket   = json.Delim('{')
	rightCurlyBracket  = json.Delim('}')
	leftSquareBracket  = json.Delim('[')
	rightSquareBracket = json.Delim(']')
	arrayName          = "products"
	idProperty = "id"
	quantityProperty = "quantity"
	errInvalidPayload = "the payload does not deserialize to SelectedProducts"
)

//decode the payload of endpoint to mount the shopping carts
func decode(body io.ReadCloser) (*applicationCheckout.SelectedProducts, error) {
	decoder := json.NewDecoder(body)
	selectedProducts := make([]*model.SelectedProduct, 0)

	var id uint32
	var quantity uint32
	var counter uint16
	var leftCurlyBracketFound, arrayNameFound, leftSquareBracketFound bool
	for {
		token, err := decoder.Token()
		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		if !(leftCurlyBracketFound && arrayNameFound && leftSquareBracketFound) {
			if counter == 3 {
				return nil, errors.New(errInvalidPayload)
			}
			counter++

			if !leftCurlyBracketFound {
				leftCurlyBracketFound = analyseToken(token, leftCurlyBracket)
				continue
			} else if !arrayNameFound {
				arrayNameFound = analyseToken(token, arrayName)
				continue
			} else if !leftSquareBracketFound {
				leftSquareBracketFound = analyseToken(token, leftSquareBracket)
				continue
			}

			continue
		}

		if analyseToken(token, idProperty) {
			token, err := decoder.Token()
			if err != nil && err != io.EOF {
				return nil, err
			} else if err == io.EOF {
				break
			}

			id = getUInt32Value(token)
		} else if analyseToken(token, quantityProperty) {
			token, err := decoder.Token()
			if err != nil && err != io.EOF {
				return nil, err
			} else if err == io.EOF {
				break
			}

			quantity = getUInt32Value(token)
		} else if analyseToken(token, rightSquareBracket) {
			break
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

	token, err := decoder.Token()
	if err != nil && err != io.EOF {
		return nil, err
	}
	if !analyseToken(token, rightCurlyBracket) {
		return nil, errors.New(errInvalidPayload)
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

func analyseToken(token json.Token, expectedToken interface{}) bool {
	switch expectedTokenValue := expectedToken.(type) {
	case json.Delim:
		if token == expectedTokenValue {
			return true
		}
	case string:
		if token == expectedTokenValue {
			return true
		}
	}

	return false
}
