package application

import (
	"gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/application/model"
	"gmarcial/eCommerce/core/checkout/application/promotional"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
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
	pickProductsUseCase      application.IPickProductsUseCase
	promotionsApplierUseCase promotional.IPromotionsApplierUseCase
}

//NewMakeCartUseCase constructor to instantiate the use case to make the shopping cart.
func NewMakeCartUseCase(pickProductsUseCase application.IPickProductsUseCase, promotionsApplierUseCase promotional.IPromotionsApplierUseCase) *MakeCartUseCase {
	return &MakeCartUseCase{
		pickProductsUseCase:      pickProductsUseCase,
		promotionsApplierUseCase: promotionsApplierUseCase,
	}
}

//Execute the use case
func (useCase *MakeCartUseCase) Execute(selectedProducts *SelectedProducts) (*CartProducts, error) {
	pickProducts := SelectedProductsToPickProducts(selectedProducts)
	if len(pickProducts.IDS) == 0 {
		return &CartProducts{Products: []*model.Product{}}, nil
	}

	pickedProducts, err := useCase.pickProductsUseCase.Execute(pickProducts)
	if err != nil {
		return nil, err
	}

	products, err := PickedProductsToPurchaseProduct(pickedProducts, selectedProducts)
	if err != nil {
		return nil, err
	}

	cart, err := purchase.NewCart(products)
	if err != nil {
		return nil, err
	}

	useCase.promotionsApplierUseCase.Apply(cart)

	return CartToCartProducts(cart), nil
}
