package application

import (
	"gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/application/model"
	"gmarcial/eCommerce/core/checkout/application/promotional"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"gmarcial/eCommerce/platform/infrastructure/log"
)

//SelectedProducts represent the products selects to add in cart, entry contract of the MakeCartUseCase.
type SelectedProducts struct {
	Products []*model.SelectedProduct `json:"products"`
}

//CartProducts represent the shopping cart, output contract of the MakeCartUseCase.
type CartProducts struct {
	TotalAmount    uint64           `json:"total_amount"`
	TotalAmountNet uint64           `json:"total_amount_with_discount"`
	TotalDiscount  uint64           `json:"total_discount"`
	Products       []*model.Product `json:"products"`
}

//IMakeCartUseCase the interface to api with behavior of construct the shopping cart
type IMakeCartUseCase interface {
	Execute(selectedProducts *SelectedProducts) (*CartProducts, error)
}

//MakeCartUseCase the api with behavior of construct the shopping cart
type MakeCartUseCase struct {
	logger                   log.Logger
	pickProductsUseCase      application.IPickProductsUseCase
	promotionsApplierUseCase promotional.IPromotionsApplierUseCase
}

//NewMakeCartUseCase constructor to instantiate the use case to make the shopping cart.
func NewMakeCartUseCase(logger log.Logger, pickProductsUseCase application.IPickProductsUseCase, promotionsApplierUseCase promotional.IPromotionsApplierUseCase) *MakeCartUseCase {
	return &MakeCartUseCase{
		logger:                   logger,
		pickProductsUseCase:      pickProductsUseCase,
		promotionsApplierUseCase: promotionsApplierUseCase,
	}
}

//Execute the use case
func (useCase *MakeCartUseCase) Execute(selectedProducts *SelectedProducts) (*CartProducts, error) {
	logger := useCase.logger
	logger.Infow("try mount the cart.")

	if len(selectedProducts.Products) == 0 {
		logger.Infow("none product was selected.",
			"selected products", selectedProducts)

		return &CartProducts{Products: []*model.Product{}}, nil
	}

	pickProducts := SelectedProductsToPickProducts(selectedProducts)

	logger.Infow("picking the selected products.",
		"pick products", pickProducts)

	pickedProducts, err := useCase.pickProductsUseCase.Execute(pickProducts)
	if err != nil {
		logger.Errorw("an error occurred to try pick the products selected.",
			"error", err.Error())
		return nil, err
	}

	products, err := PickedProductsToPurchaseProduct(pickedProducts, selectedProducts)
	if err != nil {
		logger.Errorw("an error occurred to translate the PickedProducts model to PurchaseProduct.",
			"error", err.Error())
		return nil, err
	}

	logger.Infow("mounting the cart with the picked products.",
		"products", products)

	cart, err := purchase.NewCart(products)
	if err != nil {
		logger.Errorw("an error occurred to mount the cart.",
			"error", err.Error())
		return nil, err
	}

	logger.Infow("applying the active promotions to cart.")

	useCase.promotionsApplierUseCase.Apply(cart)

	logger.Infow("cart was mounted.",
		"cart", cart)

	return CartToCartProducts(cart), nil
}
