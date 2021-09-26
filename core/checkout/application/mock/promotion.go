package mock

import "gmarcial/eCommerce/core/checkout/domain/purchase"

//Promotion the mock of a promotion
type Promotion struct {
	ApplyMock func (cart *purchase.Cart) error
}

//Apply mock of Apply.
func (promotion *Promotion) Apply(cart *purchase.Cart) error {
	return promotion.ApplyMock(cart)
}

