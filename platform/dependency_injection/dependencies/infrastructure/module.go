package infrastructure

import (
	"fmt"
	"github.com/sarulabs/di"
	"gmarcial/eCommerce/platform/infrastructure/log"
	"go.uber.org/zap"
)

const (
	errBuildInfrastructureModule = "an error occurred to build infrastructure module"
)

//Build //Build construct the graph of dependencies of infrastructure module
func Build(builder *di.Builder) {
	err := builder.Add(
		di.Def{
			Name:  "logger",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return log.Build()
			},
			Close: func(obj interface{}) error {
				return obj.(*zap.SugaredLogger).Sync()
			},
		})

	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errBuildInfrastructureModule, err.Error())
		panic(errorMessage)
	}
}