package application

import (
	"context"
	"errors"
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
	"gmarcial/eCommerce/platform/infrastructure/log"
)

//PickProducts represent the products to be picked, entry contract of the PickProductsUseCase.
type PickProducts struct {
	IDS []uint32
}

//PickedProducts represent the picked products, output contract of the PickProductsUseCase.
type PickedProducts struct {
	Products []*model.Product
}

//IPickProductsUseCase the interface to api with behavior to pick the selected products.
type IPickProductsUseCase interface {
	Execute (pickProducts *PickProducts) (*PickedProducts, error)
}

//PickProductsUseCase the api with behavior to pick the selected products.
type PickProductsUseCase struct {
	logger            log.Logger
	productRepository     products.ProductRepository
	discountServiceClient client.DiscountClient
}

//NewPickProductsUseCase constructor to instantiate the use case to pick products.
func NewPickProductsUseCase(logger log.Logger, productRepository products.ProductRepository, discountServiceClient client.DiscountClient) *PickProductsUseCase {
	return &PickProductsUseCase{
		logger: logger,
		productRepository:     productRepository,
		discountServiceClient: discountServiceClient,
	}
}

//Execute the use case
func (useCase *PickProductsUseCase) Execute(pickProducts *PickProducts) (*PickedProducts, error) {
	logger := useCase.logger
	logger.Infow("try to pick the products.")

	if pickProducts == nil {
		logger.Errorw("don't was informed the products to be picked")
		return nil, errors.New("don't was informed the products to be picked")
	}

	IDS := pickProducts.IDS
	if IDS == nil || len(IDS) == 0 {
		logger.Infow("don't have products to picked.")
		return nil, nil
	}

	productRepository := useCase.productRepository
	productsPicked, err := productRepository.GetProducts(IDS)
	if err != nil {
		logger.Errorw("an error occurred to get the products.",
			"error", err.Error())
		return nil, err
	}

	logger.Infow("the products were obtained.",
		"products picked", productsPicked)

	logger.Infow("start try to apply discount in products")

	productsModel := make([]*model.Product, 0)
	discountServiceClient := useCase.discountServiceClient
	request := &client.GetDiscountRequest{}
	for _, product := range productsPicked {
		request.ProductID = int32(product.ID)
		response, err := discountServiceClient.GetDiscount(context.Background(), request)
		if err != nil {
			logger.Errorw("an error occurred to get discount.",
				"product", product,
				"error", err.Error())
		} else {
			discountPercentage := response.GetPercentage()
			if err := product.ApplyDiscount(discountPercentage); err != nil {
				logger.Errorw("an error occurred to apply the discount.",
					"product", product,
					"error", err.Error())
			}
		}

		productsModel = append(productsModel, &model.Product{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Amount:      product.Amount,
			Discount:    product.Discount,
			IsGift:      product.IsGift,
		})
	}

	logger.Infow("the discount was applied to products.")

	logger.Infow("the products was picked.",
		"products", productsModel)

	return &PickedProducts{Products: productsModel}, nil
}
