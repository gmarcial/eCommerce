package application

import (
	"context"
	"errors"
	"gmarcial/eCommerce/core/catalog/domain/products"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
	"gmarcial/eCommerce/platform/infrastructure/test/mock"
	"google.golang.org/grpc"
	"testing"
)

const (
	address        = "localhost:50051"
	invalidAddress = "localhost:50050"
)

func buildPickProductsUseCaseTestUnit(discountResponse *client.GetDiscountResponse, err error) *PickProductsUseCase {
	productRepository := &mock.ProductRepository{
		GetProductsMock: func(IDS []uint32) ([]*products.Product, error) {
			firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)
			secondProduct, _ := products.NewProduct(2, "Ergonomic Cotton Keyboard", "Iste est ratione excepturi repellendus adipisci qui.", 93811, false)
			thirdProduct, _ := products.NewProduct(3, "Gorgeous Cotton Chips", "Nulla rerum tempore rem.", 60356, false)

			return []*products.Product{
				firstProduct,
				secondProduct,
				thirdProduct,
			}, nil
		},
	}

	discountServiceClient := &mock.DiscountClient{
		GetDiscountMock: func(ctx context.Context, in *client.GetDiscountRequest, opts ...grpc.CallOption) (*client.GetDiscountResponse, error) {
			return discountResponse, err
		}}

	return &PickProductsUseCase{
		productRepository:     productRepository,
		discountServiceClient: discountServiceClient,
	}
}

func TestUnitPickProductsUseCase_TryPickProductsWithoutInformTheProductsToBePicked(t *testing.T) {
	//Arrange
	useCase := NewPickProductsUseCase(&mock.ProductRepository{}, &mock.DiscountClient{})
	var pickProducts *PickProducts

	//Action
	pickedProducts, err := useCase.Execute(pickProducts)

	//Assert
	if pickedProducts != nil && err == nil {
		t.Error("it was possible to pick products without informing the products to be picked")
	}

	if err.Error() != "don't was informed the products to be picked" {
		t.Errorf("the occurred error is different of expected: %v", err.Error())
	}
}

func TestUnitPickProductsUseCase_TryPickProductsWithoutInformIDS(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			pickProductsUseCase *PickProductsUseCase
			pickProducts        *PickProducts
		}
	}{
		{
			"Informing the IDS as nil",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: NewPickProductsUseCase(&mock.ProductRepository{}, &mock.DiscountClient{}),
				pickProducts:        &PickProducts{IDS: nil},
			},
		},
		{
			"Informing the IDS as empty",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: NewPickProductsUseCase(&mock.ProductRepository{}, &mock.DiscountClient{}),
				pickProducts:        &PickProducts{IDS: []uint32{}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			useCase := testCase.give.pickProductsUseCase
			pickProducts := testCase.give.pickProducts
			pickedProducts, err := useCase.Execute(pickProducts)

			//Assert
			if pickedProducts != nil && err != nil {
				t.Error("it was possible to pick products without informing your ids")
			}
		})
	}
}

