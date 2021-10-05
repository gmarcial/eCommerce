package platform

import (
	platform "gmarcial/eCommerce/platform/configuration"
	dependencyInjection "gmarcial/eCommerce/platform/dependency_injection"
	"gmarcial/eCommerce/platform/http"
	"log"
)

//Run startup of platform
func Run() {
	log.Print("starting the application")
	configuration := platform.LoadConfiguration()

	container := dependencyInjection.BuildContainer(configuration)

	http.ListenAndServe(container, configuration)
}
