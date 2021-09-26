package application

import (
	"context"
	applicationCatalog "gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/catalog/application/mock"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/core/checkout/application/model"
	"gmarcial/eCommerce/core/checkout/application/promotional"
	"gmarcial/eCommerce/core/checkout/domain/promotion"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestUnitMakeCartUseCase_ConstructTheShoppingCartOutOfBlackFriday(t *testing.T) {
	//Arrange
	blackFridayDate := time.Now().AddDate(1, 0, 0)
	makeCartUseCase := buildMakeCartUseCaseTestEndToEnd(blackFridayDate)

	modelSelectedProducts := []*model.SelectedProduct{
		{
			ID:       1,
			Quantity: 1,
		},
		{
			ID:       3,
			Quantity: 2,
		},
		{
			ID:       4,
			Quantity: 1,
		},
	}
	selectedProducts := &SelectedProducts{Products: modelSelectedProducts}

	expectedTotalAmount := uint64(140869)
	expectedTotalAmountNet := uint64(136845)
	expectedTotalDiscount := uint64(4024)

	expectedCartProducts := []*model.Product{{
		ID:          1,
		Quantity:    1,
		UnitAmount:  15157,
		TotalAmount: 15157,
		Discount:    757,
		IsGift:      false,
	}, {
		ID:          3,
		Quantity:    2,
		UnitAmount:  60356,
		TotalAmount: 120712,
		Discount:    3017,
		IsGift:      false,
	}, {
		ID:          4,
		Quantity:    1,
		UnitAmount:  5000,
		TotalAmount: 5000,
		Discount:    250,
		IsGift:      true,
	}}
	expectedCartProductsQuantity := len(expectedCartProducts)

	//Action
	cart, err := makeCartUseCase.Execute(selectedProducts)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while building the shopping cart: %v", err.Error())
	}

	if cart == nil {
		t.Error("an error occurred while building the shopping cart, is nil")
	}

	totalAmount := cart.TotalAmount
	if totalAmount != expectedTotalAmount {
		t.Errorf("The total amount %v is different of expected %v", totalAmount, expectedTotalAmount)
	}

	totalAmountNet := cart.TotalAmountNet
	if cart.TotalAmountNet != expectedTotalAmountNet {
		t.Errorf("The total amount net %v is different of expected %v", totalAmountNet, expectedTotalAmountNet)
	}

	totalDiscount := cart.TotalDiscount
	if cart.TotalDiscount != expectedTotalDiscount {
		t.Errorf("The total discount %v is different of expected %v", totalDiscount, expectedTotalDiscount)
	}

	cartProducts := cart.Products
	cartProductsQuantity := len(cartProducts)
	if cartProductsQuantity != expectedCartProductsQuantity {
		t.Errorf("The quantity of products in cart %v is different of expected %v", cartProductsQuantity, expectedCartProductsQuantity)
	}

	for i := 0; i < cartProductsQuantity; i++ {
		if *cartProducts[i] != *expectedCartProducts[i] {
			t.Error("The product is different of expected")
		}
	}
}

func TestUnitMakeCartUseCase_ConstructTheShoppingCartInBlackFriday(t *testing.T) {
	//Arrange
	blackFridayDate := time.Now()
	makeCartUseCase := buildMakeCartUseCaseTestEndToEnd(blackFridayDate)

	modelSelectedProducts := []*model.SelectedProduct{
		{
			ID:       1,
			Quantity: 1,
		},
		{
			ID:       3,
			Quantity: 2,
		},
	}
	selectedProducts := &SelectedProducts{Products: modelSelectedProducts}

	expectedTotalAmount := uint64(135869)
	expectedTotalAmountNet := uint64(132095)
	expectedTotalDiscount := uint64(3774)

	expectedCartProducts := []*model.Product{{
		ID:          1,
		Quantity:    1,
		UnitAmount:  15157,
		TotalAmount: 15157,
		Discount:    757,
		IsGift:      false,
	}, {
		ID:          3,
		Quantity:    2,
		UnitAmount:  60356,
		TotalAmount: 120712,
		Discount:    3017,
		IsGift:      false,
	}, {
		ID:          4,
		Quantity:    1,
		UnitAmount:  0,
		TotalAmount: 0,
		Discount:    0,
		IsGift:      true,
	}}
	expectedCartProductsQuantity := len(expectedCartProducts)

	//Action
	cart, err := makeCartUseCase.Execute(selectedProducts)

	//Assert
	if err != nil {
		t.Errorf("an error occurred while building the shopping cart: %v", err.Error())
	}

	if cart == nil {
		t.Error("an error occurred while building the shopping cart, is nil")
	}

	totalAmount := cart.TotalAmount
	if totalAmount != expectedTotalAmount {
		t.Errorf("The total amount %v is different of expected %v", totalAmount, expectedTotalAmount)
	}

	totalAmountNet := cart.TotalAmountNet
	if cart.TotalAmountNet != expectedTotalAmountNet {
		t.Errorf("The total amount net %v is different of expected %v", totalAmountNet, expectedTotalAmountNet)
	}

	totalDiscount := cart.TotalDiscount
	if cart.TotalDiscount != expectedTotalDiscount {
		t.Errorf("The total discount %v is different of expected %v", totalDiscount, expectedTotalDiscount)
	}

	cartProducts := cart.Products
	cartProductsQuantity := len(cartProducts)
	if cartProductsQuantity != expectedCartProductsQuantity {
		t.Errorf("The quantity of products in cart %v is different of expected %v", cartProductsQuantity, expectedCartProductsQuantity)
	}

	for i := 0; i < cartProductsQuantity; i++ {
		if *cartProducts[i] != *expectedCartProducts[i] {
			t.Error("The product is different of expected")
		}
	}
}

func buildProductRepositoryTestEndToEnd() *memory.ProductRepository {
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
	return memory.NewProductRepository(productsInMemory)
}

func buildPickProductsUseCaseTestEndToEnd(productRepository products.ProductRepository) *applicationCatalog.PickProductsUseCase {
	discountServiceClient := &mock.DiscountClient{GetDiscountMock: func(ctx context.Context, in *client.GetDiscountRequest, opts ...grpc.CallOption) (*client.GetDiscountResponse, error) {
		return &client.GetDiscountResponse{Percentage: 5}, nil
	}}
	return applicationCatalog.NewPickProductsUseCase(productRepository, discountServiceClient)
}

func buildPromotionsApplierUseCaseTestEndToEnd(blackFridayDate time.Time, productRepository products.ProductRepository) *promotional.PromotionsApplierUseCase {
	getGiftProductUseCase := applicationCatalog.NewGetGiftProductUseCase(productRepository)
	blackFridayPromotion := promotional.NewBlackFridayPromotion(blackFridayDate, getGiftProductUseCase)

	activePromotions := []promotion.Promotion{blackFridayPromotion}
	return promotional.NewPromotionsApplierUseCase(activePromotions)
}

func buildMakeCartUseCaseTestEndToEnd(blackFridayDate time.Time) *MakeCartUseCase {
	productRepository := buildProductRepositoryTestEndToEnd()
	pickProductsUseCase := buildPickProductsUseCaseTestEndToEnd(productRepository)
	promotionsApplierUseCase := buildPromotionsApplierUseCaseTestEndToEnd(blackFridayDate, productRepository)

	return NewMakeCartUseCase(pickProductsUseCase, promotionsApplierUseCase)
}
