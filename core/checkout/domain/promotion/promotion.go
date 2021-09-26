package promotion

import "gmarcial/eCommerce/core/checkout/domain/purchase"

//Promotion represent the contract to apply promotion in a cart
type Promotion interface {
	Apply(cart *purchase.Cart) error
}
