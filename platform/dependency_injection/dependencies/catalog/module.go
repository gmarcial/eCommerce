package catalog

import (
	"fmt"
	"github.com/sarulabs/di"
	catalog "gmarcial/eCommerce/core/catalog/application"
	"gmarcial/eCommerce/platform/configuration"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data"
	"gmarcial/eCommerce/platform/infrastructure/adapters/catalog/data/memory"
	"gmarcial/eCommerce/platform/infrastructure/grpc/discount/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
)

const (
	errBuildCatalogModule = "an error occurred to build catalog module"
)

//Build construct the graph of dependencies of catalog module
func Build(builder *di.Builder, configuration *configuration.Configuration) {
	err := builder.Add(
		di.Def{
			Name:  "productRepository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				loadedProducts, err := data.LoadProducts()
				if err != nil {
					return nil, err
				}

				return memory.NewProductRepository(loadedProducts), nil
			},
			Close: nil,
		},
		di.Def{
			Name:  "discountServiceClient",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				address := configuration.GrpcServerAddress
				options := []grpc.DialOption{grpc.WithInsecure()}
				channel, _ := grpc.Dial(address, options...)

				return client.NewDiscountClient(channel), nil
			},
			Close: nil,
		},
		di.Def{
			Name:  "pickProductsUseCase",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				logger := ctn.Get("logger").(*zap.SugaredLogger)
				pickProductsUseCaseLogger := logger.Named("PickProductsUseCase")

				productRepository := ctn.Get("productRepository").(*memory.ProductRepository)
				discountServiceClient := ctn.Get("discountServiceClient").(client.DiscountClient)
				return catalog.NewPickProductsUseCase(pickProductsUseCaseLogger, productRepository, discountServiceClient), nil
			},
			Close: nil,
		},
		di.Def{
			Name:  "getGiftProductUseCase",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				logger := ctn.Get("logger").(*zap.SugaredLogger)
				getGiftProductUseCaseLogger := logger.Named("GetGiftProductUseCase")

				productRepository := ctn.Get("productRepository").(*memory.ProductRepository)
				return catalog.NewGetGiftProductUseCase(getGiftProductUseCaseLogger, productRepository), nil
			},
			Close: nil,
		})

	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errBuildCatalogModule, err.Error())
		log.Panic(errorMessage)
	}

	log.Print("the catalog module was constructed")
}
