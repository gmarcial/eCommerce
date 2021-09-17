package products

import "testing"

func TestUnitCreateAProductWithSuccess(t *testing.T) {
	//Arrange
	id := uint32(1)
	title := "Ergonomic Wooden Pants"
	description := "Deleniti beatae porro."
	amount := uint64(15157)
	isGift := true

	//Action
	product, err := NewProduct(id, title, description, amount, isGift)

	//Assert
	if product == nil && err != nil {
		t.Error("the product don't was created, occurred an error")
	}

	if product.ID != id {
		t.Error("the value of id does not is of expected")
	}

	if product.Title != title {
		t.Error("the value of title does not is of expected")
	}

	if product.Description != description {
		t.Error("the value of description does not is of expected")
	}

	if product.Amount != amount {
		t.Error("the value of amount does not is of expected")
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
			title       string
			description string
			amount      uint64
			isGift      bool
		}
		want string
	}{
		{
			"Don't is informed the id",
			struct {
				id          uint32
				title       string
				description string
				amount      uint64
				isGift      bool
			}{
				uint32(0),
				"Ergonomic Wooden Pants",
				"Deleniti beatae porro.",
				uint64(15157),
				false,
			},
			"don't was informed the id of product",
		},
		{
			"Don't is informed the title",
			struct {
				id          uint32
				title       string
				description string
				amount      uint64
				isGift      bool
			}{
				uint32(1),
				"",
				"Deleniti beatae porro.",
				uint64(15157),
				false,
			},
			"don't was informed the title of product",
		},
		{
			"Don't is informed the description",
			struct {
				id          uint32
				title       string
				description string
				amount      uint64
				isGift      bool
			}{
				uint32(1),
				"Ergonomic Wooden Pants",
				"",
				uint64(15157),
				false,
			},
			"don't was informed the description of product",
		},
		{
			"Don't is informed the amount",
			struct {
				id          uint32
				title       string
				description string
				amount      uint64
				isGift      bool
			}{
				uint32(1),
				"Ergonomic Wooden Pants",
				"Deleniti beatae porro.",
				uint64(0),
				false,
			},
			"don't was informed the amount of product",
		},
	}
	//Action
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			give := testCase.give
			product, err := NewProduct(give.id, give.title, give.description, give.amount, give.isGift)

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

func TestUnitProduct_ApplyDiscountPercentage(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			product            *Product
			discountPercentage float32
		}
		want uint64
	}{
		{
			"Apply 7% off discount in amount 15157",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(1),
					Title:       "Ergonomic Wooden Pants",
					Description: "Deleniti beatae porro.",
					Amount:      uint64(15157),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(7),
			},
			uint64(1060),
		},
		{
			"Apply 10.5% off discount in amount 93811",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(2),
					Title:       "Ergonomic Cotton Keyboard",
					Description: "Iste est ratione excepturi repellendus adipisci qui.",
					Amount:      uint64(93811),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(10.5),
			},
			uint64(9850),
		},
		{
			"Apply 0.8% off discount in amount 60356",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(3),
					Title:       "Gorgeous Cotton Chips",
					Description: "Nulla rerum tempore rem.",
					Amount:      uint64(60356),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(0.8),
			},
			uint64(482),
		},
		{
			"Apply 45.17% off discount in amount 56230",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(4),
					Title:       "Fantastic Frozen Chair",
					Description: "Et neque debitis omnis quam enim cupiditate.",
					Amount:      uint64(56230),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(45.17),
			},
			uint64(25399),
		},
		{
			"Apply 23.72% off discount in amount 42647",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(5),
					Title:       "Incredible Concrete Soap",
					Description: "Dolorum nobis temporibus aut dolorem quod qui corrupti.",
					Amount:      uint64(42647),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(23.72),
			},
			uint64(10115),
		},
		{
			"Apply 0.02% off discount in amount 900",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(6),
					Title:       "Handcrafted Steel Towels",
					Description: "Nam ea sed animi neque qui non quis iste.",
					Amount:      uint64(900),
					Discount:    uint64(0),
					IsGift:      true,
				},
				float32(0.02),
			},
			uint64(0),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			product := testCase.give.product
			discountPercentage := testCase.give.discountPercentage

			//Action
			err := product.ApplyDiscount(discountPercentage)

			//Assert
			if err != nil {
				t.Errorf("occurred an error, don't was possible to apply discount: %v", err.Error())
			}

			discount := product.Discount
			expectedDiscount := testCase.want
			if discount != expectedDiscount {
				t.Errorf("the value of discount calculated is different from expected, discount calculated is %v and expected %v", discount, expectedDiscount)
			}
		})

	}
}

func TestUnitProduct_TryApplyDiscountWithInvalidPercentage(t *testing.T) {
	//Arrange
	testCases := []struct {
		name string
		give struct {
			product            *Product
			discountPercentage float32
		}
		want string
	}{
		{
			"Apply 0% off discount in amount 15157",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(1),
					Title:       "Ergonomic Wooden Pants",
					Description: "Deleniti beatae porro.",
					Amount:      uint64(15157),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(0),
			},
			ErrInvalidDiscountPercentageValue,
		},
		{
			"Apply -10.5% off discount in amount 93811",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(2),
					Title:       "Ergonomic Cotton Keyboard",
					Description: "Iste est ratione excepturi repellendus adipisci qui.",
					Amount:      uint64(93811),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(-10.5),
			},
			ErrInvalidDiscountPercentageValue,
		},
		{
			"Apply -0.8% off discount in amount 60356",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(3),
					Title:       "Gorgeous Cotton Chips",
					Description: "Nulla rerum tempore rem.",
					Amount:      uint64(60356),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(-0.8),
			},
			ErrInvalidDiscountPercentageValue,
		},
		{
			"Apply -45.17% off discount in amount 56230",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(4),
					Title:       "Fantastic Frozen Chair",
					Description: "Et neque debitis omnis quam enim cupiditate.",
					Amount:      uint64(56230),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(-45.17),
			},
			ErrInvalidDiscountPercentageValue,
		},
		{
			"Apply -23.72% off discount in amount 42647",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(5),
					Title:       "Incredible Concrete Soap",
					Description: "Dolorum nobis temporibus aut dolorem quod qui corrupti.",
					Amount:      uint64(42647),
					Discount:    uint64(0),
					IsGift:      false,
				},
				float32(-23.72),
			},
			ErrInvalidDiscountPercentageValue,
		},
		{
			"Apply -0.02% off discount in amount 900",
			struct {
				product            *Product
				discountPercentage float32
			}{
				&Product{
					ID:          uint32(6),
					Title:       "Handcrafted Steel Towels",
					Description: "Nam ea sed animi neque qui non quis iste.",
					Amount:      uint64(900),
					Discount:    uint64(0),
					IsGift:      true,
				},
				float32(-0.02),
			},
			ErrInvalidDiscountPercentageValue,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			product := testCase.give.product
			discountPercentage := testCase.give.discountPercentage

			expectedDiscount := product.Discount

			//Action
			err := product.ApplyDiscount(discountPercentage)

			//Assert
			discount := product.Discount
			if discount != expectedDiscount && err == nil {
				t.Errorf("was possible to apply discount with the invalid value of a discount percentage: %v", discountPercentage)
			}
		})
	}
}
