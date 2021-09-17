package memory

import (
	"gmarcial/eCommerce/core/catalog/domain/products"
)

//ProductRepository interface to access and manipulate products data source in-memory.
type ProductRepository struct {
	products map[uint32]*products.Product
}

//NewProductRepository constructor to instantiate the repository of the product.
func NewProductRepository(products map[uint32]*products.Product) *ProductRepository {
	return &ProductRepository{
		products: products,
	}
}

//GetProducts retrieve the products through your IDS.
func (repository *ProductRepository) GetProducts(IDS []uint32) ([]*products.Product, error) {
	productsRetrieve := make([]*products.Product, 0)

	for i := 0; i < len(IDS); i++ {
		ID := IDS[i]
		if product, exist := repository.products[ID]; exist {
			productsRetrieve = append(productsRetrieve, product)
		}
	}

	return productsRetrieve, nil
}

//GetGiftProduct retrieve a product to gift.
func (repository *ProductRepository) GetGiftProduct() (*products.Product, error) {
	var productRetrieve *products.Product

	for _, product := range repository.products {
		if !product.IsGift {
			continue
		}

		productRetrieve = product
	}

	return productRetrieve, nil
}
