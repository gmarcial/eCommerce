package promotional

import (
	"errors"
	applicationCatalog "gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/core/checkout/application/mock"
	"gmarcial/eCommerce/core/checkout/domain/promotion"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"testing"
	"time"
)

func TestUnitPromotionsApplierUseCase_ApplyActivePromotions(t *testing.T){
	//Arrange
	product, _ := purchase.NewProduct(1, 1, 1000, 300, false)

	promotionFirstTestCase := &mock.Promotion{ApplyMock: func(cart *purchase.Cart) error {
		return errors.New("an error has occurred")
	}}

	promotionSecondTestCase := &mock.Promotion{ApplyMock: func(cart *purchase.Cart) error {
		product, _ := purchase.NewProduct(2, 1, 1000, 300, true)
		_ = product.WrapGift()
		_ = cart.Add(product)
		return nil
	}}

	cartFirstTestCase, promotionsApplierUseCaseFirstTestCase := buildComponentsTestCase(product, promotionFirstTestCase)
	cartSecondTestCase, promotionsApplierUseCaseSecondTestCase := buildComponentsTestCase(product, promotionSecondTestCase)

	testCases := []struct {
		name string
		give struct{
			cart *purchase.Cart
			promotionsApplierUseCase *PromotionsApplierUseCase
		}
		want struct{
			expectedQuantityProducts int
		}
	}{
		{
			"Try apply a promotion, but fail",
			struct {
				cart                     *purchase.Cart
				promotionsApplierUseCase *PromotionsApplierUseCase
			}{cart: cartFirstTestCase, promotionsApplierUseCase: promotionsApplierUseCaseFirstTestCase},
			struct{ expectedQuantityProducts int }{expectedQuantityProducts: 1},
		},
		{
			"Apply a promotion, adding a new product to cart",
			struct {
				cart                     *purchase.Cart
				promotionsApplierUseCase *PromotionsApplierUseCase
			}{cart: cartSecondTestCase, promotionsApplierUseCase: promotionsApplierUseCaseSecondTestCase},
			struct{ expectedQuantityProducts int }{expectedQuantityProducts: 2},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			cart := testCase.give.cart
			useCase := testCase.give.promotionsApplierUseCase
			useCase.Apply(cart)

			//Assert
			expectedQuantityProducts := testCase.want.expectedQuantityProducts
			quantityCartProducts := len(cart.Products)
			if quantityCartProducts != expectedQuantityProducts {
				t.Errorf("The quantity of products in cart %v is different of expected %v", quantityCartProducts, expectedQuantityProducts)
			}
		})
	}
}

func buildComponentsTestCase(product *purchase.Product, activePromotion *mock.Promotion) (*purchase.Cart, *PromotionsApplierUseCase) {
	productsFirstTestCase := []*purchase.Product{product}
	cart, _ := purchase.NewCart(productsFirstTestCase)

	activePromotionsFirstTestCase := []promotion.Promotion{activePromotion}

	promotionsApplierUseCase := NewPromotionsApplierUseCase(activePromotionsFirstTestCase)

	return cart, promotionsApplierUseCase
}

func TestEndToEndPromotionsApplierUseCase_ApplyActivePromotions(t *testing.T){
	//Arrange
	promotionsApplierUseCase := buildPromotionsApplierUseCaseTestEndToEnd()

	product, _ := purchase.NewProduct(1, 1, 1000, 300, false)
	productsFirstTestCase := []*purchase.Product{product}
	cart, _ := purchase.NewCart(productsFirstTestCase)

	const expectedQuantityCartProducts = 2

	//Action
	promotionsApplierUseCase.Apply(cart)

	//Assert
	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t.Errorf("A gift product don't was add in cart.")
	}
}

func buildPromotionsApplierUseCaseTestEndToEnd() *PromotionsApplierUseCase {
	firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)
	secondProduct, _ := products.NewProduct(2, "Ergonomic Cotton Keyboard", "Iste est ratione excepturi repellendus adipisci qui.", 93811, false)
	thirdProduct, _ := products.NewProduct(3, "Gorgeous Cotton Chips", "Nulla rerum tempore rem.", 60356, false)
	fourthProduct, _ := products.NewProduct(4, "Muro Affas Kurias", "Teirs Masss Eas.", 5000, true)

	productsInMemory := map[uint32]*products.Product{
		firstProduct.ID:  firstProduct,
		secondProduct.ID: secondProduct,
		thirdProduct.ID:  thirdProduct,
		fourthProduct.ID: fourthProduct,
	}
	productRepository := memory.NewProductRepository(productsInMemory)

	getGiftProductUseCase := applicationCatalog.NewGetGiftProductUseCase(productRepository)
	blackFridayPromotion := NewBlackFridayPromotion(time.Now(), getGiftProductUseCase)

	activePromotions := []promotion.Promotion{blackFridayPromotion}
	return NewPromotionsApplierUseCase(activePromotions)
}
