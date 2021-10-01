package checkout

import (
	"fmt"
	"github.com/sarulabs/di"
	catalog "gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/core/checkout/application"
	"gmarcial/eCommerce/core/checkout/application/promotional"
	"gmarcial/eCommerce/core/checkout/domain/promotion"
	"gmarcial/eCommerce/platform/configuration"
)

const (
	errBuildCheckoutModule = "an error occurred to build checkout module"
)

//Build //Build construct the graph of dependencies of checkout module
func Build(builder *di.Builder, configuration *configuration.Configuration) {
	err := builder.Add(
		di.Def{
			Name:  "blackFridayPromotion",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				blackFridayDate := configuration.BlackFridayDay
				getGiftProductUseCase := ctn.Get("getGiftProductUseCase").(*catalog.GetGiftProductUseCase)
				return promotional.NewBlackFridayPromotion(blackFridayDate, getGiftProductUseCase), nil
			},
			Close: nil,
		},
		di.Def{
			Name:  "promotionsApplierUseCase",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				blackFridayPromotion := ctn.Get("blackFridayPromotion").(*promotional.BlackFridayPromotion)
				activePromotions := []promotion.Promotion{blackFridayPromotion}
				return promotional.NewPromotionsApplierUseCase(activePromotions), nil
			},
			Close: nil,
		},
		di.Def{
			Name:  "makeCartUseCase",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				pickProductsUseCase := ctn.Get("pickProductsUseCase").(*catalog.PickProductsUseCase)
				promotionsApplierUseCase := ctn.Get("promotionsApplierUseCase").(*promotional.PromotionsApplierUseCase)
				return application.NewMakeCartUseCase(pickProductsUseCase, promotionsApplierUseCase), nil
			},
			Close: nil,
		})

	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errBuildCheckoutModule, err.Error())
		panic(errorMessage)
	}
}
