package purchase

import "testing"

func TestUnitTryNewCart(t *testing.T) {
	//Arrange
	var products []*Product

	//Action
	_, err := NewCart(products)

	//Assert
	if err == nil {
		t.Errorf("was possible create a cart without products: %v", err.Error())
	}
}

func TestUnitNewCart(t *testing.T) {
	//Arrange
	firstProduct, _ := NewProduct(1, 3, 500, 100, false)
	secondProduct, _ := NewProduct(2, 2, 250, 20, false)
	thirdProduct, _ := NewProduct(3, 1, 1000, 300, false)
	fourthProduct, _ := NewProduct(4, 5, 5000, 750, false)
	fifthProduct, _ := NewProduct(5, 2, 45, 0, false)
	sixthProduct, _ := NewProduct(6, 2, 88, 15, false)
	seventhProduct, _ := NewProduct(7, 3, 267, 25, false)
	eighthProduct, _ := NewProduct(8, 1, 678, 50, false)

	products := []*Product{
		firstProduct,
		secondProduct,
		thirdProduct,
		fourthProduct,
		fifthProduct,
		sixthProduct,
		seventhProduct,
		eighthProduct,
	}

	expectedTotalAmount := uint64(29745)
	expectedTotalDiscount := uint64(1260)
	expectedTotalAmountNet := (expectedTotalAmount - expectedTotalDiscount)
	expectedTotalQuantityProducts := len(products)

	//Action
	cart, err := NewCart(products)

	//Assert
	if err != nil {
		t.Errorf("the cart don't was created, occurred an error: %v", err.Error())
	}

	if cart.TotalAmount != expectedTotalAmount {
		t.Errorf("actual value of the total amount is %v, expected to %v", cart.TotalAmount, expectedTotalAmount)
	}

	if cart.TotalAmountNet != expectedTotalAmountNet {
		t.Errorf("actual value of the total amount net is %v, expected to %v", cart.TotalAmountNet, expectedTotalAmountNet)
	}

	if cart.TotalDiscount != expectedTotalDiscount {
		t.Errorf("actual value of the total of discount is %v, expected to %v", cart.TotalDiscount, expectedTotalDiscount)
	}

	quantityProducts := len(cart.Products)
	if quantityProducts != expectedTotalQuantityProducts {
		t.Errorf("actual quantity of products is %v, expected to %v", quantityProducts, expectedTotalQuantityProducts)
	}

	for _, product := range products {
		addedProduct, exist := cart.Products[product.ID]
		if !exist {
			t.Errorf("the product of ID %v don't was found", product.ID)
		}

		if addedProduct == nil {
			t.Errorf("the product of ID %v was found, but is nil", product.ID)
		}
	}
}

func BuildCart() *Cart {
	product, _ := NewProduct(1, 3, 500, 100, false)
	products := []*Product{product}
	cart, err := NewCart(products)
	if err != nil {
		panic(err.Error())
	}

	return cart
}

func BuildProduct(id uint32, quantity uint32, unitAmount, discount uint64, isGift bool) *Product {
	product, err := NewProduct(id, quantity, unitAmount, discount, isGift)
	if err != nil {
		panic(err.Error())
	}

	return product
}

func TestUnitCart_TryAddMoreProducts(t *testing.T) {
	//Arrange
	cart := BuildCart()

	//Action
	err := cart.Add(nil)

	//Assert
	if err == nil {
		t.Errorf("was possible add a product without product: %v", err.Error())
	}
}

func TestUnitCart_AddMoreProducts(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			cart            *Cart
			appendedProduct *Product
		}
		want struct {
			expectedQuantityProducts int
			expectedTotalAmountNet   uint64
			expectedTotalAmount      uint64
			expectedTotalDiscount    uint64
		}
	}{
		{
			"A new product in cart",
			struct {
				cart            *Cart
				appendedProduct *Product
			}{cart: BuildCart(), appendedProduct: BuildProduct(4, 5, 5000, 750, false)},
			struct {
				expectedQuantityProducts int
				expectedTotalAmountNet   uint64
				expectedTotalAmount      uint64
				expectedTotalDiscount    uint64
			}{
				expectedQuantityProducts: 2,
				expectedTotalAmountNet:   25650,
				expectedTotalAmount:      26500,
				expectedTotalDiscount:    850},
		},
		{
			"More of same product",
			struct {
				cart            *Cart
				appendedProduct *Product
			}{cart: BuildCart(), appendedProduct: BuildProduct(1, 3, 500, 100, false)},
			struct {
				expectedQuantityProducts int
				expectedTotalAmountNet   uint64
				expectedTotalAmount      uint64
				expectedTotalDiscount    uint64
			}{
				expectedQuantityProducts: 1,
				expectedTotalAmountNet:   2900,
				expectedTotalAmount:      3000,
				expectedTotalDiscount:    100},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			cart := testCase.give.cart
			product := testCase.give.appendedProduct
			err := cart.Add(product)

			//Assert
			if err != nil {
				t.Errorf("the product don't was added in cart, occurred an error: %v", err.Error())
			}

			want := testCase.want
			if cart.TotalAmount != want.expectedTotalAmount {
				t.Errorf("actual value of the total amount is %v, expected to %v", cart.TotalAmount, want.expectedTotalAmount)
			}

			if cart.TotalAmountNet != want.expectedTotalAmountNet {
				t.Errorf("actual value of the total amount net is %v, expected to %v", cart.TotalAmountNet, want.expectedTotalAmountNet)
			}

			if cart.TotalDiscount != want.expectedTotalDiscount {
				t.Errorf("actual value of the total of discount is %v, expected to %v", cart.TotalDiscount, want.expectedTotalDiscount)
			}

			quantityProducts := len(cart.Products)
			if quantityProducts != want.expectedQuantityProducts {
				t.Errorf("the quantity actual of products is %v, expected to %v", quantityProducts, want.expectedQuantityProducts)
			}
		})
	}
}
