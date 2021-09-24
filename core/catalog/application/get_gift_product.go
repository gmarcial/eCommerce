package application

import (
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/application/translator"
	"gmarcial/eCommerce/core/catalog/domain/products"
)

//GiftProductObtained represent the picked products, output contract of the IGetGiftProductUseCase.
type GiftProductObtained struct {
	Product *model.Product
}

//IGetGiftProductUseCase the interface to api with behavior to get a gift product.
type IGetGiftProductUseCase interface {
	Execute () (*GiftProductObtained, error)
}

//GetGiftProductUseCase the interface to api with behavior to get a gift product.
type GetGiftProductUseCase struct {
	productRepository     products.ProductRepository
}

//Execute the use case
func (useCase GetGiftProductUseCase) Execute() (*GiftProductObtained, error) {
	product, err := useCase.productRepository.GetGiftProduct()
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	giftProduct := translator.DomainProductToModelProduct(product)

	return &GiftProductObtained{Product: giftProduct}, nil
}
