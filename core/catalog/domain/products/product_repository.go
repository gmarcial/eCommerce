package products

//ProductRepository interface to access and manipulate products data source.
type ProductRepository interface {
	//GetProducts retrieve the products through your IDS.
	GetProducts(IDS []uint64) ([]*Product, error)

	//GetGiftProduct retrieve a product to gift.
	GetGiftProduct() (*Product, error)
}
