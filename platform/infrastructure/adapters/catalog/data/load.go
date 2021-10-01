package data

import (
	"encoding/json"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/filepathutil"
	"io"
	"os"
)

const (
	sourceRelativePath = "platform/infrastructure/adapters/catalog/data/source/products.json"
)

func LoadProducts() (map[uint32]*products.Product, error) {
	sourcePath, err := filepathutil.JoinWithRootDir(sourceRelativePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return nil, err
	}

	source := make(map[uint32]*products.Product, 0)
	decoder := json.NewDecoder(file)
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}

	for decoder.More() {
		productModel := &struct {
			ID          uint32 `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Amount      uint64 `json:"amount"`
			IsGift      bool   `json:"is_gift"`
		}{}

		err = decoder.Decode(productModel)
		if err != nil {
			return nil, err
		}

		product, err := products.NewProduct(productModel.ID, productModel.Title, productModel.Description, productModel.Amount, productModel.IsGift)
		if err != nil {
			return nil, err
		}

		if _, exist := source[product.ID]; !exist {
			source[product.ID] = product
		}
	}

	_, err = decoder.Token()
	if err != nil && err != io.EOF {
		return nil, err
	}

	err = file.Close()

	return source, nil
}
