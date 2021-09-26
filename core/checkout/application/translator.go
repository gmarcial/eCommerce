package application

import (
	catalog "gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/application/model"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
)

//SelectedProductsToPickProducts translates the entry contract of MakeCartUseCase to PickProductsUseCase.
func SelectedProductsToPickProducts(selectedProducts *SelectedProducts) *catalog.PickProducts {
	selectedProductsIDS := make([]uint32, 0)
	for _, selectedProduct := range selectedProducts.Products {
		selectedProductsIDS = append(selectedProductsIDS, selectedProduct.ID)
	}
	return &catalog.PickProducts{IDS: selectedProductsIDS}
}

//PickedProductsToPurchaseProduct translates output contract of the PickProductsUseCase to domain representation of product.
func PickedProductsToPurchaseProduct(pickedProducts *catalog.PickedProducts, selectedProducts *SelectedProducts) ([]*purchase.Product, error) {
	products := make([]*purchase.Product, 0)
	for i:= 0; i < len(pickedProducts.Products); i++ {
		pickedProduct := pickedProducts.Products[i]
		quantity := selectedProducts.Products[i].Quantity
		product, err := purchase.NewProduct(pickedProduct.ID, quantity, pickedProduct.Amount, pickedProduct.Discount, pickedProduct.IsGift)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

//CartToCartProducts translate the domain representation the entity of cart to output contract of the MakeCartUseCase.
func CartToCartProducts(cart *purchase.Cart) *CartProducts {
	cartProducts := make([]*model.Product, 0)
	for _, product := range cart.Products {
		cartProduct := &model.Product{
			ID:          product.ID,
			Quantity:    product.Quantity,
			UnitAmount:  product.UnitAmount,
			TotalAmount: product.TotalAmount,
			Discount:    product.Discount,
			IsGift:      product.IsGift,
		}

		cartProducts = append(cartProducts, cartProduct)
	}

	return &CartProducts{
		TotalAmount:    cart.TotalAmount,
		TotalAmountNet: cart.TotalAmountNet,
		TotalDiscount:  cart.TotalDiscount,
		Products:       cartProducts,
	}
}