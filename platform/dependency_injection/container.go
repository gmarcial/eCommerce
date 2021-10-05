package dependency_injection

import (
	"github.com/sarulabs/di"
	"gmarcial/eCommerce/platform/configuration"
	"gmarcial/eCommerce/platform/dependency_injection/dependencies/catalog"
	"gmarcial/eCommerce/platform/dependency_injection/dependencies/checkout"
	"gmarcial/eCommerce/platform/dependency_injection/dependencies/infrastructure"
	"log"
)

//BuildContainer construct the container of dependencies
func BuildContainer(configuration *configuration.Configuration) di.Container {
	builder, _ := di.NewBuilder()

	infrastructure.Build(builder)
	catalog.Build(builder, configuration)
	checkout.Build(builder, configuration)

	log.Print("the container was built.")
	return builder.Build()
}