func TestUnitPickProductsUseCase_PickProducts(t *testing.T) {
	testCases := []struct {
		name string
		give struct {
			pickProductsUseCase *PickProductsUseCase
			pickProducts        *PickProducts
		}
		want struct {
			discounts map[uint32]uint64
			err       error
		}
	}{
		{
			"Apply 0% off discount",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: buildPickProductsUseCaseTestUnit(&client.GetDiscountResponse{Percentage: 0}, nil),
				pickProducts:        &PickProducts{IDS: []uint32{1, 2, 3}}},
			struct {
				discounts map[uint32]uint64
				err       error
			}{
				discounts: map[uint32]uint64{
					1: 0,
					2: 0,
					3: 0,
				},
				err: nil},
		},
		{
			"Apply 0.05% off discount",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: buildPickProductsUseCaseTestUnit(&client.GetDiscountResponse{Percentage: 0.05}, nil),
				pickProducts:        &PickProducts{IDS: []uint32{1, 2, 3}}},
			struct {
				discounts map[uint32]uint64
				err       error
			}{
				discounts: map[uint32]uint64{
					1: 7,
					2: 46,
					3: 30,
				},
				err: nil},
		},
		{
			"Try apply -0.05% off discount, simulating the problem to apply discount with negatives values",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: buildPickProductsUseCaseTestUnit(&client.GetDiscountResponse{Percentage: -0.05}, nil),
				pickProducts:        &PickProducts{IDS: []uint32{1, 2, 3}}},
			struct {
				discounts map[uint32]uint64
				err       error
			}{
				discounts: map[uint32]uint64{
					1: 0,
					2: 0,
					3: 0,
				},
				err: nil},
		},
		{
			"Try apply discount, simulating a problem to integrate with the discount service",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{
				pickProductsUseCase: buildPickProductsUseCaseTestUnit(nil, errors.New("error")),
				pickProducts:        &PickProducts{IDS: []uint32{1, 2, 3}}},
			struct {
				discounts map[uint32]uint64
				err       error
			}{
				discounts: map[uint32]uint64{
					1: 0,
					2: 0,
					3: 0,
				},
				err: nil},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			useCase := testCase.give.pickProductsUseCase
			pickProducts := testCase.give.pickProducts
			pickedProducts, err := useCase.Execute(pickProducts)

			errExpected := testCase.want.err
			if err != errExpected {
				t.Errorf("occurred a error different of expected: %v", err.Error())
			}

			//Assert
			discounts := testCase.want.discounts
			for _, pickedProduct := range pickedProducts.Products {
				discountExpected := discounts[pickedProduct.ID]
				discount := pickedProduct.Discount
				if discount != discountExpected {
					t.Errorf("was applied the %v amount of discount, different of expected %v", discount, discountExpected)
				}
			}
		})
	}
}

func buildPickProductsUseCaseTestEndToEnd(address string) *PickProductsUseCase {
	firstProduct, _ := products.NewProduct(1, "Ergonomic Wooden Pants", "Deleniti beatae porro.", 15157, false)
	secondProduct, _ := products.NewProduct(2, "Ergonomic Cotton Keyboard", "Iste est ratione excepturi repellendus adipisci qui.", 93811, false)
	thirdProduct, _ := products.NewProduct(3, "Gorgeous Cotton Chips", "Nulla rerum tempore rem.", 60356, false)

	productsInMemory := map[uint32]*products.Product{
		firstProduct.ID:  firstProduct,
		secondProduct.ID: secondProduct,
		thirdProduct.ID:  thirdProduct,
	}
	productRepository := memory.NewProductRepository(productsInMemory)

	options := []grpc.DialOption{grpc.WithInsecure()}
	channel, _ := grpc.Dial(address, options...)

	discountServiceClient := client.NewDiscountClient(channel)
	return NewPickProductsUseCase(productRepository, discountServiceClient)
}

func TestEndToEndPickProductsUseCase_PickProducts(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			pickProductsUseCase *PickProductsUseCase
			pickProducts        *PickProducts
		}
	}{
		{
			"With the discount service available",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{pickProductsUseCase: buildPickProductsUseCaseTestEndToEnd(address), pickProducts: &PickProducts{IDS: []uint32{1, 2, 3}}},
		},
		{
			"With the discount service unavailable",
			struct {
				pickProductsUseCase *PickProductsUseCase
				pickProducts        *PickProducts
			}{pickProductsUseCase: buildPickProductsUseCaseTestEndToEnd(invalidAddress), pickProducts: &PickProducts{IDS: []uint32{1, 2, 3}}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			useCase := testCase.give.pickProductsUseCase
			pickProducts := testCase.give.pickProducts
			pickedProducts, err := useCase.Execute(pickProducts)

			//Assert
			if err != nil {
				t.Errorf("an error has occurred to pick products: %v", err.Error())
			}

			if pickedProducts == nil {
				t.Error("the products picked returned nil")
			}

			quantityPickedProducts := len(pickedProducts.Products)
			if quantityPickedProducts == 0 && quantityPickedProducts == len(pickedProducts.Products) {
				t.Error("the products picked returned empty")
			}
		})
	}
}
