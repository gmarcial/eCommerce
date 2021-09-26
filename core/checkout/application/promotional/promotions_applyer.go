package promotional

import (
	"gmarcial/eCommerce/core/checkout/domain/promotion"
	"gmarcial/eCommerce/core/checkout/domain/purchase"
)

//IPromotionsApplierUseCase the interface to api with behavior to apply all active promotions.
type IPromotionsApplierUseCase interface {
	Apply(cart *purchase.Cart)
}

//PromotionsApplierUseCase api with behavior to apply all active promotions.
type PromotionsApplierUseCase struct {
	activePromotions []promotion.Promotion
}

//NewPromotionsApplierUseCase constructor to instantiate the use case to apply the promotions.
func NewPromotionsApplierUseCase(activePromotions []promotion.Promotion) *PromotionsApplierUseCase {
	return &PromotionsApplierUseCase{activePromotions: activePromotions}
}

//Apply all active promotions
func (useCase *PromotionsApplierUseCase) Apply(cart *purchase.Cart) {
	for _, activePromotion := range useCase.activePromotions {
		err := activePromotion.Apply(cart)
		if err != nil {
			//TODO: Loggin error...
		}
	}
}
