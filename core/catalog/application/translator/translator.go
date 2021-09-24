package translator

import (
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
)

//DomainProductToModelProduct translates the product of domain representation to model.
func DomainProductToModelProduct(product *products.Product) *model.Product{
	return &model.Product{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Amount:      product.Amount,
		Discount:    product.Discount,
		IsGift:      product.IsGift,
	}
}
