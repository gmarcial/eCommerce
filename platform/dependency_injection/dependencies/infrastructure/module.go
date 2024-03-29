package infrastructure

import (
	"fmt"
	"github.com/sarulabs/di"
	logger "gmarcial/eCommerce/platform/infrastructure/log"
	"go.uber.org/zap"
	"log"
)

const (
	errBuildInfrastructureModule = "an error occurred to build infrastructure module"
)

//Build construct the graph of dependencies of infrastructure module
func Build(builder *di.Builder) {
	err := builder.Add(
		di.Def{
			Name:  "logger",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return logger.Build()
			},
			Close: func(obj interface{}) error {
				return obj.(*zap.SugaredLogger).Sync()
			},
		})

	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errBuildInfrastructureModule, err.Error())
		log.Panic(errorMessage)
	}
	log.Print("the infrastructure module was constructed")
}