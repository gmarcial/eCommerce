package application

import (
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/log"
)

//GiftProductObtained represent the picked products, output contract of the IGetGiftProductUseCase.
type GiftProductObtained struct {
	Product *model.Product
}

//IGetGiftProductUseCase the interface to api with behavior to get a gift product.
type IGetGiftProductUseCase interface {
	Execute() (*GiftProductObtained, error)
}

//GetGiftProductUseCase the interface to api with behavior to get a gift product.
type GetGiftProductUseCase struct {
	logger            log.Logger
	productRepository products.ProductRepository
}

//NewGetGiftProductUseCase constructor to instantiate the use case to pick get gift product.
func NewGetGiftProductUseCase(logger log.Logger, productRepository products.ProductRepository) *GetGiftProductUseCase {
	return &GetGiftProductUseCase{
		logger: logger,
		productRepository: productRepository,
	}
}

//Execute the use case
func (useCase GetGiftProductUseCase) Execute() (*GiftProductObtained, error) {
	logger := useCase.logger
	logger.Infow("try get a gift product.")

	product, err := useCase.productRepository.GetGiftProduct()
	if err != nil {
		logger.Errorw("an error occurred to try get a gift product.",
			"error", err.Error())
		return nil, err
	}

	if product == nil {
		logger.Infow("none gift product was founded.")
		return nil, nil
	}

	giftProduct := DomainProductToModelProduct(product)

	logger.Infow("A gift product was obtained.",
		"gift product", giftProduct)

	return &GiftProductObtained{Product: giftProduct}, nil
}
