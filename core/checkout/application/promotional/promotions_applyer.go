package promotional

import (
	"gmarcial/eCommerce/core/checkout/domain/promotion"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
	"gmarcial/eCommerce/platform/infrastructure/log"
)

//IPromotionsApplierUseCase the interface to api with behavior to apply all active promotions.
type IPromotionsApplierUseCase interface {
	Apply(cart *purchase.Cart)
}

//PromotionsApplierUseCase api with behavior to apply all active promotions.
type PromotionsApplierUseCase struct {
	logger           log.Logger
	activePromotions []promotion.Promotion
}

//NewPromotionsApplierUseCase constructor to instantiate the use case to apply the promotions.
func NewPromotionsApplierUseCase(logger log.Logger, activePromotions []promotion.Promotion) *PromotionsApplierUseCase {
	return &PromotionsApplierUseCase{
		logger:           logger,
		activePromotions: activePromotions,
	}
}

//Apply all active promotions
func (useCase *PromotionsApplierUseCase) Apply(cart *purchase.Cart) {
	logger := useCase.logger

	activePromotions := useCase.activePromotions

	logger.Infow("started to apply active promotions to cart.",
		"cart", cart,
		"quantity of active promotions", len(activePromotions))

	for _, activePromotion := range activePromotions {
		err := activePromotion.Apply(cart)
		if err != nil {
			logger.Errorw("an error occurred to apply a active promotion.",
				"active Promotion", activePromotion,
				"error", err.Error())
		}
	}

	logger.Infow("finished to apply active promotions to cart",
		"cart", cart)
}
