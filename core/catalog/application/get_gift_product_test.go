package application

import (
	"gmarcial/eCommerce/core/catalog/application/mock"
	"gmarcial/eCommerce/core/catalog/application/model"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"testing"
)

func buildGetGiftProductUseCaseTestUnit(giftProduct *products.Product) *GetGiftProductUseCase {
	productRepository := &mock.ProductRepository{
		GetGiftProductMock: func() (*products.Product, error) {
			return giftProduct, nil
		},
	}

	return &GetGiftProductUseCase{productRepository: productRepository}
}

func TestUnitGetGiftProductUseCase_TryObtainGiftProduct(t *testing.T) {
	//Arrange
	var giftProduct *products.Product
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(giftProduct)

	//Action
	giftProductObtained, err := getGiftProductUseCase.Execute()

	//Assert
	if err != nil {
		t.Error("The error returned is different of expected")
	}

	if giftProductObtained != nil {
		t.Error("The gift product obtained returned is different of expected")
	}
}

func TestUnitGetGiftProductUseCase_ObtainGiftProduct(t *testing.T) {
	//Arrange
	giftProduct, _ := products.NewProduct(1, "Gift Product", "Descrição Alada", 100, true)
	getGiftProductUseCase := buildGetGiftProductUseCaseTestUnit(giftProduct)

	expectedGiftProductObtained := &GiftProductObtained{Product: &model.Product{
		ID:          giftProduct.ID,
		Title:       giftProduct.Title,
		Description: giftProduct.Description,
		Amount:      giftProduct.Amount,
		Discount:    giftProduct.Discount,
		IsGift:      giftProduct.IsGift,
	}}

	//Action
	giftProductObtained, err := getGiftProductUseCase.Execute()

	//Assert
	if err != nil {
		t.Error("occurred an error to obtain gift product")
	}

	if giftProductObtained == nil {
		t.Error("don't was obtained a gift product")
	}

	if *giftProductObtained.Product != *expectedGiftProductObtained.Product {
		t.Error("The gift product obtained returned is different of expected")
	}
}

func buildGetGiftProductUseCaseTestEndToEnd(giftProduct *products.Product) *GetGiftProductUseCase {
	firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)
	secondProduct, _ := products.NewProduct(2, "Ergonomic Cotton Keyboard", "Iste est ratione excepturi repellendus adipisci qui.", 93811, false)
	thirdProduct, _ := products.NewProduct(3, "Gorgeous Cotton Chips", "Nulla rerum tempore rem.", 60356, false)

	productsInMemory := map[uint32]*products.Product{
		firstProduct.ID:  firstProduct,
		secondProduct.ID: secondProduct,
		thirdProduct.ID:  thirdProduct,
	}

	if giftProduct != nil {
		productsInMemory[giftProduct.ID] = giftProduct
	}

	productRepository := memory.NewProductRepository(productsInMemory)

	return &GetGiftProductUseCase{productRepository: productRepository}
}

func TestEndToEndGetGiftProductUseCase_TryObtainGiftProduct(t *testing.T) {
	//Arrange
	var giftProduct *products.Product
	getGiftProductUseCase := buildGetGiftProductUseCaseTestEndToEnd(giftProduct)

	//Action
	giftProductObtained, err := getGiftProductUseCase.Execute()

	//Assert
	if err != nil {
		t.Error("The error returned is different of expected")
	}

	if giftProductObtained != nil {
		t.Error("The gift product obtained returned is different of expected")
	}
}

func TestEndToEndGetGiftProductUseCase_ObtainGiftProduct(t *testing.T) {
	//Arrange
	giftProduct, _ := products.NewProduct(1, "Gift Product", "Descrição Alada", 100, true)
	getGiftProductUseCase := buildGetGiftProductUseCaseTestEndToEnd(giftProduct)

	expectedGiftProductObtained := &GiftProductObtained{Product: &model.Product{
		ID:          giftProduct.ID,
		Title:       giftProduct.Title,
		Description: giftProduct.Description,
		Amount:      giftProduct.Amount,
		Discount:    giftProduct.Discount,
		IsGift:      giftProduct.IsGift,
	}}

	//Action
	giftProductObtained, err := getGiftProductUseCase.Execute()

	//Assert
	if err != nil {
		t.Error("occurred an error to obtain gift product")
	}

	if giftProductObtained == nil {
		t.Error("don't was obtained a gift product")
	}

	if *giftProductObtained.Product != *expectedGiftProductObtained.Product {
		t.Error("The gift product obtained returned is different of expected")
	}
}
