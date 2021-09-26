package promotional

import (
	"gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"time"
)

const (
	quantity = 1
)

//BlackFridayPromotion represent the promotion of black friday
type BlackFridayPromotion struct {
	blackFridayDate time.Time
	getGiftProductUseCase application.IGetGiftProductUseCase
}

//NewBlackFridayPromotion constructor to instantiate the use case to apply the black friday promotion.
func NewBlackFridayPromotion(blackFridayDate time.Time, getGiftProductUseCase application.IGetGiftProductUseCase) *BlackFridayPromotion{
	return &BlackFridayPromotion{blackFridayDate: blackFridayDate, getGiftProductUseCase: getGiftProductUseCase}
}

//Apply the promotion referent to black friday
func (promotion *BlackFridayPromotion) Apply(cart *purchase.Cart) error {
	dateNow := time.Now()
	blackFridayDate := promotion.blackFridayDate
	if dateNow.Day() != blackFridayDate.Day() ||
		dateNow.Month() != blackFridayDate.Month() ||
		dateNow.Year() != blackFridayDate.Year() {
		return nil
	}


	giftProductObtained, err := promotion.getGiftProductUseCase.Execute()
	if err != nil {
		return err
	}

	if giftProductObtained == nil {
		return nil
	}

	giftProduct := giftProductObtained.Product
	product, err := purchase.NewProduct(giftProduct.ID, quantity, giftProduct.Amount, giftProduct.Discount, giftProduct.IsGift)
	if err != nil {
		return err
	}

	err = product.WrapGift()
	if err != nil {
		return err
	}

	err = cart.Add(product)
	if err != nil {
		return err
	}

	return nil
}
