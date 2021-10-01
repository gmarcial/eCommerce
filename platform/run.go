package platform

import (
	platform "gmarcial/eCommerce/platform/configuration"
	dependencyInjection "gmarcial/eCommerce/platform/dependency_injection"
	"gmarcial/eCommerce/platform/http"
)

//Run startup of platform
func Run() {
	configuration := platform.LoadConfiguration()

	container := dependencyInjection.BuildContainer(configuration)

	http.ListenAndServe(container, configuration)
}
