package purchase

import "testing"

func TestUnitCreateAProductWithSuccess(t *testing.T) {
	//Arrange
	id := uint32(1)
	quantity := uint32(5)
	unitAmount := uint64(1000)
	discount := uint64(500)
	isGift := true

	//Action
	product, err := NewProduct(id, quantity, unitAmount, discount, isGift)

	//Assert
	if product == nil && err != nil {
		t.Error("the product don't was created, occurred an error")
	}

	if product.ID != id {
		t.Error("the value of id does not is of expected")
	}

	if product.Quantity != quantity {
		t.Error("the value of quantity does not is of expected")
	}

	if product.UnitAmount != unitAmount {
		t.Error("the value of unit amount does not is of expected")
	}

	totalAmount := (unitAmount * uint64(quantity))
	if product.TotalAmount != totalAmount {
		t.Error("the value of total amount does not is of expected")
	}

	if product.Discount != discount {
		t.Error("the value of discount does not is of expected")
	}

	if product.IsGift != isGift {
		t.Error("the value of isGift does not is of expected")
	}
}

func TestUnitTryCreateAProductDontInformingMandatoryValues(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			id          uint32
			quantity    uint32
			unitAmount  uint64
			discount    uint64
			isGift      bool
		}
		want string
	}{
		{
			"Don't is informed the id",
			struct {
				id          uint32
				quantity    uint32
				unitAmount  uint64
				discount    uint64
				isGift      bool
			}{
				uint32(0),
				uint32(1),
				uint64(1000),
				uint64(500),
				false,
			},
			"don't was informed the id of product",
		},
		{
			"Don't is informed the quantity",
			struct {
				id          uint32
				quantity    uint32
				unitAmount  uint64
				discount    uint64
				isGift      bool
			}{
				uint32(1),
				uint32(0),
				uint64(1000),
				uint64(500),
				false,
			},
			"don't was informed the quantity of product",
		},
		{
			"Don't is informed the unit amount",
			struct {
				id          uint32
				quantity    uint32
				unitAmount  uint64
				discount    uint64
				isGift      bool
			}{
				uint32(1),
				uint32(2),
				uint64(0),
				uint64(500),
				false,
			},
			"don't was informed the unit amount of product",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Action
			give := testCase.give
			product, err := NewProduct(give.id, give.quantity, give.unitAmount, give.discount, give.isGift)

			//Assert
			if product != nil && err == nil {
				t.Error("the product was created the same without informing mandatory values")
			}

			if err.Error() != testCase.want {
				t.Errorf("the occurred error is different of expected: %v", err.Error())
			}
		})
	}
}

func TestUnitProduct_TryAppendTwoDifferentProducts(t *testing.T) {
	//Arrange
	product, _ := NewProduct(1, 2, 1000, 500, false)
	appendedProduct, _ := NewProduct(2, 3, 750, 100, false)

	expectedQuantity := (product.Quantity + appendedProduct.Quantity)
	expectedTotalAmount := (product.TotalAmount + appendedProduct.TotalAmount)

	//Action
	err := product.append(appendedProduct)

	//Arrange
	if err == nil {
		t.Errorf("was possible to append two different products.")
	}

	if product.Quantity == expectedQuantity {
		t.Errorf("contains the quantity of %v product, expected to contain %v", product.Quantity, expectedQuantity)
	}

	if product.TotalAmount == expectedTotalAmount {
		t.Errorf("the total amount is %v, expected %v", product.TotalAmount, expectedTotalAmount)
	}
}

func TestUnitProduct_AppendMoreProducts(t *testing.T) {
	//Arrange
	product, _ := NewProduct(1, 2, 1000, 500, false)
	appendedProduct, _ := NewProduct(1, 3, 1000, 500, false)

	expectedQuantity := (product.Quantity + appendedProduct.Quantity)
	expectedTotalAmount := (product.TotalAmount + appendedProduct.TotalAmount)

	//Action
	err := product.append(appendedProduct)

	//Arrange
	if err != nil {
		t.Errorf("occurred an error, don't was possible to append: %v", err.Error())
	}

	if product.Quantity != expectedQuantity {
		t.Errorf("contains the quantity of %v product, expected to contain %v", product.Quantity, expectedQuantity)
	}

	if product.TotalAmount != expectedTotalAmount {
		t.Errorf("the total amount is %v, expected %v", product.TotalAmount, expectedTotalAmount)
	}
}

func TestUnitProduct_AppendAProductForGift(t *testing.T) {
	//Arrange
	product, _ := NewProduct(2, 2, 5000, 750, true)
	appendedGift, _ := NewProduct(2, 2, 5000, 750, true)
	_ = appendedGift.WrapGift()

	expectedQuantity := (product.Quantity + appendedGift.Quantity)
	expectedTotalAmount := product.TotalAmount

	//Action
	err := product.append(appendedGift)

	//Arrange
	if err != nil {
		t.Errorf("occurred an error, don't was possible to append: %v", err.Error())
	}

	if product.Quantity != expectedQuantity {
		t.Errorf("contains the quantity of %v product, expected to contain %v", product.Quantity, expectedQuantity)
	}

	if product.TotalAmount != expectedTotalAmount {
		t.Errorf("the total amount is %v, expected %v", product.TotalAmount, expectedTotalAmount)
	}
}

func TestUnitProduct_TryAProductForGiftWhichDontIsClassifiedAsAGift(t *testing.T) {
	//Arrange
	product, _ := NewProduct(2, 2, 5000, 750, false)

	//Action
	err := product.WrapGift()

	//Assert
	if err == nil {
		t.Errorf("was possible to make a product how gift.")
	}
}

func TestUnitProduct_AProductForGift(t *testing.T) {
	//Arrange
	product, _ := NewProduct(2, 2, 5000, 750, true)

	//Action
	err := product.WrapGift()

	//Assert
	if err != nil {
		t.Errorf("occurred an error, don't was possible to make a product how gift: %v", err.Error())
	}

	if product.isGiftWrapped != true {
		t.Error("the product is not like a gift")
	}

	if product.UnitAmount != 0 {
		t.Error("a gift product don't cost, unit amount should be zero")
	}

	if product.TotalAmount != 0 {
		t.Error("a gift product don't cost, total amount should be zero")
	}

	if product.Discount != 0 {
		t.Error("a gift product don't cost, discount should be zero")
	}
}
