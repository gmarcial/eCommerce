package promotional

import (
	"gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"gmarcial/eCommerce/platform/infrastructure/log"
	"time"
)

const (
	quantity = 1
)

//BlackFridayPromotion represent the promotion of black friday
type BlackFridayPromotion struct {
	logger                log.Logger
	blackFridayDay        time.Time
	getGiftProductUseCase application.IGetGiftProductUseCase
}

//NewBlackFridayPromotion constructor to instantiate the use case to apply the black friday promotion.
func NewBlackFridayPromotion(logger log.Logger, blackFridayDate time.Time, getGiftProductUseCase application.IGetGiftProductUseCase) *BlackFridayPromotion {
	return &BlackFridayPromotion{
		logger: logger,
		blackFridayDay: blackFridayDate,
		getGiftProductUseCase: getGiftProductUseCase,
	}
}

//Apply the promotion referent to black friday
func (promotion *BlackFridayPromotion) Apply(cart *purchase.Cart) error {
	logger := promotion.logger
	logger.Infow("try apply black friday promotion.")

	dateNow := time.Now()
	blackFridayDate := promotion.blackFridayDay
	if dateNow.Day() != blackFridayDate.Day() ||
		dateNow.Month() != blackFridayDate.Month() ||
		dateNow.Year() != blackFridayDate.Year() {

		logger.Infow("promotion not applied, today don't is black Friday day.",
			"today", dateNow,
			"black friday day", blackFridayDate)

		return nil
	}

	giftProductObtained, err := promotion.getGiftProductUseCase.Execute()
	if err != nil {

		logger.Errorw("an error occurred to try obtain a gift product.",
			"error", err)

		return err
	}

	if giftProductObtained == nil {
		logger.Infow("none gift product was obtained.")
		return nil
	}

	giftProduct := giftProductObtained.Product
	product, err := purchase.NewProduct(giftProduct.ID, quantity, giftProduct.Amount, giftProduct.Discount, giftProduct.IsGift)
	if err != nil {
		logger.Errorw("an error occurred to translate the GiftProductObtained to Product.",
			"error", err)
		return err
	}

	logger.Infow("the product to gift was obtained.",
		"product", giftProduct)

	err = product.WrapGift()
	if err != nil {
		logger.Errorw("an error occurred to wrap the product for gift.",
			"error", err)
		return err
	}

	logger.Infow("the product was wrapped to gift.",
		"gift product", product)

	err = cart.Add(product)
	if err != nil {
		logger.Errorw("an error occurred to add the gift to cart.",
			"error", err)
		return err
	}

	logger.Infow("the gift product was add to cart.",
		"cart", cart)

	logger.Infow("black friday promotion applied.")
	return nil
}
