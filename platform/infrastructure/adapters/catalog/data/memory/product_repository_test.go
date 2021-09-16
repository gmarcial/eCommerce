package memory

import (
	"gmarcial/eCommerce/core/catalog/domain/products"
	"testing"
)

func TestProductRepository_GetProductsThroughYourIDS(t *testing.T) {
	//Arrange
	firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)
	secondProduct, _ := products.NewProduct(2, "Ergonomic Cotton Keyboard", "Iste est ratione excepturi repellendus adipisci qui.", 93811, false)
	thirdProduct, _ := products.NewProduct(3, "Gorgeous Cotton Chips", "Nulla rerum tempore rem.", 60356, false)

	productsInMemory := map[uint64]*products.Product{
		firstProduct.ID:  firstProduct,
		secondProduct.ID: secondProduct,
		thirdProduct.ID:  thirdProduct,
	}

	productRepository := &ProductRepository{products: productsInMemory}

	testCases := []struct {
		name string
		give []uint64
		want []*products.Product
	}{
		{
			"Try get products which don't exist",
			[]uint64{423455234234, 34537665456, 896777686867},
			[]*products.Product{},
		},
		{
			"Get products which exist",
			[]uint64{1, 2, 3},
			[]*products.Product{firstProduct, secondProduct, thirdProduct},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			IDS := testCase.give
			productsRetrieve, err := productRepository.GetProducts(IDS)

			//Assert
			if err != nil {
				t.Errorf("an error has occurred to try to retrieve the products: %v", err.Error())
			}

			quantityProductsRetrieved := len(productsRetrieve)
			quantityWant := len(testCase.want)
			if quantityProductsRetrieved != quantityWant {
				t.Errorf("the number of products retrieve is %v, different of expected %v", quantityProductsRetrieved, quantityWant)
			}

			for i := 0; i < len(testCase.want); i++ {
				productRetrieve := *productsRetrieve[i]
				expectedProduct := *testCase.want[i]

				if productRetrieve != expectedProduct {
					t.Errorf("a product retrieved is %v, different of expected %v:", productRetrieve, expectedProduct)
				}
			}
		})
	}
}

func TestProductRepository_GetAGiftProduct(t *testing.T) {
	//Arrange
	giftProduct, _ := products.NewProduct(6, "Handcrafted Steel Towels", "Nam ea sed animi neque qui non quis iste.", 900, true)

	productsInMemory := map[uint64]*products.Product{
		giftProduct.ID: giftProduct,
	}

	productRepository := &ProductRepository{products: productsInMemory}

	//Action
	giftProductRetrieved, err := productRepository.GetGiftProduct()

	//Assert
	if err != nil {
		t.Errorf("an error has occurred to try to retrieve a gift product: %v", err.Error())
	}

	if !giftProductRetrieved.IsGift {
		t.Error("the product retrieved don't is a gift")
	}
}

func TestProductRepository_TryGetAGiftProductWhenDontExist(t *testing.T) {
	//Arrange
	firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)

	productsInMemory := map[uint64]*products.Product{
		firstProduct.ID: firstProduct,
	}

	productRepository := &ProductRepository{products: productsInMemory}

	//Action
	giftProductRetrieved, err := productRepository.GetGiftProduct()

	//Assert
	if err != nil {
		t.Errorf("an error has occurred to try to retrieve a gift product: %v", err.Error())
	}

	if giftProductRetrieved != nil {
		t.Error("was retrieved a product as a gift, but this product don't is a gift")
	}
}
