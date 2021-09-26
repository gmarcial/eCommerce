package mock

import "gmarcial/eCommerce/core/catalog/domain/products"

//ProductRepository interface to mock access and manipulate products data source.
type ProductRepository struct {
	GetProductsMock func (IDS []uint32) ([]*products.Product, error)
	GetGiftProductMock func () (*products.Product, error)
}

//GetProducts mock retrieve the products through your IDS.
func (repository *ProductRepository) GetProducts(IDS []uint32) ([]*products.Product, error) {
	return repository.GetProductsMock(IDS)
}

//GetGiftProduct mock retrieve a product to gift.
func (repository *ProductRepository) GetGiftProduct() (*products.Product, error) {
	return repository.GetGiftProductMock()
}