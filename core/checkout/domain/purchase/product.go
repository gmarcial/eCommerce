package purchase

import (
	"errors"
	"fmt"
)

var (
	ErrDontInformedValue = "don't was informed the %v of product"
)

//Product represents the entity of product in the checkout context.
type Product struct {
	ID            uint32
	Quantity      uint32
	UnitAmount    uint64
	TotalAmount   uint64
	Discount      uint64
	IsGift        bool
	isGiftWrapped bool
}

//NewProduct constructor to instantiate a product.
func NewProduct(id uint32, quantity uint32, unitAmount, discount uint64, isGift bool) (*Product, error) {
	if id == 0 {
		message := fmt.Sprintf(ErrDontInformedValue, "id")
		return nil, errors.New(message)
	}

	if quantity == 0 {
		message := fmt.Sprintf(ErrDontInformedValue, "quantity")
		return nil, errors.New(message)
	}

	if unitAmount == 0 {
		message := fmt.Sprintf(ErrDontInformedValue, "unit amount")
		return nil, errors.New(message)
	}

	totalAmount := (unitAmount * uint64(quantity))

	return &Product{
		ID:          id,
		Quantity:    quantity,
		UnitAmount:  unitAmount,
		TotalAmount: totalAmount,
		Discount:    discount,
		IsGift:      isGift,
	}, nil
}

//append add more products of same.
func (product *Product) append(appendedProduct *Product) error {
	if product.ID != appendedProduct.ID {
		return errors.New("the product to be appended doesn't is the same of target")
	}

	product.Quantity += appendedProduct.Quantity
	if !appendedProduct.isGiftWrapped {
		product.TotalAmount += appendedProduct.TotalAmount
	}

	return nil
}

//WrapGift alter the state of product to gift, case it is classified how a gift.
func (product *Product) WrapGift() error {
	if !product.IsGift {
		return errors.New("this product cannot be used how a gift")
	}

	product.isGiftWrapped = true
	product.UnitAmount = 0
	product.TotalAmount = 0
	product.Discount = 0

	return nil
}
