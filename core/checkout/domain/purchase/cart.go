package purchase

import (
	"errors"
)

//Cart represents the entity of cart in the checkout context.
type Cart struct {
	TotalAmount    uint64
	TotalAmountNet uint64
	TotalDiscount uint64
	Products      map[uint32]*Product
}

//NewCart constructor to instantiate a cart.
func NewCart(products []*Product) (*Cart, error) {
	if products == nil && len(products) == 0 {
		return nil, errors.New("don't was informed the products")
	}

	addedProducts := make(map[uint32]*Product, 0)
	var totalAmount, totalAmountNet, totalDiscount uint64
	for _, product := range products{
		totalAmount += product.TotalAmount
		totalAmountNet += (product.TotalAmount - product.Discount)
		totalDiscount += product.Discount

		addedProducts[product.ID] = product
	}

	return &Cart{
		TotalAmount:    totalAmount,
		TotalAmountNet: totalAmountNet,
		TotalDiscount:  totalDiscount,
		Products:       addedProducts,
	}, nil
}

//Add more products to cart
func (cart *Cart) Add(product *Product) error {
	if product == nil {
		return errors.New("don't was informed the product")
	}

	ID := product.ID
	if containedProduct, exist := cart.Products[ID]; exist {
		if err := containedProduct.append(product); err != nil {
			return err
		}
	} else {
		cart.Products[product.ID] = product
		cart.TotalAmountNet -= product.Discount
		cart.TotalDiscount += product.Discount
	}

	cart.TotalAmount += product.TotalAmount
	cart.TotalAmountNet += product.TotalAmount

	return nil
}
