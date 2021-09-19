package application

import (
	"context"
	"errors"
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
)

//PickProductsUseCase the api with behavior to pick the selected products.
type PickProductsUseCase struct {
	productRepository     products.ProductRepository
	discountServiceClient client.DiscountClient
}

//NewPickProductsUseCase constructor to instantiate the use case to pick products.
func NewPickProductsUseCase(productRepository products.ProductRepository, discountServiceClient client.DiscountClient) *PickProductsUseCase {
	return &PickProductsUseCase{
		productRepository:     productRepository,
		discountServiceClient: discountServiceClient,
	}
}

//PickProducts represent the products to be picked, entry contract of the PickProductsUseCase.
type PickProducts struct {
	IDS []uint32
}

//PickedProducts represent the picked products, output contract of the PickProductsUseCase.
type PickedProducts struct {
	Products []*model.Product
}

//Execute ...
func (useCase *PickProductsUseCase) Execute(pickProducts *PickProducts) (*PickedProducts, error) {
	if pickProducts == nil {
		return nil, errors.New("don't was informed the products to be picked")
	}

	IDS := pickProducts.IDS
	if IDS == nil || len(IDS) == 0 {
		return nil, nil
	}

	productRepository := useCase.productRepository
	productsPicked, err := productRepository.GetProducts(IDS)
	if err != nil {
		return nil, err
	}

	productsModel := make([]*model.Product, 0)
	discountServiceClient := useCase.discountServiceClient
	request := &client.GetDiscountRequest{}
	for _, product := range productsPicked {
		request.ProductID = int32(product.ID)
		response, err := discountServiceClient.GetDiscount(context.Background(), request)
		if err != nil {
			//TODO: log error...
		} else {
			discountPercentage := response.GetPercentage()
			if err := product.ApplyDiscount(discountPercentage); err != nil {
				//TODO: log error...
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

	return &PickedProducts{Products: productsModel}, nil
}
