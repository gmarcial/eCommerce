package products

import (
	"errors"
	"fmt"
)

var (
	ErrDontInformedValue = "don't was informed the %v of product"
	ErrInvalidDiscountPercentageValue = "a discount percentage of value negative or zero don't is accept"
)

//Product represents the entity of product in the catalog context.
type Product struct {
	ID          uint32
	Title       string
	Description string
	Amount      uint64
	Discount    uint64
	IsGift      bool
}

//NewProduct constructor to instantiate a product.
func NewProduct(id uint32, title, description string, amount uint64, isGift bool) (*Product, error) {
	if id == 0 {
		message := fmt.Sprintf(ErrDontInformedValue, "id")
		return nil, errors.New(message)
	}

	if title == "" {
		message := fmt.Sprintf(ErrDontInformedValue, "title")
		return nil, errors.New(message)
	}

	if description == "" {
		message := fmt.Sprintf(ErrDontInformedValue, "description")
		return nil, errors.New(message)
	}

	if amount == 0 {
		message := fmt.Sprintf(ErrDontInformedValue, "amount")
		return nil, errors.New(message)
	}

	return &Product{
		ID:          id,
		Title:       title,
		Description: description,
		Amount:      amount,
		IsGift:      isGift,
	}, nil
}

//ApplyDiscount apply a discount percentage to the cost of the product, calculating and store the discount value.
func (product *Product) ApplyDiscount(discountPercentage float32) error {
	if discountPercentage < 0 {
		return errors.New(ErrInvalidDiscountPercentageValue)
	}

	amount := float32(product.Amount)
	discount := (amount * discountPercentage) / 100
	product.Discount = uint64(discount)

	return nil
}
