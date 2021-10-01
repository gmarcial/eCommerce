package promotional

import (
	"errors"
	"gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/core/checkout/application/mock"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"testing"
	"time"
)

func TestUnitBlackFridayPromotion_ApplyPromotionWhenBlackFriday(t *testing.T) {
	//Arrange
	product := &model.Product{
		ID:          4,
		Title:       "Pasodk Qooosa",
		Description: "Pslqq dasas qwe asdasd",
		Amount:      10000,
		Discount:    5000,
		IsGift:      true,
	}
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(product, nil)
	cart, blackFridayPromotion := buildComponentsTestUnit(time.Now(), getGiftProductUseCase)

	const expectedQuantityCartProducts = 2
	const expectedProductAddedId = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while applying the black friday promotion: %v", err.Error())
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}

	if _, exist := cart.Products[expectedProductAddedId]; !exist {
		t.Error("the expected product added don't exist in products cart")
	}
}

func TestUnitBlackFridayPromotion_TryApplyPromotionWhenDontIsBlackFriday(t *testing.T) {
	//Arrange
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(nil, nil)
	backFridayDate := time.Now().AddDate(1, 0, 0)
	cart, blackFridayPromotion := buildComponentsTestUnit(backFridayDate, getGiftProductUseCase)

	const expectedQuantityCartProducts = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while applying the black friday promotion: %v", err.Error())
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}
}

func TestUnitBlackFridayPromotion_TryApplyPromotionFailToGetGiftProduct(t *testing.T) {
	//Arrange
	expectedError := errors.New("an error has occurred")
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(nil, expectedError)
	cart, blackFridayPromotion := buildComponentsTestUnit(time.Now(), getGiftProductUseCase)

	const expectedQuantityCartProducts = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != expectedError {
		t. Errorf("the error %v, is different of expected %v", err, expectedError)
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}
}

func TestUnitBlackFridayPromotion_ApplyPromotionWhenBlackFridayAndDontFoundGiftProducts(t *testing.T) {
	//Arrange
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(nil, nil)
	cart, blackFridayPromotion := buildComponentsTestUnit(time.Now(), getGiftProductUseCase)

	const expectedQuantityCartProducts = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while applying the black friday promotion: %v", err.Error())
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}
}

func buildGetGiftProductUseCaseTestUnit(product *model.Product, err error) *mock.GetGiftProductUseCase {
	return &mock.GetGiftProductUseCase{ExecuteMock: func() (*application.GiftProductObtained, error) {
		if product != nil {
			return &application.GiftProductObtained{Product: product}, err
		}

		return nil, err
	}}
}

func buildComponentsTestUnit(blackFridayDate time.Time, getGiftProductUseCase *mock.GetGiftProductUseCase) (*purchase.Cart, *BlackFridayPromotion) {
	product, _ := purchase.NewProduct(1, 1, 1000, 300, false)
	productsFirstTestCase := []*purchase.Product{product}
	cart, _ := purchase.NewCart(productsFirstTestCase)

	blackFridayPromotion := NewBlackFridayPromotion(blackFridayDate, getGiftProductUseCase)

	return cart, blackFridayPromotion
}

func TestEndToEndBlackFridayPromotion_ApplyPromotionWhenBlackFriday(t *testing.T) {
	//Arrange
	cart, blackFridayPromotion := buildComponentsTestEndToEnd(time.Now())
	const expectedQuantityCartProducts = 2
	const expectedProductAddedId = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while applying the black friday promotion: %v", err.Error())
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}

	if _, exist := cart.Products[expectedProductAddedId]; !exist {
		t.Error("the expected product added don't exist in products cart")
	}
}

func TestEndToEndBlackFridayPromotion_ApplyPromotionWhenDontIsBlackFriday(t *testing.T) {
	//Arrange
	blackFridayDate := time.Now().AddDate(1, 0, 0)
	cart, blackFridayPromotion := buildComponentsTestEndToEnd(blackFridayDate)
	const expectedQuantityCartProducts = 1

	//Action
	err := blackFridayPromotion.Apply(cart)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while applying the black friday promotion: %v", err.Error())
	}

	quantityCartProducts := len(cart.Products)
	if quantityCartProducts != expectedQuantityCartProducts {
		t. Errorf("the quantity of products %v, is different of expected %v", quantityCartProducts, expectedQuantityCartProducts)
	}
}

func buildComponentsTestEndToEnd(blackFridayDate time.Time) (*purchase.Cart, *BlackFridayPromotion) {
	product, _ := purchase.NewProduct(1, 1, 1000, 300, false)
	productsFirstTestCase := []*purchase.Product{product}
	cart, _ := purchase.NewCart(productsFirstTestCase)

	firstProduct, _ := products.NewProduct(4, "Muro Affas Kurias", "Teirs Masss Eas.", 5000, true)
	productsInMemory := map[uint32]*products.Product{
		firstProduct.ID:  firstProduct,
	}
	productRepository := memory.NewProductRepository(productsInMemory)

	getGiftProductUseCase := application.NewGetGiftProductUseCase(productRepository)

	blackFridayPromotion := NewBlackFridayPromotion(blackFridayDate, getGiftProductUseCase)

	return cart, blackFridayPromotion
}